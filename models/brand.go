package models

type Brand struct {
	Data       BrandData `json:"data,omitempty"`
	Page       int
	TotalPages int
}

type Brands struct {
	Data []BrandData `json:"data,omitempty"`
}

type BrandData struct {
	ID         int       `json:"id,omitempty"`
	Attributes BrandAttr `json:"attributes,omitempty"`
}

type BrandAttr struct {
	Name         string `json:"name,omitempty"`
	Description  string `json:"description,omitempty"`
	Slug         string `json:"slug,omitempty"`
	ExternalLink string `json:"externalLink,omitempty"`

	Logo     Img           `json:"logo,omitempty"`
	Products ShortProducts `json:"products,omitempty"`
}
