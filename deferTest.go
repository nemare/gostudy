package main

import (
	"fmt"
	"time"
)

func main() {
	foo()
}

func trace(funcName string) func(){
	start := time.Now()
	fmt.Printf("function %s enter at %s \n",funcName,start)

	return func(){
		fmt.Printf("function %s exit at %s(elapsed %s)",funcName,time.Now(),time.Since(start))
	}
}

func foo(){
	fmt.Printf("foo begin at %s \n",time.Now())

	defer trace("foo")()
	time.Sleep(5*time.Second)

	fmt.Printf("foo end at %s \n",time.Now())
}


