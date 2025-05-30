package api

import (
	"fmt"
	"os/exec"
	goruntime "runtime"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

func (a *App) CheckDependencies() (isReady bool) {
	path, err := exec.LookPath("ollama")
	if err != nil {
		fmt.Println("Ollama not found, starting async installation...")
		go a.InstallOllamaAndPhi3()
		return false
	}
	runtime.LogPrintf(a.ctx, "Ollama found at: %s", path)
	return true
}

func (a *App) InstallOllamaAndPhi3() {
	var installCmd *exec.Cmd

	switch goruntime.GOOS {
	case "windows":
		installCmd = exec.Command("powershell", "iwr https://ollama.com/download/OllamaSetup.exe -OutFile OllamaSetup.exe; Start-Process .\\OllamaSetup.exe -Wait")
	case "linux", "darwin":
		installCmd = exec.Command("sh", "-c", "curl -fsSL https://ollama.com/install.sh | sh")
	default:
		runtime.LogError(a.ctx, "Unsupported OS for Ollama installation.")
		runtime.EventsEmit(a.ctx, "ollama:install_failed", "Unsupported OS")
		return
	}

	out, err := installCmd.CombinedOutput()
	if err != nil {
		runtime.LogError(a.ctx, fmt.Sprintf(" Ollama install failed: %v\nOutput: %s", err, string(out)))
		runtime.EventsEmit(a.ctx, "ollama:install_failed", string(out))
		return
	}

	runtime.LogPrintf(a.ctx, "Ollama installed. Output: %s", string(out))
	runtime.EventsEmit(a.ctx, "ollama:installed", "Ollama installed successfully.")

	runtime.LogPrintf(a.ctx, "Pulling phi3 model...")
	pullCmd := exec.Command("ollama", "pull", "phi3")
	pullOut, pullErr := pullCmd.CombinedOutput()

	if pullErr != nil {
		runtime.LogError(a.ctx, fmt.Sprintf("Failed to pull phi3: %v\nOutput: %s", pullErr, string(pullOut)))
		runtime.EventsEmit(a.ctx, "phi3:pull_failed", string(pullOut))
		return
	}

	runtime.LogPrintf(a.ctx, "phi3 pulled successfully. Output: %s", string(pullOut))
	runtime.EventsEmit(a.ctx, "phi3:installed", "Phi3 model pulled and ready.")
}
