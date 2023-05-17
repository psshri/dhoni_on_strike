# libraries #########################################################

import datetime
import schedule
import time

from checkSchedule import get_schedule, evaluate_schedule
from checkLiveScore import get_live_score
from checkLiveScore import is_batting

# functions #########################################################

def func_today_string():
    today = datetime.date.today()
    today_string = today.strftime("%Y-%m-%d")
    return today_string

# constants #########################################################

fixtures_url = "https://cricket-live-data.p.rapidapi.com/fixtures-by-date/"
live_score_url = "https://cricket-live-data.p.rapidapi.com/match/"
X_RapidAPI_Host = "cricket-live-data.p.rapidapi.com"

api_key_path = 'config/apiKey.txt'

today_string = func_today_string()
ipl_series_id = 1430
team_name = "Delhi Capitals" 
team_id = 120252
player_name = "Prithvi Shaw"
player_id = 3210516
fixture_data_path = "fixtures.json"
live_score_data_path = "live_score.json"

# functions ########################################################

# counter = None
counter = {'value': None}

def update_counter():
    global counter
    counter['value'] = is_batting(live_score_data_path, player_id)

def print_status(counter):
    print("\n")
    if counter['value'] == 1:
        print(player_name, "is on strike!")
    else:
        print(team_name +"'" ,"match is today,", player_name, "is yet to bat!")
    print("\n")

# function calls ###################################################

get_schedule(fixtures_url, X_RapidAPI_Host, api_key_path, today_string, fixture_data_path)

match_today, match_time_730, match_time_330, match_info = evaluate_schedule(fixture_data_path, ipl_series_id, team_id, today_string)

if match_today == 1:
    # get_live_score(match_info, live_score_url, api_key_path, X_RapidAPI_Host, live_score_data_path)
    # counter = is_batting(live_score_data_path, player_id)

    schedule.every(2).seconds.do(get_live_score, match_info=match_info,
                                  url=live_score_url,
                                  api_key_path=api_key_path,
                                  X_RapidAPI_Host=X_RapidAPI_Host,
                                  live_score_data_path=live_score_data_path)

    schedule.every(2).seconds.do(update_counter)

    schedule.every(2).seconds.do(print_status, counter=counter)

    while True:
        schedule.run_pending()
        time.sleep(1)
        # print(counter)


    # if counter == 1:
    #     print(player_name, "is on strike!")
    # else:
    #     print(team_name +"'" ,"match is today,", player_name, "is yet to bat!")

else:
    print("No", team_name ,"match today!")