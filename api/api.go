package api

import (
	"fmt"
	"mvcgolangproject/model"
)

type Request struct {
	Transid string
	Orderid int
}

type Response struct {
	Status string
}

func Controller(req *Request) interface{} {
	var mysql model.Model
	fmt.Println("Controller or API")
	// mysql.Operation("IS123", 123)
	response := mysql.Operation(req.Transid, req.Orderid)
	return response
}
