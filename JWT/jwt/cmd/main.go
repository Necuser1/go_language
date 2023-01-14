package main

import (
	"LJWT/pkg/database"
	routes "LJWT/pkg/router"
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	_ = initDataBase()
	router := gin.Default()
	router.POST("insert/", routes.Insert)
	router.GET("data/", routes.GetData)
	router.GET("data/:id", routes.GetDataByID)
	router.POST("login/", routes.Login)
	fmt.Println("Server started at 8080")
	router.Run(":8080")
}

func initDataBase() error {
	database.Connect()
	db := database.GetDB()
	err := database.CreateTable(db)
	return err
}
