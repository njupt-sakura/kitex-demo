# IDL

IDL, Interface Definition Language 接口定义语言

thrift 基本类型

1. bool
2. byte
3. i8
4. i16
5. i32
6. i64
7. double
8. string

thrift 特殊类型: binary

thrift 容器类型

1. `list<t1>`
2. `set<t1>`
3. `map<k1, v2>`

thrift 类型定义

```thrift
// 类似 C++ 的类型定义
typedef i32 MyInteger
```

thrift 枚举类型

```thrift
enum TweetType {
  Zero // 0 默认从 0 开始赋值
  Two = 2
  Ten = 0xa
}
```

thrift 命名空间: 类似 Java, Go 的 package

```thrift
namespace cpp  com.github.njupt-sakura.kitex-demo
namespace java com.github.njupt-sakura.kitex-demo
namespace go   com.github.njupt-sakura.kitex-demo
```

允许 include 其他 thrift 文件

```thrift
include "base.thrift"
```

thrift 定义常量

```thrift
const i32 INT_CONST = 123;

const map<string, string> MAP_CONST = {
  "k1": "v1",
  "k2": "v2"
}
```

thrift 的 struct

struct 的 field 有唯一 ID

- required 必填字段
- optional 选填字段
- default 字段有默认值

```thrift
struct Location {
  1: required double lat;
  2: required double lng;
}

struct Tweet {
  1: required i32 userId;
  2: required string userName;
  3: required string msg;
  4: optional Location location;
  16: optional string lang = "English" // 默认值
}
```

thrift 的 service, 类似 Java, Go 的接口

```thrift
service Twitter {
  void ping();
  bool postTweet(1: Tweet tweet);
  list<Tweet> searchTweet(1: string query);
  // The 'oneway' modifier indicates that the client only makes a request and does not wait for any response at all. Oneway methods MUST be void.
  oneway void zip();
}
```
