package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"src/checkLiveScore"
	"src/checkSchedule"
	"strconv"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/secretsmanager"
)

func print_status(counter int, player_name string, team_name string, telegram_bot_token string, telegram_chat_id string) {
	fmt.Println()
	if counter == 1 {
		fmt.Println(player_name, "is on strike!")
		message := player_name + " is on strike!"
		telegram_bot_send_text(message, telegram_bot_token, telegram_chat_id)
	} else {
		fmt.Println(team_name+"'", "match is today,", player_name, "is yet to bat!")
		message := team_name + "'" + " match is today, " + player_name + " is yet to bat!"
		telegram_bot_send_text(message, telegram_bot_token, telegram_chat_id)
	}
	fmt.Println()
}

func get_secret(secret_name string) (string, string, string) {
	region := "us-east-1"
	sess := session.Must(session.NewSession())
	svc := secretsmanager.New(sess, aws.NewConfig().WithRegion(region))

	result, err := svc.GetSecretValue(&secretsmanager.GetSecretValueInput{SecretId: &secret_name})
	if err != nil {
		log.Fatal(err.Error())
	}
	secrets := *result.SecretString
	// fmt.Println(reflect.TypeOf(secrets))

	var secret map[string]interface{}
	err = json.Unmarshal([]byte(secrets), &secret)
	if err != nil {
		fmt.Println("Failed to decode secret JSON:", err)
	}

	telegram_chat_id, ok := secret["telegram_chat_id"].(string)
	if !ok {
		fmt.Println("Failed to get telegram_chat_id from secret")
	}

	telegram_bot_token, ok := secret["telegram_bot_token"].(string)
	if !ok {
		fmt.Println("Failed to get telegram_bot_token from secret")
	}

	rapidAPI_api_key, ok := secret["rapidAPI_api_key"].(string)
	if !ok {
		fmt.Println("Failed to get rapidAPI_api_key from secret")
	}

	// return telegram_bot_token, telegram_chat_id, rapidAPI_api_key
	return telegram_bot_token, telegram_chat_id, rapidAPI_api_key
}

// func telegramKeys(info_file_path string) (string, string, error) {
// 	file, err := ioutil.ReadFile(info_file_path)
// 	if err != nil {
// 		if e, ok := err.(*os.PathError); ok && e.Err == syscall.ENOENT {
// 			return "", "", fmt.Errorf("Error: File '%s' not found.", info_file_path)
// 		}
// 		return "", "", fmt.Errorf("Error: Unable to read file '%s'.", info_file_path)
// 	}

// 	var data map[string]interface{}
// 	err = json.Unmarshal(file, &data)
// 	if err != nil {
// 		return "", "", err
// 	}

// 	bot_token, ok := data["telegram_bot_token"].(string)
// 	if !ok {
// 		return "", "", fmt.Errorf("Error: 'telegram_bot_token' not found in JSON data.")
// 	}

// 	chat_id, ok := data["telegram_chat_id"].(string)
// 	if !ok {
// 		return "", "", fmt.Errorf("Error: 'telegram_chat_id' not found in JSON data.")
// 	}

// 	return bot_token, chat_id, nil
// }

func telegram_bot_send_text(message string, telegram_bot_token string, telegram_chat_id string) {
	// bot_token, chat_id, err := telegramKeys(info_file_path)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	telegram_chat_id_int, err := strconv.ParseInt(telegram_chat_id, 10, 64)
	// telegram_chat_id_int, err := strconv.Atoi(telegram_chat_id)
	// fmt.Println(telegram_chat_id_int)
	// fmt.Println(telegram_chat_id)

	// Create a new bot instance
	bot, err := tgbotapi.NewBotAPI(telegram_bot_token)
	if err != nil {
		log.Fatal(err)
	}

	// Set up the message configuration
	// message := player_name + " is on strike!"
	msg := tgbotapi.NewMessage(telegram_chat_id_int, message)

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
	// info_file_path := "config/info.json"
	today_string := time.Now().Format("2006-01-02")
	// today_string := "2023-05-19"
	fixtures_data_path := "/tmp/fixtures.json"
	live_score_data_path := "/tmp/live_score.json"
	ipl_series_id := 1430

	// team_id := 101742
	team_id_str := os.Getenv("TEAM_ID")
	team_id, err := strconv.Atoi(team_id_str)
	if err != nil {
		team_id = 101742
	}

	// player_id := 84255
	player_id_str := os.Getenv("PLAYER_ID")
	player_id, err := strconv.Atoi(player_id_str)
	if err != nil {
		player_id = 84255
	}

	// player_name := "MS Dhoni"
	// team_name := "Chennai Super Kings"
	player_name := os.Getenv("PLAYER_NAME")
	team_name := os.Getenv("TEAM_NAME")

	// interval := 2
	interval_str := os.Getenv("INTERVAL")
	interval, err := strconv.Atoi(interval_str)
	if err != nil {
		interval = 300
	}

	// secret_name := "dhoni_on_strike"
	secret_name := os.Getenv("SECRET_NAME")
	telegram_bot_token, telegram_chat_id, rapidAPI_api_key := get_secret(secret_name)

	checkSchedule.Get_schedule(fixtures_url, xRapidAPIHost, rapidAPI_api_key, today_string, fixtures_data_path)

	match_today, match_time_330, match_time_730, matchInfo, err := checkSchedule.Evaluate_schedule(fixtures_data_path, ipl_series_id, team_id, today_string)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	if match_today == 1 {

		// Create a ticker that ticks every 2 seconds

		fmt.Println()
		fmt.Println(team_name + "' match is today!")
		fmt.Println()

		ticker := time.NewTicker(time.Duration(interval) * time.Second)
		// ticker := time.NewTicker(2 * time.Second)

		// Start a goroutine to execute the functions periodically
		go func() {
			for {
				select {
				case <-ticker.C:
					checkLiveScore.Get_live_score(matchInfo, live_score_url, rapidAPI_api_key, xRapidAPIHost, live_score_data_path)
					result := checkLiveScore.Is_batting(live_score_data_path, player_id)
					print_status(result, player_name, team_name, telegram_bot_token, telegram_chat_id)
				}
			}
		}()

		// Keep the main goroutine running
		select {}

	} else {
		fmt.Println("No " + team_name + "' match today!")
		message := "No " + team_name + "' match today!"
		telegram_bot_send_text(message, telegram_bot_token, telegram_chat_id)
	}

	// fmt.Println("player_name: ", player_name)
	// fmt.Println("team_name: ", team_name)
	// fmt.Println("player_id: ", player_id)
	// fmt.Println("team_id: ", team_id)
	fmt.Println("match_time_330: ", match_time_330)
	fmt.Println("match_time_730: ", match_time_730)
}
