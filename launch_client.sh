#!/bin/bash

while netstat -tln | grep -q "${SERVER_PORT}"; do
  echo "waiting for server to start on ${SERVER_PORT}"
  sleep 1
done

# Start client as soon as server is ready
echo 'client is ready to launch'
cd /go/bin/ || exit
./multigoprograms client