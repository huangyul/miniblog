COMMON_SELF_DIR := $(dir $(lastword $(MAKEFILE_LIST)))
# 项目根目录
ROOT_DIR := $(abspath $(COMMON_SELF_DIR))
# 构建产物、临时文件存放目录
OUTPUT_DIR := $(ROOT_DIR)/_output

# 打包window可执行文件
.PHONY: build_window
build_window: clean tidy
	@GOOS=windows GOARCH=amd64 go build -v -o $(OUTPUT_DIR)/miniblog.exe $(ROOT_DIR)/miniblog/cmd/main.go

# 打包
.PHONY: build
build: clean tidy
	@go build -v -o $(OUTPUT_DIR)/miniblog $(ROOT_DIR)/miniblog/cmd/main.go

# 安装依赖
.PHONY: tidy
tidy:
	@go mod tidy

# 清空输出目录
.PHONY: clean
clean:
	@rm -vrf $(OUTPUT_DIR)
