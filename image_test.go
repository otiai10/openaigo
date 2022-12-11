package openaigo

import (
	"os"
	"testing"

	. "github.com/otiai10/mint"
)

func TestClient_EditImage(t *testing.T) {
	f, err := os.Open("./testdata/baby-sea-otter.png")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	res, err := (&Client{BaseURL: mockserver.URL}).EditImage(nil, ImageEditRequestBody{
		Image:  f,
		Prompt: "make it cuter",
	})
	Expect(t, err).ToBe(nil)
	Expect(t, res).TypeOf("openaigo.ImageEditResponse")
}

func TestClient_CreateImageVariation(t *testing.T) {
	f, err := os.Open("./testdata/baby-sea-otter.png")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	res, err := (&Client{BaseURL: mockserver.URL}).CreateImageVariation(nil, ImageVariationRequestBody{
		Image: f,
	})
	Expect(t, err).ToBe(nil)
	Expect(t, res).TypeOf("openaigo.ImageVariationResponse")
}
