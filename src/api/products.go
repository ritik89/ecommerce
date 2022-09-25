package rest

import (
	"github.com/gin-gonic/gin"
)

func ProductsRoute(routerGroup *gin.RouterGroup) *gin.RouterGroup {
	rg := routerGroup.Group("/products")
	return rg
}
