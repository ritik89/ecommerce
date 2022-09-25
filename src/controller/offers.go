package controller

import (
	"github.com/gin-gonic/gin"
	"google.golang.org/appengine/log"
	"strconv"
	"zamp/src/response"
	"zamp/src/service"
)

func GetProductsOnOffers(ctx *gin.Context) {
	offerType := ctx.Query("offer_type")
	discountPercentage := ctx.Query("discount_percentage")
	discountPercentageInt, err := strconv.Atoi(discountPercentage)

	results, err := service.GetProductsOnOffer(ctx, offerType, discountPercentageInt)
	if err != nil {
		log.Errorf(ctx, "error while getting products with discount %s", discountPercentage)
		response.Fail(ctx, 500, err, "")
	}
	response.Success(ctx, 200, results, "")
}
