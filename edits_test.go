package openaigo

import (
	"testing"

	. "github.com/otiai10/mint"
)

func TestClient_CreateEdit(t *testing.T) {
	client := NewClient("")
	client.BaseURL = mockserver.URL
	res, err := client.CreateEdit(nil, EditCreateRequestBody{})
	Expect(t, err).ToBe(nil)
	Expect(t, res).TypeOf("openaigo.EditCreateResponse")
}
