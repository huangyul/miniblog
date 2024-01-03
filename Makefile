OUTPUT := _output

.PHONY: all
all: tidy build


.PHONY: build
build:
	@go build -v -o _output/miniblog.exe cmd/miniblog/main.go

.PHONY: tidy
tidy:
	@go mod tidy

.PHONY: clean
clean:
	@rm -rf $(OUTPUT)