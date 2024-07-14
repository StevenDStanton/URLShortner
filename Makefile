# Define the source and build directories
SRC_DIR=src
BUILD_DIR=build

build:
	@mkdir -p $(BUILD_DIR)
	@GOOS=linux GOARCH=amd64 go build -o $(BUILD_DIR)/bootstrap $(SRC_DIR)/main.go
	@zip -j $(BUILD_DIR)/function.zip $(BUILD_DIR)/bootstrap
	@rm $(BUILD_DIR)/bootstrap

clean:
	@rm -f $(BUILD_DIR)/bootstrap $(BUILD_DIR)/function.zip

.PHONY: build clean deploy
