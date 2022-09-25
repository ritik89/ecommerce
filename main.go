package main

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	restServer "zamp/src/api"
	"zamp/src/configs"
	"zamp/src/database/postgres"
)

func init() {
	fmt.Println("loading app configs")
	configs.SetupConfig()

	log.Info("initializing db connection")
	postgres.InitDB()
}

func main() {
	log.Info("starting server.....")
	err := restServer.BuildServer().Run("0.0.0.0:8080")
	if err != nil {
		log.Fatalf("Error occured while running server. %s", err.Error())
	}
}
