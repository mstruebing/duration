GO_FLAGS = GO111MODULE=on CGO_ENABLED=0

default: build
	
build:
	 $(GO_FLAGS) go build -o ./bin/duration ./...

run: build
	./bin/duration ./test-script/script.sh

test: 
	go test -cover ./...

clean:
	rm ./bin/duration
