package llmclient

type IProviderClient interface {
	SendRequest(prompt string) ([]byte, error)
	GetAnswerFromResponse(apiResponse []byte) (string, error)
}
