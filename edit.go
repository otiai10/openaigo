package openaigo

type EditRequestBody struct {
	Model       string  `json:"model"`
	Instruction string  `json:"instruction"`
	Input       string  `json:"input,omitempty"`
	N           int     `json:"n,omitempty"` // FIXME: If 0 given, omitted.
	Temperature float32 `json:"temperature,omitempty"`
	TopP        float32 `json:"top_p,omitempty"`
}

type EditResponse struct {
	Object  ObjectType `json:"object"`
	Created int64      `json:"craeted"`
	Choices []Choice   `json:"choices"`
	Usage   Usage      `json:"usage"`
}
