package main

import (
	"encoding/json"
	"fmt"
	"reflect"

	openai "github.com/oskareek/llmgo/llmclient/providers/openai"
	manager "github.com/oskareek/llmgo/manager"
)

type Book struct {
	Name   string
	Author string
	Rating int
}

func main() {
	//test1()
	test2()
}

func test1() {
	client := openai.OpenAIClient{
		APIKey:   "",
		GptModel: openai.GptModels.Gpt_35_turbo,
	}

	llmManager := manager.ClientManager{Client: &client}

	response, err := llmManager.GenerateResponse("Can you give me 10 popular books?")
	if err != nil {
		fmt.Printf("%s", err)
		return
	}

	fmt.Print("\n" + string(response) + "\n")
}

func test2() {
	client := openai.OpenAIClient{
		APIKey:   "",
		GptModel: openai.GptModels.Gpt_35_turbo,
	}

	llmManager := manager.ClientManager{Client: &client}

	var obj []Book
	response, err := llmManager.GenerateJsonResponse("Can you give me 10 popular books?", reflect.TypeOf(obj))
	if err != nil {
		fmt.Printf("%s", err)
		return
	}

	fmt.Print("\n" + string(response) + "\n")

	err = json.Unmarshal(response, &obj)
	if err != nil {
		fmt.Printf("%s", err)
		return
	}

	fmt.Printf("%+v\n", obj)
}
