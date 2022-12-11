package openaigo

import (
	"testing"

	. "github.com/otiai10/mint"
)

func TestClient_CreateModeration(t *testing.T) {
	client := NewClient("")
	client.BaseURL = mockserver.URL
	res, err := client.CreateModeration(nil, ModerationCreateRequestBody{
		Input: "I want to kixx you.",
	})
	Expect(t, err).ToBe(nil)
	Expect(t, res).TypeOf("openaigo.ModerationCreateResponse")
}
