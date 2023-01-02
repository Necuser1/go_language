package database

import (
	"fmt"

	"GO_PROJECT/pkg/comman"
	"GO_PROJECT/pkg/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	db *gorm.DB
)

func Connect() {
	dbConfig := comman.DataBase{}
	config := config.GetDbConfig(dbConfig)
	postgresqlDbInfo := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s",
		config.IP, config.PORT, config.UserName, config.PASSWORD, config.DATABASE,
	)
	fmt.Println(postgresqlDbInfo)
	var err error
	db, err = gorm.Open(postgres.Open(postgresqlDbInfo), &gorm.Config{})
	if err != nil {
		panic("Not able to connect database")
	}
}
func DbConnection() *gorm.DB {
	return db
}
