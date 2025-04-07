package dto

type TemplateCreateDto struct {
	Name        string       `json:"name" db:"name"`
	Description *string      `json:"description" db:"description"`
	Canvases    *interface{} `json:"canvases" db:"canvases"`
}

type TemplateUpdateDto struct {
	Name        *string      `json:"name" db:"name"`
	Description *string      `json:"description" db:"description"`
	Canvases    *interface{} `json:"canvases" db:"canvases"`
	IsDeleted   *bool        `json:"isDeleted" db:"is_deleted"`
}
