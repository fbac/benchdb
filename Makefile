BINARY = benchdb
BIN_FOLDER = bin
TEST_FOLDER = test

.PHONY: build clean test docs

#######################
# benchdb cmd targets #
#######################

build:
	go build -o ${BIN_FOLDER}

clean:
	rm -f ${BIN_FOLDER}/${BINARY}

test:
	go test -v

coverage:
	go test -cover -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out
	
doc:
	godoc -http localhost:8888

#######################
# timescaledb targets #
#######################