package models

type MiscPages struct {
	Data []MiscPageData `json:"data,omitempty"`
}

type MiscPage struct {
	Data MiscPageData `json:"data,omitempty"`
}

type MiscPageData struct {
	ID         int          `json:"id,omitempty"`
	Attributes MiscPageAttr `json:"attributes,omitempty"`
}

type MiscPageAttr struct {
	Shop        string `json:"shop,omitempty"`
	Slug        string `json:"slug,omitempty"`
	Title       string `json:"title,omitempty"`
	Description string `json:"description,omitempty"`
	Content     string `json:"content,omitempty"`
	UpdatedAt   string `json:"updatedAt,omitempty"`
}
