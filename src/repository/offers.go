package repository

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"zamp/src/database/postgres"
	"zamp/src/models"
)

func GetProductsWithDiscounts(ctx *gin.Context, discountPerc int) (interface{}, error) {
	var products []models.Product
	db := postgres.GetDB(ctx)
	whereClause := buildWhereQuery(discountPerc)
	query := db.Table("product_offers").Select("*").Joins("left join products" +
		" on product_offers.product_id = products.id")
	if whereClause != "" {
		query.Where(whereClause)
	}
	query.Scan(&products)
	return products, nil
}

func buildWhereQuery(discountPerc int) string {
	whereClause := ""
	if discountPerc > 0 {
		whereClause = fmt.Sprintf("product_offers.discount_percentage = %d", discountPerc)
	}
	return whereClause
}
