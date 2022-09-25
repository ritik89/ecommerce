package models

type Cart struct {
	Id                 string  `json:"id" sql:"type:uuid;primary_key;"`
	Description        string  `json:"description"`
	DiscountPercentage string  `json:"discount_percentage"`
	CashbackAmount     float32 `json:"cashback_amount"`
	VoucherCode        float32 `json:"voucher_code"`
	MinPurchaseAmount  float32 `json:"min_purchase_amount"`
	MaxDiscountAllowed float32 `json:"max_discount_allowed"`
	IsActive           float32 `json:"is_active"`
}

type CartItems struct {
	ProductId string `json:"product_id"`
	OfferType string `json:"offer_type"`
	OfferId   string `json:"offer_id"`
}
