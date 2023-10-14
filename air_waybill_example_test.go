package tracking51_test

import (
	"context"
	"fmt"
	"github.com/51tracking/51tracking-sdk-go"
)

func ExampleClient_CreateAnAirWayBill() {
	key := "you api key"
	cli, err := tracking51.NewClient(key)

	if err != nil {
		fmt.Println(err)
		return
	}

	params := tracking51.AirWaybillParams{
		AwbNumber: "235-69030430",
	}
	result, err := cli.CreateAnAirWayBill(context.Background(), params)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(result)

	var airWaybillItem, ok = result.Data.(*tracking51.AirWaybillItem)
	if !ok {
		fmt.Println("Structure type conversion failed")
		return
	}
	fmt.Printf("awb_number:%s destination:%s\n", airWaybillItem.AwbNumber, airWaybillItem.Destination)

}
