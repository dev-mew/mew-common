package models

type Doc struct {
	Data DocData `json:"data,omitempty"`
}

type Docs struct {
	Data []DocData `json:"data,omitempty"`
}

type DocData struct {
	ID         int     `json:"id,omitempty"`
	Attributes DocAttr `json:"attributes,omitempty"`
}

type DocAttr struct {
	Name            string `json:"name,omitempty"`
	AlternativeText string `json:"alternativeText,omitempty"`
	URL             string `json:"url,omitempty"`
}
