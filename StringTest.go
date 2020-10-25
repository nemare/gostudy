package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

func main() {
	a := "abcd"
	fmt.Println(a[2])
	fmt.Println("%c", a[2])
	fmt.Println(strings.TrimRight("01.1元", "元"))

	var b []string
	fmt.Println(len(b))

	var bigint int64
	bigint = 144115188075865872
	fmt.Println(strconv.FormatInt(bigint, 10))

	//ab := make(map[string]interface{})
	var ab map[string]interface{}
	ab["resp"] = 12
	c, ok := ab["resp"].(string)
	if ok {
		fmt.Println(c)
	} else {
		fmt.Println(ok)
	}

	fmt.Println(time.Unix(0, 1603294383368098192))

	var num []int
	fmt.Println(len(num))
}
