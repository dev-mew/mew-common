package models

type Block struct {
	ID           int    `json:"id,omitempty"`
	Component    string `json:"__component,omitempty"`
	Title        string `json:"title,omitempty"`
	Description  string `json:"description,omitempty"`
	TextPosition string `json:"textPosition,omitempty"`
	TextColor    string `json:"textColor,omitempty"`
	YoutubeID    string `json:"youtubeID,omitempty"`

	Slide      []BoxImageResponsive `json:"slide,omitempty"`
	Cta        CTA                  `json:"cta,omitempty"`
	CtaBtn     CTABtn               `json:"ctaButton,omitempty"`
	Image      Img                  `json:"image,omitempty"`
	Categories Categories           `json:"categories,omitempty"`
	Products   ShortProducts        `json:"products,omitempty"`
	Product    ShortProduct         `json:"product,omitempty"`
	TwinCards  []FeatureCard        `json:"twinCards,omitempty"`
}

type CTA struct {
	ID           int    `json:"id,omitempty"`
	Text         string `json:"text,omitempty"`
	Action       string `json:"action,omitempty"`
	CSS          string `json:"css,omitempty"`
	IconPosition string `json:"iconPosition,omitempty"`
	Icon         Img    `json:"icon,omitempty"`
}

type CTABtn struct {
	Text  string `json:"text,omitempty"`
	Link  string `json:"link,omitempty"`
	Color string `json:"color,omitempty"`
}

type FeatureCard struct {
	Title        string `json:"title,omitempty"`
	Description  string `json:"description,omitempty"`
	Link         string `json:"link,omitempty"`
	TextPosition string `json:"textPosition,omitempty"`
	TextColor    string `json:"textColor,omitempty"`
	Image        Img    `json:"image,omitempty"`
}
