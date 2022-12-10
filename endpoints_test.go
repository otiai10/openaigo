package openaigo

import (
	"os"
	"testing"

	. "github.com/otiai10/mint"
)

func TestMain(m *testing.M) {
	code := m.Run()
	os.Exit(code)
}

func TestClient_ListModels(t *testing.T) {
	Expect(t, true).ToBe(true) // TODO:
}
