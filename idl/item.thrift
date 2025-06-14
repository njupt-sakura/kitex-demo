// TODO namespace go com.github.njupt-sakura.kitex-demo
namespace go item

include "base.thrift"

struct Item {
  1: i64 id
  2: string title
  3: string desc
  4: i64 stock
}

struct GetItemReq {
  1: required i64 id
}

struct GetItemRes {
  1: Item item

  255: base.BaseRes baseRes
}

service ItemService {
  GetItemRes GetItem(1: GetItemReq req)
}
