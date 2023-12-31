package tracking51_test

import (
	"context"
	"fmt"
	"github.com/51tracking/51tracking-sdk-go"
)

func ExampleClient_GetCouriers() {
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
