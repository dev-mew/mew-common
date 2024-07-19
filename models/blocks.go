package models

type Block struct {
	ID          int    `json:"id,omitempty"`
	Component   string `json:"__component,omitempty"`
	Title       string `json:"title,omitempty"`
	Description string `json:"description,omitempty"`
	YoutubeID   string `json:"youtubeID,omitempty"`

	Slide      []BoxImageResponsive `json:"slide,omitempty"`
	Cta        CTA                  `json:"cta,omitempty"`
	Image      Img                  `json:"image,omitempty"`
	Categories Categories           `json:"categories,omitempty"`
	Products   ShortProducts        `json:"products,omitempty"`
}

type CTA struct {
	ID           int    `json:"id,omitempty"`
	Text         string `json:"text,omitempty"`
	Action       string `json:"action,omitempty"`
	CSS          string `json:"css,omitempty"`
	IconPosition string `json:"iconPosition,omitempty"`
	Icon         Img    `json:"icon,omitempty"`
}
