package constants

var LiveOfferTypes = []string{"product", "product_category", "cart", "custom"}

var OfferTypeToTableMapping = map[string]string{
	"product":          "product_offers",
	"product_category": "product_category_offers",
	"cart":             "cart_offers",
	"custom":           "custom_offers",
}
