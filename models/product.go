package models

import (
	"strings"
)

type ShortProducts struct {
	Data []ShortProductData `json:"data,omitempty"`
}

type ShortProduct struct {
	Data ShortProductData `json:"data,omitempty"`
}

type ShortProductData struct {
	ID         *int             `json:"id,omitempty"`
	Attributes ShortProductAttr `json:"attributes,omitempty"`
}

func (sp ShortProductData) Filter(afs []string) bool {
	mathces := 0
	isCompleteMatch := true

	for _, af := range afs {
		if strings.Contains(strings.ToLower(sp.Attributes.Title), strings.ToLower(af)) {
			mathces += 1
		} else {
			isCompleteMatch = false
		}
	}

	return isCompleteMatch
}

type ShortProductAttr struct {
	Sku            string   `json:"sku,omitempty"`
	IIQID          string   `json:"iiqID,omitempty"`
	AllowedForSale bool     `json:"allowedForSale,omitempty"`
	Stock          int32    `json:"stock,omitempty"`
	Price          float64  `json:"price,omitempty"`
	PriceInCents   int32    `json:"priceInCents,omitempty"`
	Discount       int32    `json:"discount,omitempty"`
	Shipping       string   `json:"shipping,omitempty"`
	Shop           []string `json:"shops,omitempty"`
	IsNew          bool     `json:"isNew,omitempty"`
	Title          string   `json:"title,omitempty"`
	Slug           string   `json:"slug,omitempty"`
	EnergyClass    string   `json:"energyClass,omitempty"`
	EprelLabel     Img      `json:"eprelLabel,omitempty"`
	Images         Imgs     `json:"images,omitempty"`
}

type AppProds struct {
	Data AppProdsData `json:"data,omitempty"`
}

type AppProdsData struct {
	ID         int32              `json:"id,omitempty"`
	Attributes AppProdsAttributes `json:"attributes,omitempty"`
}

type AppProdsAttributes struct {
	Products ProdsData `json:"products"`
}

type ProdsData struct {
	Data []ProdAttr `json:"data"`
}

type ProdAttr struct {
	ID         int32      `json:"id"`
	Attributes ProdFields `json:"attributes"`
}

type ProdFields struct {
	Sku  string `json:"sku"`
	Slug string `json:"slug"`
}
