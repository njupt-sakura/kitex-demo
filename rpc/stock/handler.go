package main

import (
	"context"
	stock "github.com/njupt-sakura/kitex-demo/kitex_gen/stock"
)

// StockServiceImpl implements the last service interface defined in the IDL.
type StockServiceImpl struct{}

// GetItemStock implements the StockServiceImpl interface.
func (s *StockServiceImpl) GetItemStock(ctx context.Context, req *stock.GetItemStockReq) (res *stock.GetItemStockRes, err error) {
	res = stock.NewGetItemStockRes()
	res.Stock = req.GetItemId()
	return // res, nil
}
