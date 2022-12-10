package openaigo

import (
	"testing"

	. "github.com/otiai10/mint"
)

func TestNewClient(t *testing.T) {
	Expect(t, NewClient("xxx")).TypeOf("*openaigo.Client")
}
