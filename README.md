# kitex-demo

```sh
# Ubuntu
sudo apt install golang-go

# MacOS
brew install go

# Windows
scoop install go

# env
export GO111MODULE=on
export GOPATH="$HOME/go"
export GOBIN="$GOPATH/bin"
export PATH="$GOBIN:$PATH"
export GOPROXY=https://goproxy.cn

go mod init github.com/njupt-sakura/kitex-demo
go install github.com/cloudwego/kitex/tool/cmd/kitex@latest

which kitex
kitex --version
```

编写 [base.thrift](./idl/base.thrift), [item.thrift](./idl/item.thrift), [stock.thrift](./idl/stock.thrift)

执行 kitex 命令

```bash
pwd # /path/to/kitex-demo

kitex -module github.com/njupt-sakura/kitex-demo idl/item.thrift
kitex -module github.com/njupt-sakura/kitex-demo idl/stock.thrift

mkdir -p ./rpc/item ./rpc/stock

cd ./rpc/item &&                                        \
kitex -module github.com/njupt-sakura/kitex-demo        \
      -service item                                     \
      -use github.com/njupt-sakura/kitex-demo/kitex_gen \
      ../../idl/item.thrift

cd ../stock &&                                          \
kitex -module github.com/njupt-sakura/kitex-demo        \
      -service stock                                    \
      -use github.com/njupt-sakura/kitex-demo/kitex_gen \
      ../../idl/stock.thrift

cd ../.. && go mod tidy
```

### 编写 item 商品 RPC 服务 (8888 端口)

[rpc/item/handler.go](./rpc/item/handler.go)

```bash
cd ./rpc/item
sh ./build.sh && sh ./output/bootstrap.sh # 等价于 go run .
```

### 编写 API HTTP 服务 (8889 端口)

编写 [api/main.go](./api/main.go) 暴露 HTTP 接口

```sh
cd ./api
go run .

curl http://localhost:8889/api/item
# GetItemRes({Item:Item({Id:1024 Title:Title of the item Desc:Description of the item Stock:0}) BaseRes:BaseRes({Code: Msg:})})
```

### 编写 stock 库存 RPC 服务

[rpc/stock/handler.go](./rpc/stock/handler.go)

- item 商品 RPC 服务占用了 8888 端口 (默认 8888 端口)
- API HTTP 服务占用了 8889 端口
- 需要在 stock 库存 RPC 服务中修改监听的端口

```bash
cd ./rpc/stock
go run . # 等价于 sh ./build.sh && sh ./output/bootstrap.sh
```

### 补充 item 商品 RPC 服务

[rpc/item/handler](./rpc/item/handler.go)

现在

- item 商品 RPC 服务: 8888 端口
- stock 库存 RPC 服务: 8890 端口
- API HTTP 服务: 8889 端口

补充 item 商品 RPC 服务, 需要创建 stock 库存 RPC 客户端, 创建 RPC 请求, 构造 RPC 请求参数, 发起 RPC 调用

- 在 [rpc/item/main.go](./rpc/item/main.go) 中创建 stock 库存 RPC 客户端
- 在 [rpc/item/handler.go](./rpc/item/handler.go) 中创建 RPC 请求, 构造 RPC 请求参数, 发起 RPC 调用

```bash
cd ./rpc/item
sh ./build.sh && sh ./output/bootstrap.sh

curl http://localhost:8889/api/item
# GetItemRes({Item:Item({Id:1024 Title:Title of the item Desc:Description of the item Stock:1024}) BaseRes:BaseRes({Code: Msg:})})
```

### 服务注册

```bash
# MacOS
brew install docker-compose

docker-compose up -d etcd
go get github.com/kitex-contrib/registry-etcd
```

- 在 [rpc/stock/main.go](./rpc/stock/main.go) 中注册 stock 库存 RPC 服务
- 在 [rpc/item/main.go](./rpc/item/main.go) 中注册 item 商品 RPC 服务

```bash
cd ./rpc/item
sh ./build.sh && sh ./output/bootstrap.sh

cd ../stock
sh ./build.sh && sh ./output/bootstrap.sh

# 验证
docker exec -it etcd /bin/bash
etcdctl get --prefix 'kitex'
```

### 服务发现

- 在 [rpc/item/handler.go](./rpc/item/handler.go) 中使用服务发现
- 在 [api/main.go](./api/main.go) 中使用服务发现

```bash
cd ./rpc/stock
sh ./build.sh && sh ./output/bootstrap.sh

cd ../item
sh ./build.sh && sh ./output/bootstrap.sh

cd ../../api
go run .

curl http://localhost:8889/api/item
# GetItemRes({Item:Item({Id:1024 Title:Title of the item Desc:Description of the item Stock:1024}) BaseRes:BaseRes({Code: Msg:})})
```
