# Define the source and build directories
SRC_DIR=src
BUILD_DIR=build

build:
	@mkdir -p $(BUILD_DIR)
	@cd $(SRC_DIR) && GOOS=linux GOARCH=amd64 go build -o ../$(BUILD_DIR)/bootstrap main.go
	@cd $(BUILD_DIR) && zip -j function.zip bootstrap
	@rm $(BUILD_DIR)/bootstrap

clean:
	@rm -f $(BUILD_DIR)/bootstrap $(BUILD_DIR)/function.zip

.PHONY: build clean deploy
