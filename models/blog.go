package models

type Blogs struct {
	Data []BlogData `json:"data,omitempty"`
}

type Blog struct {
	Data BlogData `json:"data,omitempty"`
}

type BlogData struct {
	ID         int      `json:"id,omitempty"`
	Attributes BlogAttr `json:"attributes,omitempty"`
}

type BlogAttr struct {
	Title     string `json:"title,omitempty"`
	Slug      string `json:"slug,omitempty"`
	Intro     string `json:"intro,omitempty"`
	Content   string `json:"content,omitempty"`
	UpdatedAt string `json:"updatedAt,omitempty"`

	Image  Img     `json:"image,omitempty"`
	Blocks []Block `json:"blocks,omitempty"`
}
