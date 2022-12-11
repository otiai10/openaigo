package openaigo

import (
	"context"
	"testing"

	. "github.com/otiai10/mint"
)

func TestNewClient(t *testing.T) {
	Expect(t, NewClient("xxx")).TypeOf("*openaigo.Client")
}

func TestClient_internal(t *testing.T) {
	client := NewClient("invalid")
	client.Organization = "org-xxx"
	_, err := client.ListModels(context.TODO())
	Expect(t, err).Not().ToBe(nil)

	client.BaseURL = mockserver.URL
	_, err = client.RetrieveModel(nil, "notfound")
	Expect(t, err).Not().ToBe(nil)
}
