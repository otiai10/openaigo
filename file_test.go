package openaigo

import (
	"os"
	"testing"

	. "github.com/otiai10/mint"
)

func TestClient_UploadFile(t *testing.T) {
	client := NewClient("")
	client.BaseURL = mockserver.URL
	_, err := client.UploadFile(nil, FileUploadRequestBody{})
	Expect(t, err).Not().ToBe(nil)
	f, _ := os.Open("./testdata/train.jsonl")
	res, err := client.UploadFile(nil, FileUploadRequestBody{File: f})
	Expect(t, err).ToBe(nil)
	Expect(t, res).TypeOf("openaigo.FileUploadResponse")
}

func TestClient_ListFiles(t *testing.T) {
	client := NewClient("")
	client.BaseURL = mockserver.URL
	res, err := client.ListFiles(nil)
	Expect(t, err).ToBe(nil)
	Expect(t, res).TypeOf("openaigo.FileListResponse")
}

func TestClient_RetrieveFileContent(t *testing.T) {
	client := NewClient("")
	client.BaseURL = mockserver.URL
	res, err := client.RetrieveFileContent(nil, "abcdefg")
	Expect(t, err).ToBe(nil)
	Expect(t, res).TypeOf("*http.bodyEOFSignal")
	res.Close()
}

func TestClient_RetrieveFile(t *testing.T) {
	client := NewClient("")
	client.BaseURL = mockserver.URL
	res, err := client.RetrieveFile(nil, "abcdefg")
	Expect(t, err).ToBe(nil)
	Expect(t, res).TypeOf("openaigo.FileRetrieveResponse")
}

func TestClient_DeleteFile(t *testing.T) {
	client := NewClient("")
	client.BaseURL = mockserver.URL
	res, err := client.DeleteFile(nil, "abcdefg")
	Expect(t, err).ToBe(nil)
	Expect(t, res).TypeOf("openaigo.FileDeleteResponse")
}
