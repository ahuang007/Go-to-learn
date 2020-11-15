package main

import (
	"encoding/json"
	"fmt"
)

type Product struct {
	Name      string  `json:Name`
	ProductId int64   `json:ProductId`
	No        int     `jsoon:No`
	Price     float64 `json:Price`
	IsOnSale  bool    `json:IsOnSale`
}

func main() {
	p := &Product{}
	p.Name = "iphone 8"
	p.IsOnSale = true
	p.No = 10001
	p.ProductId = 1
	p.Price = 4999.0

	data, _ := json.Marshal(p)
	strJson := string(data)
	fmt.Println(strJson)
	//fmt.Println(*p)

	q := &Product{}
	err := json.Unmarshal([]byte(strJson), q)
	fmt.Println(err)
	fmt.Println(*q)
}
