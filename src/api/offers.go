package rest

import (
	"github.com/gin-gonic/gin"
	"zamp/src/controller"
)

func offers(routerGroup *gin.RouterGroup) *gin.RouterGroup {
	rg := routerGroup.Group("/offers")
	rg.GET("/products", controller.GetProductsOnOffers)
	return rg
}
