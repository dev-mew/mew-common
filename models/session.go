package models

type Session struct {
	ID       string    `json:"id"`
	Shop     string    `json:"shop,omitempty"`
	Cart     RDCart    `json:"cart,omitempty"`
	Customer *Customer `json:"customer,omitempty"`
}
