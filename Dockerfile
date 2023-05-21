# use an official python runtime as the base image, python:3.9-alpine is lightweight will not increase the size of the image a lot
FROM python:3.9-alpine

# set the working directory inside the container
WORKDIR /app

# copy the requirements file to the container
COPY requirements.txt .

# install the python dependencies
RUN pip install --no-cache-dir -r requirements.txt

# copy the relevant files from src to the container
COPY src/config/info.json config/
COPY src/checkLiveScore/checkLiveScore.py checkLiveScore/
COPY src/checkSchedule/checkSchedule.py checkSchedule/
COPY src/main.py .

# set the default values for environment variables
ENV TEAM_ID=145221
ENV PLAYER_ID=84717

# specify the command to run your application
CMD ["python", "main.py"]