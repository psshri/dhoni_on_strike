# libraries #########################################################

import datetime
import schedule
import time
import requests
import json
import os

from checkSchedule.checkSchedule import get_schedule, evaluate_schedule
from checkLiveScore.checkLiveScore import get_live_score, is_batting

# helper functions ##################################################

def readAPIkey(file_path):
    try:
        with open(file_path, 'r') as file:
            data = json.load(file)
        return data['rapidAPI_api_key']
    except FileNotFoundError:
        print(f"Error: File '{file_path}' not found.")
    except IOError:
        print(f"Error: Unable to read file '{file_path}'.")

def func_today_string():
    today = datetime.date.today()
    today_string = today.strftime("%Y-%m-%d")
    return today_string

def telegram_bot_send_text(player_name, info_file_path):
    with open(info_file_path, 'r') as file:
        data = json.load(file)

    bot_token = data['telegram_bot_token']
    chat_id = data['telegram_chat_id']

    message = player_name + ' is on strike\\!'
    
    url = 'https://api.telegram.org/bot' + bot_token + '/sendMessage?chat_id=' \
          + chat_id + '&parse_mode=MarkdownV2&text=' + message

    requests.get(url)

# constants #########################################################

fixtures_url = "https://cricket-live-data.p.rapidapi.com/fixtures-by-date/"
live_score_url = "https://cricket-live-data.p.rapidapi.com/match/"
X_RapidAPI_Host = "cricket-live-data.p.rapidapi.com"

info_file_path = 'config/info.json'

# today_string = func_today_string()
today_string = "2023-05-19"
ipl_series_id = 1430
# team_name = "Punjab Kings" 
# team_id = 145221
# player_name = "Shikhar Dhawan"
# player_id = 84717
team_id = os.environ.get('TEAM_ID')
team_id = int(team_id)
team_name = os.environ.get('TEAM_NAME')
player_id = os.environ.get('PLAYER_ID')
player_id = int(player_id)
player_name = os.environ.get('PLAYER_NAME')
interval = os.environ.get('INTERVAL')
interval = int(interval)
fixture_data_path = "checkSchedule/fixtures.json"
live_score_data_path = "checkLiveScore/live_score.json"

# functions ########################################################

# counter = None
counter = {'value': None}

def update_counter():
    global counter
    counter['value'] = is_batting(live_score_data_path, player_id)

def print_status(counter):
    print("\n")
    if counter['value'] == 1:
        print(player_name + " is on strike!")
        telegram_bot_send_text(player_name, info_file_path)

    else:
        print(team_name +"'" + " match is today, " + player_name + " is yet to bat!")
    print("\n")

# function calls ###################################################

get_schedule(fixtures_url, X_RapidAPI_Host, info_file_path, today_string, fixture_data_path)

match_today, match_time_730, match_time_330, match_info = evaluate_schedule(fixture_data_path, ipl_series_id, team_id, today_string)

if match_today == 1:
    # get_live_score(match_info, live_score_url, info_file_path, X_RapidAPI_Host, live_score_data_path)
    # counter = is_batting(live_score_data_path, player_id)

    schedule.every(interval).seconds.do(get_live_score, match_info=match_info,
                                  url=live_score_url,
                                  info_file_path=info_file_path,
                                  X_RapidAPI_Host=X_RapidAPI_Host,
                                  live_score_data_path=live_score_data_path)

    schedule.every(interval).seconds.do(update_counter)

    schedule.every(interval).seconds.do(print_status, counter=counter)

    while True:
        schedule.run_pending()
        time.sleep(1)

else:
    print("No" + team_name + "' match today!")