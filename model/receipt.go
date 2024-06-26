package model

type Receipt struct {
	Retailer     string `json:"retailer"`
	PurchaseDate string `json:"purchaseDate"`
	PurchaseTime string `json:"purchaseTime"`
	Items        []Item `json:"items"`
	Total        string `json:"total"` // Note that Total is a string here
	ID           string `json:"id"`    // To store the generated ID
	Points       int    `json:"points"`
}

type Item struct {
	ShortDescription string `json:"shortDescription"`
	Price            string `json:"price"` // Price is also a string
}
