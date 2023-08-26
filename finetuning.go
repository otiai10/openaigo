package openaigo

type FineTuningJob struct {
	ID              string          `json:"id"`
	Object          string          `json:"object"`
	CreatedAt       int64           `json:"created_at"`
	FinishedAt      int64           `json:"finished_at"`
	Model           string          `json:"model"`
	FineTunedModel  string          `json:"fine_tuned_model,omitempty"`
	OrganizationID  string          `json:"organization_id"`
	Status          string          `json:"status"`
	Hyperparameters Hyperparameters `json:"hyperparameters"`
	TrainingFile    string          `json:"training_file"`
	ValidationFile  string          `json:"validation_file,omitempty"`
	ResultFiles     []string        `json:"result_files"`
	TrainedTokens   int             `json:"trained_tokens"`
}

type Hyperparameters struct {
	Epochs int `json:"n_epochs"`
}

type FineTuningCreateRequestBody struct {
	TrainingFile    string           `json:"training_file"`
	ValidationFile  string           `json:"validation_file,omitempty"`
	Model           string           `json:"model,omitempty"`
	Hyperparameters *Hyperparameters `json:"hyperparameters,omitempty"`
	Suffix          string           `json:"suffix,omitempty"`
}

type FineTuningListEventsResponse struct {
	Object  string          `json:"object"`
	Data    []FineTuneEvent `json:"data"`
	HasMore bool            `json:"has_more"`
}

type FineTuningEvent struct {
	Object    string `json:"object"`
	ID        string `json:"id"`
	CreatedAt int    `json:"created_at"`
	Level     string `json:"level"`
	Message   string `json:"message"`
	Data      any    `json:"data"`
	Type      string `json:"type"`
}
