package main

import (
	"fmt"
	"github.com/shopspring/decimal"
)

func main() {

	var balance *Balance
	balance = &Balance{
		Entries: make([]Entry, 2),
	}
	amount := decimal.NewFromInt(0)
	balance.Amount = amount
	fmt.Println(balance.Amount)
}
