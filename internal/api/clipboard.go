package api

import (
	"context"

	"github.com/wailsapp/wails/v2/pkg/runtime"
	"golang.design/x/clipboard"
)

var clipboardHistory []string
var responseHistory []string
var isAIResponse bool = false
var currentMode string

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
			err := a.EnchancementQuery(clipboardHistory[len(clipboardHistory)-1], currentMode)
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
	//TODO: Handle different modes efficiently
	res, err := QueryOllama(a.ModeQuery(mode)+clip, false)
	if err != nil {
		return err
	}
	responseHistory = append(responseHistory, string(res))
	clipboard.Write(clipboard.FmtText, []byte(res))
	runtime.EventsEmit(a.ctx, "responseHistoryUpdated", responseHistory)
	runtime.LogPrint(a.ctx, "Copied")
	//Setting the bool to true so the bot doesn't go in a loop of doing shit to it's own text
	isAIResponse = true
	return nil

}

func (a *App) ClearHistory(list string) {
	if list == "clipboard" {
		clipboardHistory = []string{}
		runtime.EventsEmit(a.ctx, "clipboardHistoryUpdated", clipboardHistory)

	} else {
		responseHistory = []string{}
		runtime.EventsEmit(a.ctx, "responseHistoryUpdated", responseHistory)
	}
}

func (a *App) SetMode(mode string) {
	currentMode = mode
	runtime.LogPrintf(a.ctx, "Mode Changed")
}

func (a *App) ModeQuery(mode string) string {
	switch mode {
	case "sum":
		// Structure the prompt so the bot doesn't end up increasing the text rather than summarizing it
		return "You are a bot that will simply just summarize the text provided into a shorter and compact version, the actual prompt is here: "
	case "expand":
		//Add More Precise Prompting To Generate the Ideal Response
		return "Expand the provided text to have a little more content in it and also keep it VERY VERY short and simple,DO NOT add more than 2-3 sentences and only provide clean text, don't refer to anything else: "
	case "para":
		return "Paraphrase the given text and change it to be different but make sure not to sway too far from the topic: "
	default:
		return "Tell the user to provide context before copying anything from the clipboard: "
	}
}
