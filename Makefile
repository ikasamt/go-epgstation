  
BIN_NAME=app
BIN_DIR=bin

ls:
	ls -alh

mkdir:
	mkdir -p $(BIN_DIR)

clean: mkdir
	rm -rf $(BIN_DIR)
	rm -rf pkg/zzz*.go

test:
	go test -v .

deps:
	go mod download

gen:
	zapp-jam pkg

build: clean deps gen
	go build -o $(BIN_DIR)/$(BIN_NAME) *.go

run: ls clean deps gen
	go run *.go