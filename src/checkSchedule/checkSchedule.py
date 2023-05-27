# This script fetches the match schedule for today and finds out whether there
# will be any match of ${team_name}, if there is a match then at what time

# libraries #########################################################

import requests
import json
from collections import namedtuple

#####################################################################


# get today's match schedule
def get_schedule(url, X_RapidAPI_Host, rapidAPI_api_key, today_string, fixture_data_path):
    
    url = url + today_string
    headers = {
	    "X-RapidAPI-Key": rapidAPI_api_key,
	    "X-RapidAPI-Host": X_RapidAPI_Host
    }

    response = requests.get(url, headers=headers)
    json_data = response.json()
    with open(fixture_data_path, "w") as file:
        json.dump(json_data, file, indent=4)

    print("Today's match schedule fetched successfully")

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
        if item['series_id'] == ipl_series_id and (item['home']['id'] == team_id or item['away']['id'] == team_id):
            match_today = 1
            match_date = today_string
            match_time = todays_matches['results'][5]['date'][11:19]
            match_info = matchInfo(item['id'], item['home']['name'], item['away']['name'], match_date, match_time)
            if match_time == '14:00:00':
                match_time_730 = 1
            else:
                match_time_330 = 1

    print("Today's schedule evaluated successfully!")

    return match_today, match_time_730, match_time_330, match_info