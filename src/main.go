package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
)

var fileContents, err = ioutil.ReadFile("helper/apiKey.txt")
var apiKey string = string(fileContents)
var apiHost string = "cricket-live-data.p.rapidapi.com"

func hitAPI(url string) []uint8 {
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("X-RapidAPI-Key", apiKey)
	req.Header.Add("X-RapidAPI-Host", apiHost)
	res, _ := http.DefaultClient.Do(req)
	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)
	return body
}

func main() {

	// IPL fixtures
	url_fixtures := "https://cricket-live-data.p.rapidapi.com/fixtures-by-series/1430"
	body_fixtures := hitAPI(url_fixtures)

	// create a file to store the data
	file_fixtures, err_fixtures1 := os.Create("raw_output/output_fixtures.json")
	if err_fixtures1 != nil {
		fmt.Println("Error creating file: ", err_fixtures1)
		return
	}
	defer file_fixtures.Close()

	// convert the body data to json format for better readability
	var data_fixtures interface{}
	err_fixtures2 := json.Unmarshal(body_fixtures, &data_fixtures)
	if err_fixtures2 != nil {
		fmt.Println("Error unmarshaling body JSON: ", err_fixtures2)
		return
	}
	formattedData_fixtures, _ := json.MarshalIndent(data_fixtures, "", "    ")

	// writing formattedData_fixtures to the file
	_, err_fixtures3 := file_fixtures.Write(formattedData_fixtures)
	if err_fixtures3 != nil {
		fmt.Println("Error writing to file: ", err_fixtures3)
		return
	}
	fmt.Println("Output written to file successfully!")
}
