GO_ENVS = GO111MODULE=on CGO_ENABLED=0

default: build
	
build:
	 $(GO_ENVS) go build -o ./bin/duration cmd/duration/duration.go

run: build
	./bin/duration ./script.sh

test: 
	go test -cover ./...

clean:
	rm ./bin/duration
