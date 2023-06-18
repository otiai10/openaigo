package openaigo

type CompletionRequestBody struct {

	// Model: ID of the model to use.
	// You can use the List models API to see all of your available models, or see our Model overview for descriptions of them.
	// See https://beta.openai.com/docs/api-reference/completions/create#completions/create-model
	Model string `json:"model"`

	// Prompt: The prompt(s) to generate completions for, encoded as a string, array of strings, array of tokens, or array of token arrays.
	// Note that <|endoftext|> is the document separator that the model sees during training, so if a prompt is not specified the model will generate as if from the beginning of a new document.
	// See https://beta.openai.com/docs/api-reference/completions/create#completions/create-prompt
	Prompt []string `json:"prompt"`

	// MaxTokens: The maximum number of tokens to generate in the completion.
	// The token count of your prompt plus max_tokens cannot exceed the model's context length. Most models have a context length of 2048 tokens (except for the newest models, which support 4096).
	// See https://beta.openai.com/docs/api-reference/completions/create#completions/create-max_tokens
	MaxTokens int `json:"max_tokens,omitempty"`

	// Temperature: What sampling temperature to use. Higher values means the model will take more risks. Try 0.9 for more creative applications, and 0 (argmax sampling) for ones with a well-defined answer.
	// We generally recommend altering this or top_p but not both.
	// See https://beta.openai.com/docs/api-reference/completions/create#completions/create-temperature
	Temperature float32 `json:"temperature,omitempty"`

	// Suffix: The suffix that comes after a completion of inserted text.
	// See https://beta.openai.com/docs/api-reference/completions/create#completions/create-suffix
	Suffix string `json:"suffix,omitempty"`

	// TopP: An alternative to sampling with temperature, called nucleus sampling,
	// where the model considers the results of the tokens with top_p probability mass.
	// So 0.1 means only the tokens comprising the top 10% probability mass are considered.
	// We generally recommend altering this or temperature but not both.
	// See https://beta.openai.com/docs/api-reference/completions/create#completions/create-top_p
	TopP float32 `json:"top_p,omitempty"`

	// N: How many completions to generate for each prompt.
	// Note: Because this parameter generates many completions, it can quickly consume your token quota.
	// Use carefully and ensure that you have reasonable settings for max_tokens and stop.
	// See https://beta.openai.com/docs/api-reference/completions/create#completions/create-n
	N int `json:"n,omitempty"`

	// Stream: Whether to stream back partial progress.
	// If set, tokens will be sent as data-only server-sent events as they become available,
	// with the stream terminated by a data: [DONE] message.
	// See https://beta.openai.com/docs/api-reference/completions/create#completions/create-stream
	Stream bool `json:"stream,omitempty"`

	// LogProbs: Include the log probabilities on the logprobs most likely tokens, as well the chosen tokens.
	// For example, if logprobs is 5, the API will return a list of the 5 most likely tokens. The API will always return the logprob of the sampled token, so there may be up to logprobs+1 elements in the response.
	// The maximum value for logprobs is 5. If you need more than this, please contact us through our Help center and describe your use case.
	// See https://beta.openai.com/docs/api-reference/completions/create#completions/create-logprobs
	LogProbs int `json:"logprobs,omitempty"`

	// Echo: Echo back the prompt in addition to the completion.
	// See https://beta.openai.com/docs/api-reference/completions/create#completions/create-echo
	Echo bool `json:"echo,omitempty"`

	// Stop: Up to 4 sequences where the API will stop generating further tokens. The returned text will not contain the stop sequence.
	// See https://beta.openai.com/docs/api-reference/completions/create#completions/create-stop
	Stop []string `json:"stop,omitempty"`

	// PresencePenalty: Number between -2.0 and 2.0.
	// Positive values penalize new tokens based on whether they appear in the text so far, increasing the model's likelihood to talk about new topics.
	// See more information about frequency and presence penalties.
	// See https://beta.openai.com/docs/api-reference/completions/create#completions/create-presence_penalty
	PresencePenalty float32 `json:"presence_penalty,omitempty"`

	// FrequencyPenalty: Number between -2.0 and 2.0.
	// Positive values penalize new tokens based on their existing frequency in the text so far, decreasing the model's likelihood to repeat the same line verbatim.
	// See more information about frequency and presence penalties.
	// See https://beta.openai.com/docs/api-reference/completions/create#completions/create-frequency_penalty
	FrequencyPenalty float32 `json:"frequency_penalty,omitempty"`

	// BestOf: Generates best_of completions server-side and returns the "best" (the one with the highest log probability per token). Results cannot be streamed.
	// When used with n, best_of controls the number of candidate completions and n specifies how many to return – best_of must be greater than n.
	// Note: Because this parameter generates many completions, it can quickly consume your token quota. Use carefully and ensure that you have reasonable settings for max_tokens and stop.
	// See https://beta.openai.com/docs/api-reference/completions/create#completions/create-best_of
	BestOf int `json:"best_of,omitempty"`

	// LogitBias: Modify the likelihood of specified tokens appearing in the completion.
	// Accepts a json object that maps tokens (specified by their token ID in the GPT tokenizer) to an associated bias value from -100 to 100. You can use this tokenizer tool (which works for both GPT-2 and GPT-3) to convert text to token IDs. Mathematically, the bias is added to the logits generated by the model prior to sampling. The exact effect will vary per model, but values between -1 and 1 should decrease or increase likelihood of selection; values like -100 or 100 should result in a ban or exclusive selection of the relevant token.
	// As an example, you can pass {"50256": -100} to prevent the <|endoftext|> token from being generated.
	// See https://beta.openai.com/docs/api-reference/completions/create#completions/create-logit_bias
	LogitBias map[string]int `json:"logit_bias,omitempty"`

	// User: A unique identifier representing your end-user, which can help OpenAI to monitor and detect abuse. Learn more.
	// See https://beta.openai.com/docs/api-reference/completions/create#completions/create-user
	User string `json:"user,omitempty"`
}

type CompletionResponse struct {
	ID      string             `json:"id"`
	Object  ObjectType         `json:"object"`
	Created int64              `json:"created"`
	Model   string             `json:"model"`
	Choices []CompletionChoice `json:"choices"`
	Usage   Usage
}
