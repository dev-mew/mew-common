package models

type Categories struct {
	Data []CategoryData `json:"data,omitempty"`
}

type Category struct {
	Data       CategoryData `json:"data,omitempty"`
	Page       int
	TotalPages int
}

type CategoryData struct {
	ID         *int         `json:"id,omitempty"`
	Attributes CategoryAttr `json:"attributes,omitempty"`
}

type CategoryAttr struct {
	Name        string `json:"name,omitempty"`
	Slug        string `json:"slug,omitempty"`
	Description string `json:"description,omitempty"`

	Filters   map[string][]string `json:"filters,omitempty"`
	Image     Img                 `json:"image,omitempty"`
	ImageWide Img                 `json:"imageWide,omitempty"`
	Products  ShortProducts       `json:"products,omitempty"`
	Blocks    []Block             `json:"blocks,omitempty"`
	Metadata  MetaInfo            `json:"metadata"`
}

type ReqFilterCategoryProducts struct {
	Page          int
	ActiveFilters []string
}
