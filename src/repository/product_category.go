package repository

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"zamp/src/database/postgres"
	"zamp/src/models"
)

func GetProductCategoryOffers(ctx *gin.Context, productId string, discountPerc int) (interface{}, error) {
	var offers []models.Offer
	db := postgres.GetDB(ctx)
	whereClause := buildWhereQueryForProductCategories(productId, discountPerc)
	db.Table("product_category_offers").Select("*").Joins(
		"left join product_categories" +
			" on product_category_offers.product_category_id = product_categories.id").Joins(
		"left join products on product_categories.id = products.product_category").Where(
		whereClause).Scan(&offers)
	return offers, nil
}

func buildWhereQueryForProductCategories(productId string, discountPerc int) string {
	whereClause := "product_category_offers.is_active = true "
	if discountPerc > 0 {
		whereClause += fmt.Sprintf("AND products.id = %d", productId)
	}
	if discountPerc > 0 {
		whereClause += fmt.Sprintf("AND product_offers.discount_percentage = %d", discountPerc)
	}
	return whereClause
}
