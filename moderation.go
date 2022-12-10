package openaigo

type ModerationCreateRequestBody struct {
	Input string `json:"input"`
	Model string `json:"model,omitempty"`
}

type ModerationCreateResponse struct {
	ID      string           `json:"id"`
	Model   string           `json:"model"`
	Results []ModerationData `json:"results"`
}

type ModerationData struct {
	Categories struct {
		Hate            bool `json:"hate"`
		HateThreatening bool `json:"hate/threatening"`
		SelfHarm        bool `json:"self-harm"`
		Sexual          bool `json:"sexual"`
		SexualMinors    bool `json:"sexual/minors"`
		Violence        bool `json:"violence"`
		ViolenceGraphic bool `json:"violence/graphic"`
	} `json:"categories"`
	CategoryScores struct {
		Hate            float32 `json:"hate"`
		HateThreatening float32 `json:"hate/threatening"`
		SelfHarm        float32 `json:"self-harm"`
		Sexual          float32 `json:"sexual"`
		SexualMinors    float32 `json:"sexual/minors"`
		Violence        float32 `json:"violence"`
		ViolenceGraphic float32 `json:"violence/graphic"`
	} `json:"category_scores"`
}
