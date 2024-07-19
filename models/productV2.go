package models

type FullProduct2 struct {
	Data FullProductData2 `json:"data"`
}

type FullProductData2 struct {
	Attributes FullProductAttr2 `json:"attributes,omitempty"`
}

type FullProductAttr2 struct {
	Product FullProductSubAttr    `json:"product,omitempty"`
	Related []ShortProductSubAttr `json:"related,omitempty"`
}

type FullProductSubAttr struct {
	ID               int32                  `json:"id"`
	Locale           string                 `json:"locale,omitempty"`
	AllowedForSale   bool                   `json:"allowedForSale,omitempty"`
	StockingStatus   string                 `json:"stockingStatus,omitempty"`
	Sku              string                 `json:"sku,omitempty"`
	IIQID            string                 `json:"iiqID,omitempty"`
	Name             string                 `json:"name,omitempty"`
	Stock            int32                  `json:"stock,omitempty"`
	Price            float64                `json:"price,omitempty"`
	PriceInCents     int32                  `json:"priceInCents,omitempty"`
	Discount         int32                  `json:"discount,omitempty"`
	Shipping         string                 `json:"shipping,omitempty"`
	PricedPer        string                 `json:"pricedPer,omitempty"`
	Mpn              string                 `json:"mpn,omitempty"`
	Gtin             string                 `json:"gtin,omitempty"`
	Weee             bool                   `json:"weee,omitempty"`
	Title            string                 `json:"title,omitempty"`
	Slug             string                 `json:"slug,omitempty"`
	Shop             []string               `json:"shop,omitempty"`
	IsNew            bool                   `json:"isNew,omitempty"`
	LongDescription  string                 `json:"longDescription"`
	ShortDescription string                 `json:"shortDescription"`
	Collection       string                 `json:"collection,omitempty"`
	Compatible       string                 `json:"compatible,omitempty"`
	BoxContent       []string               `json:"boxContent,omitempty"`
	EnergyClass      string                 `json:"energyClass,omitempty"`
	EprelLabel       ProductEprelLabel      `json:"eprelLabel,omitempty"`
	Features         map[string]interface{} `json:"features,omitempty"`
	FAQ              map[string]string      `json:"faq,omitempty"`
	Images           []Img2                 `json:"images,omitempty"`
	Documents        []ProductDocument      `json:"documents,omitempty"`
	Categories       []ProductCategory      `json:"categories,omitempty"`
	Brand            ProductBrand           `json:"brand,omitempty"`
	Blocks           []ProductBlock         `json:"blocks,omitempty"`
	Metadata         MetaInfo               `json:"metadata"`
}

type ProductBlock struct {
	ID        int                   `json:"id,omitempty"`
	Component string                `json:"__component,omitempty"`
	Title     string                `json:"title,omitempty"`
	Products  []ShortProductSubAttr `json:"products,omitempty"`
}

type ShortProductSubAttr struct {
	ID             int32             `json:"id,omitempty"`
	AllowedForSale bool              `json:"allowedForSale,omitempty"`
	Sku            string            `json:"sku,omitempty"`
	IIQID          string            `json:"iiqID,omitempty"`
	Stock          int32             `json:"stock,omitempty"`
	Price          float64           `json:"price,omitempty"`
	PriceInCents   int32             `json:"priceInCents,omitempty"`
	Discount       int32             `json:"discount,omitempty"`
	Shipping       string            `json:"shipping,omitempty"`
	Shop           []string          `json:"shops,omitempty"`
	IsNew          bool              `json:"isNew,omitempty"`
	Title          string            `json:"title,omitempty"`
	Slug           string            `json:"slug,omitempty"`
	EnergyClass    string            `json:"energyClass,omitempty"`
	EprelLabel     ProductEprelLabel `json:"eprelLabel,omitempty"`
	Images         []Img2            `json:"images,omitempty"`
	Categories     []ProductCategory `json:"categories,omitempty"`
}

type ProductEprelLabel struct {
	AlternativeText string `json:"alternativeText,omitempty"`
	URL             string `json:"url,omitempty"`
}

type ProductDocument struct {
	Name            string `json:"name,omitempty"`
	AlternativeText string `json:"alternativeText,omitempty"`
	URL             string `json:"url,omitempty"`
}

type ProductCategory struct {
	ID   *int32 `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
	Slug string `json:"slug,omitempty"`
}

type ProductBrand struct {
	ID   int              `json:"id,omitempty"`
	Name string           `json:"name,omitempty"`
	Slug string           `json:"slug,omitempty"`
	Logo ProductBrandLogo `json:"logo,omitempty"`
}

type ProductBrandLogo struct {
	AlternativeText string `json:"alternativeText,omitempty"`
	URL             string `json:"url,omitempty"`
}
