package models

type BoxImage struct {
	ID           int    `json:"id,omitempty"`
	InternalLink string `json:"internalLink,omitempty"`
	Name         string `json:"name,omitempty"`
	ButtonText   string `json:"buttonText,omitempty"`
	Campaign     string `json:"campaign,omitempty"`
	Image        Img    `json:"image,omitempty"`
}

type BoxImageResponsive struct {
	ID           int    `json:"id,omitempty"`
	InternalLink string `json:"internalLink,omitempty"`
	Name         string `json:"name,omitempty"`
	ButtonText   string `json:"buttonText,omitempty"`
	Campaign     string `json:"campaign,omitempty"`
	TextColor    string `json:"textColor,omitempty"`
	ImageDesktop Img    `json:"imageDesktop,omitempty"`
	ImageMobile  Img    `json:"imageMobile,omitempty"`
}
