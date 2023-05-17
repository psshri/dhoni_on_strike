import schedule
import time

def func1():
    print("Function 1")

def func2():
    print("Function 2")

def func3(count):
    result = count[0] + 1
    count[0] = result
    print(f"Function 3: {result}")

# Schedule func1 and func2 to run every 10 seconds
schedule.every(2).seconds.do(func1)
schedule.every(2).seconds.do(func2)

# Variable to store the count
count = [0]

# Function to update the count using func3
def update_count():
    func3(count)

# Schedule update_count to run every 1 minute
schedule.every(2).seconds.do(update_count)

# Run the scheduled tasks continuously
while True:
    schedule.run_pending()
    time.sleep(1)
