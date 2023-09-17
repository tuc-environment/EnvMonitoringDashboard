#!/bin/bash

# Start the first process
./EnvMonitoringDashboard.exe 2>&1 &

# Start the second process
uvicorn main:app --port 5000 2>&1 &

# Exit with status of process that exited first
while [[ true ]]; do
    sleep 1
done