package openaigo

type EditCreateRequestBody struct {
	Model       string  `json:"model"`
	Instruction string  `json:"instruction"`
	Input       string  `json:"input,omitempty"`
	N           int     `json:"n,omitempty"`
	Temperature float32 `json:"temperature,omitempty"`
	TopP        float32 `json:"top_p,omitempty"`
}

type EditCreateResponse struct {
	Object  ObjectType         `json:"object"`
	Created int64              `json:"created"`
	Choices []CompletionChoice `json:"choices"`
	Usage   Usage              `json:"usage"`
}
