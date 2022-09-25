package repository

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"zamp/src/database/postgres"
	"zamp/src/models"
)

func PrintUser() {
	user := models.User{uuid.New().String(), "ritik2"}
	db := postgres.GetDB(context.TODO())
	err := db.Create(&user).Error
	if err != nil {
		fmt.Printf("error %v", err)
		return
	}
	fmt.Printf("created user are %v", user)
}
