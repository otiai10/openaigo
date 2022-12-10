package openaigo

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type ErrorResponseBody struct {
	Error ErrorEntry `json:"error"`
}

type ErrorEntry struct {
	Message string      `json:"message"`
	Type    string      `json:"type"`
	Param   interface{} `json:"param"` // FIXME: typing
	Code    interface{} `json:"code"`  // FIXME: typing
}

func (err ErrorEntry) Error() string {
	return fmt.Sprintf("%v: %v, %v (%v)", err.Type, err.Message, err.Param, err.Code)
}

func (client *Client) apiError(res *http.Response) error {
	errbody := ErrorResponseBody{}
	if err := json.NewDecoder(res.Body).Decode(&errbody); err != nil {
		return fmt.Errorf("failed to decode error body: %v", err)
	}
	return errbody.Error
}
