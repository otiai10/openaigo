package openaigo

import (
	"sync"
	"testing"

	. "github.com/otiai10/mint"
)

func TestClient_ChatCompletion(t *testing.T) {
	client := NewClient("")
	client.BaseURL = mockserver.URL
	res, err := client.ChatCompletion(nil, ChatCompletionRequestBody{
		Model: GPT3_5Turbo,
	})
	Expect(t, err).ToBe(nil)
	Expect(t, res).TypeOf("openaigo.ChatCompletionResponse")
}

func TestClient_ChatCompletion_Stream(t *testing.T) {
	client := NewClient("")
	client.BaseURL = mockserver.URL
	wg := sync.WaitGroup{}
	wg.Add(2)
	res, err := client.ChatCompletion(nil, ChatCompletionRequestBody{
		Model:  GPT3_5Turbo,
		Stream: true,
		StreamCallback: func(res ChatCompletionResponse, done bool, err error) {
			Expect(t, err).ToBe(nil)
			wg.Done()
		},
	})
	Expect(t, err).ToBe(nil)
	Expect(t, res).TypeOf("openaigo.ChatCompletionResponse")
	wg.Wait()
}
