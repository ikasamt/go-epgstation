  
BIN_NAME=app
BIN_DIR=bin

ls:
	ls -alh

mkdir:
	mkdir -p $(BIN_DIR)

clean: mkdir
	rm -rf $(BIN_DIR)

test:
	go test -v .

deps:
	go mod download

build: clean deps
	go build -o $(BIN_DIR)/$(BIN_NAME) *.go

run: ls clean deps
	go run *.go