## wxgamevp

[![Build Status](https://travis-ci.com/Birjemin/wxgamevp.svg?branch=master)](https://travis-ci.com/Birjemin/wxgamevp) [![Go Report Card](https://goreportcard.com/badge/github.com/birjemin/wxgamevp)](https://goreportcard.com/report/github.com/birjemin/wxgamevp) [![codecov](https://codecov.io/gh/Birjemin/wxgamevp/branch/master/graph/badge.svg)](https://codecov.io/gh/Birjemin/wxgamevp)


[开发者中心](https://developers.weixin.qq.com/minigame/dev/api-backend/midas-payment/midas.cancelPay.html)

### 引入方式
```
go get github.com/birjemin/wxgamevp
```

### 接口列表

- [cancelPay](https://developers.weixin.qq.com/minigame/dev/api-backend/midas-payment/midas.cancelPay.html) ✅
- [getBalance](https://developers.weixin.qq.com/minigame/dev/api-backend/midas-payment/midas.getBalance.html) ✅
- [pay](https://developers.weixin.qq.com/minigame/dev/api-backend/midas-payment/midas.pay.html) ✅
- [present](https://developers.weixin.qq.com/minigame/dev/api-backend/midas-payment/midas.present.html) ✅

### 使用方式

- 示例

```golang
httpClient := &utils.HTTPClient{
    Client: &http.Client{
        Timeout: 5 * time.Second,
    },
}
```

### 备注
- 测试
    ```
    go test
    ```
- 格式化代码
    ```
    golint
    ```
- 覆盖率
    ```
    go test -cover
    go test -coverprofile=coverage.out 
    go tool cover -html=coverage.out
    ```