package dto

type CreateReceiptResponse struct {
	ID string `json:"id"`
}

type GetPointsResponse struct {
	Points int `json:"points"`
}

type ResponseErrorModel struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Details string `json:"details"`
}
