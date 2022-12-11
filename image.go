package openaigo

import (
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
)

const (
	Size256  string = "256x256"
	Size512  string = "512x512"
	Size1024 string = "1024x1024"
)

type ImageGenerationRequestBody struct {
	Prompt         string `json:"prompt"`
	N              int    `json:"n,omitempty"`
	Size           string `json:"size,omitempty"`
	ResponseFormat string `json:"response_format,omitempty"`
	User           string `json:"user,omitempty"`
}

type ImageResponse struct {
	Created int64       `json:"created"`
	Data    []ImageData `json:"data"`
}

type ImageData struct {
	Base64 string `json:"b64_json"`
	URL    string `json:"url"`
}

type ImageGenerationResponse ImageResponse

type ImageEditRequestBody struct {
	// image Required
	// The image to use as the basis for the variation(s). Must be a valid PNG file, less than 4MB, and square.
	// User MUST close it if it's like ReadCloser.
	Image io.Reader

	// n integer Optional Defaults to 1
	// The number of images to generate. Must be between 1 and 10.
	N int

	// size string Optional Defaults to 1024x1024
	// The size of the generated images. Must be one of 256x256, 512x512, or 1024x1024.
	Size string

	// response_format string Optional Defaults to url
	// The format in which the generated images are returned. Must be one of url or b64_json.
	ResponseFormat string

	// user string Optional
	// A unique identifier representing your end-user, which can help OpenAI to monitor and detect abuse.
	// Learn more: https://beta.openai.com/docs/guides/safety-best-practices/end-user-ids
	User string

	// mask string Optional
	// An additional image whose fully transparent areas (e.g. where alpha is zero) indicate where image should be edited.
	// Must be a valid PNG file, less than 4MB, and have the same dimensions as image.
	// User MUST close it if it's like ReadCloser.
	Mask io.Reader

	// prompt string Required
	// A text description of the desired image(s). The maximum length is 1000 characters.
	Prompt string
}

func (body ImageEditRequestBody) ToMultipartFormData() (buf *bytes.Buffer, contenttype string, err error) {
	if body.Image == nil {
		return nil, "", fmt.Errorf("body.Image must not be nil")
	}
	buf = bytes.NewBuffer(nil)
	w := multipart.NewWriter(buf)
	imgw, err := w.CreateFormFile("image", "image.png")
	if err != nil {
		return nil, "", fmt.Errorf("failed to create FormFile: %v", err)
	}
	if _, err := io.Copy(imgw, body.Image); err != nil {
		return nil, "", fmt.Errorf("failed to copy io.Reader to buffer: %v", err)
	}

	if body.Mask != nil {
		maskw, err := w.CreateFormFile("mask", "mask.png")
		if err != nil {
			return nil, "", err
		}
		if _, err := io.Copy(maskw, body.Mask); err != nil {
			return nil, "", err
		}
	}

	// prompt is required for image edit.
	w.WriteField("prompt", body.Prompt)

	if body.N > 1 {
		w.WriteField("n", fmt.Sprintf("%d", body.N))
	}
	if body.Size != "" {
		w.WriteField("size", body.Size)
	}
	if body.ResponseFormat != "" {
		w.WriteField("response_format", body.ResponseFormat)
	}
	if body.User != "" {
		w.WriteField("user", body.User)
	}

	if err = w.Close(); err != nil {
		return nil, "", err
	}

	return buf, w.FormDataContentType(), err
}

type ImageEditResponse ImageResponse

type ImageVariationRequestBody struct {
	// image Required
	// The image to use as the basis for the variation(s). Must be a valid PNG file, less than 4MB, and square.
	// User MUST close it if it's like ReadCloser.
	Image io.Reader

	// n integer Optional Defaults to 1
	// The number of images to generate. Must be between 1 and 10.
	N int

	// size string Optional Defaults to 1024x1024
	// The size of the generated images. Must be one of 256x256, 512x512, or 1024x1024.
	Size string

	// response_format string Optional Defaults to url
	// The format in which the generated images are returned. Must be one of url or b64_json.
	ResponseFormat string

	// user string Optional
	// A unique identifier representing your end-user, which can help OpenAI to monitor and detect abuse.
	// Learn more: https://beta.openai.com/docs/guides/safety-best-practices/end-user-ids
	User string
}

func (body ImageVariationRequestBody) ToMultipartFormData() (buf *bytes.Buffer, contenttype string, err error) {
	if body.Image == nil {
		return nil, "", fmt.Errorf("body.Image must not be nil")
	}
	buf = bytes.NewBuffer(nil)
	w := multipart.NewWriter(buf)
	defer w.Close()
	imgw, err := w.CreateFormFile("image", "image.png")
	if err != nil {
		return nil, "", err
	}
	if _, err := io.Copy(imgw, body.Image); err != nil {
		return nil, "", err
	}
	if body.N > 1 {
		w.WriteField("n", fmt.Sprintf("%d", body.N))
	}
	if body.Size != "" {
		w.WriteField("size", body.Size)
	}
	if body.ResponseFormat != "" {
		w.WriteField("response_format", body.ResponseFormat)
	}
	if body.User != "" {
		w.WriteField("user", body.User)
	}
	return buf, w.FormDataContentType(), err
}

type ImageVariationResponse ImageResponse
