package config

import (
	"fmt"

	"github.com/neerajsrivastav7/GO_PROJECT/pkg/comman"
	"github.com/spf13/viper"
)

func GetConfig() comman.Config {
	viper.AddConfigPath("../")
	viper.SetConfigName("config")
	viper.SetConfigType("json")
	viper.ReadInConfig()
	config := comman.Config{}
	err := viper.Unmarshal(&config)
	if err != nil {
		fmt.Println("not able to unmarshal config file")
	}
	return config
}
func GetDbConfig(dbConfig comman.DataBase) comman.DataBase {
	config := GetConfig()
	dbConfig = config.DB
	return dbConfig
}

func GetApplicationConfig(appConfig comman.Application) comman.Application {
	config := GetConfig()
	appConfig = config.App
	return appConfig
}
