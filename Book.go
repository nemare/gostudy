package main

import (
	"fmt"
	"strings"
)

type Book struct {
	title   string
	author  string
	subject string
	book_id int64
}

func main() {
	newBook := Book{"我是", "piaozhikang", "化学", 12398839947384734}
	fmt.Println(newBook.author, newBook.book_id, newBook.author, newBook.subject)

	var a float64
	fmt.Println(a)

	b := ""
	coordinatorList := strings.Split(b, ",")

	fmt.Println(len(coordinatorList))
}
