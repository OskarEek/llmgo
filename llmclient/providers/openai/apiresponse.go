package openai

type Response struct {
	ID     string `json:"id"`
	Object string `json:"object"`
	Output []struct {
		Content []struct {
			Type string `json:"type"`
			Text string `json:"text"`
		} `json:"content"`
	} `json:"output"`
}
