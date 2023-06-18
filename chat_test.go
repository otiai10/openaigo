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

func TestClient_ChatCompletion_FunctionCall(t *testing.T) {
	client := NewClient("")
	client.BaseURL = mockserver.URL
	res, err := client.Chat(nil, ChatRequest{
		Model: GPT3_5Turbo,
		Messages: []Message{
			{
				Role: "user", Content: "Hello, I'm John.",
			},
		},
		Functions: []Function{
			{
				Name: "test_method",
				Parameters: Parameters{
					Type: "object",
					Properties: map[string]map[string]any{
						"arg_0": {
							"type":        "string",
							"description": "This is a test",
						},
					},
					Required: []string{"arg_0"},
				},
			},
		},
		FunctionCall: "auto",
	})
	Expect(t, err).ToBe(nil)
	Expect(t, res).TypeOf("openaigo.ChatCompletionResponse")
}
