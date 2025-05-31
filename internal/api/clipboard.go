package api

import (
	"context"
	"os"

	"github.com/wailsapp/wails/v2/pkg/runtime"
	"golang.design/x/clipboard"
)

var clipboardHistory []string
var responseHistory []string
var isAIResponse bool = false

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
	//TODO: Add more
	ch := clipboard.Watch(context.TODO(), clipboard.FmtText)
	for data := range ch {
		if !isAIResponse {
			clipboardHistory = append(clipboardHistory, string(data))
			runtime.EventsEmit(a.ctx, "clipboardHistoryUpdated", clipboardHistory)
			err := a.EnchancementQuery(clipboardHistory[len(clipboardHistory)-1], "sum")
			if err != nil {
				runtime.LogPrint(a.ctx, "Enhancement failed: "+err.Error())
			}
		} else {
			isAIResponse = false
		}
	}

	return nil
}

func (a *App) EnchancementQuery(clip string, mode string) error {
	runtime.LogPrintf(a.ctx, clip)
	runtime.LogPrintf(a.ctx, mode)
	//TODO: Add better handling of different cases and add more cases
	switch mode {
	case "sum":
		res, err := QueryOllama("Expand the provided text to have a little more content in it and also keep it very short and simple, do not expand more than 3-4 sentences :"+clip, false)
		if err != nil {
			return err
		}
		responseHistory = append(responseHistory, string(res))
		clipboard.Write(clipboard.FmtText, []byte(res))
		//TODO: Implement this in the frontend and add a different panel for responses
		runtime.EventsEmit(a.ctx, "responseHistoryUpdated", responseHistory)
		runtime.LogPrint(a.ctx, "Copied")
		isAIResponse = true
	default:
		QueryOllama("Expand the provided text to have context to it, keep it short and simple", false)
	}
	return nil

}
