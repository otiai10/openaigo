package openaigo

import (
	"context"
	"testing"

	. "github.com/otiai10/mint"
)

func TestClient_CreateFineTuning(t *testing.T) {
	client := NewClient("")
	client.BaseURL = mockserver.URL
	res, err := client.CreateFineTuning(context.TODO(), FineTuningCreateRequestBody{
		TrainingFile: "file-XGinujblHPwGLSztz8cPS8XY",
	})
	Expect(t, err).ToBe(nil)
	Expect(t, res).TypeOf("openaigo.FineTuningJob")
}

func TestClient_RetrieveFineTuning(t *testing.T) {
	client := NewClient("")
	client.BaseURL = mockserver.URL
	res, err := client.RetrieveFineTuning(context.TODO(), "abcdefghi")
	Expect(t, err).ToBe(nil)
	Expect(t, res).TypeOf("openaigo.FineTuningJob")
}

func TestClient_CancelFineTuning(t *testing.T) {
	client := NewClient("")
	client.BaseURL = mockserver.URL
	res, err := client.CancelFineTuning(context.TODO(), "abcdefghi")
	Expect(t, err).ToBe(nil)
	Expect(t, res).TypeOf("openaigo.FineTuningJob")
}

func TestClient_ListFineTuningEvents(t *testing.T) {
	client := NewClient("")
	client.BaseURL = mockserver.URL
	res, err := client.ListFineTuningEvents(context.TODO(), "abcdefghi")
	Expect(t, err).ToBe(nil)
	Expect(t, res).TypeOf("openaigo.FineTuningListEventsResponse")
}
