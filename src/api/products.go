package rest

import (
	"github.com/gin-gonic/gin"
	"zamp/src/controller"
)

func products(routerGroup *gin.RouterGroup) *gin.RouterGroup {
	rg := routerGroup.Group("/products")
	rg.GET("/:productId", controller.GetProductDetails)
	return rg
}
