package openaigo

import (
	"testing"

	. "github.com/otiai10/mint"
)

func TestClient_ChatCompletion(t *testing.T) {
	client := NewClient("")
	client.BaseURL = mockserver.URL
	res, err := client.ChatCompletion(nil, ChatCompletionRequestBody{})
	Expect(t, err).ToBe(nil)
	Expect(t, res).TypeOf("openaigo.ChatCompletionResponse")
	Expect(t, res.Stream()).ToBe(nil)
}

func TestClient_ChatCompletion_stream(t *testing.T) {
	client := NewClient("")
	client.BaseURL = mockserver.URL
	res, err := client.ChatCompletion(nil, ChatCompletionRequestBody{
		Stream: true,
	})
	Expect(t, err).ToBe(nil)
	Expect(t, res).TypeOf("openaigo.ChatCompletionResponse")
	Expect(t, res.Stream()).Not().ToBe(nil)
}
