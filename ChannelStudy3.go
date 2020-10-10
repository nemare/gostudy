package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan int)
	defer close(c)

	go func() {
		c <- 3 + 4
		fmt.Println("send 7 to c end")
	}()
	time.Sleep(3 * time.Second)

	i := <-c
	fmt.Println(i)
	fmt.Println("main end.")
}
