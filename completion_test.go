package openaigo

import (
	"testing"

	. "github.com/otiai10/mint"
)

func TestClient_Completion(t *testing.T) {
	client := NewClient("")
	client.BaseURL = mockserver.URL
	res, err := client.Completion_Legacy(nil, CompletionRequestBody{})
	Expect(t, err).ToBe(nil)
	Expect(t, res).TypeOf("openaigo.CompletionResponse")

	// client.BaseURL = "xxxxx"
	// _, err = client.Completion(nil, CompletionRequestBody{})
	// Expect(t, err).Not().ToBe(nil)
}
