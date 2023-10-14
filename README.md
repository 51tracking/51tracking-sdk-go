51tracking-sdk-go
=================

The Go SDK of 51Tracking API

Contact: <service@51tracking.org>

## Official document

[Document](https://www.51tracking.com/v4/api-index/API-)

## Index
1. [Installation](https://github.com/51tracking/51tracking-sdk-go#installation)
2. [Testing](https://github.com/51tracking/51tracking-sdk-go#testing)
3. [Error Handling](https://github.com/51tracking/51tracking-sdk-go#error-handling)
4. SDK
   1. [Couriers](https://github.com/51tracking/51tracking-sdk-go#couriers)
   2. [Trackings](https://github.com/51tracking/51tracking-sdk-go#trackings)
   3. [Air Waybill](https://github.com/51tracking/51tracking-sdk-go#air-waybill)


## Installation

51tracking-sdk-go 要求 Go 版本支持 [模块](https://github.com/golang/go/wiki/Modules)，并使用导入版本控制。因此，请确保在安装 51tracking-sdk-go 之前初始化 Go 模块：

```
go mod init github.com/my/repo
go get github.com/51tracking/51tracking-sdk-go
```

插件引入:

``` go
import "github.com/51tracking/51tracking-sdk-go"
```

## Quick Start

```go
package main

import (
   "context"
   "fmt"
   "github.com/51tracking/51tracking-sdk-go"
)

func main() {
   key := "you api key"
   cli, err := tracking51.NewClient(key)

   if err != nil {
      fmt.Println(err)
      return
   }

   result, err := cli.GetAllCouriers(context.Background())
   if err != nil {
      fmt.Println(err)
      return
   }

   fmt.Println(result)

   var couriers, ok = result.Data.(*[]tracking51.Courier)
   if !ok {
      fmt.Println("Structure type conversion failed")
      return
   }
   for _, item := range *couriers {
      fmt.Printf("courier_name:%s courier_code:%s\n", item.CourierName, item.CourierCode)
   }
}

```

## Testing
```
go test
```

## Error handling

**Throw** by the new SDK client

```go
cli, err := tracking51.NewClient("")

if err != nil {
fmt.Println(err)
return
}

/*
API Key is missing
*/
```

**Throw** by the parameter validation in function

```go
cli, err := tracking51.NewClient("you api key")

if err != nil {
   fmt.Println(err)
   return
}

params := tracking51.DetectParams{
    TrackingNumber: "",
}
result, err := cli.Detect(context.Background(), params)
if err != nil {
   fmt.Println(err)
   return
}

/*
Tracking number cannot be empty
*/
```
## Examples

## Couriers
##### Return a list of all supported couriers.
https://api.51Tracking.com/v4/couriers/all
```go
result, err := cli.GetAllCouriers(context.Background())
if err != nil {
  fmt.Println(err)
  return
}

fmt.Println(result)
```

##### Return a list of matched couriers based on submitted tracking number.
https://api.51Tracking.com/v4/couriers/detect
```go
params := tracking51.DetectParams{
  TrackingNumber: "92612903029511573030094531",
}

result, err := cli.Detect(context.Background(), params)
if err != nil {
  fmt.Println(err)
  return
}

fmt.Println(result)
```

## Trackings
##### Create a tracking.
https://api.51Tracking.com/v4/trackings/create
```go
params := tracking51.CreateTrackingParams{
  TrackingNumber: "9400111899562537683144",
  CourierCode:    "usps",
}

result, err := cli.CreateTracking(context.Background(), params)
if err != nil {
  fmt.Println(err)
  return
}

fmt.Println(result)
```

##### Get tracking results of multiple trackings.
https://api.51Tracking.com/v4/trackings/get
```go
// Perform queries based on various conditions
//params := tracking51.GetTrackingResultsParams{
//	TrackingNumbers: "92612903029511573030094532",
//	CourierCode:     "usps",
//}
//
//params := tracking51.GetTrackingResultsParams{
//	TrackingNumbers: "92612903029511573030094531,9400111899562539126562",
//	CourierCode:     "usps",
//}

currentTime := time.Now()
zeroTime := currentTime.UTC()
layout := "2006-01-02T15:04:05-07:00"
formattedTime := zeroTime.Format(layout)
params := tracking51.GetTrackingResultsParams{
  CreatedDateMin: "2023-08-23T06:00:00+00:00",
  CreatedDateMax: formattedTime,
}

result, err := cli.GetTrackingResults(context.Background(), params)
if err != nil {
  fmt.Println(err)
  return
}

fmt.Println(result)
```

##### Create multiple trackings (Max. 40 tracking numbers create in one call).
https://api.51Tracking.com/v4/trackings/batch
```go
params := []tracking51.CreateTrackingParams{
  {
      TrackingNumber: "92632903279511573030094832",
      CourierCode:    "usps",
  },
  {
      TrackingNumber: "92642903289511563030094932",
      CourierCode:    "usps",
  },
}

result, err := cli.BatchCreateTrackings(context.Background(), params)
if err != nil {
  fmt.Println(err)
  return
}

fmt.Println(result)
```

##### Update a tracking by ID.
https://api.51Tracking.com/v4/trackings/update/{id}
```go
params := tracking51.UpdateTrackingParams{
   CustomerName: "New name",
   Note:         "New tests order note",
}

idString := "9a1d3844a50f3851e76e3ee347881588"
result, err := cli.UpdateTrackingByID(context.Background(), idString, params)
if err != nil {
   fmt.Println(err)
   return
}

fmt.Println(result)
```

##### Delete a tracking by ID.
https://api.51Tracking.com/v4/trackings/delete/{id}
```go
idString := "9a1d3844a50f3851e76e3ee347881588"
result, err := cli.DeleteTrackingByID(context.Background(), idString)
if err != nil {
   fmt.Println(err)
   return
}

fmt.Println(result)
```

##### Retrack expired tracking by ID.
https://api.51Tracking.com/v4/trackings/retrack/{id}
```go
idString := "99ff2ce10105aeb8627ec0c03e1773bd"
result, err := cli.RetrackTrackingByID(context.Background(), idString)
if err != nil {
   fmt.Println(err)
   return
}

fmt.Println(result)
```
## Air Waybill
##### Create an air waybill.
https://api.51Tracking.com/v4/awb
```go
params := tracking51.AirWaybillParams{
AwbNumber: "235-69030430",
}

result, err := cli.CreateAnAirWayBill(context.Background(), params)
if err != nil {
fmt.Println(err)
return
}

fmt.Println(result)
```

## 响应状态码

51Tracking 使用传统的HTTP状态码来表明 API 请求的状态。通常，2xx形式的状态码表示请求成功，4XX形式的状态码表请求发生错误（比如：必要参数缺失），5xx格式的状态码表示 51tracking 的服务器可能发生了问题。

| Http CODE | META CODE | TYPE               | MESSAGE                                   |
|-----------|-----------|--------------------|-------------------------------------------|
| 200       | 200       | <code>成功</code>    | 请求响应成功。                                   |
| 400       | 400       | <code>错误请求</code>  | 请求类型错误。请查看 API 文档以了解此 API 的请求类型。          |
| 400       | 4101      | <code>错误请求</code>  | 物流单号已存在。                                  |
| 400       | 4102      | <code>错误请求</code>  | 物流单号不存在。请先使用「Create接口」将单号添加至系统。           |
| 400       | 4103      | <code>错误请求</code>  | 您已超出 API 调用的创建数量。每次创建的最大数量为 40 个快递单号。     |
| 400       | 4110      | <code>错误请求</code>  | 物流单号(tracking_number) 不符合规则。              |
| 400       | 4111      | <code>错误请求</code>  | 物流单号(tracking_number)为必填字段。               |
| 400       | 4112      | <code>错误请求</code>  | 查询ID无效。                                   |
| 400       | 4113      | <code>错误请求</code>  | 不允许重新查询。您只能重新查询过期的物流单号。                   |
| 400       | 4120      | <code>错误请求</code>  | 物流商简码(courier_code)的值无效。                  |
| 400       | 4121      | <code>错误请求</code>  | 无法识别物流商。                                  |
| 400       | 4122      | <code>错误请求</code>  | 特殊物流商字段缺失或填写不符合规范。                        |
| 400       | 4130      | <code>错误请求</code>  | 请求参数的格式无效。                                |
| 400       | 4160      | <code>错误请求</code>  | 空运单号(awb_number)是必需的或有效的格式。               |
| 400       | 4161      | <code>错误请求</code>  | 此空运航空不支持查询。                               |
| 400       | 4165      | <code>错误请求</code>  | 查询失败：未创建空运单号。                             |
| 400       | 4166      | <code>错误请求</code>  | 删除未创建的空运单号失败。                             |
| 400       | 4167      | <code>错误请求</code>  | 空运单号已存在，无需再次创建。                           |
| 400       | 4190      | <code>错误请求</code>  | 当前查询额度不足。                                 |
| 401       | 401       | <code>未经授权</code>  | 身份验证失败或没有权限。请检查并确保您的 API 密钥正确无误。          |
| 403       | 403       | <code>禁止</code>    | 禁止访问。请求被拒绝或不允许访问。                         |
| 404       | 404       | <code>未找到</code>   | 页面不存在。请检查并确保您的链接正确无误。                     |
| 429       | 429       | <code>太多请求</code>  | 超出 API 请求限制，请稍后重试。请查看 API 文档以了解此 API 的限制。 |
| 500       | 511       | <code>服务器错误</code> | 服务器错误。请联系我们： service@51Tracking.org。      |
| 500       | 512       | <code>服务器错误</code> | 服务器错误。请联系我们：service@51Tracking.org。       |
| 500       | 513       | <code>服务器错误</code> | 服务器错误。请联系我们： service@51Tracking.org。      |