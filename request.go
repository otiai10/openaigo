package openaigo

import "bytes"

type MultipartFormDataRequestBody interface {
	ToMultipartFormData() (*bytes.Buffer, string, error)
}
