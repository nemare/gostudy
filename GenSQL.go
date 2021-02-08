package main

import "fmt"

func main() {

	sql := ""
	for i := 700; i <= 999; i++ {
		fmt.Println(fmt.Sprintf(sql, i))
	}
}
