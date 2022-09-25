package rest

import (
	"github.com/gin-gonic/gin"
)

func products(routerGroup *gin.RouterGroup) *gin.RouterGroup {
	rg := routerGroup.Group("/products")
	return rg
}
