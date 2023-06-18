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
	return call(ctx, client, http.MethodGet, p, nil, resp, nil)
}

// RetrieveModel: GET /models/{model}
// Retrieves a model instance, providing basic information about the model such as the owner and permissioning.
// See https://beta.openai.com/docs/api-reference/models/retrieve
func (client *Client) RetrieveModel(ctx context.Context, model string) (resp ModelRetrieveResponse, err error) {
	p := fmt.Sprintf("/models/%s", model)
	return call(ctx, client, http.MethodGet, p, nil, resp, nil)
}

// Completion: POST https://api.openai.com/v1/completions
// Creates a completion for the provided prompt and parameters
// See https://beta.openai.com/docs/api-reference/completions/create
func (client *Client) Completion(ctx context.Context, body CompletionRequestBody) (resp CompletionResponse, err error) {
	p := "/completions"
	return call(ctx, client, http.MethodPost, p, body, resp, nil)
}

// Edit: POST https://api.openai.com/v1/edits
// Creates a new edit for the provided input, instruction, and parameters.
// See https://beta.openai.com/docs/api-reference/edits/create
func (client *Client) CreateEdit(ctx context.Context, body EditCreateRequestBody) (resp EditCreateResponse, err error) {
	p := "/edits"
	return call(ctx, client, http.MethodPost, p, body, resp, nil)
}

// CreateImage: POST https://api.openai.com/v1/images/generations
// Creates an image given a prompt.
// See https://beta.openai.com/docs/api-reference/images/create
func (client *Client) CreateImage(ctx context.Context, body ImageGenerationRequestBody) (resp ImageGenerationResponse, err error) {
	p := "/images/generations"
	return call(ctx, client, http.MethodPost, p, body, resp, nil)
}

func (client *Client) EditImage(ctx context.Context, body ImageEditRequestBody) (resp ImageEditResponse, err error) {
	p := "/images/edits"
	return call(ctx, client, http.MethodPost, p, body, resp, nil)
}

// CreateImageVariation: POST https://api.openai.com/v1/images/variations
// Creates a variation of a given image.
// See https://beta.openai.com/docs/api-reference/images/create-variation
func (client *Client) CreateImageVariation(ctx context.Context, body ImageVariationRequestBody) (resp ImageVariationResponse, err error) {
	p := "/images/variations"
	return call(ctx, client, http.MethodPost, p, body, resp, nil)
}

// CreateEmbedding: POST https://api.openai.com/v1/embeddings
// Creates an embedding vector representing the input text.
// See https://beta.openai.com/docs/api-reference/embeddings/create
func (client *Client) CreateEmbedding(ctx context.Context, body EmbeddingCreateRequestBody) (resp EmbeddingCreateResponse, err error) {
	p := "/embeddings"
	return call(ctx, client, http.MethodPost, p, body, resp, nil)
}

// ListFiles: GET https://api.openai.com/v1/files
// Returns a list of files that belong to the user's organization.
// See https://beta.openai.com/docs/api-reference/files/list
func (client *Client) ListFiles(ctx context.Context) (resp FileListResponse, err error) {
	p := "/files"
	return call(ctx, client, http.MethodGet, p, nil, resp, nil)
}

// UploadFile: POST https://api.openai.com/v1/files
// Upload a file that contains document(s) to be used across various endpoints/features.
// Currently, the size of all the files uploaded by one organization can be up to 1 GB.
// Please contact us if you need to increase the storage limit.
// See https://beta.openai.com/docs/api-reference/files/upload
func (client *Client) UploadFile(ctx context.Context, body FileUploadRequestBody) (resp FileUploadResponse, err error) {
	p := "/files"
	return call(ctx, client, http.MethodPost, p, body, resp, nil)
}

// DeleteFile: DELETE https://api.openai.com/v1/files/{file_id}
// Delete a file.
// See https://beta.openai.com/docs/api-reference/files/delete
func (client *Client) DeleteFile(ctx context.Context, id string) (resp FileDeleteResponse, err error) {
	p := fmt.Sprintf("/files/%s", id)
	return call(ctx, client, http.MethodDelete, p, nil, resp, nil)
}

// RetrieveFile: GET https://api.openai.com/v1/files/{file_id}
// Returns information about a specific file.
// See https://beta.openai.com/docs/api-reference/files/retrieve
func (client *Client) RetrieveFile(ctx context.Context, id string) (resp FileRetrieveResponse, err error) {
	p := fmt.Sprintf("/files/%s", id)
	return call(ctx, client, http.MethodGet, p, nil, resp, nil)
}

// RetrieveFileContent: GET https://api.openai.com/v1/files/{file_id}/content
// Returns the contents of the specified file.
// User must Close response after used.
// See https://beta.openai.com/docs/api-reference/files/retrieve-content
func (client *Client) RetrieveFileContent(ctx context.Context, id string) (res io.ReadCloser, err error) {
	endpoint, err := client.endpoint(fmt.Sprintf("/files/%s/content", id))
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
	return call(ctx, client, http.MethodPost, p, body, resp, nil)
}

// CreateFineTune: POST https://api.openai.com/v1/fine-tunes
// Creates a job that fine-tunes a specified model from a given dataset.
// Response includes details of the enqueued job including job status and the name of the fine-tuned models once complete.
// Learn more about Fine-tuning: https://beta.openai.com/docs/guides/fine-tuning
// See https://beta.openai.com/docs/api-reference/fine-tunes/create
func (client *Client) CreateFineTune(ctx context.Context, body FineTuneCreateRequestBody) (resp FineTuneCreateResponse, err error) {
	p := "/fine-tunes"
	return call(ctx, client, http.MethodPost, p, body, resp, nil)
}

// ListFineTunes: GET https://api.openai.com/v1/fine-tunes
// List your organization's fine-tuning jobs.
// See https://beta.openai.com/docs/api-reference/fine-tunes/list
func (client *Client) ListFineTunes(ctx context.Context) (resp FineTuneListResponse, err error) {
	p := "/fine-tunes"
	return call(ctx, client, http.MethodGet, p, nil, resp, nil)
}

// RetrieveFineTune: GET https://api.openai.com/v1/fine-tunes/{fine_tune_id}
// Gets info about the fine-tune job.
// Learn more about Fine-tuning https://beta.openai.com/docs/guides/fine-tuning
// See https://beta.openai.com/docs/api-reference/fine-tunes/retrieve
func (client *Client) RetrieveFineTune(ctx context.Context, id string) (resp FineTuneRetrieveResponse, err error) {
	p := fmt.Sprintf("/fine-tunes/%s", id)
	return call(ctx, client, http.MethodGet, p, nil, resp, nil)
}

// CancelFineTune: POST https://api.openai.com/v1/fine-tunes/{fine_tune_id}/cancel
// Immediately cancel a fine-tune job.
// See https://beta.openai.com/docs/api-reference/fine-tunes/cancel
func (client *Client) CancelFineTune(ctx context.Context, id string) (resp FineTuneCancelResponse, err error) {
	p := fmt.Sprintf("/fine-tunes/%s/cancel", id)
	return call(ctx, client, http.MethodPost, p, nil, resp, nil)
}

// ListFineTuneEvents: GET https://api.openai.com/v1/fine-tunes/{fine_tune_id}/events
// Get fine-grained status updates for a fine-tune job.
// See https://beta.openai.com/docs/api-reference/fine-tunes/events
func (client *Client) ListFineTuneEvents(ctx context.Context, id string) (resp FineTuneListEventsResponse, err error) {
	p := fmt.Sprintf("/fine-tunes/%s/events", id)
	return call(ctx, client, http.MethodGet, p, nil, resp, nil)
}

// DeleteFineTuneModel: DELETE https://api.openai.com/v1/models/{model}
// Delete a fine-tuned model. You must have the Owner role in your organization.
// See https://beta.openai.com/docs/api-reference/fine-tunes/delete-model
func (client *Client) DeleteFineTuneModel(ctx context.Context, id string) (resp FineTuneDeleteModelResponse, err error) {
	p := fmt.Sprintf("/models/%s", id)
	return call(ctx, client, http.MethodDelete, p, nil, resp, nil)
}

// Chat, short-hand of ChatCompletion.
// Creates a completion for the chat message.
func (client *Client) Chat(ctx context.Context, body ChatRequest) (resp ChatCompletionResponse, err error) {
	return client.ChatCompletion(ctx, ChatCompletionRequestBody(body))
}

// ChatCompletion: POST https://api.openai.com/v1/chat/completions
// Creates a completion for the chat message.
// See https://platform.openai.com/docs/api-reference/chat/create
func (client *Client) ChatCompletion(ctx context.Context, body ChatCompletionRequestBody) (resp ChatCompletionResponse, err error) {
	p := "/chat/completions"
	if body.StreamCallback != nil {
		body.Stream = true // Nosy ;)
		return call(ctx, client, http.MethodPost, p, body, resp, body.StreamCallback)
	}
	return call(ctx, client, http.MethodPost, p, body, resp, nil)
}
