package models

import (
	"math"
	"time"
)

type NewOrder struct {
	Data NewOrderData `json:"data"`
}

type NewOrderData struct {
	ID                  int             `json:"id,omitempty"`
	Status              string          `json:"status,omitempty"`
	Shop                string          `json:"shop,omitempty"`
	CustomerID          int             `json:"customerID,omitempty"`
	CartID              string          `json:"cartID,omitempty"`
	FullName            string          `json:"fullName,omitempty"`
	Email               string          `json:"email,omitempty"`
	SubTotal            int32           `json:"subTotal,omitempty"`
	Tax                 int32           `json:"tax,omitempty"`
	Shipping            int32           `json:"shipping,omitempty"`
	Items               []ValidCartItem `json:"items,omitempty"`
	StripePID           string          `json:"stripePID,omitempty"`
	StripeAmount        int32           `json:"stripeAmount,omitempty"`
	StripeCurrency      string          `json:"stripeCurrency,omitempty"`
	StripeStatus        string          `json:"stripeStatus,omitempty"`
	City                string          `json:"city,omitempty"`
	Country             string          `json:"country,omitempty"`
	State               string          `json:"state,omitempty"`
	Line1               string          `json:"line1,omitempty"`
	Line2               string          `json:"line2,omitempty"`
	PostalCode          string          `json:"postalCode,omitempty"`
	Name                string          `json:"name,omitempty"`
	Phone               string          `json:"phone,omitempty"`
	Carrier             string          `json:"carrier,omitempty"`
	TrackingID          string          `json:"trackingID,omitempty"`
	InvoiceLink         string          `json:"invoiceLink,omitempty"`
	StripeInvoiceNumber string          `json:"stripeInvoiceNumber,omitempty"`
	InvoiceItems        []InvoiceItem   `json:"invoiceItems,omitempty"`
	StripeInvoiceID     string          `json:"stripeInvoiceID,omitempty"`
	CreditNoteLink      string          `json:"creditNoteLink,omitempty"`
	CreditReason        string          `json:"creditReason,omitempty"`
}

func (no *NewOrder) CalcValues(vat int32) {
	for _, item := range no.Data.Items {
		no.Data.SubTotal += item.ItemTotal
		tf := float64(item.ItemTotal) * float64(vat) / 100
		tax := int32(math.RoundToEven(tf))
		no.Data.Tax += tax
	}
	no.Data.Shipping = 750
	if no.Data.SubTotal > 5000 {
		no.Data.Shipping = 0
	}
}

type SaveOrders struct {
	Data []SaveOrderData `json:"data,omitempty"`
}

type SavedOrder struct {
	Data SaveOrderData `json:"data,omitempty"`
}

type SaveOrderData struct {
	ID         int            `json:"id,omitempty"`
	Attributes SavedOrderAttr `json:"attributes,omitempty"`
}

type SavedOrderAttr struct {
	Status          string          `json:"status,omitempty"`
	Shop            string          `json:"shop,omitempty"`
	CustomerID      int             `json:"customerID,omitempty"`
	CartID          string          `json:"cartID,omitempty"`
	FullName        string          `json:"fullName,omitempty"`
	Email           string          `json:"email,omitempty"`
	SubTotal        int32           `json:"subTotal,omitempty"`
	Tax             int32           `json:"tax,omitempty"`
	Shipping        int32           `json:"shipping,omitempty"`
	Items           []ValidCartItem `json:"items,omitempty"`
	StripePID       string          `json:"stripePID,omitempty"`
	StripeAmount    int32           `json:"stripeAmount,omitempty"`
	StripeCurrency  string          `json:"stripeCurrency,omitempty"`
	StripeStatus    string          `json:"stripeStatus,omitempty"`
	City            string          `json:"city,omitempty"`
	Country         string          `json:"country,omitempty"`
	State           string          `json:"state,omitempty"`
	Line1           string          `json:"line1,omitempty"`
	Line2           string          `json:"line2,omitempty"`
	PostalCode      string          `json:"postalCode,omitempty"`
	Name            string          `json:"name,omitempty"`
	Phone           string          `json:"phone,omitempty"`
	Carrier         string          `json:"carrier,omitempty"`
	TrackingID      string          `json:"trackingID,omitempty"`
	InvoiceLink     string          `json:"invoiceLink,omitempty"`
	CreatedAt       time.Time       `json:"createdAt,omitempty"`
	InvoiceItems    []InvoiceItem   `json:"invoiceItems,omitempty"`
	StripeInvoiceID string          `json:"stripeInvoiceID,omitempty"`
	CreditNoteLink  string          `json:"creditNoteLink,omitempty"`
	CreditReason    string          `json:"creditReason,omitempty"`
}

type InvoiceItem struct {
	Description string `json:"description,omitempty"`
	UnitAmount  int32  `json:"unitAmount"`
	Quantity    int32  `json:"quantity"`
}
