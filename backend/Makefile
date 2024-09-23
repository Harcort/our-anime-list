GOPATH := $(shell go env GOPATH):$(shell pwd)
build:
	go build

run: build
	./backend serve &

request: run
	curl -X POST -H "Content-Type: application/json" -d '{"watchlist_id":1, "user_id": 1}' http://localhost:8080/watchlist
	killall backend