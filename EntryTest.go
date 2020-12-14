package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	//allAmount := decimal.NewFromInt(0)

	balance := &Balance{
		Entries: make([]Entry, 0),
	}

	allEntries := []*Entry{
		{ID: 21, Value: "8.92元"},
		{ID: 22, Value: "16.02元"},
	}

	if len(allEntries) == 0 {
		return
	}

	for _, entry := range allEntries {
		if entry.ID == 22 {
			entry.Value = "0元"
		}
	}

	for _, entry := range allEntries {
		fmt.Println(entry)
	}

	//balance.Amount = allAmount
	//fmt.Println(balance.Amount)
	fmt.Println(balance.Entries)
	fmt.Println(len(balance.Entries))

	balance1 := &Balance{}
	fmt.Println(balance1.Value)

	fmt.Println(BuildWjActBudgetKey(103, 1234, "1dasda"))
}

type Balance struct {
	Title  string `json:"title,omitempty"`
	Remark string `json:"remark,omitempty"`
	Value  string `json:"value,omitempty"`
	//Amount  decimal.Decimal `json:"amount,omitempty"`
	Entries []Entry `json:"entries"`
}
type Entry struct {
	ID    int    `json:"id"`
	Icon  string `json:"icon"`
	Name  string `json:"name"`
	Desc  string `json:"desc"`
	Url   string `json:"url"`
	Value string `json:"value"`
}

func BuildWjActBudgetKey(scene int, actId int64, subFactorId string) string {
	actStr := "{" + strconv.Itoa(int(actId)) + "}"
	return strings.Join([]string{"activity_budget", strconv.Itoa(int(scene)), actStr, subFactorId}, ":")
}
