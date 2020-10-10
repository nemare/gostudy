package main

import (
	"fmt"
	"time"
)

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}

	c <- sum
}

func main() {

	s := []int{7, 2, 8, -9, 4, 0}
	c := make(chan int)
	go sum(s[len(s)/2:], c)
	go sum(s[:len(s)/2], c)

	x, y := <-c, <-c
	fmt.Println(x, y, x+y)

	go func() {
		time.Sleep(time.Hour)
	}()

	d := make(chan int, 10)
	close(d)
	i, ok := <-d
	fmt.Println(i, ok)
}
