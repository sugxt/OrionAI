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
	// Optionally start Ollama here:
	err := a.StartOllamaModel()
	if err != nil {
		runtime.LogErrorf(a.ctx, "Failed to start Ollama:")
	}
	GetUserDetails()
}
func NewApp() *App {
	app := &App{}

	return app
}

func (a *App) StartOllamaModel() error {
	cmd := exec.Command("ollama", "run", "phi3")

	// Optional: for debugging, you can attach stdout/stderr
	// cmd.Stdout = os.Stdout
	// cmd.Stderr = os.Stderr

	if err := cmd.Start(); err != nil {
		return err
	}

	a.ollamaCmd = cmd
	return nil
}

func (a *App) SaveConfig(data []byte) error {
	path, err := getConfigPath()
	if err != nil {
		return err
	}

	return os.WriteFile(path, data, 0644)
}

func LoadConfig() (string, error) {
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

func (a *App) Shutdown(ctx context.Context) {
	if a.ollamaCmd != nil && a.ollamaCmd.Process != nil {
		_ = a.ollamaCmd.Process.Kill()
	}
}

func getConfigPath() (string, error) {
	// Get the system's user config directory
	baseDir, err := os.UserConfigDir()
	if err != nil {
		return "", err
	}

	// Build path to your app-specific folder
	cliprDir := filepath.Join(baseDir, "clipr")

	// Make sure the directory exists
	err = os.MkdirAll(cliprDir, os.ModePerm)
	if err != nil {
		return "", err
	}

	// Final path: ~/.config/clipr/config.json
	return filepath.Join(cliprDir, "config.json"), nil
}

func GetUserDetails() (userInfo string) {
	//TODO: Read the config file of the user and send it to the PrePrompt Function to Append Before Querying the Model
	userDetails, err := LoadConfig()
	if err != nil {
		return "No User Data"
	}
	// runtime.LogPrint(a.ctx, userDetails)
	return userDetails
}
