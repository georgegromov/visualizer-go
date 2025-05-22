package templates

type TemplateCreateDto struct {
	Name        string  `json:"name" db:"name"`
	Description *string `json:"description" db:"description"`
}

type TemplateUpdateDto struct {
	Name        *string `json:"name" db:"name"`
	Description *string `json:"description" db:"description"`
	IsDeleted   *bool   `json:"isDeleted" db:"is_deleted"`
}

type TemplateSaveAsDTO struct {
	Canvases []AnalysisCanvas `json:"canvases" validate:"required"`
}
