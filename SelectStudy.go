package main

import (
	"fmt"
	"os"
	"time"
)

func launch() {
	fmt.Println("发射！")
}

func abort(in chan<- struct{}) {
	os.Stdin.Read(make([]byte, 1))
	in <- struct{}{}
}

func main() {
	fmt.Println("准备发射火箭！")
	tick := time.Tick(time.Second)
	ab := make(chan struct{})
	go abort(ab)

	for countdown := 5; countdown > 0; countdown-- {
		fmt.Println(countdown)
		select {
		case <-tick:
		case <-ab:
			fmt.Println("发射终止...")
			return

		}
	}

	launch()

}
