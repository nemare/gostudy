package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	// 1

	// 2

	var a map[string]string
	fmt.Println(len(a))

	b := "-1"
	c, _ := strconv.Atoi(strings.Split(b, ",")[0])
	fmt.Println(c)
}
