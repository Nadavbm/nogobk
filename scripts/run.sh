#!/bin/sh
export DATABASE_USER=nogobk
export DATABASE_PASSWORD=N0g0bAck
export DATABASE_DB=nogobk
export DATABASE_PORT=5432
export DATABASE_HOST=localhost

docker-compose down
docker-compose up -d database

sleep 8

go run api/run/main.go