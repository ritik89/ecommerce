package service

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"zamp/src/models"
	"zamp/src/repository"
)

func GetUserCart(ctx *gin.Context, userId string) (interface{}, error) {
	var result = make(map[string]interface{})
	cartItems, err := repository.GetCartItems(ctx, userId)
	result["cartItems"] = cartItems
	if err != nil {
		log.Errorf("couldn't fetch cartdetails of user %s", userId)
		return nil, err
	}
	cartTotal, err := GetCartTotal(ctx, cartItems)
	if err != nil {
		log.Errorf("couldn't fetch cartdetails of user %s", userId)
		return nil, err
	}
	result["cartTotal"] = cartTotal["cartTotal"]
	result["cartMinSaleValue"] = cartTotal["cartMinSaleValue"]
	return result, err
}

func GetCartTotal(ctx *gin.Context, cartItems []models.CartItems) (map[string]interface{}, error) {
	var cartTotal float32
	var cartMinSaleValue float32
	for _, cartItem := range cartItems {
		productDetails, err := repository.GetProductDetails(ctx, cartItem.ProductId)
		if err != nil {
			log.Errorf("error while calculating cart total")
			return nil, err
		}
		offerDetails, err := repository.GetOfferDetails(ctx, cartItem.OfferType, cartItem.OfferId)
		cartTotal += getEffectiveProductPrice(productDetails.Price, productDetails.MinSalePrice,
			offerDetails.DiscountPercentage, offerDetails.CashbackAmount, cartTotal)
		cartMinSaleValue += productDetails.MinSalePrice
	}
	return map[string]interface{}{
		"cartTotal":        cartTotal,
		"cartMinSaleValue": cartMinSaleValue,
	}, nil
}

func getEffectiveProductPrice(price float32, productMinSalePrice float32, discountPerc int, cashback float32, cartTotal float32) float32 {
	if cartTotal+price < productMinSalePrice {
		return price
	}
	discountedPrice := price - cashback - float32(discountPerc)*0.01
	if discountedPrice < productMinSalePrice {
		return productMinSalePrice
	}
	return discountedPrice
}
