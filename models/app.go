package models

type App struct {
	Data AppData `json:"data,omitempty"`
}

type AppData struct {
	ID         int32         `json:"id,omitempty"`
	Attributes AppAttributes `json:"attributes,omitempty"`
}

type AppAttributes struct {
	Name           string `json:"name,omitempty"`
	Description    string `json:"description,omitempty"`
	Currency       string `json:"currency,omitempty"`
	ExchangeRate   int32  `json:"exchangeRate,omitempty"`
	Vat            int32  `json:"vat,omitempty"`
	AllowSale      bool   `json:"allowSale,omitempty"`
	Menu           string `json:"menu,omitempty"`
	Footer         string `json:"footer,omitempty"`
	GAC            string `json:"gac"`
	EmailOrderList string `json:"emailOrderList"`

	Logo        Img         `json:"logo,omitempty"`
	LandingPage LandingPage `json:"landing_page"`
}
