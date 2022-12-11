package openaigo

import (
	"testing"

	. "github.com/otiai10/mint"
)

func TestClient_ListModels(t *testing.T) {
	client := NewClient("")
	client.BaseURL = mockserver.URL
	res, err := client.ListModels(nil)
	Expect(t, err).ToBe(nil)
	Expect(t, res).TypeOf("openaigo.ModelsListResponse")
}
