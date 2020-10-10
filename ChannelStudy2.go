package main

import (
	"fmt"
	"time"
)

func Counter(out chan<- int) {
	for x := 0; x < 5; x++ {
		out <- x
		time.Sleep(time.Second)
	}
	close(out)
}

func Squarer(out chan<- int, in <-chan int) {
	for x := range in {
		out <- x * x
	}
	close(out)
}

func Printer(in <-chan int) {
	for x := range in {
		fmt.Println(x)
	}
}

func main() {
	naturals := make(chan int)
	squares := make(chan int)

	ch := make(chan string, 3)
	fmt.Println(cap(ch))
	ch <- "piao"
	fmt.Println(len(ch))

	go Counter(naturals)
	go Squarer(squares, naturals)
	Printer(squares)

	fmt.Println("done!")
}
