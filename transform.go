package main

import (
	"fmt"
	"strconv"
)

func main() {
	var a uint64
	a = 18446744073709551615
	intStr := strconv.FormatUint(a, 10)

	intNum, _ := strconv.Atoi(intStr)

	fmt.Println(intNum)

	fmt.Println(int64(a))
}
