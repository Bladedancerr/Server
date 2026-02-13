#variables
BIN_DIR=bin
BINARY_NAME=server.exe
GO_FILES=main.go

#ensure that commands are treated as commands
.PHONY: run clean build

#default "make"
all: run

#"make build"
build:
	@echo "Building $(BINARY_NAME) into $(BIN_DIR)..."
	@mkdir -p $(BIN_DIR)
	go build -o $(BIN_DIR)/$(BINARY_NAME) $(GO_FILES)

#"make run"
run: build
	@echo "Running $(BINARY_NAME)..."
	$(BIN_DIR)/$(BINARY_NAME)

#"make clean"
clean:
	@echo "Cleaning up..."
	rm -rf $(BIN_DIR)/$(BINARY_NAME)