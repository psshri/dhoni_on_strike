package main

import (
	"fmt"
	"src/checkLiveScore"
	"src/checkSchedule"
)

func main() {

	// fmt.Println("API Key:", apiKey)

	fixtures_url := "https://cricket-live-data.p.rapidapi.com/fixtures-by-date/"
	live_score_url := "https://cricket-live-data.p.rapidapi.com/match/"
	xRapidAPIHost := "cricket-live-data.p.rapidapi.com"

	info_file_path := "config/info.json"
	// today_string := time.Now().Format("2006-01-02")
	today_string := "2023-05-19"
	fixtures_data_path := "checkSchedule/fixtures.json"
	live_score_data_path := "checkLiveScore/live_score.json"

	checkSchedule.Get_schedule(fixtures_url, xRapidAPIHost, info_file_path, today_string, fixtures_data_path)

	ipl_series_id := 1430
	team_id := 145221
	player_id := 84717

	match_today, match_time_330, match_time_730, matchInfo, err := checkSchedule.Evaluate_schedule(fixtures_data_path, ipl_series_id, team_id, today_string)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println(match_today)
	fmt.Println(match_time_330)
	fmt.Println(match_time_730)
	fmt.Println("Match Info:", matchInfo)

	checkLiveScore.Get_live_score(matchInfo, live_score_url, info_file_path, xRapidAPIHost, live_score_data_path)

	result := checkLiveScore.Is_batting(live_score_data_path, player_id)
	fmt.Println("Result:", result)

}
