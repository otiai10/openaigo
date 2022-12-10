package openaigo

import (
	"context"
	"fmt"
	"io"
	"net/http"
)

// ListModels: GET /models
// Lists the currently available models, and provides basic information about each one such as the owner and availability.
// See https://beta.openai.com/docs/api-reference/models/list
func (client *Client) ListModels(ctx context.Context) (resp ModelsListResponse, err error) {
	p := "/models"
	return call(ctx, client, http.MethodGet, p, nil, resp)
}

// RetrieveModel: GET /models/{model}
// Retrieves a model instance, providing basic information about the model such as the owner and permissioning.
// See https://beta.openai.com/docs/api-reference/models/retrieve
func (client *Client) RetrieveModel(ctx context.Context, model string) (resp ModelRetrieveResponse, err error) {
	p := fmt.Sprintf("/models/%s", model)
	return call(ctx, client, http.MethodGet, p, nil, resp)
}

// Completion: POST https://api.openai.com/v1/completions
// Creates a completion for the provided prompt and parameters
// See https://beta.openai.com/docs/api-reference/completions/create
func (client *Client) Completion(ctx context.Context, body CompletionRequestBody) (resp CompletionResponse, err error) {
	p := "/completions"
	return call(ctx, client, http.MethodPost, p, body, resp)
}

// Edit: POST https://api.openai.com/v1/edits
// Creates a new edit for the provided input, instruction, and parameters.
// See https://beta.openai.com/docs/api-reference/edits/create
func (client *Client) Edit(ctx context.Context, body EditRequestBody) (resp EditResponse, err error) {
	p := "/edits"
	return call(ctx, client, http.MethodPost, p, body, resp)
}

// CreateImage: POST https://api.openai.com/v1/images/generations
// Creates an image given a prompt.
// See https://beta.openai.com/docs/api-reference/images/create
func (client *Client) CreateImage(ctx context.Context, body ImageGenerationRequestBody) (resp ImageGenerationResponse, err error) {
	p := "/images/generations"
	return call(ctx, client, http.MethodPost, p, body, resp)
}

func (client *Client) EditImage(ctx context.Context, body ImageEditRequestBody) (resp ImageEditResponse, err error) {
	p := "/images/edits"
	return call(ctx, client, http.MethodPost, p, body, resp)
}

// CreateImageVariation: POST https://api.openai.com/v1/images/variations
// Creates a variation of a given image.
// See https://beta.openai.com/docs/api-reference/images/create-variation
func (client *Client) CreateImageVariation(ctx context.Context, body ImageVariationRequestBody) (resp ImageVariationResponse, err error) {
	p := "/images/variations"
	return call(ctx, client, http.MethodPost, p, body, resp)
}

// CreateEmbedding: POST https://api.openai.com/v1/embeddings
// Creates an embedding vector representing the input text.
// See https://beta.openai.com/docs/api-reference/embeddings/create
func (client *Client) CreateEmbedding(ctx context.Context, body EmbeddingCreateRequestBody) (resp EmbeddingCreateResponse, err error) {
	p := "/embeddings"
	return call(ctx, client, http.MethodPost, p, body, resp)
}

// ListFiles: GET https://api.openai.com/v1/files
// Returns a list of files that belong to the user's organization.
// See https://beta.openai.com/docs/api-reference/files/list
func (client *Client) ListFiles(ctx context.Context) (resp ListFilesResponse, err error) {
	p := "/files"
	return call(ctx, client, http.MethodGet, p, nil, resp)
}

// UploadFile: POST https://api.openai.com/v1/files
// Upload a file that contains document(s) to be used across various endpoints/features.
// Currently, the size of all the files uploaded by one organization can be up to 1 GB.
// Please contact us if you need to increase the storage limit.
// See https://beta.openai.com/docs/api-reference/files/upload
func (client *Client) UploadFile(ctx context.Context, body FileUploadRequestBody) (resp FileUploadResponse, err error) {
	p := "/files"
	return call(ctx, client, http.MethodPost, p, body, resp)
}

// DeleteFile: DELETE https://api.openai.com/v1/files/{file_id}
// Delete a file.
// See https://beta.openai.com/docs/api-reference/files/delete
func (client *Client) DeleteFile(ctx context.Context, id string) (resp FileDeleteResponse, err error) {
	p := fmt.Sprintf("/files/%s", id)
	return call(ctx, client, http.MethodDelete, p, nil, resp)
}

// RetrieveFile: GET https://api.openai.com/v1/files/{file_id}
// Returns information about a specific file.
// See https://beta.openai.com/docs/api-reference/files/retrieve
func (client *Client) RetrieveFile(ctx context.Context, id string) (resp FileRetrieveResponse, err error) {
	p := fmt.Sprintf("/files/%s", id)
	return call(ctx, client, http.MethodGet, p, nil, resp)
}

// RetrieveFileContent: GET https://api.openai.com/v1/files/{file_id}/content
// Returns the contents of the specified file.
// User must Close response after used.
// See https://beta.openai.com/docs/api-reference/files/retrieve-content
func (client *Client) RetrieveFileContent(ctx context.Context, id string) (res io.ReadCloser, err error) {
	endpoint, err := client.endpoint(fmt.Sprintf("/files/%s", id))
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest(http.MethodGet, endpoint, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", client.APIKey))
	if ctx != nil {
		req = req.WithContext(ctx)
	}
	if client.HTTPClient == nil {
		client.HTTPClient = http.DefaultClient
	}
	response, err := client.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	if response.StatusCode >= 400 {
		return nil, client.apiError(response)
	}
	return response.Body, nil
}

// CreateModeration: POST https://api.openai.com/v1/moderations
// Classifies if text violates OpenAI's Content Policy.
// See https://beta.openai.com/docs/api-reference/moderations/create
func (client *Client) CreateModeration(ctx context.Context, body ModerationCreateRequestBody) (resp ModerationCreateResponse, err error) {
	p := "/moderations"
	return call(ctx, client, http.MethodPost, p, body, resp)
}
