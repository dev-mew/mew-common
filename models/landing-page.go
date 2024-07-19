package models

type LandingPage struct {
	Data LandingPageData `json:"data,omitempty"`
}

type LandingPageData struct {
	ID         int             `json:"id,omitempty"`
	Attributes LandingPageAttr `json:"attributes,omitempty"`
}

type LandingPageAttr struct {
	Title       string   `json:"title,omitempty"`
	Description string   `json:"description,omitempty"`
	Locale      string   `json:"locale,omitempty"`
	Metadata    MetaInfo `json:"metadata,omitempty"`
	Blocks      []Block  `json:"blocks,omitempty"`
}
