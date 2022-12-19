package main

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/lib/pq"
)

type Config struct {
	Host     string
	Port     string
	Password string
	User     string
	DBName   string
}

func NewConnection(config *Config) (*sql.DB, error) {
	postgresqlDbInfo := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		config.Host, config.Port, config.User, config.Password, config.DBName)
	var db *sql.DB
	var err error
	for {
		db, _ = sql.Open("postgres", postgresqlDbInfo)
		err = db.Ping()
		if err != nil {
			fmt.Println("waitting for postgress to up and running")
			time.Sleep(2 * time.Second)
			continue
		} else {
			fmt.Println("postgres started")
			break
		}
	}
	return db, err

}
