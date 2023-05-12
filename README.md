# dhoni_on_strike
This application notifies the end user whenever MS Dhoni is out on the pitch to bat!


todos:
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

todos in application code
- use packages in go to make your appliction scalable and reusable, for example a user should be able to enter the player name(s) of his choice and should be notified for his list of player(s). 
- you should fill the match data also by api endpoints and not manually


features:
- should be scalable, i should be able to add users to that notification list
- notification should be sent via amazon sns and the notification should be sent to telegram
- store the match details in amazon dynamodb and the code running on aws lambda should be able to fetch those details, make sure security and RBAC is properly delegated
- code security
- the application code uses concurrency in go
- application code has multiple packages, so as to make the code reusable, 
- user notification logic will be mentioned in the application code but all the tasks (storing user data, etc) related to notification will be handled by sns
- user should also be notified via a phone call
