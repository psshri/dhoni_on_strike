import schedule
import time

def func1():
    print("Function 1")

def func2():
    print("Function 2")

# Schedule func1 and func2 to run every 10 seconds
schedule.every(2).seconds.do(func1)
schedule.every(2).seconds.do(func2)
# schedule.every(5).minutes.do(func1)

# Run the scheduled tasks continuously
while True:
    schedule.run_pending()
    # time.sleep()
