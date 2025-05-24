package api

import (
	"fmt"
	"os/exec"
	goruntime "runtime" // Alias the Go runtime

	"github.com/wailsapp/wails/v2/pkg/runtime" // Wails runtime
)

func CheckFunction() {
	fmt.Println("This is working")
}

func (a *App) CheckDependencies() (isReady bool) {
	// Check if Ollama is installed by trying to get its version
	cmd := exec.Command("ollama", "--version")
	err := cmd.Run()
	if err != nil {
		fmt.Println("Ollama is not installed. Attempting installation...")
		InstallOllama()
		return false
	}
	runtime.LogPrintf(a.ctx, "Ollama is Already Installed!")
	return true
}

func InstallOllama() {
	var cmd *exec.Cmd

	switch goruntime.GOOS {
	case "windows":
		cmd = exec.Command("powershell", "iwr https://ollama.com/download/OllamaSetup.exe -OutFile OllamaSetup.exe; Start-Process .\\OllamaSetup.exe -Wait")
	case "linux":
		cmd = exec.Command("sh", "-c", "curl -fsSL https://ollama.com/install.sh | sh")
	case "darwin":
		cmd = exec.Command("sh", "-c", "curl -fsSL https://ollama.com/install.sh | sh")
	default:
		fmt.Println("Unsupported OS")
		return
	}

	err := cmd.Run()
	if err != nil {
		fmt.Printf("Failed to install Ollama: %v\n", err)
	} else {
		fmt.Println("Ollama installation attempted.")
	}
}
