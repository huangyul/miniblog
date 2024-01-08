OUTPUT_DIR := _output

.PHONY: all
all: clean tidy build

.PHONY: build
build:
	@go build -o $(OUTPUT_DIR)/miniblog.exe -v cmd/miniblog/main.go

.PHONY: tidy
tidy:
	@go mod tidy

.PHONY: clean
clean:
	@rm -rf $(OUTPUT_DIR)