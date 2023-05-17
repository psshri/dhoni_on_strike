# This script fetches the live score every 5 mins and checks whether
# ${player_id} is out on crease to bat

# libraries #########################################################

import requests
import json
from collections import namedtuple

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

# get live score
def get_live_score(match_info, url, api_key_path, X_RapidAPI_Host, live_score_data_path):
    match_id = match_info.id
    url = url + str(match_id)
    X_RapidAPI_Key = readAPIkey(api_key_path)

    headers = {
        "X-RapidAPI-Key": X_RapidAPI_Key,
        "X-RapidAPI-Host": X_RapidAPI_Host
    }

    response = requests.get(url, headers=headers)
    json_data = response.json()
    with open(live_score_data_path, "w") as file:
        json.dump(json_data, file, indent=4)

    return 

# check if ${player_id} is on crease or not
def is_batting(live_score_data_path, player_id):
    counter = 0

    with open(live_score_data_path, "r") as file:
        live_score = json.load(file)

    list_batsmen = []
    try:
        scorecard = live_score['results']['live_details']['scorecard']
    except:
        scorecard = None

    if scorecard != None:
        for item1 in scorecard:
            for item2 in item1['batting']:
                list_batsmen.append(item2['player_id'])

    if player_id in list_batsmen:
        counter = 1

    return counter