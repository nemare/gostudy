package main

import "fmt"

func main() {
	// 1
	fmt.Println(f() == f())
	fmt.Println(*f() == *f())

	// 2
	v := 1
	incr(&v)
	fmt.Println(v)
	fmt.Println(incr(&v))
}

func f() *int {
	v := 1
	fmt.Println(&v)
	return &v
}

func incr(p *int) int {
	*p++
	return *p
}
