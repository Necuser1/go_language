package routes

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/neerajsrivastav7/GO_PROJECT/pkg/comman"
	"github.com/neerajsrivastav7/GO_PROJECT/pkg/controllers"
)

func Start(config comman.Application) {
	router := gin.Default()
	router.POST("/book/", controllers.CreateBook)
	router.GET("/getbook/", controllers.GetAllBook)
	router.GET("/book/:bookID", controllers.GetByBookID)
	router.DELETE("/book/:bookID", controllers.DeleteBook)
	router.Run(":" + strconv.Itoa(config.MYPORT))
}
