# This script fetches the match schedule for today and finds out whether there
# will be any match of ${team_name}, if there is a match then at what time

# libraries #########################################################

import requests
import json
from collections import namedtuple

#####################################################################

# helper functions ##################################################

def readAPIkey(file_path):
    try:
        with open(file_path, 'r') as file:
            api_key = file.read().strip()
        return api_key
    except FileNotFoundError:
        print(f"Error: File '{file_path}' not found.")
    except IOError:
        print(f"Error: Unable to read file '{file_path}'.")

#####################################################################


# get today's match schedule
def get_schedule(url, X_RapidAPI_Host, api_key_path, today_string, fixture_data_path):

    X_RapidAPI_Key = readAPIkey(api_key_path)
    
    url = url + today_string
    headers = {
	    "X-RapidAPI-Key": X_RapidAPI_Key,
	    "X-RapidAPI-Host": X_RapidAPI_Host
    }

    response = requests.get(url, headers=headers)
    json_data = response.json()
    with open(fixture_data_path, "w") as file:
        json.dump(json_data, file, indent=4)

    return

# find out if there is a match of ${team_name} today
def evaluate_schedule(fixture_data_path, ipl_series_id, team_id, today_string):

    match_today = 0 ## changes to 1 if there is a match today
    match_time_330 = 0
    match_time_730 = 0


    with open(fixture_data_path, "r") as file:
        todays_matches = json.load(file)
    
    matchInfo = namedtuple('matchInfo', ['id', 'team1', 'team2', 'date', 'time'])
    match_info = None

    for item in todays_matches['results']:
        if item['series_id'] == ipl_series_id:
            if item['home']['id'] == team_id or item['away']['id'] == team_id:
                match_today = 1
                match_date = today_string
                match_time = todays_matches['results'][5]['date'][11:19]
                match_info = matchInfo(item['id'], item['home']['name'], item['away']['name'], match_date, match_time)
                if match_time == '14:00:00':
                    match_time_730 = 1
                else:
                    match_time_330 = 1

    return match_today, match_time_730, match_time_330, match_info