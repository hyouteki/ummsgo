#!/bin/bash

go run server.go &
SERVER_PID=$!

go run client.go

sleep 5
sudo kill -9 $SERVER_PID
sudo wait $SERVER_PID 2>/dev/null
