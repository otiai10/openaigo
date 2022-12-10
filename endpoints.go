package openaigo

import (
	"context"
	"fmt"
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
