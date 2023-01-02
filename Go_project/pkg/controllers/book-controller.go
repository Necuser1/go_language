package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"GO_PROJECT/pkg/comman"
	"GO_PROJECT/pkg/database"

	"github.com/gin-gonic/gin"
)

func CreateBook(c *gin.Context) {
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusBadRequest, "Not a valid data")
		return
	}
	//var data interface{}
	data := comman.Book{}
	err = json.Unmarshal(body, &data)
	result := database.InsertData(data)
	if result == true {
		c.JSON(http.StatusCreated, "Data inserted into the table")
	} else {
		c.JSON(http.StatusBadRequest, "not able to insert data into table")
	}

}

func GetAllBook(c *gin.Context) {
	book := database.GetAllBook()
	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": book,
	})
}

func GetByBookID(c *gin.Context) {
	id := c.Param("bookID")
	fmt.Println(id)
	if id == "" {
		c.JSON(http.StatusBadRequest, "BookID not provided")
	}
	book := database.GetById(id)
	_, err := json.Marshal(&book)
	if err != nil {
		c.JSON(http.StatusBadRequest, "not able to parsed Data")
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code":    http.StatusOK,
			"message": book,
		})
	}
}

func UpdateBook(c *gin.Context) {
	id := c.Param("bookID")
	fmt.Println(id)
}
func DeleteBook(c *gin.Context) {
	id := c.Param("bookID")
	deleteStatus := database.DeleteBook(id)
	if deleteStatus == true {
		c.JSON(http.StatusAccepted, "Deleted")
	} else {
		c.JSON(http.StatusAccepted, "Not Deleted")
	}
}
