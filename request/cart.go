package request

type PostCartRequest struct {
	UserId    string `json:"user_id" binding:"required""`
	ProductId string `json:"product_id" binding:"required""`
	OfferType string `json:"offer_type"`
	OfferId   string `json:"offer_id"`
}
