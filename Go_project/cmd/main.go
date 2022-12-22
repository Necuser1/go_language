package main

import (
	"fmt"

	"github.com/neerajsrivastav7/GO_PROJECT/pkg/config"
	"github.com/neerajsrivastav7/GO_PROJECT/pkg/database"
	routes "github.com/neerajsrivastav7/GO_PROJECT/pkg/routers"
)

func main() {
	AppConf := config.GetConfig()
	fmt.Println("Start router")
	database.Connect()
	db := database.DbConnection()
	err := database.Migrate(db)
	if err != nil {
		fmt.Println("Not able to create table in the database")
	} else {
		fmt.Println("Started app on following config")
		fmt.Println(AppConf.App)
		routes.Start(AppConf.App)
	}
}
