package manager

import (
	"reflect"

	llm "github.com/OskarEek/llmgo/LLMClient"
	jsonHelper "github.com/OskarEek/llmgo/utilities/jsonHelper"
)

type ClientManager struct {
	Client llm.IProviderClient
}

func (cm *ClientManager) GenerateResponse(prompt string) (string, error) {
	resp, err := cm.Client.GenerateResponse(prompt)
	if err != nil {
		return "", err
	}
	return string(resp), nil
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

	result, err := cm.Client.GenerateResponse(prompt)
	if err != nil {
		var zero []byte
		return zero, err
	}

	return result, nil
}
