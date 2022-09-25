package service

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"zamp/src/repository"
)

// TODO to use offerType for showing different offer types on homepage
func GetProductsOnOffer(ctx *gin.Context, offerType string, discount int) (interface{}, error) {
	var result interface{}
	var err error
	result, err = repository.GetDiscountedProducts(ctx, "", discount)
	return result, err
}

func GetOffers(ctx *gin.Context, offerType string, productId string, discount int) (interface{}, error) {
	var result interface{}
	var err error
	switch offerType {
	case "product":
		result, err = repository.GetProductOfferDetails(ctx, productId, discount)
	case "product_category":
		result, err = repository.GetProductCategoryOffers(ctx, productId, discount)
	case "cart":
		result, err = nil, nil
	case "custom":
		result, err = nil, nil
	default:
		result, err = nil, fmt.Errorf("unkwon offer type")
	}
	return result, err
}
