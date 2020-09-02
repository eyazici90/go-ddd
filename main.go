package main

import (
	"fmt"
	"reflect"

	"orderContext/domain/order"
)

func main() {

	fmt.Println("App Started!!!")

	order := order.NewOrder("1", "21", "312")

	order.Pay()

	events := order.Events()

	order.ClearEvents()

	for _, e := range events {
		fmt.Println(reflect.TypeOf(e).Name())
	}

}
