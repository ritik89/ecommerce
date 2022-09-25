package repository

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"zamp/src/database/postgres"
	"zamp/src/models"
)

func GetDiscountedProducts(ctx *gin.Context, productId string, discountPerc int) (interface{}, error) {
	var products []models.Product
	db := postgres.GetDB(ctx)
	whereClause := buildWhereQueryForProducts(productId, discountPerc)
	db.Table("product_offers").Select("*").Joins("left join products" +
		" on product_offers.product_id = products.id").Where(whereClause).Scan(&products)
	return products, nil
}

func GetProductOfferDetails(ctx *gin.Context, productId string, discountPerc int) (interface{}, error) {
	var offers []models.Offer
	db := postgres.GetDB(ctx)
	whereClause := buildWhereQueryForProducts(productId, discountPerc)
	db.Table("product_offers").Select("*").Joins("left join products" +
		" on product_offers.product_id = products.id").Where(whereClause).Scan(&offers)
	return offers, nil
}

func buildWhereQueryForProducts(productId string, discountPerc int) string {
	whereClause := "product_offers.is_active = true "
	if discountPerc > 0 {
		whereClause += fmt.Sprintf("AND products.id = %d", productId)
	}
	if discountPerc > 0 {
		whereClause += fmt.Sprintf("AND product_offers.discount_percentage = %d", discountPerc)
	}
	return whereClause
}
