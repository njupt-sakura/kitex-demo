package main

import (
	"log"
	"net"

	rpcinfo "github.com/cloudwego/kitex/pkg/rpcinfo"
	server "github.com/cloudwego/kitex/server"
	etcd "github.com/kitex-contrib/registry-etcd"
	stock "github.com/njupt-sakura/kitex-demo/kitex_gen/stock/stockservice"
)

func main() {
	// 创建服务注册中心
	r, err := etcd.NewEtcdRegistry([]string{"127.0.0.1:2379"})
	if err != nil {
		log.Fatalln(err)
	}

	addr, _ := net.ResolveTCPAddr("tcp", "127.0.0.1:8890")
	svr := stock.NewServer(new(StockServiceImpl), server.WithServiceAddr(addr),
		// 指定服务注册中心和服务基本信息
		server.WithRegistry(r),
		server.WithServerBasicInfo(
			&rpcinfo.EndpointBasicInfo{
				ServiceName: "stock",
			},
		),
	)

	if err := svr.Run(); err != nil {
		log.Println(err.Error())
	}
}
