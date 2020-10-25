package main

import (
	"fmt"
	"strings"

	"github.com/shopspring/decimal"
)

func main() {
	allAmount := decimal.NewFromInt(0)

	balance := &Balance{
		Entries: make([]Entry, 0),
	}

	allEntries := []Entry{
		{ID: 21, Value: "8.92元"},
		{ID: 22, Value: "16.02元"},
	}

	if len(allEntries) == 0 {
		return
	}

	for _, entry := range allEntries {
		if entry.ID == 21 || entry.ID == 22 {
			amountStr := strings.TrimRight(entry.Value, "元")
			amount, err := decimal.NewFromString(amountStr)
			if err != nil {
				continue
			}
			balance.Entries = append(balance.Entries, entry)
			allAmount = allAmount.Add(amount)
		}
	}

	balance.Amount = allAmount
	fmt.Println(balance.Amount)
	fmt.Println(balance.Entries)
	fmt.Println(len(balance.Entries))
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
