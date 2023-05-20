package main

import (
	"fmt"
	"src/checkLiveScore"
	"src/checkSchedule"
	"time"
)

func print_status(counter int, player_name string, team_name string) {
	fmt.Println()
	if counter == 1 {
		fmt.Println(player_name, "is on strike!")
	} else {
		fmt.Println(team_name+"'", "match is today,", player_name, "is yet to bat!")
	}
	fmt.Println()
}

func main() {

	fixtures_url := "https://cricket-live-data.p.rapidapi.com/fixtures-by-date/"
	live_score_url := "https://cricket-live-data.p.rapidapi.com/match/"
	xRapidAPIHost := "cricket-live-data.p.rapidapi.com"
	info_file_path := "config/info.json"
	// today_string := time.Now().Format("2006-01-02")
	today_string := "2023-05-19"
	fixtures_data_path := "checkSchedule/fixtures.json"
	live_score_data_path := "checkLiveScore/live_score.json"
	ipl_series_id := 1430
	team_name := "Punjab Kings"
	team_id := 145221
	player_name := "Shikhar Dhawan"
	player_id := 84717

	checkSchedule.Get_schedule(fixtures_url, xRapidAPIHost, info_file_path, today_string, fixtures_data_path)

	match_today, match_time_330, match_time_730, matchInfo, err := checkSchedule.Evaluate_schedule(fixtures_data_path, ipl_series_id, team_id, today_string)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	if match_today == 1 {

		// Create a ticker that ticks every 2 seconds
		ticker := time.NewTicker(2 * time.Second)

		// Start a goroutine to execute the functions periodically
		go func() {
			for {
				select {
				case <-ticker.C:
					checkLiveScore.Get_live_score(matchInfo, live_score_url, info_file_path, xRapidAPIHost, live_score_data_path)
					result := checkLiveScore.Is_batting(live_score_data_path, player_id)
					print_status(result, player_name, team_name)
				}
			}
		}()

		// Keep the main goroutine running
		select {}

	} else {
		print("No", team_name, "match today!")
	}

	fmt.Println(match_time_330)
	fmt.Println(match_time_730)

}
