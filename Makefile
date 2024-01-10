TARGET_DIR=bin
GOBUILD=go build

BINARY_NAME=go_htmx_templ

local:
	go run main.go

build:
	$(GOBUILD) -o $(TARGET_DIR)/$(BINARY_NAME) 

run:
	./$(TARGET_DIR)/$(BINARY_NAME)

clean:
	go clean
	rm -f $(TARGET_DIR)/$(BINARY_NAME)

release: build run