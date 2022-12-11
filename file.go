package openaigo

import (
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
)

type FileListResponse struct {
	Object string     `json:"object"`
	Data   []FileData `json:"data"`
}

type FileData struct {
	ID        string `json:"id"`
	Object    string `json:"object"`
	Bytes     int64  `json:"bytes"`
	CreatedAt int64  `json:"created_at"`
	Filename  string `json:"filename"`
	Purpose   string `json:"purpuse"`
}

type FileUploadRequestBody struct {
	File    io.Reader
	Purpose string
}

func (body FileUploadRequestBody) ToMultipartFormData() (*bytes.Buffer, string, error) {
	if body.File == nil {
		return nil, "", fmt.Errorf("body.File must not be nil")
	}
	buf := bytes.NewBuffer(nil)
	w := multipart.NewWriter(buf)
	defer w.Close()
	filew, err := w.CreateFormFile("file", "file.jsonl")
	if err != nil {
		return nil, "", err
	}
	if _, err := io.Copy(filew, body.File); err != nil {
		return nil, "", err
	}
	w.WriteField("purpose", body.Purpose)
	return buf, w.FormDataContentType(), err

}

type FileUploadResponse FileData

type FileDeleteResponse struct {
	ID      string `json:"id"`
	Object  string `json:"object"`
	Deleted bool   `json:"deleted"`
}

type FileRetrieveResponse FileData
