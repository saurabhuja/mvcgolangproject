package model

import "fmt"

type Model struct {
	Transid string
	Orderid int
}

func (p *Model) Operation(tid string, oid int) string {
	fmt.Println("Model or Database operation")
	p.Transid = tid
	p.Orderid = oid
	return "apisuccess"
}
