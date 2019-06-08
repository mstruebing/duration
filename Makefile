build:
	GO111MODULE=on CGO_ENABLED=0 go build -o ./bin/duration cmd/duration/duration.go

run: build
	./bin/duration ./script.sh
