BUILD_DIR=./cmd
BUILD_FILE=risponse.go
BINARY_NAME=risponse

build:
	go build -o ./$(BINARY_NAME) $(BUILD_DIR)/$(BUILD_FILE)

clean:
	rm ./$(BINARY_NAME)
