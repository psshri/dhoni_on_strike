# dhoni_on_strike
This application notifies the end user whenever MS Dhoni is out on the pitch to bat!

NOTE
make sure to update the today_string at last

### todos:
- research the capabilities of aws lambda
- research the capabilities of aws sns
- research how to send notifications to telegram
- research on the integration between amazon dynamodb and aws lambda with proper rbac and security
- use any open source tool for code security scanning 
- use the amazon service for code security (from free tier)
- check if you could run the code as an image in aws lamda
- if your code is using any secret then store the secret in some key vault (explore aws secret service and hashi corp vault)
- use terraform to deploy infra in aws
- also use aws SDKs to deploy infra in aws (kiran's requirement)

### todos in application code
- use packages in go to make your appliction scalable and reusable, for example a user should be able to enter the player name(s) of his choice and should be notified for his list of player(s). 
- you should fill the match data also by api endpoints and not manually
- right now i am getting telegram notification continuously, and also even after the batsman is on strike, the function is running continuously, i have to include a logic to stop executing the function once the batsmen is on strike


### features:
- should be scalable, i should be able to add users to that notification list
- notification should be sent via amazon sns and the notification should be sent to telegram
- store the match details in amazon dynamodb and the code running on aws lambda should be able to fetch those details, make sure security and RBAC is properly delegated
- code security
- the application code uses concurrency in go
- application code has multiple packages, so as to make the code reusable, 
- user notification logic will be mentioned in the application code but all the tasks (storing user data, etc) related to notification will be handled by sns
- user should also be notified via a phone call



### imp links
- https://github.com/mskian/cricket-cli


### aws lambda
- mera ek lambda function roz run hoga aur updated fixtures ki list rds me update karega


### dockerfile

go get -u github.com/go-sql-driver/mysql
the above line of code is required to setup the connection between mysql and golang. see how you can run the above line of code within the container image in dockerfile


### flow 
- parse the fixtures and store it in mysql db/table
- hit the score api every few mins and get the current score and evaluate whether a particular player is on strike or not


### code scheduling in aws
- a code should be executed each day that fetches the fixtures info and updates it in rds database
- the main code that runs and fetches live score for comparison should be executed whenever csk's match is played, so the trigger for that code has to be created based on the rds data

### fixtures.json

table columns:
- id
- date
- 

### optimizations in code
- add a line to output when the api limit is reached, usme try except wala block bhi dalo
- secret keys ko as an environment variables pass karo aur aws lambda me env variable create karo unke liye and encrypt bhi karo


### steps followed to  containerize this python app
- create the dockerfile
- create the requirements.txt file. this file lists all the python packages required by your application
- copy only the relevant files from src to your container
- docker build -t dhonionstrike:python .  (build the image)
- docker run -it dhonionstrike:python (run the container)
- the resultant image was 900MB large, so i used alpine version of base image as they are lightweight, then the size came down to 57MB
- you can use tools like docker-slim to find out how to optimize your container
- use env variables to provide values of variable like player_id, team_id during the container run command, instead of hardcoding them; edit the main.py and add ENV to dockerfile
- after editing the files, rebuild the docker image using the above command and now use the below command to run the container
- docker run -it -e PLAYER_ID=84717 -e TEAM_ID=145221 dhonionstrike:python
- alternatively, you can create a file for all the env variables and pass this file during docker run command
- docker run -it --env-file env.list dhonionstrike:python

run the following commands to tag the image and push it to dockerhub
docker tag dhonionstrike:python psshri/dhoni_on_strike:python-v1.0
docker login
docker push psshri/dhoni_on_strike:python-v1.0
docker pull psshri/dhoni_on_strike:python-v1.0

docker build -t dhoni_on_strike:python-v1.0 .
docker run -it --env-file env.list dhoni_on_strike:python-v1.0
docker tag dhoni_on_strike:python-v1.0 psshri/dhoni_on_strike:python-v1.0
docker push psshri/dhoni_on_strike:python-v1.0


### running python code in aws lambda
- adding a layer for request module in lambda using this link (https://aws.amazon.com/blogs/compute/upcoming-changes-to-the-python-sdk-in-aws-lambda/)
- also, as per above article, request module is to be imported via boto library


### pushing the python container to amazon ecr
- create a public repository
- install aws cli in ubuntu, follow this link (https://docs.aws.amazon.com/cli/latest/userguide/getting-started-install.html)
- to authorize, you will need Access Key ID and Secret Access Key
- create a user (from iam and not identity center) just for cli access and get access key id and secret access key
- tag the image using the below command
<!-- docker tag dhoni_on_strike:python-v1.0 public.ecr.aws/x6i9k3w4/dhoni_on_strike:python-v1.0 -->
docker tag dhoni_on_strike:python-v1.0 014935736506.dkr.ecr.us-east-1.amazonaws.com/dhoni_on_strike:python-v1.0
- retrieve an authentication token and authenticate your docker client to your registry, run the following command
<!-- aws ecr-public get-login-password --region us-east-1 | docker login --username AWS --password-stdin public.ecr.aws/x6i9k3w4 -->
aws ecr get-login-password --region us-east-1 | docker login --username AWS --password-stdin 014935736506.dkr.ecr.us-east-1.amazonaws.com
- push the image using the following command
docker push 014935736506.dkr.ecr.us-east-1.amazonaws.com/dhoni_on_strike:python-v1.0

### running python container in aws lambda
- write functionality is not available during aws lambda invocation, you have to use /tmp directory if you want to store some data, so we had to modify the fixtures_data_path and live score data path to /tmp, otherwise you can store the file in s3,
- modify the code and create container 
- the lambda function stops after 3 seconds of running, to resolve this do the following
lambda fn > configuration > general configuration > edit timeout
- aws lambda can run for a maximum of 15 mins, so use step functions to invoke the lambda function repeatedly after 15 mins, now there are multiple areas of optimization that are possible now
- to provide the environment variables go to configuration > environment variables


### create a secret in aws secrets manager
- write an aws sdk script to do this thing
- create a secret in secrets manager
- modify the python code to read secrets from secrets manager secret_name
- create img, push, import it into lambda
- grant the lambda with permission to read the secret from secrets manager
