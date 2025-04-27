package openai

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

type OpenAIClient struct {
	APIKey string
	Model  GptModel
}

func (oc *OpenAIClient) SendRequest(prompt string) ([]byte, error) {
	url := "https://api.openai.com/v1/responses"
	requestBody := map[string]interface{}{
		"model": oc.Model,
		"input": prompt,
	}

	jsonData, err := json.Marshal(requestBody)
	if err != nil {
		var zero []byte
		return zero, err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		var zero []byte
		return zero, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+oc.APIKey)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		var zero []byte
		return zero, err
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		var zero []byte
		return zero, err
	}
	return body, nil
}

func (oc *OpenAIClient) GetAnswerFromResponse(apiResponse []byte) (string, error) {
	var response Response
	err := json.Unmarshal(apiResponse, &response)
	if err != nil {
		return "", err
	}

	if len(response.Output) > 0 && len(response.Output[0].Content) > 0 {
		text := cleanupJson(response.Output[0].Content[0].Text)
		return text, nil
	}

	return "", fmt.Errorf("no valid content found in the response")
}

func cleanupJson(dirtyJson string) string {
	dirtyJson = strings.TrimSpace(dirtyJson)
	dirtyJson = strings.TrimPrefix(dirtyJson, "```json")
	dirtyJson = strings.TrimSuffix(dirtyJson, "```")
	return dirtyJson
}
