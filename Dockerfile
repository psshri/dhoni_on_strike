FROM public.ecr.aws/lambda/python:3.10

# set the working directory inside the container
# WORKDIR /app

# copy the requirements file to the container
COPY requirements.txt .

# install the python dependencies
RUN pip3 install -r requirements.txt --target "${LAMBDA_TASK_ROOT}"

# copy the relevant files from src to the container
# COPY src/config/info.json config/
COPY src/checkLiveScore/checkLiveScore.py ${LAMBDA_TASK_ROOT}/checkLiveScore/
COPY src/checkSchedule/checkSchedule.py ${LAMBDA_TASK_ROOT}/checkSchedule/
COPY src/main.py ${LAMBDA_TASK_ROOT}

# set the default values for environment variables
ENV PLAYER_ID=84255
ENV TEAM_ID=101742
ENV PLAYER_NAME="MS Dhoni"
ENV TEAM_NAME="Chennai Super Kings"
ENV INTERVAL=2
ENV SECRET_NAME="dhoni_on_strike"

# specify the command to run your application
CMD ["main.handler"]