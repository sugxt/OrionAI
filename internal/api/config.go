package api

import (
	"context"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

type App struct {
	ollamaCmd *exec.Cmd
	ctx       context.Context
}

func (a *App) Startup(ctx context.Context) {
	a.ctx = ctx
	var isReady = a.CheckDependencies()
	if isReady {
		runtime.LogError(a.ctx, "Dependencies Installed")
	}

	err := a.StartClipBoard()
	if err != nil {
		runtime.LogError(a.ctx, "It is not working")
	}

	err = a.StartOllamaModel()
	if err != nil {
		runtime.LogErrorf(a.ctx, "Failed to start Ollama:")
	}
	a.GetUserDetails()
}
func NewApp() *App {
	app := &App{}

	return app
}

func (a *App) StartOllamaModel() error {
	cmd := exec.Command("ollama", "run", "phi3")

	if err := cmd.Start(); err != nil {
		return err
	}

	a.ollamaCmd = cmd
	return nil
}

func (a *App) SaveConfig(data string) error {
	path, err := getConfigPath()
	if err != nil {
		return err
	}
	return os.WriteFile(path, []byte(data), 0644)
}

func (a *App) LoadConfig() (string, error) {
	path, err := getConfigPath()
	if err != nil {
		return "", err
	}

	content, err := os.ReadFile(path)
	if err != nil {
		if os.IsNotExist(err) {
			return "", nil
		}
		return "", err
	}

	return string(content), nil
}

func (a *App) DeleteConfig() (string, error) {
	path, err := getConfigPath()
	if err != nil {
		return "", err
	}
	err = os.Remove(path)
	if err != nil {
		return "", err
	}
	return "deleted", nil
}

func (a *App) Shutdown(ctx context.Context) {
	if a.ollamaCmd != nil && a.ollamaCmd.Process != nil {
		_ = a.ollamaCmd.Process.Kill()
	}
}

func getConfigPath() (string, error) {
	baseDir, err := os.UserConfigDir()
	if err != nil {
		return "", err
	}

	cliprDir := filepath.Join(baseDir, "clipr")

	err = os.MkdirAll(cliprDir, os.ModePerm)
	if err != nil {
		return "", err
	}

	return filepath.Join(cliprDir, "config.json"), nil
}

func (a *App) GetUserDetails() (userInfo string) {
	userDetails, err := a.LoadConfig()
	if err != nil {
		return ""
	}
	return userDetails
}
