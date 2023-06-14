package openaigo

type (
	ModelData struct {
		ID         string            `json:"id"`
		Object     ObjectType        `json:"object"`
		Created    int64             `json:"created"`
		OwnedBy    string            `json:"owned_by"`
		Permission []ModelPermission `json:"permission"`
		Root       string            `json:"root"`
		Parent     string            `json:"parent"`
	}
	ModelPermission struct {
		ID                 string     `json:"id"`
		Object             ObjectType `json:"object"`
		Created            int64      `json:"created"`
		AllowCreateEngine  bool       `json:"allow_create_engine"`
		AllowSampling      bool       `json:"allow_sampling"`
		AllowLogProbs      bool       `json:"allow_logprobs"`
		AllowSearchIndices bool       `json:"allow_search_indices"`
		AllowView          bool       `json:"allow_view"`
		AllowFineTuning    bool       `json:"allow_fine_tuning"`
		Organization       string     `json:"organization"`
		Group              string     `json:"group"`
		IsBlocking         bool       `json:"is_blocking"`
	}
)

type ModelsListResponse struct {
	Data   []ModelData `json:"data"`
	Object ObjectType
}

type ModelRetrieveResponse ModelData

// https://beta.openai.com/docs/models/overview
const (
	// {{{ https://beta.openai.com/docs/models/gpt-3
	TextDavinci003 = "text-davinci-003"
	TextCurie001   = "text-curie-001"
	TextBabbage001 = "text-babbage-001"
	TextAda001     = "text-ada-001"
	// }}}

	// {{{ https://platform.openai.com/docs/models/gpt-3-5
	GPT3_5Turbo          = "gpt-3.5-turbo"
	GPT3_5Turbo_0301     = "gpt-3.5-turbo-0301"
	GPT3_5Turbo_0613     = "gpt-3.5-turbo-0613"
	GPT3_5Turbo_16K      = "gpt-3.5-turbo-16k"
	GPT3_5Turbo_16K_0613 = "gpt-3.5-turbo-16k-0613"
	// }}}

	// {{{ https://platform.openai.com/docs/models/gpt-4
	GPT4          = "gpt-4"
	GPT4_0314     = "gpt-4-0314"
	GPT4_0613     = "gpt-4-0613"
	GPT4_32K      = "gpt-4-32k"
	GPT4_32K_0314 = "gpt-4-32k-0314"
	GPT4_32K_0613 = "gpt-4-32k-0613"
	// }}}
)
