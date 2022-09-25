package service

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"zamp/src/repository"
)

func GetOffers(ctx *gin.Context, offerType string, discount int) (interface{}, error) {
	var result interface{}
	var err error
	switch offerType {
	case "product":
		result, err = repository.GetProductsWithDiscounts(ctx, discount)
	default:
		result, err = nil, fmt.Errorf("unkwon offer type")
	}
	return result, err
}
