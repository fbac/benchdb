BINARY = benchdb
BIN_FOLDER = bin
TEST_FOLDER = test

.PHONY: build clean test docs

build:
	go build -o ${BIN_FOLDER}

clean:
	rm -f ${BIN_FOLDER}/${BINARY}

test:
	go test -v
	go test -coverprofile coverage.out
	
docs:
	godoc -http localhost:8888