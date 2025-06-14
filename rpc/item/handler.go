package main

import (
	"context"
	"log"

	"github.com/cloudwego/kitex/client"
	etcd "github.com/kitex-contrib/registry-etcd"
	item "github.com/njupt-sakura/kitex-demo/kitex_gen/item"
	"github.com/njupt-sakura/kitex-demo/kitex_gen/stock"
	"github.com/njupt-sakura/kitex-demo/kitex_gen/stock/stockservice"
)

// ItemServiceImpl implements the last service interface defined in the IDL.
type ItemServiceImpl struct {
	stockCli stockservice.Client
}

func NewStockClient(addr string) (stockservice.Client, error) {
	// // return stockservice.NewClient("stock" /* destService */, client.WithHostPorts(addr))

	// 使用服务发现
	resolver, err := etcd.NewEtcdResolver([]string{"127.0.0.1:2379"})
	if err != nil {
		log.Fatalln(err)
	}
	return stockservice.NewClient("stock", client.WithResolver(resolver))
}

// GetItem implements the ItemServiceImpl interface.
func (s *ItemServiceImpl) GetItem(ctx context.Context, req *item.GetItemReq) (res *item.GetItemRes, err error) {
	res = item.NewGetItemRes()
	res.Item = item.NewItem()
	res.Item.Id = req.GetId()
	res.Item.Title = "Title of the item"
	res.Item.Desc = "Description of the item"

	// 创建 RPC 请求
	stockReq := stock.NewGetItemStockReq()
	// 构造 RPC 请求参数
	stockReq.ItemId = req.GetId()
	// 发起 RPC 调用
	stockRes, err := s.stockCli.GetItemStock(context.Background(), stockReq)
	if err != nil {
		log.Println(err)
		stockRes.Stock = 0
	}
	res.Item.Stock = stockRes.GetStock()

	return res, nil
}
