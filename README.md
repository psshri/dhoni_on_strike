# Dhoni on Strike!

This application notifies the end user whenever MS Dhoni is out on the pitch to bat!

### Docker

Run the following command to containerize the application

```bash
$ docker build -t dhoni_on_strike:golang-v1.0 .
```

### Areas of optimizations

* Use Amazon SNS for notification instead of Telegram
* Store the src/checkLiveScore/live_score.json and src/checkSchedule/fixtures.json in an S3 bucket instead of /tmp directory within the image
* Create Lambda function, Secret, ECR repo using Terraform/AWS SDK