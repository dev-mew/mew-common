package models

type Img struct {
	Data ImgData `json:"data,omitempty"`
}

type Imgs struct {
	Data []ImgData `json:"data,omitempty"`
}

type ImgData struct {
	ID         int     `json:"id,omitempty"`
	Attributes ImgAttr `json:"attributes,omitempty"`
}

type ImgAttr struct {
	AlternativeText string     `json:"alternativeText,omitempty"`
	URL             string     `json:"url,omitempty"`
	Formats         ImgFormats `json:"formats,omitempty"`
}

type Img2 struct {
	AlternativeText string     `json:"alternativeText,omitempty"`
	URL             string     `json:"url,omitempty"`
	Width           int        `json:"width,omitempty"`
	Height          int        `json:"height,omitempty"`
	Formats         ImgFormats `json:"formats,omitempty"`
}

type ImgFormats struct {
	Thumbnail ImgFormatData `json:"thumbnail,omitempty"`
	Small     ImgFormatData `json:"small,omitempty"`
}

type ImgFormatData struct {
	URL    string `json:"url,omitempty"`
	Width  int    `json:"width,omitempty"`
	Height int    `json:"height,omitempty"`
}
