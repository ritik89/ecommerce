package service

import (
	"github.com/gin-gonic/gin"
	"github.com/google/martian/log"
	"zamp/src/constants"
	"zamp/src/repository"
)

func GetProductDetails(ctx *gin.Context, productId string) (interface{}, error) {
	result := make(map[string]interface{})
	productDetails, err := repository.GetProductDetails(ctx, productId)
	result["productDetails"] = productDetails

	offers := getEligibleProductOffers(ctx, productId)
	result["offers"] = offers
	return result, err
}

func getEligibleProductOffers(ctx *gin.Context, productId string) map[string]interface{} {
	offers := make(map[string]interface{})
	var err error
	for _, offerType := range constants.LiveOfferTypes {
		offers[offerType], err = GetOffers(ctx, offerType, productId, 0)
		if err != nil {
			log.Errorf("Error while getting offer of type %s", offerType)
		}
	}
	return offers
}
