package rest

import (
	"github.com/gin-gonic/gin"
	"zamp/src/controller"
)

func cart(routerGroup *gin.RouterGroup) *gin.RouterGroup {
	rg := routerGroup.Group("/cart")
	rg.GET("", controller.GetUserCart)
	return rg
}
