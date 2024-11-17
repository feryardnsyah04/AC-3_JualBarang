package models

type CartItem struct {
	ID       int    `json:"id"`
	Product  string `json:"product"`
	Category string `json:"category"`
	Price    int    `json:"price"`
	Quantity int    `json:"quantity"`
}
