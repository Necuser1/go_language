package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Handler(c *gin.Context) {

	data, err := ioutil.ReadAll(c.Request.Body)
	var unmarshalData interface{}
	err = json.Unmarshal(data, &unmarshalData)
	fmt.Println(unmarshalData)
	fmt.Println(err)
	c.JSON(http.StatusAccepted, "data data")

}
func main() {
	app := gin.Default()
	app.POST("/ngsi-ld/v1/entities/:id/attrs/:attribute", Handler)
	fmt.Println("Server Started:")
	app.Run(":8080")
}
