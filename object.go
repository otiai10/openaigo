package openaigo

type ObjectType string

const (
	OTModel           ObjectType = "model"
	OTModelPermission ObjectType = "model_permission"
	OTList            ObjectType = "list"
	OTEdit            ObjectType = "edit"
	OTTextCompletion  ObjectType = "text_completion"
	OTEEmbedding      ObjectType = "embedding"
	OTFile            ObjectType = "file"
	OTFineTune        ObjectType = "fine-tune"
	OTFineTuneEvent   ObjectType = "fine-tune-event"
)
