package main

import (
	"context"
	"log"
	"time"

	client "github.com/cloudwego/kitex/client"
	callopt "github.com/cloudwego/kitex/client/callopt"
	etcd "github.com/kitex-contrib/registry-etcd"
	item "github.com/njupt-sakura/kitex-demo/kitex_gen/item"
	itemservice "github.com/njupt-sakura/kitex-demo/kitex_gen/item/itemservice"

	app "github.com/cloudwego/hertz/pkg/app"
	server "github.com/cloudwego/hertz/pkg/app/server"
)

var cli itemservice.Client

func main() {
	// // 创建 client
	// c, err := itemservice.NewClient("item" /* destService */, client.WithHostPorts("0.0.0.0:8888"))
	// if err != nil {
	// 	log.Fatalln(err)
	// }
	// cli = c

  // 使用服务发现
	resolver, err := etcd.NewEtcdResolver([]string{"127.0.0.1:2379"})
	if err != nil {
		log.Fatalln(err)
	}
	c, err := itemservice.NewClient("item", client.WithResolver(resolver))
	if err != nil {
		log.Fatalln(err)
	}
	cli = c

	// 使用 net/http 或其他 HTTP 框架暴露 HTTP 接口
	// 这里使用 hertz HTTP 框架
	httpServer := server.New(server.WithHostPorts("0.0.0.0:8889"))
	httpServer.GET("/api/item", ItemHandler)

	if err := httpServer.Run(); err != nil {
		log.Fatalln(err)
	}
}

func ItemHandler(ctx context.Context, reqContext *app.RequestContext) {
	// 调用 RPC 服务
	rpcReq := item.NewGetItemReq()
	rpcReq.Id = 1024
	rpcRes, err := cli.GetItem(context.Background(), rpcReq, callopt.WithRPCTimeout(3*time.Second))
	if err != nil {
		log.Fatalln(err)
	}
	reqContext.String(200, rpcRes.String())
}
