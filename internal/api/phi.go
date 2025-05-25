package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

type OllamaRequest struct {
	Model  string `json:"model"`
	Prompt string `json:"prompt"`
}

type OllamaResponse struct {
	Response string `json:"response"`
}

var conversationHistory []string

func QueryOllama(prompt string, isExec bool) (string, error) {
	body := map[string]interface{}{
		"model":  "phi3",
		"prompt": prompt,
		"stream": false,
	}
	jsonBody, _ := json.Marshal(body)

	resp, err := http.Post("http://localhost:11434/api/generate", "application/json", bytes.NewReader(jsonBody))
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	data, _ := io.ReadAll(resp.Body)

	var result struct {
		Response string `json:"response"`
	}
	err = json.Unmarshal(data, &result)
	if err != nil {
		return "", err
	}

	if isExec {
		TaskExecution(result.Response)
	}

	return result.Response, nil
}

func IsOllamaRunning() bool {
	client := http.Client{
		Timeout: 1 * time.Second,
	}
	resp, err := client.Get("http://localhost:11434")
	if err != nil {
		return false
	}
	defer resp.Body.Close()
	return resp.StatusCode == http.StatusOK
}

func (a *App) AskOllama(prompt string) (string, error) {
	fullPrompt := strings.Join(GetPastConversation(conversationHistory), "\n") + "\n" + PrePrompt("general") + "\nUser: " + "new prompt = " + prompt + "\nAssistant:"
	runtime.LogPrintf(a.ctx, fullPrompt)
	response, err := QueryOllama(fullPrompt, false)
	if err == nil {
		SaveLastInteraction(prompt, response)
	}
	return response, err
}

func PrePrompt(activity string) string {
	userDetails := GetUserDetails()
	var initialPrompt = "This is pre-prompt to structure your responses,DO NOT mention any of the contexts the actual prompt starts from the 'new prompt =' part, ONLY  reply to the last prompt. The past responses of the user will be provided so keep it in context. Keep the answers fairly short. DO NOT mention anything about the chat history that is provided before the prompt \n These are the user details, you can greet the user to make it more personalized: " + userDetails
	switch activity {
	case "code":
		return initialPrompt + `You are a coding assistant. Respond with code when asked and explain briefly.`
	case "alarm":
		return initialPrompt + `You are a time management assistant. Respond with commands to set alarms or reminders.`
	case "tasks":
		return initialPrompt + `You are a task execution assistaant. I will provide a list of tasks and you are ONLY to return the keyword of what category the task execution lies in. Here are the categories = \n "code","browser","alarm","reboot","shutdown"`
	case "general":
		fallthrough
	default:
		return initialPrompt + `You are a helpful desktop assistant. Answer questions in a short and sweet manner`
	}
}

func GetPastConversation(pastConvo []string) []string {
	// Limit to last 10 messages to stay within token limit
	//TODO: Make the last 10 messages persist by saving it inside a prompt file
	if len(pastConvo) > 10 {
		return pastConvo[len(pastConvo)-10:]
	}
	return pastConvo
}

func TaskExecution(taskType string) (isExecuted bool) {
	fmt.Sprintln(taskType)
	return false
}

func SaveLastInteraction(prompt, response string) {
	conversationHistory = append(conversationHistory, "User: "+prompt, "Assistant: "+response)
}
