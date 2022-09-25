package repository

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"zamp/src/database/postgres"
	"zamp/src/models"
)

func GetCartItems(ctx *gin.Context, userId string) ([]models.CartItems, error) {
	var cartItems []models.CartItems
	db := postgres.GetDB(ctx)
	db.Table("cart_items").Select("product_id, offer_type, offer_id").Joins("left join cart"+
		" on cart_items.user_id = cart.user_id").Where("cart_items.user_id = ?", userId).Scan(&cartItems)
	return cartItems, nil
}

func buildWhereQueryForCart(userId string) string {
	whereClause := fmt.Sprintf("cart_items.user_id = %s", userId)
	return whereClause
}
