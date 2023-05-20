// This script fetches the match schedule for today and finds out whether there
// will be any match of ${team_name}, if there is a match then at what time

package checkSchedule

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"syscall"
)

// structs //////////////////////////////////////////////////////////
type MatchInfo struct {
	ID    int    `json:"id"`
	Team1 string `json:"team1"`
	Team2 string `json:"team2"`
	Date  string `json:"date"`
	Time  string `json:"time"`
}

type Match struct {
	Away struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	} `json:"away"`
	Date string `json:"date"`
	Home struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	} `json:"home"`
	ID       int `json:"id"`
	SeriesID int `json:"series_id"`
}

type MatchData struct {
	Results []Match `json:"results"`
}

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
func Get_schedule(url string, xRapidAPIHost string, info_file_path string, today_string string, fixtures_data_path string) {
	xRapidAPIKey, err := readAPIKey(info_file_path)
	if err != nil {
		log.Fatal(err)
	}

	url = url + today_string
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
	err = ioutil.WriteFile(fixtures_data_path, formattedData_fixtures, 0644)
	if err != nil {
		fmt.Println("Error writing to file: ", err)
		return
	}

	fmt.Println("Today's match schedule fetched successfully")
}

func Evaluate_schedule(fixtures_data_path string, ipl_series_id int, team_id int, today_string string) (int, int, int, MatchInfo, error) {
	matchToday := 0
	matchTime330 := 0
	matchTime730 := 0

	fixturesData, err := ioutil.ReadFile(fixtures_data_path)
	if err != nil {
		return 0, 0, 0, MatchInfo{}, fmt.Errorf("error reading file: %v", err)
	}

	var matchData MatchData
	err = json.Unmarshal(fixturesData, &matchData)
	if err != nil {
		return 0, 0, 0, MatchInfo{}, fmt.Errorf("error decoding JSON: %v", err)
	}

	var matchInfo MatchInfo
	for _, item := range matchData.Results {
		if item.SeriesID == ipl_series_id && (item.Home.ID == team_id || item.Away.ID == team_id) {
			matchToday = 1
			matchDate := today_string
			matchTime := item.Date[11:19]
			matchInfo = MatchInfo{
				ID:    item.ID,
				Team1: item.Home.Name,
				Team2: item.Away.Name,
				Date:  matchDate,
				Time:  matchTime,
			}
			if matchTime == "14:00:00" {
				matchTime730 = 1
			} else {
				matchTime330 = 1
			}
		}
	}

	fmt.Println("Today's schedule evaluated successfully!")

	return matchToday, matchTime730, matchTime330, matchInfo, nil
}
