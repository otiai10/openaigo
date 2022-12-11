package openaigo

import (
	"context"
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

func TestClient_RetrieveModel(t *testing.T) {
	client := NewClient("")
	client.BaseURL = mockserver.URL
	res, err := client.RetrieveModel(nil, "text-davinci-003")
	Expect(t, err).ToBe(nil)
	Expect(t, res).TypeOf("openaigo.ModelRetrieveResponse")

	_, err = client.RetrieveModel(context.TODO(), "200-but-invalidjson")
	Expect(t, err).Not().ToBe(nil)
}
