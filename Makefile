CURRENT_DIR = $(shell pwd)

clean:
	rm -rf client/client

build: clean
	go build -o client/client $(CURRENT_DIR)/client/client.go

run: build   
	./client/client
