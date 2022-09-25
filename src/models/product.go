package models

type Product struct {
	Id              string  `json:"name" sql:"type:uuid;primary_key;"`
	Name            string  `json:"name"`
	Description     string  `json:"description"`
	ProductCategory string  `json:"product_category" sql:"type:uuid"`
	Units           int     `json:"units"`
	Price           float32 `json:"price"`
	MinSalePrice    float32 `json:"min_sale_price"`
}
