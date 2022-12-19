package main

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type Repository struct {
	DB *sql.DB
}

func (repo *Repository) createTable() {
	stml := `CREATE TABLE jsonData (data json)`

	_, err := repo.DB.Exec(stml)

	if err != nil {
		fmt.Println(err)
	}

}
func (repo *Repository) insertEvent(w http.ResponseWriter, r *http.Request) {

	databytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "not a valid data", http.StatusBadRequest)
		return
	}
	if err != nil {
		http.Error(w, "not a valid data", http.StatusBadRequest)
		return
	}
	fmt.Println("comming here")
	sqlStatement := `INSERT INTO  jsonData (data) VALUES ($1)`
	createData := string(databytes)

	_, err = repo.DB.Exec(sqlStatement, createData)
	if err != nil {
		panic(err)
	} else {
		fmt.Println("\nRow inserted successfully!")
	}

}
func (repo *Repository) listAllEvent(w http.ResponseWriter, r *http.Request) {
	rows, err := repo.DB.Query("SELECT data FROM jsonData")
	if err != nil {
		fmt.Println(err)
		http.Error(w, "not a valid Query", http.StatusBadRequest)
		return
	}
	result := make([]interface{}, 0)
	for rows.Next() {
		var data []uint8
		err = rows.Scan(&data)
		result = append(result, string(data))
	}
	fmt.Fprint(w, "Recorded Data:", result)
}

func (repo *Repository) setRouter() {
	http.HandleFunc("/insertEvent", repo.insertEvent)
	http.HandleFunc("/listAllEvent", repo.listAllEvent)
}
func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}
	config_data := Config{
		Host:     os.Getenv("DB_HOST"),
		Port:     "5432",
		Password: "postgres",
		User:     "postgres",
		DBName:   "golangdb",
	}
	fmt.Println(config_data)
	db, err := NewConnection(&config_data)
	if err != nil {
		log.Fatal("could not migrate db")
	}
	r := Repository{
		DB: db,
	}
	r.createTable()
	r.setRouter()
	if err := http.ListenAndServe(":9999", nil); err != nil {
		log.Fatal(err)
	}
}
