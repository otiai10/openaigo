package openaigo

import (
	"testing"

	. "github.com/otiai10/mint"
)

func TestClient_CreateFineTune(t *testing.T) {
	client := NewClient("")
	client.BaseURL = mockserver.URL
	res, err := client.CreateFineTune(nil, FineTuneCreateRequestBody{
		TrainingFile: "file-XGinujblHPwGLSztz8cPS8XY",
	})
	Expect(t, err).ToBe(nil)
	Expect(t, res).TypeOf("openaigo.FineTuneCreateResponse")
}

func TestClient_ListFineTunes(t *testing.T) {
	client := NewClient("")
	client.BaseURL = mockserver.URL
	res, err := client.ListFineTunes(nil)
	Expect(t, err).ToBe(nil)
	Expect(t, res).TypeOf("openaigo.FineTuneListResponse")
}

func TestClient_RetrieveFineTune(t *testing.T) {
	client := NewClient("")
	client.BaseURL = mockserver.URL
	res, err := client.RetrieveFineTune(nil, "abcdefghi")
	Expect(t, err).ToBe(nil)
	Expect(t, res).TypeOf("openaigo.FineTuneRetrieveResponse")
}

func TestClient_CancelFineTune(t *testing.T) {
	client := NewClient("")
	client.BaseURL = mockserver.URL
	res, err := client.CancelFineTune(nil, "abcdefghi")
	Expect(t, err).ToBe(nil)
	Expect(t, res).TypeOf("openaigo.FineTuneCancelResponse")
}

func TestClient_ListFineTuneEvents(t *testing.T) {
	client := NewClient("")
	client.BaseURL = mockserver.URL
	res, err := client.ListFineTuneEvents(nil, "abcdefghi")
	Expect(t, err).ToBe(nil)
	Expect(t, res).TypeOf("openaigo.FineTuneListEventsResponse")
}

func TestClient_DeleteFineTune(t *testing.T) {
	client := NewClient("")
	client.BaseURL = mockserver.URL
	res, err := client.CancelFineTune(nil, "abcdefghi")
	Expect(t, err).ToBe(nil)
	Expect(t, res).TypeOf("openaigo.FineTuneCancelResponse")
}
