package repository

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"zamp/src/constants"
	"zamp/src/database/postgres"
	"zamp/src/models"
)

func GetDiscountedProducts(ctx *gin.Context, productId string, discountPerc int) (interface{}, error) {
	var products []models.Product
	db := postgres.GetDB(ctx)
	whereClause := buildWhereQueryForProducts(productId, discountPerc)
	db.Table("product_offers").Select("*").Joins("left join products" +
		" on product_offers.product_id = products.id").Where(whereClause).Scan(&products)
	return products, nil
}

func GetProductOfferDetails(ctx *gin.Context, productId string, discountPerc int) (interface{}, error) {
	var offers []models.Offer
	db := postgres.GetDB(ctx)
	whereClause := buildWhereQueryForProducts(productId, discountPerc)
	db.Table("product_offers").Select("*").Joins("left join products" +
		" on product_offers.product_id = products.id").Where(whereClause).Scan(&offers)
	return offers, nil
}

func GetOfferDetails(ctx *gin.Context, offerType string, offerId string) (models.Offer, error) {
	var offer models.Offer
	db := postgres.GetDB(ctx)
	offerTable, _ := getOfferTable(offerType)
	db.Table(offerTable).Select("*").Where("offer_id = ?", offerId).Scan(&offer)
	return offer, nil
}

func getOfferTable(offerType string) (string, error) {
	table, _ := constants.OfferTypeToTableMapping[offerType]
	return table, nil
}

func buildWhereQueryForProducts(productId string, discountPerc int) string {
	whereClause := "product_offers.is_active = true "
	if productId != "" {
		whereClause += fmt.Sprintf("AND products.id = %s", productId)
	}
	if discountPerc > 0 {
		whereClause += fmt.Sprintf("AND product_offers.discount_percentage = %d", discountPerc)
	}
	return whereClause
}
