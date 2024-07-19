package models

type RDCart struct {
	Items     []CartItem `json:"items"`
	ItemCount int32      `json:"itemCount"`
}

type CartItem struct {
	IIQID        string `json:"iiqID"`
	Sku          string `json:"sku"`
	PriceInCents int32  `json:"priceInCents"`
	Discount     int32  `json:"discount"`
	Quantity     int32  `json:"quantity"`
	Shipping     string `json:"shipping"`
	Title        string `json:"title"`
	URL          string `json:"url"`
	Img          string `json:"img"`
	Stock        int32  `json:"stock"`
}

func (rc *RDCart) CountItems(vat int32) {
	rc.ItemCount = 0
	for i := range rc.Items {
		existingItem := &rc.Items[i]
		rc.ItemCount += existingItem.Quantity
	}
}

func (rc *RDCart) UpdateInsertCartItem(item CartItem) {
	for i := range rc.Items {
		existingItem := &rc.Items[i]
		if existingItem.IIQID == item.IIQID {
			existingItem.Sku = item.Sku
			existingItem.PriceInCents = item.PriceInCents
			existingItem.Discount = item.Discount
			existingItem.Quantity = item.Quantity
			existingItem.Shipping = item.Shipping
			existingItem.Title = item.Title
			existingItem.URL = item.URL
			existingItem.Img = item.Img
			existingItem.Stock = item.Stock
			return
		}
	}
	rc.Items = append(rc.Items, item)
}

func (rc *RDCart) RemoveCartItem(iiqID string) {
	for i, existingItem := range rc.Items {
		if existingItem.IIQID == iiqID {
			rc.Items = append(rc.Items[:i], rc.Items[i+1:]...)
			return
		}
	}
}

type CheckoutRequestShoppingCart struct {
	CartID    string     `json:"cartID"`
	CartItems []CartItem `json:"cartItems"`
}

type ValidCartItem struct {
	ID           int     `json:"id"`
	IIQID        string  `json:"iiqID"`
	Sku          string  `json:"sku"`
	Stock        int32   `json:"stock"`
	Price        float64 `json:"price"`
	PriceInCents int32   `json:"priceInCents"`
	Discount     int32   `json:"discount"`
	Quantity     int32   `json:"quantity"`
	Shipping     string  `json:"shipping"`
	ItemTotal    int32   `josn:"itemTotal"`
	Name         string  `json:"name"`
	Title        string  `json:"title"`
	URL          string  `json:"url"`
	Img          string  `json:"img"`
}

type ValidCart struct {
	ID         string          `json:"ID"`
	CustomerID int             `json:"customerID,omitempty"`
	StripePID  string          `json:"stripePID,omitempty"`
	Prods      []ValidCartItem `json:"prods"`
}
