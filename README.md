# dhoni_on_strike
This application notifies the end user whenever MS Dhoni is out on the pitch to bat!


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
go get github.com/go-telegram-bot-api/telegram-bot-api
the above line of code is required to setup the connection between mysql and golang. see how you can run the above line of code within the container image in dockerfile
- when you will create an image, see how to handle the secrets, find out the best ways to handle secrets for an application when you want to containerize it
- also find out how you can build your container so that it is not tied to any team_name or player_name, the user should be able to pass on those values



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



### initialize a go project using
go mod init <module-name>
go mod init src

go get github.com/go-telegram-bot-api/telegram-bot-api

### aws 

- free tier of aws lambda; understand it very well
- can i send phone notification through aws sns
- free tier of aws sns
- can i schedule my lambda function to execute on a particular time in a day and can the next trigger be defined by the code itself?
- see how to store secrets in aws key vault and how would aws lambda will access those secrets securely

### things to do in aws
- setup everything manually first
- then setup everything using github actions
- then setup everything using aws sdk
- containerize your code and then run it in aws lambda


### how to containerize the golang application
- create the docker file
- run the following command to build the image
docker build -t dhonionstrike:golang .
- run the following command to run the container
docker run -it dhonionstrike:golang
- create env.list file to store all the hardcoded values
- execute the following command to run the container
docker run -it --env-file env.list dhonionstrike:golang