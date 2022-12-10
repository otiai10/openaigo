package openaigo

type FineTuneCreateRequestBody struct {
	TrainingFile                 string    `json:"training_file"`
	ValidationFile               string    `json:"validation_file,omitempty"`
	Model                        string    `json:"model,omitempty"`
	NEpochs                      int       `json:"n_epochs,omitempty"`
	BatchSize                    int       `json:"batch_size,omitempty"`
	LearningRateMultiplier       float32   `json:"learning_rate_multiplier,omitempty"`
	PromptLossWeight             float32   `json:"prompt_loss_weight,omitempty"`
	ComputeClassificationMetrics bool      `json:"compute_classification_metrics,omitempty"`
	ClassificationNClasses       int       `json:"classification_n_classes,omitempty"`
	ClassificationPositiveClass  string    `json:"classification_positive_class,omitempty"`
	ClassificationBetas          []float32 `json:"classification_betas,omitempty"`
	Suffix                       string    `json:"suffix,omitempty"`
}

type FineTuneData struct {
	ID              string          `json:"id"`
	Object          string          `json:"object"`
	Model           string          `json:"model"`
	CreatedAt       int64           `json:"created_at"`
	Events          []FineTuneEvent `json:"events"`
	FineTunedModel  interface{}     `json:"fine_tuned_model"` // TODO: typing
	Hyperparams     Hyperparams     `json:"hyperparams"`
	OrganizationID  string          `json:"organization_id"`
	ResultFiles     []FileData      `json:"result_files"`
	Status          string          `json:"status"`
	ValidationFiles []FileData      `json:"validation_files"`
	TrainingFiles   []FileData      `json:"training_files"`
	UpdatedAt       int64           `json:"updated_at"`
}

type FineTuneCreateResponse struct {
	Events       []FineTuneEvent `json:"events"`
	FineTuneData `json:",inline"`
}

type FineTuneEvent struct {
	Object    string `json:"object"`
	CreatedAt int64  `json:"created_at"`
	Level     string `json:"level"`
	Message   string `json:"message"`
}

type Hyperparams struct {
	BatchSize              int     `json:"batch_size"`
	LearningRateMultiplier float32 `json:"learning_rate_multiplier"`
	NEpochs                int     `json:"n_epochs"`
	PromptLossWeight       float32 `json:"prompt_loss_weight"`
}

type FineTuneListResponse struct {
	Object string         `json:"object"`
	Data   []FineTuneData `json:"data"`
}

type FineTuneRetrieveResponse struct {
	Events       []FineTuneEvent `json:"events"`
	FineTuneData `json:",inline"`
}

type FineTuneCancelResponse struct {
	Events       []FineTuneEvent `json:"events"`
	FineTuneData `json:",inline"`
}

type FineTuneListEventsResponse struct {
	Object string          `json:"object"`
	Data   []FineTuneEvent `json:"data"`
}

type FineTuneDeleteModelResponse struct {
	ID      string `json:"string"`
	Object  string `json:"object"`
	Deleted bool   `json:"deleted"`
}
