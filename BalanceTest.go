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

type Balance struct {
	Title   string          `json:"title,omitempty"`
	Remark  string          `json:"remark,omitempty"`
	Value   string          `json:"value,omitempty"`
	Amount  decimal.Decimal `json:"amount,omitempty"`
	Entries []Entry         `json:"entries"`
}
type Entry struct {
	ID    int    `json:"id"`
	Icon  string `json:"icon"`
	Name  string `json:"name"`
	Desc  string `json:"desc"`
	Url   string `json:"url"`
	Value string `json:"value"`
}
