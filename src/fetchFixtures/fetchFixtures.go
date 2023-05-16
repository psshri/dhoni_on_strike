// This script downloads IPL fixtures and creates a fixtures.json file

package fetchFixtures

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
)

var fileContents, err = ioutil.ReadFile("config/apiKey.txt")

//	if err != nil {
//		fmt.Println("error reading file: ", err)
//	}
var apiKey string = string(fileContents)

var apiHost string = "cricket-live-data.p.rapidapi.com"

func HitAPI() {

	url_fixtures := "https://cricket-live-data.p.rapidapi.com/fixtures-by-series/1430"
	req_fixtures, _ := http.NewRequest("GET", url_fixtures, nil)
	req_fixtures.Header.Add("X-RapidAPI-Key", apiKey)
	req_fixtures.Header.Add("X-RapidAPI-Host", apiHost)
	res_fixtures, _ := http.DefaultClient.Do(req_fixtures)
	defer res_fixtures.Body.Close()
	body_fixtures, _ := io.ReadAll(res_fixtures.Body)

	// create a file to store the data
	file_fixtures, err_fixtures1 := os.Create("fetchFixtures/fixtures.json")
	if err_fixtures1 != nil {
		fmt.Println("Error creating fixtures.json: ", err_fixtures1)
		return
	}
	defer file_fixtures.Close()

	// convert the body data to json format for better readability
	var data_fixtures interface{}
	err_fixtures2 := json.Unmarshal(body_fixtures, &data_fixtures)
	if err_fixtures2 != nil {
		fmt.Println("Error unmarshaling body_fixtures: ", err_fixtures2)
		return
	}
	formattedData_fixtures, _ := json.MarshalIndent(data_fixtures, "", "    ")

	// writing formattedData_fixtures to the file
	_, err_fixtures3 := file_fixtures.Write(formattedData_fixtures)
	if err_fixtures3 != nil {
		fmt.Println("Error writing to file: ", err_fixtures3)
		return
	}
	fmt.Println("fixtures.json created successfully!")

	// return body
}
