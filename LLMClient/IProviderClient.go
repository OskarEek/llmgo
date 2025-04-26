package llmclient

type IProviderClient interface {
	GenerateResponse(prompt string) ([]byte, error)
}
