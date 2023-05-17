# libraries #########################################################

import requests
import datetime
import json
from collections import namedtuple

from checkSchedule import get_schedule, evaluate_schedule

# functions #########################################################

def func_today_string():
    today = datetime.date.today()
    today_string = today.strftime("%Y-%m-%d")
    return today_string

# constants #########################################################

fixtures_url = "https://cricket-live-data.p.rapidapi.com/fixtures-by-date/"
X_RapidAPI_Host = "cricket-live-data.p.rapidapi.com"

api_key_path = 'config/apiKey.txt'

today_string = func_today_string()
ipl_series_id = 1430
team_name = "Delhi Capitals" ## delhi capitals
player_name = "d" ## mitchell marsh
json_data_path = "fixtures.json"



# function calls ###################################################

get_schedule(fixtures_url, X_RapidAPI_Host, api_key_path, today_string, ipl_series_id, team_name, json_data_path)

match_today, match_time_730, match_time_330 = evaluate_schedule(json_data_path, ipl_series_id, team_name, today_string)

print("match_today: ", match_today)
print("match_time_730: ", match_time_730)
print("match_time_330: ", match_time_330)