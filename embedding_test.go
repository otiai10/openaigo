package openaigo

import (
	"testing"

	. "github.com/otiai10/mint"
)

func TestClient_CreateEmbedding(t *testing.T) {
	client := NewClient("")
	client.BaseURL = mockserver.URL
	res, err := client.CreateEmbedding(nil, EmbeddingCreateRequestBody{
		Model: "text-similarity-babbage-001",
		Input: []string{"The food was delicious and the waiter..."},
	})
	Expect(t, err).ToBe(nil)
	Expect(t, res).TypeOf("openaigo.EmbeddingCreateResponse")
}
