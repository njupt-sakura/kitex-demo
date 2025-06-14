package main

import (
	"log"

	rpcinfo "github.com/cloudwego/kitex/pkg/rpcinfo"
	server "github.com/cloudwego/kitex/server"
	etcd "github.com/kitex-contrib/registry-etcd"
	item "github.com/njupt-sakura/kitex-demo/kitex_gen/item/itemservice"
)

func main() {
	// 创建服务注册中心
	r, err := etcd.NewEtcdRegistry([]string{"127.0.0.1:2379"})
	if err != nil {
		log.Fatalln(err)
	}

	itemServiceImpl := new(ItemServiceImpl)

	// 创建 stock 库存 RPC 客户端
	stockCli, err := NewStockClient("0.0.0.0:8890")
	if err != nil {
		log.Fatalln(err)
	}
	itemServiceImpl.stockCli = stockCli

	svr := item.NewServer(itemServiceImpl,
		// 指定服务注册中心和服务基本信息
		server.WithRegistry(r),
		server.WithServerBasicInfo(
			&rpcinfo.EndpointBasicInfo{
				ServiceName: "item",
			},
		))

	if err := svr.Run(); err != nil {
		log.Println(err.Error())
	}
}
