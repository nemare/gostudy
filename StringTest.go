package main

import (
	"fmt"
	"strings"
)

func main() {
	a := "abcd"
	fmt.Println(a[2])
	fmt.Println("%c", a[2])
	fmt.Println(strings.TrimRight("01.1元", "元"))

	var b []string
	fmt.Println(len(b))
}
