#!/bin/bash

# The exit status for grep
# 0 if a line is selected
# 1 if no lines were selected
# 2 if an error occurred
while netstat -tln | grep -q "${SERVER_PORT}"; do
  echo "waiting for server to start on ${SERVER_PORT}"
  sleep 1
done

# Start client as soon as server is ready
echo 'client is ready to launch'

# If directory does not exist, simply exist the script
cd /go/bin/ || exit

./multigoprograms client