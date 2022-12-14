# enmooy-account-sdk

## 使用方法

``` shell
go get -u github.com/jack0829/enmooy-account-sdk
```

### JWT 解析

``` go
package main

import (
    "fmt"
    sdk "github.com/jack0829/enmooy-account-sdk/jwt"
)

var (

    // JWT 签名密钥
    key []byte = []byte("your key")
    
    // 数据加密密钥
    salt string = "your salt"
)

func main() {

    // 请求中的 token
    token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiJsb2dpbiIsImV4cCI6MTY3MTI1Mzc0MiwiaWF0IjoxNjcwOTk0NTQyLCJkYXRhIjp7InVpZCI6IncxeWVWN1k3Uk9LbGc5T0UiLCJlaWQiOiJlbzQyTG1LcXFWS0pxOGI5IiwibmljayI6ImFiYyIsImF2YXRhciI6Inh4LmpwZyIsInN2YyI6WyJvc3MiLCJzbGIiXX19.tgLmrY0N1RFzEum8xo6lWVJyuFW2XeTldfjBlHaM2qE"
    
    jwt := sdk.New(key, salt)
    if err := jwt.Decode(token); err != nil {
        fmt.Println(err)
        return
    }
    
    // 解析成功
    data := jwt.Data()
    fmt.Printf("加密后数据 %#+v\n", data)
    
    // 获取原始 uid 和 eid
    if uid, err := data.GetUid(jwt.GetEncoder()); err == nil {
        fmt.Printf("原始用户ID = %d\n", uid)
    }
    if eid, err := data.GetEid(jwt.GetEncoder()); err == nil {
        fmt.Printf("原始企业ID = %d\n", eid)
    }
}
```

如果一切没有错误，将会输出：

``` log
加密后数据 &jwt.Data{Uid:"w1yeV7Y7ROKlg9OE", Eid:"eo42LmKqqVKJq8b9", Nick:"abc", Avatar:"xx.jpg", Svc:[]string{"oss", "slb"}}
原始用户ID = 123
原始企业ID = 456
```

