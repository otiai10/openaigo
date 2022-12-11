package openaigo

import (
	"os"
	"testing"

	. "github.com/otiai10/mint"
)

func TestClient_CreateImage(t *testing.T) {
	client := NewClient("")
	client.BaseURL = mockserver.URL
	res, err := client.CreateImage(nil, ImageGenerationRequestBody{
		Prompt: "A cute baby sea otter",
	})
	Expect(t, err).ToBe(nil)
	Expect(t, res).TypeOf("openaigo.ImageGenerationResponse")
}

func TestClient_EditImage(t *testing.T) {
	f, err := os.Open("./testdata/baby-sea-otter.png")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	mask, err := os.Open("./testdata/baby-sea-otter.png")
	if err != nil {
		panic(err)
	}
	defer mask.Close()

	res, err := (&Client{BaseURL: mockserver.URL}).EditImage(nil, ImageEditRequestBody{
		Image:  f,
		Prompt: "make it cuter",
	})
	Expect(t, err).ToBe(nil)
	Expect(t, res).TypeOf("openaigo.ImageEditResponse")

	res, err = (&Client{BaseURL: mockserver.URL}).EditImage(nil, ImageEditRequestBody{
		Image:          f,
		Prompt:         "make it cuter",
		Mask:           mask,
		N:              6,
		Size:           Size512,
		ResponseFormat: "url",
		User:           "otiai20",
	})
	Expect(t, err).ToBe(nil)
	Expect(t, res).TypeOf("openaigo.ImageEditResponse")

	_, err = (&Client{BaseURL: mockserver.URL}).EditImage(nil, ImageEditRequestBody{
		Image:  nil,
		Prompt: "make it cuter",
	})
	Expect(t, err).Not().ToBe(nil)
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

	res, err = (&Client{BaseURL: mockserver.URL}).CreateImageVariation(nil, ImageVariationRequestBody{
		Image:          f,
		N:              4,
		Size:           Size256,
		ResponseFormat: "b64_json",
		User:           "otiai10",
	})
	Expect(t, err).ToBe(nil)
	Expect(t, res).TypeOf("openaigo.ImageVariationResponse")

	_, err = (&Client{BaseURL: mockserver.URL}).CreateImageVariation(nil, ImageVariationRequestBody{
		Image: nil,
	})
	Expect(t, err).Not().ToBe(nil)
}
