package models

type CartItem struct {
	ID       int    `json:"id"`
	Product  string `json:"product"`
	Variant  string `json:"variant"`
	Price    int    `json:"price"`
	Quantity int    `json:"quantity"`
}
