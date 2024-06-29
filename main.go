package main

import (
	"fmt"
	"mvcgolangproject/api"
)

func main() {
	input := api.Request{
		Transid: "IS234",
		Orderid: 3435,
	}
	ret := api.Controller(&input)
	fmt.Println(ret)
}
