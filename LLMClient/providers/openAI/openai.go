package OpenAI

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
)

type OpenAIClient struct {
	APIKey   string
	GptModel GptModel
}

func (oc *OpenAIClient) GenerateResponse(prompt string) ([]byte, error) {
	url := "https://api.openai.com/v1/responses"
	contentType := "application/json"
	requestBody := map[string]interface{}{
		"model": oc.GptModel,
		"input": prompt,
	}
	jsonData, err := json.Marshal(requestBody)

	if err != nil {
		var zero []byte
		return zero, err
	}

	resp, err := http.Post(url, contentType, bytes.NewBuffer(jsonData))

	if err != nil {
		var zero []byte
		return zero, err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		var zero []byte
		return zero, err
	}

	return body, nil
}
