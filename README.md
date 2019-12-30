# huobiapi

火币网 API Go 客户端

## 安装

本模块使用 [godep](https://github.com/golang/dep) 作为包管理工具


```go
package main

import (
    "fmt"
    "github.com/merkles/huobiapi"
)

func main() {
    // 创建客户端实例
    market, err := huobiapi.NewMarket()
    if err != nil {
        panic(err)
    }
    // 订阅主题
    market.Subscribe("market.eosusdt.trade.detail", func(topic string, json *huobiapi.JSON) {
        // 收到数据更新时回调
        fmt.Println(topic, json)
    })
    // 请求数据
    json, err := market.Request("market.eosusdt.detail")
    if err != nil {
        panic(err)
    } else {
        fmt.Println(json)
    }
    // 进入阻塞等待，这样不会导致进程退出
    market.Loop()
}
```

## RESTful 版行情和交易查询

```go
package main

import (
    "fmt"
)

func main() {
    // 创建客户端实例
    client, err := huobiapi.NewClient("key id", "key secret")
    if err != nil {
        panic(err)
    }
    // 发送请求
    ret, err := client.Request("GET", "/market/history/trade", huobiapi.ParamsData{
        "symbol": "eosusdt",
        "size": "10",
    })
    data, err := ret.Get("data").Array()
    if err != nil {
        panic(err)
    }
    for _, v := range data {
        fmt.Println(v)
    }
}
```
