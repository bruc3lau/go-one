#!/bin/bash

# 设置应用名称
APP_NAME="go-one"

# 创建 build 目录
mkdir -p build

echo "开始构建多架构二进制文件..."

# 构建 AMD64 (x86_64) 版本
echo "构建 AMD64 版本..."
GOOS=darwin GOARCH=amd64 go build -o build/$APP_NAME-amd64 ./cmd/main.go

# 构建 ARM64 版本
echo "构建 ARM64 版本..."
GOOS=darwin GOARCH=arm64 go build -o build/$APP_NAME-arm64 ./cmd/main.go

# 合并为通用二进制文件（仅在 macOS 上可用）
echo "合并为通用二进制文件..."
lipo -create -output build/$APP_NAME-universal \
    build/$APP_NAME-amd64 \
    build/$APP_NAME-arm64

echo "构建完成！"
echo "生成的文件:"
echo "- build/$APP_NAME-amd64 (仅 Intel)"
echo "- build/$APP_NAME-arm64 (仅 Apple Silicon)"
echo "- build/$APP_NAME-universal (通用二进制)"

# 显示文件信息
echo -e "\n文件信息:"
file build/$APP_NAME-universal
