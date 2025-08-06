Init

```shell
go mod tidy && go mod vendor
```

## TODO 模块测试

核心框架和库:

* github.com/gin-gonic/gin: 一个流行的 Go Web 框架。
* github.com/jinzhu/gorm: 一个强大的 Go ORM 库，用于数据库操作。
* github.com/go-redis/redis: Go 的 Redis 客户端。
* github.com/spf13/viper: 用于处理配置文件的库。
* github.com/rs/zerolog: 一个快速、零分配的 JSON 日志库。
* google.golang.org/grpc: 用于 gRPC 通信。

消息队列和事件处理:

* github.com/Shopify/sarama: Sarama 是一个用于 Apache Kafka 的 Go 客户端库。
* github.com/eclipse/paho.mqtt.golang: 一个 Go 的 MQTT 客户端库。
* github.com/r3labs/sse: 用于 Server-Sent Events (SSE) 的库。

云服务和存储:

* github.com/aliyun/aliyun-oss-go-sdk: 阿里云对象存储 (OSS) 的 Go SDK。
* github.com/baidubce/bce-sdk-go: 百度云的 Go SDK。
* github.com/tencentyun/cos-go-sdk-v5: 腾讯云对象存储 (COS) 的 Go SDK。

工具和实用程序:

* bou.ke/monkey: 一个在 Go 中进行猴子补丁的库，常用于测试。
* github.com/robfig/cron/v3: 一个用于定时任务的库。
* github.com/dgrijalva/jwt-go: 用于处理 JSON Web Tokens (JWT) 的库。
* github.com/stretchr/testify: 一个用于 Go 测试的断言和模拟工具包。
* github.com/urfave/cli: 一个用于构建命令行应用程序的库。

图像和绘图:

* github.com/fogleman/gg: 一个用于在 Go 中进行 2D 图形渲染的库。
* github.com/llgcode/draw2d: 另一个 2D 绘图库。
* github.com/nfnt/resize: 用于图像缩放的库。

这只是 require 部分列出的直接依赖项的一部分。还有很多间接依赖项，它们是这些主要模块所依赖的库。

---

#### GIN

```text
核心包:
github.com/gin-gonic/gin v1.10.1 (主框架)
JSON 相关:
github.com/bytedance/sonic (字节跳动开发的高性能 JSON 库)
github.com/goccy/go-json (JSON 编解码库)
github.com/json-iterator/go (高性能 JSON 迭代器)
验证相关:
github.com/go-playground/validator/v10 (参数验证库)
github.com/go-playground/locales (本地化支持)
github.com/go-playground/universal-translator (翻译器)
其他工具包:
github.com/gin-contrib/sse (Server-Sent Events 支持)
github.com/mattn/go-isatty (终端检测)
github.com/ugorji/go/codec (高性能编解码器)
gopkg.in/yaml.v3 (YAML 支持)
标准库扩展:
golang.org/x/crypto (加密相关)
golang.org/x/net (网络扩展)
golang.org/x/sys (系统调用)
golang.org/x/text (文本处理)

```

```shell
go list -m -json github.com/gin-gonic/gin    # 查看 gin 的详细信息
go mod graph | grep gin                      # 查看与 gin 相关的依赖图

```

---

```shell
# 单次编译使用
go build -mod=vendor cmd/main.go
# 或运行时使用
go run -mod=vendor cmd/main.go
# 查看构建详情
go build -mod=vendor -x cmd/main.go

# 更新或创建 vendor 目录
go mod vendor
# 验证 vendor 目录
go mod verify
```

---
grpc
> 1.创建一个 .proto 文件来定义服务
> pkg/grpc/helloworld/helloworld.proto
```text
syntax = "proto3";                                                                                                                                                                            │
                                                                                                                                                                                              │
option go_package = "go-one/pkg/grpc/helloworld";                                                                                                                                             │
                                                                                                                                                                                              │
package helloworld;                                                                                                                                                                           │
                                                                                                                                                                                              │
// The greeting service definition.                                                                                                                                                           │
service Greeter {                                                                                                                                                                             │
  // Sends a greeting                                                                                                                                                                         │
  rpc SayHello (HelloRequest) returns (HelloReply) {}                                                                                                                                         │
}                                                                                                                                                                                             │
                                                                                                                                                                                              │
// The request message containing the user's name.                                                                                                                                            │
message HelloRequest {                                                                                                                                                                        │
  string name = 1;                                                                                                                                                                            │
}                                                                                                                                                                                             │
                                                                                                                                                                                              │
// The response message containing the greetings                                                                                                                                              │
message HelloReply {                                                                                                                                                                          │
  string message = 1;                                                                                                                                                                         │
}
```
> 2.安装 protoc 编译器和 Go gRPC 插件来生成 Go 代码。使用 go get 来安装这些工具
```shell
 go get -u google.golang.org/protobuf/cmd/protoc-gen-go                                       

```
> 3.使用 protoc 编译器生成 Go 代码
 ```shell
protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative pkg/grpc/helloworld/helloworld.proto                   

```
> 4.实现 gRPC 服务
```text
   1. 启动服务器
      打开一个新的终端并运行以下命令：

   go run pkg/grpc/server/main.go

   2. 运行客户端
      打开另一个终端并运行以下命令：

   go run pkg/grpc/client/main.go
```
> 5.测试 gRPC 服务
```text
 pip install grpcio grpcio-tools
 python -m grpc_tools.protoc -I. --python_out=. --grpc_python_out=. ./pkg/grpc/helloworld/helloworld.proto
```
```text
go run ./pkg/pubsub/server/main.go
go run ./pkg/pubsub/client/main.go
```