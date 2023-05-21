package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"src/checkLiveScore"
	"src/checkSchedule"
	"strconv"
	"syscall"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func print_status(counter int, player_name string, team_name string, info_file_path string) {
	fmt.Println()
	if counter == 1 {
		fmt.Println(player_name, "is on strike!")
		telegram_bot_send_text(player_name, info_file_path)
	} else {
		fmt.Println(team_name+"'", "match is today,", player_name, "is yet to bat!")
	}
	fmt.Println()
}

func telegramKeys(info_file_path string) (string, string, error) {
	file, err := ioutil.ReadFile(info_file_path)
	if err != nil {
		if e, ok := err.(*os.PathError); ok && e.Err == syscall.ENOENT {
			return "", "", fmt.Errorf("Error: File '%s' not found.", info_file_path)
		}
		return "", "", fmt.Errorf("Error: Unable to read file '%s'.", info_file_path)
	}

	var data map[string]interface{}
	err = json.Unmarshal(file, &data)
	if err != nil {
		return "", "", err
	}

	bot_token, ok := data["telegram_bot_token"].(string)
	if !ok {
		return "", "", fmt.Errorf("Error: 'telegram_bot_token' not found in JSON data.")
	}

	chat_id, ok := data["telegram_chat_id"].(string)
	if !ok {
		return "", "", fmt.Errorf("Error: 'telegram_chat_id' not found in JSON data.")
	}

	return bot_token, chat_id, nil
}

func telegram_bot_send_text(player_name string, info_file_path string) {
	bot_token, chat_id, err := telegramKeys(info_file_path)
	if err != nil {
		log.Fatal(err)
	}
	chat_id_int, err := strconv.ParseInt(chat_id, 10, 64)

	// Create a new bot instance
	bot, err := tgbotapi.NewBotAPI(bot_token)
	if err != nil {
		log.Fatal(err)
	}

	// Set up the message configuration
	message := player_name + " is on strike!"
	msg := tgbotapi.NewMessage(chat_id_int, message)

	// Send the message
	_, err = bot.Send(msg)
	if err != nil {
		log.Fatal(err)
	}
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

	// team_id := 145221
	team_id_str := os.Getenv("TEAM_ID")
	team_id, err := strconv.Atoi(team_id_str)
	if err != nil {
		team_id = 145221
	}

	// player_id := 84717
	player_id_str := os.Getenv("PLAYER_ID")
	player_id, err := strconv.Atoi(player_id_str)
	if err != nil {
		player_id = 84717
	}

	// player_name := "Shikhar Dhawan"
	player_name := os.Getenv("PLAYER_NAME")
	team_name := os.Getenv("TEAM_NAME")

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
					print_status(result, player_name, team_name, info_file_path)
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
