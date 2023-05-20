package checkLiveScore

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"src/checkSchedule"
	"syscall"
)

// structs ///////////////////////////////////////////////////////////

type LiveScoreData struct {
	Results struct {
		LiveDetails struct {
			Scorecard []struct {
				Batting []struct {
					PlayerID   int    `json:"player_id"`
					PlayerName string `json:"player_name"`
				} `json:"batting"`
			} `json:"scorecard"`
		} `json:"live_details"`
	} `json:"results"`
}

//////////////////////////////////////////////////////////////////////

// helper functions /////////////////////////////////////////////////////////

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

// logic functions /////////////////////////////////////////////////////////

// get live score
func Get_live_score(matchInfo checkSchedule.MatchInfo, url string, info_file_path string, xRapidAPIHost string, live_score_data_path string) {
	matchID := matchInfo.ID
	url = fmt.Sprintf("%s%d", url, matchID)
	// xRapidAPIKey := readAPIKey(infoFilePath)
	xRapidAPIKey, err := readAPIKey(info_file_path)
	if err != nil {
		log.Fatal(err)
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("Error creating HTTP request:", err)
		return
	}
	req.Header.Add("X-RapidAPI-Key", xRapidAPIKey)
	req.Header.Add("X-RapidAPI-Host", xRapidAPIHost)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	body_liveScore, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	// convert the body_fixtures to json format for better readability
	var data_liveScore interface{}
	err = json.Unmarshal(body_liveScore, &data_liveScore)
	if err != nil {
		fmt.Println("Error unmarshaling body_fixtures: ", err)
		return
	}
	formattedData_liveScore, _ := json.MarshalIndent(data_liveScore, "", "    ")

	// writing formattedData_liveScore to the file
	err = ioutil.WriteFile(live_score_data_path, formattedData_liveScore, 0644)
	if err != nil {
		fmt.Println("Error writing to file: ", err)
		return
	}

	fmt.Println("Live score updated successfully!")
}

// check if ${player_id} is on crease or not
func Is_batting(live_score_data_path string, player_id int) int {
	counter := 0

	data_liveScore, err := ioutil.ReadFile(live_score_data_path)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return counter
	}

	var liveScore LiveScoreData
	err = json.Unmarshal(data_liveScore, &liveScore)
	if err != nil {
		fmt.Println("Error parsing JSON:", err)
		return counter
	}

	listBatsmen := []int{}
	scorecard := liveScore.Results.LiveDetails.Scorecard
	for _, item1 := range scorecard {
		for _, item2 := range item1.Batting {
			listBatsmen = append(listBatsmen, item2.PlayerID)
		}
	}

	for _, id := range listBatsmen {
		if player_id == id {
			counter = 1
			break
		}
	}

	fmt.Println("Live score evaluated successfully")

	return counter
}
