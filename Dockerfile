# use the official go image as the base image
FROM golang:alpine

# set the working directory inside the container
WORKDIR /go/src

# copy the source code into the container
COPY src/config/info.json config/
COPY src/checkLiveScore/checkLiveScore.go checkLiveScore/
COPY src/checkSchedule/checkSchedule.go checkSchedule/
COPY src/main.go .

# initialize the go module
RUN go mod init src

# install the necessary dependencies
RUN go get -d -v github.com/go-telegram-bot-api/telegram-bot-api
RUN go get github.com/aws/aws-sdk-go/aws
RUN go get github.com/aws/aws-sdk-go/aws/session
RUN go get github.com/aws/aws-sdk-go/service/secretsmanager

# build the go application
RUN go build -o main main.go

# set the entry point for the container
CMD ["./main"]