#!/bin/sh
source env.sh

docker-compose down
docker-compose up -d database

sleep 10

go run api/run/main.go