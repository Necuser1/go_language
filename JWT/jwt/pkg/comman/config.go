package comman

import (
	"time"

	"github.com/golang-jwt/jwt"
)

type Config struct {
	App Application `json:"app"`
	DB  DataBase    `json:"db"`
	Jwt Auth        `json:"auth"`
}

type Auth struct {
	UserName string `json:"username"`
	Password string `json:"password"`
}
type DataBase struct {
	IP       string `json:"ip"`
	PORT     string `json:"port"`
	UserName string `json:"username"`
	PASSWORD string `json:"password"`
	DATABASE string `json:"database"`
}

type Application struct {
	MYIP   string `json:"myip"`
	MYPORT int    `json:"myport"`
}

type Data struct {
	ID   string      `json:"ID"`
	Data interface{} `json:"data"`
}

type Result struct {
	Id       int
	EntityId string
	Detail   interface{}
}
type Claims struct {
	Username string `json:"username"`
	Password string `json:"password"`
	jwt.StandardClaims
}

type Token struct {
	Name    string
	Value   string
	Expires time.Time
}
