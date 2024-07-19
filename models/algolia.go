package models

type AlgoliaResults struct {
	Hits []AlgoliaProduct `json:"hits"`
}

type AlgoliaProduct struct {
	Sku    string                `json:"sku"`
	Title  string                `json:"title"`
	Slug   string                `json:"slug"`
	Images []AlgoliaProductImage `json:"images,omitempty"`
}

type AlgoliaProductImage struct {
	AlternativeText string `json:"alternativeText,omitempty"`
	URL             string `json:"url,omitempty"`
}
