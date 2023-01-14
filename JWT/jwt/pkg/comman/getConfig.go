package comman

import (
	"fmt"

	"github.com/spf13/viper"
)

func GetConfig() Config {
	viper.AddConfigPath("../")
	viper.SetConfigName("config")
	viper.SetConfigType("json")
	viper.ReadInConfig()
	config := Config{}
	err := viper.Unmarshal(&config)
	if err != nil {
		fmt.Println("not able to unmarshal config file")
	}
	return config
}
func GetDbConfig(dbConfig DataBase) DataBase {
	config := GetConfig()
	dbConfig = config.DB
	return dbConfig
}

func GetApplicationConfig(appConfig Application) Application {
	config := GetConfig()
	appConfig = config.App
	return appConfig
}

func AuthConfig(authConfig Auth) Auth {
	config := GetConfig()
	authConfig = config.Jwt
	return authConfig
}
