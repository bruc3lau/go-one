#!/bin/bash

# 设置应用名称和版本
APP_NAME="go-one"
VERSION="1.0.0"

# 创建 build 目录
mkdir -p build

echo "开始构建多平台二进制文件..."

# macOS builds
echo "构建 macOS 版本..."
# macOS AMD64 (Intel)
GOOS=darwin GOARCH=amd64 go build -o build/$APP_NAME-darwin-amd64 ./cmd/main.go
# macOS ARM64 (Apple Silicon)
GOOS=darwin GOARCH=arm64 go build -o build/$APP_NAME-darwin-arm64 ./cmd/main.go

# Windows builds
echo "构建 Windows 版本..."
# Windows AMD64 (64-bit)
GOOS=windows GOARCH=amd64 go build -o build/$APP_NAME-windows-amd64.exe ./cmd/main.go
# Windows 386 (32-bit)
GOOS=windows GOARCH=386 go build -o build/$APP_NAME-windows-386.exe ./cmd/main.go

# 如果在 macOS 上，创建通用二进制文件
if [[ "$OSTYPE" == "darwin"* ]]; then
    echo "创建 macOS 通用二进制文件..."
    lipo -create -output build/$APP_NAME-darwin-universal \
        build/$APP_NAME-darwin-amd64 \
        build/$APP_NAME-darwin-arm64
fi

echo "构建完成！"
echo "生成的文件:"
echo "macOS:"
echo "- build/$APP_NAME-darwin-amd64 (Intel Mac)"
echo "- build/$APP_NAME-darwin-arm64 (Apple Silicon)"
if [[ "$OSTYPE" == "darwin"* ]]; then
    echo "- build/$APP_NAME-darwin-universal (Universal Mac)"
fi
echo "Windows:"
echo "- build/$APP_NAME-windows-amd64.exe (64-bit Windows)"
echo "- build/$APP_NAME-windows-386.exe (32-bit Windows)"

# 创建压缩包
echo "创建发布包..."
mkdir -p build/release

# macOS 压缩包
if [[ "$OSTYPE" == "darwin"* ]]; then
    tar -czf build/release/$APP_NAME-$VERSION-darwin-universal.tar.gz -C build $APP_NAME-darwin-universal
fi
tar -czf build/release/$APP_NAME-$VERSION-darwin-amd64.tar.gz -C build $APP_NAME-darwin-amd64
tar -czf build/release/$APP_NAME-$VERSION-darwin-arm64.tar.gz -C build $APP_NAME-darwin-arm64

# Windows ZIP 包
if command -v zip >/dev/null 2>&1; then
    cd build && zip -r release/$APP_NAME-$VERSION-windows-amd64.zip $APP_NAME-windows-amd64.exe
    cd build && zip -r release/$APP_NAME-$VERSION-windows-386.zip $APP_NAME-windows-386.exe
    cd ..
fi

echo "完成！"
echo "发布包位于 build/release/ 目录"
