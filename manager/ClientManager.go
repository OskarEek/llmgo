package manager

import (
	"fmt"
	"reflect"

	llm "github.com/oskareek/llmgo/llmclient"
	jsonHelper "github.com/oskareek/llmgo/utilities/jsonhelper"
)

type ClientManager struct {
	Client llm.IProviderClient
}

func (cm *ClientManager) GenerateResponse(prompt string) (string, error) {
	resp, err := cm.Client.SendRequest(prompt)
	if err != nil {
		return "", err
	}

	text, err := cm.Client.GetTextFromResponse(resp)
	if err != nil {
		return "", err
	}

	return string(text), nil
}

func (cm *ClientManager) GenerateJsonResponse(prompt string, mapType reflect.Type) ([]byte, error) {
	jsonFormat, err := jsonHelper.GetJsonStructureFromType(mapType)
	if err != nil {
		var zero []byte
		return zero, err
	}

	//Tell LLM to answer with JSON data based on the provided type
	jsonPrompt :=
		"\n I want you to strictly and only respond with json data than can be mapped to the following structure: \n" +
			jsonFormat +
			"\nDo not add any additional words. Highest priority is to be able to map your response directly to the provided structure"
	prompt += jsonPrompt

	resp, err := cm.Client.SendRequest(prompt)
	fmt.Print("\n" + string(resp) + "\n")
	if err != nil {
		var zero []byte
		return zero, err
	}

	text, err := cm.Client.GetTextFromResponse(resp)
	if err != nil {
		var zero []byte
		return zero, err
	}

	return []byte(text), nil
}
