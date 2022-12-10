package openaigo

type EmbeddingCreateRequestBody struct {
	Model string   `json:"model"`
	Input []string `json:"input"`
	User  string   `json:"user,omitempty"`
}

type EmbeddingCreateResponse struct {
	Object string          `json:"object"`
	Data   []EmbeddingData `json:"data"`
	Usage  Usage           `json:"usage"`
}

type EmbeddingData struct {
	Object    string    `json:"object"`
	Embedding []float32 `json:"embedding"`
	Index     int       `json:"index"`
}
