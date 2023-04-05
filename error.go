package openaigo

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type APIErrorType string

const (
	ErrorInsufficientQuota APIErrorType = "insufficient_quota"
	ErrorInvalidRequest    APIErrorType = "invalid_request_error"
)

type APIError struct {
	Message string       `json:"message"`
	Type    APIErrorType `json:"type"`
	Param   interface{}  `json:"param"` // TODO: typing
	Code    interface{}  `json:"code"`  // TODO: typing

	Status     string
	StatusCode int
}

func (err APIError) Error() string {
	return fmt.Sprintf("openai API error: %v: %v (param: %v, code: %v)", err.Type, err.Message, err.Param, err.Code)
}

func (client *Client) apiError(res *http.Response) error {
	errbody := struct {
		Error APIError `json:"error"`
	}{APIError{Status: res.Status, StatusCode: res.StatusCode}}
	if err := json.NewDecoder(res.Body).Decode(&errbody); err != nil {
		return fmt.Errorf("failed to decode error body: %v", err)
	}
	return errbody.Error
}
