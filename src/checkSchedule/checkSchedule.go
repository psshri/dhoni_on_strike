// This script fetches the match schedule for today and finds out whether there
// will be any match of ${team_name}, if there is a match then at what time

package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"syscall"
	"time"
)

// functions ////////////////////////////////////////////////////////

func readAPIKey(filePath string) (string, error) {
	file, err := ioutil.ReadFile(filePath)
	if err != nil {
		if e, ok := err.(*os.PathError); ok && e.Err == syscall.ENOENT {
			return "", fmt.Errorf("Error: File '%s' not found.", filePath)
		}
		return "", fmt.Errorf("Error: Unable to read file '%s'.", filePath)
	}

	var data map[string]interface{}
	err = json.Unmarshal(file, &data)
	if err != nil {
		return "", err
	}

	apiKey, ok := data["rapidAPI_api_key"].(string)
	if !ok {
		return "", fmt.Errorf("Error: 'rapidAPI_api_key' not found in JSON data.")
	}

	return apiKey, nil
}

//////////////////////////////////////////////////////////////////////////

// get today's match schedule
func getSchedule(url string, xRapidAPIHost string, infoFilePath string, todayString string, fixtureDataPath string) {
	xRapidAPIKey, err := readAPIKey(infoFilePath)
	if err != nil {
		log.Fatal(err)
	}

	url = url + todayString
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Add("X-RapidAPI-Key", xRapidAPIKey)
	req.Header.Add("X-RapidAPI-Host", xRapidAPIHost)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	body_fixtures, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	// convert the body_fixtures to json format for better readability
	var data_fixtures interface{}
	err = json.Unmarshal(body_fixtures, &data_fixtures)
	if err != nil {
		fmt.Println("Error unmarshaling body_fixtures: ", err)
		return
	}
	formattedData_fixtures, _ := json.MarshalIndent(data_fixtures, "", "    ")

	// writing formattedData_fixtures to the file
	err = ioutil.WriteFile(fixtureDataPath, formattedData_fixtures, 0644)
	if err != nil {
		fmt.Println("Error writing to file: ", err)
		return
	}

	fmt.Println("Today's match schedule fetched successfully")
}

func main() {
	apiKey, err := readAPIKey("../config/info.json")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("API Key:", apiKey)

	url := "https://cricket-live-data.p.rapidapi.com/fixtures-by-date/"
	xRapidAPIHost := "cricket-live-data.p.rapidapi.com"
	infoFilePath := "../config/info.json"
	todayString := time.Now().Format("2006-01-02")
	// todayString := "2023-05-19"
	fixtureDataPath := "fixtures.json"

	getSchedule(url, xRapidAPIHost, infoFilePath, todayString, fixtureDataPath)

	iplSeriesID := 1430
	teamID := 145221

}
