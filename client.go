package openaigo

import (
	"bufio"
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
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

func (client *Client) build(ctx context.Context, method, p string, body interface{}, stream bool) (req *http.Request, err error) {
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
	
	if stream { // support SSE  https://developer.mozilla.org/en-US/docs/Web/API/Server-sent_events/Using_server-sent_events#event_stream_format
		req.Header.Set("Accept", "text/event-stream")
		req.Header.Set("Cache-Control", "no-cache")
		req.Header.Set("Connection", "keep-alive")
	}
	
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
	defer httpres.Body.Close()
	if httpres.StatusCode >= 400 {
		return client.apiError(httpres)
	}
	if err := json.NewDecoder(httpres.Body).Decode(response); err != nil {
		return fmt.Errorf("failed to decode response to %T: %v", response, err)
	}
	return nil
}

func call[T any](ctx context.Context, client *Client, method string, p string, body interface{}, resp T) (T, error) {
	req, err := client.build(ctx, method, p, body, false)
	if err != nil {
		return resp, err
	}
	err = client.execute(req, &resp)
	return resp, err
}

func callForStream(ctx context.Context, client *Client, method string, p string, body interface{}, callBack func(response ChatCompletionStreamResponse, err error)){
	req, err := client.build(ctx, method, p, body, true)
	if err != nil {
		callBack(ChatCompletionStreamResponse{}, err)
		return
	}
	
	if client.HTTPClient == nil {
		client.HTTPClient = http.DefaultClient
	}

	resp, err := client.HTTPClient.Do(req)
	reader := bufio.NewReader(resp.Body)

	for {

		s, err := reader.ReadString('\n')
		if err == io.EOF {
			callBack(ChatCompletionStreamResponse{}, err)
			break
		}

		if strings.EqualFold(s, "data: [DONE]") {
			callBack(ChatCompletionStreamResponse{}, io.EOF)
			break
		}

		if strings.HasPrefix(s, "error: ") {
			callBack(ChatCompletionStreamResponse{}, errors.New(s))
			break
		}

		s = strings.TrimSpace(s)
		s = strings.TrimPrefix(s, "data: ")
		
		v := ChatCompletionStreamResponse{}
		err = json.Unmarshal([]byte(s), &v)
		callBack(v, err)
	}
}
