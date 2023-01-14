package router

import (
	"LJWT/pkg/comman"
	"LJWT/pkg/database"
	"LJWT/pkg/jwt_token"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Insert(c *gin.Context) {
	if validateToken(c.GetHeader("token")) == false {
		c.JSON(http.StatusBadGateway, "Not a valid token")
	}
	body, _ := ioutil.ReadAll(c.Request.Body)
	goData := comman.Data{}
	err := json.Unmarshal(body, &goData)
	if err != nil {
		c.JSON(http.StatusBadRequest, "Not provide correct data")
	}
	db := database.GetDB()
	insertError := database.InsertIntoTable(db, goData)
	fmt.Println(insertError)
	c.JSON(http.StatusAccepted, "data inserted Successfully")
}

func GetData(c *gin.Context) {
	if validateToken(c.GetHeader("token")) == false {
		c.JSON(http.StatusBadGateway, "Not a valid token")
	}
	db := database.GetDB()
	data := database.GetAllPostedData(db)
	c.JSON(http.StatusAccepted, data)
}

func validateToken(token string) bool {
	if token == jwt_token.GetToken() {
		return true
	} else {
		return false
	}
}
func GetDataByID(c *gin.Context) {
	if validateToken(c.GetHeader("token")) == false {
		c.JSON(http.StatusBadGateway, "Not a valid token")
	}
	db := database.GetDB()
	id := c.Param("id")
	data, err := database.GetPostedDataByID(id, db)
	if err != nil {
		c.JSON(http.StatusBadGateway, err)
	} else {
		c.JSON(http.StatusAccepted, data)
	}
}

func Login(c *gin.Context) {
	credential := &comman.Auth{}
	err := json.NewDecoder(c.Request.Body).Decode(credential)
	if err != nil {
		fmt.Println(err)
	}

	configCredential := comman.AuthConfig(comman.Auth{})
	fmt.Println(configCredential)
	var ok bool = false
	if credential.UserName == configCredential.UserName && credential.Password == configCredential.UserName {
		ok = true
	}

	if ok == false {
		c.JSON(http.StatusUnauthorized, "Not a valid credential")
	}
	token, err := jwt_token.GenerateJWT(credential)
	if err != nil {
		c.JSON(http.StatusBadGateway, "Not able to create Token")
	}
	c.JSON(http.StatusAccepted, token)
}
