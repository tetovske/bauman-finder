package config

import (
	"github.com/spf13/viper"
	"log"
)

const Directory = "./config"

func getConfigFiles() []string {
	return []string{"bot_config", "bot_messages"}
}

func Init() {
	viper.AddConfigPath(Directory)

	for _, filePath := range getConfigFiles() {
		viper.SetConfigName(filePath)
		err := viper.MergeInConfig()
		if err != nil {
			log.Fatal(err)
		}
	}
}
