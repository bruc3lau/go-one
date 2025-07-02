# 设置应用名称和版本
APP_NAME := go-one
VERSION := 1.0.0

# 明确指定默认目标
.DEFAULT_GOAL := build

# 获取当前系统信息
UNAME_S := $(shell uname -s)
UNAME_M := $(shell uname -m)

# 默认目标平台和架构
ifeq ($(UNAME_S),Darwin)
    DEFAULT_OS := darwin
    ifeq ($(UNAME_M),arm64)
        DEFAULT_ARCH := arm64
    else
        DEFAULT_ARCH := amd64
    endif
else
    DEFAULT_OS := windows
    DEFAULT_ARCH := amd64
endif

# 构建目录
BUILD_DIR := build
RELEASE_DIR := $(BUILD_DIR)/release

# 定义所有可能的构建目标
DARWIN_TARGETS := $(BUILD_DIR)/$(APP_NAME)-darwin-amd64 $(BUILD_DIR)/$(APP_NAME)-darwin-arm64
WINDOWS_TARGETS := $(BUILD_DIR)/$(APP_NAME)-windows-amd64.exe $(BUILD_DIR)/$(APP_NAME)-windows-386.exe

# 清理构建目录
.PHONY: clean
clean:
	rm -rf $(BUILD_DIR)

# 创建必要的目录
.PHONY: init
init:
	mkdir -p $(BUILD_DIR) $(RELEASE_DIR)

# 默认目标：构建当前平台版本
.PHONY: build
build: init $(BUILD_DIR)/$(APP_NAME)-$(DEFAULT_OS)-$(DEFAULT_ARCH)$(if $(filter windows,$(DEFAULT_OS)),.exe,)

# 构建所有版本
.PHONY: build-all
build-all: $(DARWIN_TARGETS) $(WINDOWS_TARGETS)
ifeq ($(UNAME_S),Darwin)
	@echo "创建 macOS 通用二进制文件..."
	lipo -create -output $(BUILD_DIR)/$(APP_NAME)-darwin-universal \
		$(BUILD_DIR)/$(APP_NAME)-darwin-amd64 \
		$(BUILD_DIR)/$(APP_NAME)-darwin-arm64
endif

# 各个平台的构建规则
$(BUILD_DIR)/$(APP_NAME)-darwin-amd64: init
	@echo "构建 macOS AMD64 版本..."
	GOOS=darwin GOARCH=amd64 go build -o $@ ./cmd/main.go

$(BUILD_DIR)/$(APP_NAME)-darwin-arm64: init
	@echo "构建 macOS ARM64 版本..."
	GOOS=darwin GOARCH=arm64 go build -o $@ ./cmd/main.go

$(BUILD_DIR)/$(APP_NAME)-windows-amd64.exe: init
	@echo "构建 Windows AMD64 版本..."
	GOOS=windows GOARCH=amd64 go build -o $@ ./cmd/main.go

$(BUILD_DIR)/$(APP_NAME)-windows-386.exe: init
	@echo "构建 Windows 386 版本..."
	GOOS=windows GOARCH=386 go build -o $@ ./cmd/main.go

# 打包规则
.PHONY: package-darwin package-windows package-all

package-darwin: $(DARWIN_TARGETS)
	@echo "打包 macOS 版本..."
	tar -czf $(RELEASE_DIR)/$(APP_NAME)-$(VERSION)-darwin-amd64.tar.gz -C $(BUILD_DIR) $(APP_NAME)-darwin-amd64
	tar -czf $(RELEASE_DIR)/$(APP_NAME)-$(VERSION)-darwin-arm64.tar.gz -C $(BUILD_DIR) $(APP_NAME)-darwin-arm64
ifeq ($(UNAME_S),Darwin)
	tar -czf $(RELEASE_DIR)/$(APP_NAME)-$(VERSION)-darwin-universal.tar.gz -C $(BUILD_DIR) $(APP_NAME)-darwin-universal
endif

package-windows: $(WINDOWS_TARGETS)
	@echo "打包 Windows 版本..."
	cd $(BUILD_DIR) && zip -r release/$(APP_NAME)-$(VERSION)-windows-amd64.zip $(APP_NAME)-windows-amd64.exe
	cd $(BUILD_DIR) && zip -r release/$(APP_NAME)-$(VERSION)-windows-386.zip $(APP_NAME)-windows-386.exe

package-all: package-darwin package-windows

# 帮助信息
.PHONY: help
help:
	@echo "可用的 make 目标："
	@echo "  make              - 构建当前系统的版本"
	@echo "  make -j [N]      - 并行构建，N 是并行任务数"
	@echo "  make build-all   - 构建所有平台的版本"
	@echo "  make package-all - 构建并打包所有版本"
	@echo "  make clean       - 清理构建目录"
	@echo "  make help        - 显示此帮助信息"
