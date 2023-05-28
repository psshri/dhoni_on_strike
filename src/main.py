# libraries #########################################################

import datetime
import schedule
import time
import requests
import json
import os
import boto3
from botocore.exceptions import ClientError

from checkSchedule.checkSchedule import get_schedule, evaluate_schedule
from checkLiveScore.checkLiveScore import get_live_score, is_batting

# global variables ##################################################

counter = {'value': None}
alerted = {'value': None}

# helper functions ##################################################

def func_today_string():
    today = datetime.date.today()
    today_string = today.strftime("%Y-%m-%d")
    return today_string

def telegram_bot_send_text(message, telegram_bot_token, telegram_chat_id):
    
    url = 'https://api.telegram.org/bot' + telegram_bot_token + '/sendMessage?chat_id=' \
          + telegram_chat_id + '&parse_mode=MarkdownV2&text=' + message

    requests.get(url)

def get_secret(secret_name):

    # secret_name = "dhoni_on_strike"
    region_name = "us-east-1"

    # Create a Secrets Manager client
    session = boto3.session.Session()
    client = session.client(
        service_name='secretsmanager',
        region_name=region_name
    )

    try:
        get_secret_value_response = client.get_secret_value(SecretId=secret_name)
    except ClientError as e:
        print(e.response['Error']['Code'])

    # Decrypts secret using the associated KMS key.
    secret = get_secret_value_response['SecretString']
    secret = json.loads(secret)
    telegram_chat_id = secret['telegram_chat_id']
    telegram_bot_token = secret["telegram_bot_token"]
    rapidAPI_api_key = secret["rapidAPI_api_key"]

    return telegram_chat_id, telegram_bot_token, rapidAPI_api_key

# functions ########################################################

def update_counter(live_score_data_path, player_id):
    global counter
    counter['value'] = is_batting(live_score_data_path, player_id)

def print_status(counter, player_name, telegram_bot_token, telegram_chat_id, team_name):
    print("\n")
    if counter['value'] == 1:
        print(player_name + " is on strike!")
        message = player_name + ' is on strike\\!'
        telegram_bot_send_text(message, telegram_bot_token, telegram_chat_id)
        global alerted
        alerted['value'] = 1


    else:
        print(team_name +"'" + " match is today, " + player_name + " is yet to bat!")
        # message = team_name +"'" + " match is today, " + player_name + " is yet to bat\\!"
        # telegram_bot_send_text(message, telegram_bot_token, telegram_chat_id)
    print("\n")

# function calls ###################################################

def main():
    fixtures_url = "https://cricket-live-data.p.rapidapi.com/fixtures-by-date/"
    live_score_url = "https://cricket-live-data.p.rapidapi.com/match/"
    X_RapidAPI_Host = "cricket-live-data.p.rapidapi.com"

    today_string = func_today_string()
    # today_string = "2023-05-19"
    ipl_series_id = 1430

    team_id = os.environ.get('TEAM_ID')
    team_id = int(team_id)
    team_name = os.environ.get('TEAM_NAME')
    player_id = os.environ.get('PLAYER_ID')
    player_id = int(player_id)
    player_name = os.environ.get('PLAYER_NAME')
    interval = os.environ.get('INTERVAL')
    interval = int(interval)

    secret_name = os.environ.get('SECRET_NAME')
    telegram_chat_id, telegram_bot_token, rapidAPI_api_key = get_secret(secret_name)

    fixture_data_path = "/tmp/fixtures.json"
    live_score_data_path = "/tmp/live_score.json"

    get_schedule(fixtures_url, X_RapidAPI_Host, rapidAPI_api_key, today_string, fixture_data_path)
    match_today, match_time_730, match_time_330, match_info = evaluate_schedule(fixture_data_path, ipl_series_id, team_id, today_string)

    if match_today == 1:

        print(team_name + "' match is today!")
        # message = team_name + "' match is today\\!"
        # telegram_bot_send_text(message, telegram_bot_token, telegram_chat_id)

        schedule.every(interval).seconds.do(get_live_score, match_info=match_info,
                                    url=live_score_url,
                                    rapidAPI_api_key=rapidAPI_api_key,
                                    X_RapidAPI_Host=X_RapidAPI_Host,
                                    live_score_data_path=live_score_data_path)

        schedule.every(interval).seconds.do(update_counter, 
                                            live_score_data_path=live_score_data_path, 
                                            player_id=player_id)

        schedule.every(interval).seconds.do(print_status, 
                                            counter=counter, 
                                            player_name=player_name, 
                                            telegram_bot_token=telegram_bot_token, 
                                            telegram_chat_id=telegram_chat_id, 
                                            team_name=team_name)

        while True:
            schedule.run_pending()
            time.sleep(1)
            global alerted
            if alerted['value'] == 1:
                break

    else:
        print("No " + team_name + "' match today!")
        message = 'No ' + team_name + ' match today\\!'
        telegram_bot_send_text(message, telegram_bot_token, telegram_chat_id)

    return 'done!'

def handler(event, context):
    return main()