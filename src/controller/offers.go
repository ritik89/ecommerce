package controller

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"zamp/src/response"
	"zamp/src/service"
)

func GetProductsOnOffers(ctx *gin.Context) {
	offerType := ctx.Query("offer_type")
	discountPercentage := ctx.Query("discount_percentage")
	discountPercentageInt, err := strconv.Atoi(discountPercentage)

	results, err := service.GetOffers(ctx, offerType, discountPercentageInt)
	if err != nil {
		response.Fail(ctx, 500, err, "")
	}
	response.Success(ctx, 200, results, "")
}
