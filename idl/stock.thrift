// TODO namespace go com.github.njupt-sakura.kitex-demo
namespace go stock

include "base.thrift"

struct GetItemStockReq {
  1: required i64 itemId
}

struct GetItemStockRes {
  1: i64 stock

  255: base.BaseRes baseRes
}

service StockService {
  GetItemStockRes GetItemStock(1: GetItemStockReq req)
}
