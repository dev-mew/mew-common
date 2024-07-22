package models

type MetaInfo struct {
	ID              int32   `json:"id,omitempty"`
	MetaTitle       string  `json:"metaTitle,omitempty"`
	MetaDescription string  `json:"metaDescription,omitempty"`
	MetaImage       MetaImg `json:"metaImage,omitempty"`
}

type MetaImg struct {
	AlternativeText string `json:"alternativeText,omitempty"`
	URL             string `json:"url,omitempty"`
}

type MetaData struct {
	ID              int32  `json:"id,omitempty"`
	MetaTitle       string `json:"metaTitle,omitempty"`
	MetaDescription string `json:"metaDescription,omitempty"`
	MetaImage       Img    `json:"metaImage,omitempty"`
}
