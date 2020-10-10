package main

import (
	"fmt"
	"reflect"
)

func sqlQuote(x interface{}) string {
	switch x := x.(type) {
	case nil:
		return "null"
	case int, uint:
		fmt.Println(reflect.TypeOf(x))
		return fmt.Sprintf("%d", x)
	case bool:
		if x {
			return "TRUE"
		}
		return "FALSE"
	case string:
		return x
	default:
		panic(fmt.Sprintf("unexpected type %T : %v", x, x))
	}
}

func main() {
	x := 10
	fmt.Println(sqlQuote(x))
}
