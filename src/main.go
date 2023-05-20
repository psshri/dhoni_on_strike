package main

import (
	"fmt"
	"src/checkSchedule"
	"time"
)

func main() {

	// fmt.Println("API Key:", apiKey)

	url := "https://cricket-live-data.p.rapidapi.com/fixtures-by-date/"
	xRapidAPIHost := "cricket-live-data.p.rapidapi.com"
	info_file_path := "config/info.json"
	today_string := time.Now().Format("2006-01-02")
	// todayString := "2023-05-19"
	fixtures_data_path := "checkSchedule/fixtures.json"

	checkSchedule.Get_schedule(url, xRapidAPIHost, info_file_path, today_string, fixtures_data_path)

	ipl_series_id := 1430
	team_id := 101751

	match_today, match_time_330, match_time_730, matchInfo, err := checkSchedule.Evaluate_schedule(fixtures_data_path, ipl_series_id, team_id, today_string)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println(match_today)
	fmt.Println(match_time_330)
	fmt.Println(match_time_730)
	fmt.Println("Match Info:", matchInfo)

}
