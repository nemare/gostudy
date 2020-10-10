package main

import (
	"fmt"
	"reflect"
	"time"
)

type People struct {
	Age  int
	Name string
}

const b = 1

func test(a time.Duration) {
	fmt.Println(a)
}

func main() {

	p := People{1, "张三"}
	var p2 = p
	var p1 = p

	fmt.Printf("%p, %p, %p", &p, &p1, &p2)
	fmt.Println("")

	c := b + 1
	fmt.Println(reflect.TypeOf(b))
	test(b)
	fmt.Println(reflect.TypeOf(c))
	//test(b)
}
