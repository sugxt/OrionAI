package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type OllamaRequest struct {
	Model  string `json:"model"`
	Prompt string `json:"prompt"`
}

type OllamaResponse struct {
	Response string `json:"response"`
}

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

	// This response is JSON with a "response" field
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
	return QueryOllama(prompt)
}

func (a *App) CopyToClipboard(content string) string {
	return fmt.Sprintf("Copied %s", content)
}
