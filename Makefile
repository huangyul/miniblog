COMMON_SELF_DIR := $(dir $(lastword $(MAKEFILE_LIST)))

ROOT_DIR := $(abspath $(shell cd $(COMMON_SELF_DIR)/ && pwd -P))

OUTPUT_DIR := $(ROOT_DIR)/_OUTPUT

.PHONY: all
all: format build

.PHONY: build
build: tidy # 编译源码，依赖 tidy 目标自动添加/移除依赖包.
	@go build -v -o $(OUTPUT_DIR)/miniblog $(ROOT_DIR)/cmd/miniblog/main.go
.PHONY: format
format: # 格式化 Go 源码.
	@gofmt -s -w ./
.PHONY: add-copyright
add-copyright: # 添加版权头信息.
	@addlicense -v -f $(ROOT_DIR)/scripts/boilerplate.txt $(ROOT_DIR) --skip-dirs=third_party,vendor,$(OUTPUT_DIR)
.PHONY: swagger
swagger: # 启动 swagger 在线文档.
	@swagger serve -F=swagger --no-open --port 65534 $(ROOT_DIR)/api/openapi/openapi.yaml
.PHONY: tidy
tidy: # 自动添加/移除依赖包.
	@go mod tidy
.PHONY: clean
clean: # 清理构建产物、临时文件等.
	@-rm -vrf $(OUTPUT_DIR)