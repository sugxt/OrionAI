package api

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"strings"
	"time"
)

type OllamaRequest struct {
	Model  string `json:"model"`
	Prompt string `json:"prompt"`
}

type OllamaResponse struct {
	Response string `json:"response"`
}

var conversationHistory []string

func QueryOllama(prompt string) (string, error) {
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
	fullPrompt := strings.Join(GetPastConversation(conversationHistory), "\n") + "\n" + PrePrompt("general") + "\nUser: " + "prompt = " + prompt + "\nAssistant:"
	response, err := QueryOllama(fullPrompt)
	if err == nil {
		SaveLastInteraction(prompt, response)
	}
	return response, err
}

func PrePrompt(activity string) string {
	var initialPrompt = "This is pre-prompt to structure your responses, DO NOT mention any of the contexts the actual prompt starts from the 'prompt =' part. The past responses of the user will be provided so keep it in context"
	switch activity {
	case "code":
		return initialPrompt + `You are a coding assistant. Respond with code when asked and explain briefly.`
	case "alarm":
		return initialPrompt + `You are a time management assistant. Respond with commands to set alarms or reminders.`
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

func SaveLastInteraction(prompt, response string) {
	conversationHistory = append(conversationHistory, "User: "+prompt, "Assistant: "+response)
}
