package llmclient

type IProviderClient interface {
	SendRequest(prompt string) ([]byte, error)
	GetTextFromResponse(apiResponse []byte) (string, error)
}
