// This script creates a database and a table in MySQL

package createDBTable

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func CreateDBTable() {

	// open the deDetails.json file
	dbDetails, err := os.Open("config/dbDetails.json")
	if err != nil {
		fmt.Println(err)
	}

	// defer the closing of dbDetails file so that we can parse it later on
	defer dbDetails.Close()

	// read the opened dbDetails file as a byte array
	byteValue, _ := ioutil.ReadAll(dbDetails)

	var result map[string]interface{}

	json.Unmarshal([]byte(byteValue), &result)

	rds1 := result["rds"]
	username1 := result["username"]
	password1 := result["password"]
	dbName1 := result["dbName"]
	tableName1 := result["tableName"]

	rds, ok := rds1.(string)
	if !ok {
		fmt.Println("Error converting rds1 to string")
		return
	}

	username, ok := username1.(string)
	if !ok {
		fmt.Println("Error converting username1 to string")
		return
	}

	password, ok := password1.(string)
	if !ok {
		fmt.Println("Error converting password1 to string")
		return
	}

	dbName, ok := dbName1.(string)
	if !ok {
		fmt.Println("Error converting dbName1 to string")
		return
	}

	tableName, ok := tableName1.(string)
	if !ok {
		fmt.Println("Error converting tableName1 to string")
		return
	}

	// open a db connection
	db, err := sql.Open(rds, username+":"+password+"@tcp(127.0.0.1:3306)/")

	// check for errors
	if err != nil {
		log.Fatal(err)
		fmt.Println("error in connecting to db")
	}

	// the below line of code makes sure that the db connection is closed once
	// the function is executed completely
	defer db.Close()

	// if the database exists delete it first
	_, err = db.Exec("DROP DATABASE IF EXISTS " + dbName)

	// create a new database
	_, err = db.Exec("CREATE DATABASE IF NOT EXISTS " + dbName)
	if err != nil {
		log.Fatal(err)
		fmt.Println("error in creating database")
	}

	// switch to the new database
	_, err = db.Exec("USE " + dbName)
	if err != nil {
		log.Fatal(err)
		fmt.Println("error in using the database")
	}

	// create a new table in the database
	_, err = db.Exec("CREATE TABLE IF NOT EXISTS " + tableName + " (id INT AUTO_INCREMENT PRIMARY KEY, date DATE, time TIME, match_id INT, match_subtitle VARCHAR(255), home_team VARCHAR(255), away_team VARCHAR(255), result VARCHAR(255), status VARCHAR(255))")
	if err != nil {
		log.Fatal(err)
		fmt.Println("error in creating a table")
	}

	fmt.Println("database " + dbName + " and table " + tableName + " created successfully!")
}
