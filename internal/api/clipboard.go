package api

import (
	"os"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

func (a *App) CreateTempDir() (string, error) {

	tempdir, err := os.MkdirTemp("", "clipboard_dir")
	if err != nil {
		return "", err
	}
	runtime.LogPrint(a.ctx, tempdir)

	return "", nil
}

func (a *App) GetFromClipBoard() bool {
	runtime.LogPrint(a.ctx, "The clipboard is live")
	return true
}
