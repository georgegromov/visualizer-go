package dto

type TemplateCreateDto struct {
	Name        string  `json:"name" db:"name"`
	Description *string `json:"description" db:"description"`
}

type TemplateUpdateDto struct {
	Name        *string      `json:"name" db:"name"`
	Description *string      `json:"description" db:"description"`
	Canvases    *interface{} `json:"canvases" db:"canvases"`
	IsDeleted   *bool        `json:"is_deleted" db:"is_deleted"`
}
