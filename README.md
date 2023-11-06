## Overview

This repository is my solution for the problem 1.1 Write Hello world.

### Task:  

Write Go program what should write to stdout "Hello World" on startup. It should finish in 10 seconds with the message "Goodbye world".

If the user interrupts the execution earlier, then you should print the message "Stopped by the user after x seconds"

### Solution:

I used channels and select statement to implement the solution.
First channel is the timer channel that sends information once the time is up. In this case select statement prints "Goodbye world" and finishes the app.
Second channel is interruption channel that sends information once it's getting interruption signal(or SIGTERM and SIGINT).

## How to run
Clone the repository:

    git clone https://git.foxminded.ua/foxstudent106264/task-1.1.git

Build and application:

    go build 
    
Run the .exe file:

    ./task-1.1.exe