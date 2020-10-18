
# postgres connect
psql postgres://nogobk:N0g0bAck@localhost:5432/nogobk?sslmode=disable

# login request
curl -X POST -H "Content-Type: application/json" -d '{"email":"nadav@dav.com", "password":"N0g0bAck"}' http://localhost:8081/login

# signup request
curl -X POST -H "Content-Type: application/json" --data '{"name":"Nadav Ben Mazia","email":"nadav@dav.com", "password":"N0g0bAck", "confirm":"N0g0bAck"}' http://localhost:8081/signup