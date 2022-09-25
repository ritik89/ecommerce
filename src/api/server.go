package rest

import (
	"github.com/gin-gonic/gin"
	//"us-stocks-service/api/rest/endpoints"
	//"us-stocks-service/api/rest/endpoints/fund_transfer"
	//"us-stocks-service/api/rest/endpoints/identity"
	//"us-stocks-service/api/rest/endpoints/price_alert_endpoints"
	//"us-stocks-service/internal/pkg/new_relic"
)

func BuildServer() *gin.Engine {
	server := gin.Default()

	v1 := server.Group("/api/v1")
	ProductsRoute(v1)

	server.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "server is running",
		})
	})
	return server
}
