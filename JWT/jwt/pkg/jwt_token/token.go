package jwt_token

import (
	"LJWT/pkg/comman"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
)

var (
	jwtKey     = comman.Token{}
	token_data = []byte("secret_key")
)

func GenerateJWT(credentials *comman.Auth) (comman.Token, error) {

	expirationTime := time.Now().Add(time.Minute * 5)

	claims := &comman.Claims{
		Username: credentials.UserName,
		Password: credentials.Password,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(token_data)

	fmt.Println(err)

	tokenResponse := comman.Token{}
	tokenResponse.Expires = expirationTime
	tokenResponse.Name = "Token"
	tokenResponse.Value = tokenString
	jwtKey = tokenResponse
	return tokenResponse, nil

}

func GetToken() string {
	return jwtKey.Value
}
