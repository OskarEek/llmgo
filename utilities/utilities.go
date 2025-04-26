package utilities

import (
	"encoding/json"
	"fmt"
	"reflect"

	llm "github.com/OskarEek/llmgo/LLMClient"
)

func GenerateJsonResponse[T any](client llm.LLMProviderClient, prompt string) (T, error) {
	var obj T
	t := reflect.TypeOf(obj)

	var jsonFormat = ""
	//Handle structs
	if t.Kind() == reflect.Struct {
		structure, err := GetJsonStructureFromStructOfType(t)

		if err != nil {
			var zero T
			return zero, err
		}

		jsonFormat += structure

	} else { //Handle all other types
		jsonFormat = fmt.Sprintf("%s", t)
	}

	//TODO: Handle list of structs?

	//Tell LLM to answere with JSON data
	jsonPrompt := "\n I want you to strictly and only respond with json data than can be mapped to the following structure: \n" +
		jsonFormat + "\n" +
		"Do not add any additional words."

	prompt += jsonPrompt

	result, err := client.GenerateResponse(prompt)

	err := json.Marshal()
	client.GenerateResponse()

}

func GetJsonStructureFromStructOfType(t reflect.Type) (string, error) {
	result := fmt.Sprintf("{\n")

	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		// Recursively handle nested struct
		if f.Type.Kind() == reflect.Struct {

			if f.Type.Kind() != t.Kind() {
				return "", fmt.Errorf("Self referential structs is not allowed")
			}

			subFieldStructure, err := GetJsonStructureFromStructOfType(f.Type)

			if err != nil {
				return "", err
			}

			result += fmt.Sprintf("    %s: %s,\n", f.Name, subFieldStructure)

		} else {
			result += fmt.Sprintf("    %s: (%s),\n", f.Name, f.Type)
		}
	}

	result += "}"

	return result, nil
}
