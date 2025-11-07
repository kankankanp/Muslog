package request

type UpdateUserSettingRequest struct {
	EditorType string `json:"editorType" validate:"required,oneof=markdown wysiwyg"`
}
