GOPATH := $(shell go env GOPATH):$(shell pwd)
build:
	go build

run: build
	./backend serve &

request: run
	echo "Creating a new user"
	curl -XPOST localhost:8080/api/v1/auth/signup -H "Content-Type: application/json" -d '{"username": "myUsername","email": "user@example.com","password": "mySecurePassword"}'
	echo "Logging in"
	curl -XPOST localhost:8080/api/v1/auth/login -H "Content-Type: application/json" -d '{"username": "myUsername","password": "mySecurePassword"}'
	echo "getting the user"
	curl localhost:8080/api/v1/user/myUsername -H "Content-Type: application/json" -H "Authorization: $(curl -XPOST localhost:8080/api/v1/auth/login -H "Content-Type: application/json" -d '{"username": "myUsername","password": "mySecurePassword"}' | jq -r .data)"

	curl -X POST -H "Content-Type: application/json" -d '{"watchlist_id":1, "user_id": 1}' http://localhost:8080/watchlist
	killall backend