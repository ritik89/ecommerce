package repository

import (
	"github.com/gin-gonic/gin"
	"zamp/src/database/postgres"
	"zamp/src/models"
)

func GetProductDetails(ctx *gin.Context, productId string) (models.Product, error) {
	var products models.Product
	db := postgres.GetDB(ctx)
	db.Table("products").Select("*").First(&products)
	return products, nil
}
