package database

import (
	"LJWT/pkg/comman"
	"database/sql"
	"encoding/json"
	"fmt"
	"time"

	_ "github.com/lib/pq"
)

var (
	db *sql.DB
)

func Connect() {
	dbConfig := comman.DataBase{}
	config := comman.GetDbConfig(dbConfig)
	postgresqlDbInfo := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		config.IP, config.PORT, config.UserName, config.PASSWORD, config.DATABASE)
	fmt.Println(postgresqlDbInfo)
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
	fmt.Println("Successfully connected!")
}

func GetDB() *sql.DB {
	return db
}

func CreateTable(db *sql.DB) error {
	stmt := `CREATE TABLE Employee 
					( id serial PRIMARY KEY,
					entityId varchar, 
					detail json)`
	_, err := db.Exec(stmt)
	fmt.Println(err)
	if err != nil {
		return err
	}
	fmt.Println("Table created Succesfully")
	return err
}

func InsertIntoTable(DB *sql.DB, data comman.Data) error {
	smt := `INSERT INTO Employee (entityId, detail) VALUES ($1, $2)`
	id := data.ID
	detail := data.Data
	json_data, err := json.Marshal(detail)
	err_result, _ := db.Exec(smt, id, json_data)

	if err_result != nil {
		fmt.Println(err_result)
	}
	return err
}

func GetAllPostedData(db *sql.DB) []comman.Result {
	rows, _ := db.Query("SELECT * FROM Employee")
	ResultData := make([]comman.Result, 0)
	for rows.Next() {
		snb := comman.Result{}
		err := rows.Scan(&snb.Id, &snb.EntityId, &snb.Detail)
		if err != nil {
			fmt.Println(err)
		}
		snb_detail := string(snb.Detail.([]byte))
		var jsonMap map[string]interface{}
		json.Unmarshal([]byte(snb_detail), &jsonMap)
		snb.Detail = jsonMap
		ResultData = append(ResultData, snb)

	}
	return ResultData
}

func GetPostedDataByID(id string, db *sql.DB) (comman.Result, error) {
	result := comman.Result{}
	Query := fmt.Sprintf("SELECT * FROM employee WHERE entityId = '%s'", id)
	fmt.Println(Query)
	rows, err := db.Query(Query)
	if err != nil {
		return result, err
	}
	for rows.Next() {
		err := rows.Scan(&result.Id, &result.EntityId, &result.Detail)
		if err != nil {
			fmt.Println(err)
		}
		snb_detail := string(result.Detail.([]byte))
		var jsonMap map[string]interface{}
		json.Unmarshal([]byte(snb_detail), &jsonMap)
		result.Detail = jsonMap
	}
	fmt.Println(result)
	return result, err
}
