package routes

import (
	"strconv"

	"GO_PROJECT/pkg/comman"
	"GO_PROJECT/pkg/controllers"

	"github.com/gin-gonic/gin"
)

func Start(config comman.Application) {
	router := gin.Default()
	router.POST("/book/", controllers.CreateBook)
	router.GET("/getbook/", controllers.GetAllBook)
	router.GET("/book/:bookID", controllers.GetByBookID)
	router.DELETE("/book/:bookID", controllers.DeleteBook)
	router.Run(":" + strconv.Itoa(config.MYPORT))
}
