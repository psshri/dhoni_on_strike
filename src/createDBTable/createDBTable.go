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

	rds, ok := rds1.(string)
	if !ok {
		fmt.Println("Error: username is not a string type")
		return
	}

	username, ok := username1.(string)
	if !ok {
		fmt.Println("Error: username is not a string type")
		return
	}

	password, ok := password1.(string)
	if !ok {
		fmt.Println("Error: username is not a string type")
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

	// create a new database
	_, err = db.Exec("CREATE DATABASE IF NOT EXISTS IPL_Fixtures")
	if err != nil {
		log.Fatal(err)
		fmt.Println("error in creating database")
	}

	// switch to the new database
	_, err = db.Exec("USE IPL_Fixtures")
	if err != nil {
		log.Fatal(err)
		fmt.Println("error in using the database")
	}

	// create a new table in the database
	_, err = db.Exec("CREATE TABLE IF NOT EXISTS fixtures_2023 (id INT AUTO_INCREMENT PRIMARY KEY, date DATE, match_id INT, match_subtitle VARCHAR(255), home_team VARCHAR(255), away_team VARCHAR(255), result VARCHAR(255), status VARCHAR(255))")
	if err != nil {
		log.Fatal(err)
		fmt.Println("error in creating a table")
	}

	fmt.Println("database and table created successfully!")
}
