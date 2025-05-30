package api

import (
	"context"
	"os"

	"github.com/wailsapp/wails/v2/pkg/runtime"
	"golang.design/x/clipboard"
)

var clipboardHistory []string

func (a *App) CreateTempDir() (string, error) {

	tempdir, err := os.MkdirTemp("", "clipboard_dir")
	if err != nil {
		return "", err
	}
	runtime.LogPrint(a.ctx, tempdir)

	return tempdir, nil
}

func (a *App) StartClipBoard() error {
	err := clipboard.Init()
	if err != nil {
		runtime.LogPrint(a.ctx, "Clipboard Initialization Failed")
		return err
	}

	ch := clipboard.Watch(context.TODO(), clipboard.FmtText)
	for data := range ch {
		runtime.LogPrint(a.ctx, string(data))
		//TODO: Append all copied data to an array and make the data persist over a few instances
		clipboardHistory = append(clipboardHistory, string(data))
		runtime.LogPrint(a.ctx, clipboardHistory[0])
		//TODO: Make changes in the frontend if the clipboard history changes

	}

	return nil
}

func (a *App) GetClipBoardHistory() []string {
	return clipboardHistory
}

func (a *App) EnchancementQuery(clip string, mode string) {

}
