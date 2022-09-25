package configs

import (
	"fmt"
	"github.com/spf13/viper"
)

func SetupConfig() {
	configName := "local"
	viper.SetConfigName(configName)
	viper.AddConfigPath("./configs")
	viper.SetConfigType("yml")

	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("Error reading config file, %s\n", err)
	}
}
