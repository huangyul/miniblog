COMMON_SELF_DIR := $(dir $(lastword $(MAKEFILE_LIST)))
# 项目根目录
ROOT_DIR := $(abspath $(shell cd $(COMMON_SELF_DIR) && cd && echo %cd%))
# 构建产物、临时文件存放目录
OUTPUT_DIR := $(ROOT_DIR)/_output

ROOT_DIR_WIN := $(subst /,\,$(ROOT_DIR))
OUTPUT_DIR_WIN := $(subst /,\,$(OUTPUT_DIR))

.PHONY: all
all: build

.PHONY: build
build:
	@go build -o $(OUTPUT_DIR_WIN)/miniblog.exe $(ROOT_DIR_WIN)/cmd/miniblog/main.go
