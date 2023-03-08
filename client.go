package openaigo

import (
	"bufio"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

var (
	StreamPrefixDATA  = []byte("data: ")
	StreamSignalDONE  = []byte("[DONE]")
	StreamPrefixERROR = []byte("error: ")
)

const DefaultOpenAIAPIURL = "https://api.openai.com/v1"

// Client for api.openai.com API endpoints.
type Client struct {

	// APIKey issued by OpenAI console.
	// See https://beta.openai.com/account/api-keys
	APIKey string

	// BaseURL of API including the version.
	// e.g., https://api.openai.com/v1
	BaseURL string

	// Organization
	Organization string

	// HTTPClient (optional) to proxy HTTP request.
	// If nil, *http.DefaultClient will be used.
	HTTPClient *http.Client
}

func NewClient(apikey string) *Client {
	return &Client{
		APIKey: apikey,
		// Organization: org-GXjGDRs5UuJ4CvQ2u9d5uy0k
		// BaseURL: DefaultOpenAIAPIURL,
		// HTTPClient: http.DefaultClient,
	}
}

func (client *Client) endpoint(p string) (string, error) {
	if client.BaseURL == "" {
		client.BaseURL = DefaultOpenAIAPIURL
	}
	u, err := url.Parse(client.BaseURL)
	if err != nil {
		return "", err
	}
	u.Path = strings.Join([]string{
		strings.TrimRight(u.Path, "/"),
		strings.TrimLeft(p, "/"),
	}, "/")
	return u.String(), nil
}

func (client *Client) build(ctx context.Context, method, p string, body interface{}) (req *http.Request, err error) {
	endpoint, err := client.endpoint(p)
	if err != nil {
		return nil, err
	}
	r, contenttype, err := client.bodyToReader(body)
	if err != nil {
		return nil, fmt.Errorf("failed to build request buf from given body: %v", err)
	}
	req, err = http.NewRequest(method, endpoint, r)
	if err != nil {
		return nil, fmt.Errorf("failed to init request: %v", err)
	}
	req.Header.Add("Content-Type", contenttype)
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", client.APIKey))
	if client.Organization != "" {
		req.Header.Add("OpenAI-Organization", client.Organization)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}
	return req, nil
}

func (client *Client) bodyToReader(body interface{}) (io.Reader, string, error) {
	var r io.Reader
	switch v := body.(type) {
	// case io.Reader:
	// 	r = v
	case nil:
		r = nil
	case MultipartFormDataRequestBody: // TODO: Refactor
		buf, ct, err := v.ToMultipartFormData()
		if err != nil {
			return nil, "", err
		}
		return buf, ct, nil
	default:
		b, err := json.Marshal(body)
		if err != nil {
			return nil, "", err
		}
		r = bytes.NewBuffer(b)
	}
	return r, "application/json", nil
}

func (client *Client) execute(req *http.Request, response interface{}) error {
	if client.HTTPClient == nil {
		client.HTTPClient = http.DefaultClient
	}
	httpres, err := client.HTTPClient.Do(req)
	if err != nil {
		return err
	}
	if httpres.StatusCode >= 400 {
		return client.apiError(httpres)
	}

	switch v := (response).(type) {
	case *ChatCompletionResponse:
		if v.stream != nil {
			go handle(v.stream, httpres.Body)
		}
	case *CompletionResponse:
		if v.stream != nil {
			go handle(v.stream, httpres.Body)
		}
	case *FineTuneListEventsResponse:
		if v.stream != nil {
			go handle(v.stream, httpres.Body)
		}
	default:
		return decode(response, httpres.Body)
	}
	return nil
}

func call[T any](ctx context.Context, client *Client, method string, p string, body interface{}, resp T) (T, error) {
	req, err := client.build(ctx, method, p, body)
	if err != nil {
		return resp, err
	}
	err = client.execute(req, &resp)
	return resp, err
}

func decode(response any, body io.ReadCloser) error {
	defer body.Close()
	if err := json.NewDecoder(body).Decode(response); err != nil {
		return fmt.Errorf("failed to decode response to %T: %v", response, err)
	}
	return nil
}

// handle handles data-only server-sent events from HTTP response body.
// This is used only for Create Completion, Create Chat Completion and List FineTune events.
func handle[T any](stream chan<- T, body io.ReadCloser) {
	defer body.Close()
	defer close(stream)
	s := bufio.NewScanner(body)
	for s.Scan() {
		b := s.Bytes()
		if len(b) == 0 {
			continue
		}
		if bytes.HasPrefix(b, StreamPrefixDATA) {
			if bytes.HasSuffix(b, StreamSignalDONE) {
				return
			}
			r := new(T)
			if err := json.Unmarshal(b[len(StreamPrefixDATA):], r); err != nil {
				// TODO: Error handling in stream mode
				// r.Error = err
			}
			stream <- *r
		} else if bytes.HasPrefix(b, StreamPrefixERROR) {
			return
		}
	}
	return
}
