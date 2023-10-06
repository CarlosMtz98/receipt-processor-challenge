package dto

type CreateReceiptRequest struct {
	Retailer     string       `json:"retailer" validate:"required"`
	PurchaseDate string       `json:"purchaseDate" validate:"required,datetime=2006-01-02"`
	PurchaseTime string       `json:"purchaseTime" validate:"required,datetime=15:04"`
	Items        []CreateItem `json:"items" validate:"required,min=1,dive"`
	Total        string       `json:"total" validate:"required,currency"`
}

type CreateItem struct {
	ShortDescription string `json:"shortDescription" validate:"required"`
	Price            string `json:"price" validate:"required,currency"`
}
