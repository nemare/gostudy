package main

import (
	"fmt"
	"time"
)

func main() {
	//conn, err := net.Dial("tcp", "127.0.0.1:8001")
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//done := make(chan string)
	//go func() {
	//	_, _ = io.Copy(os.Stdout, conn)
	//	log.Println("goroutine done...")
	//	done <- "I am done"
	//}()
	//
	//_, _ = io.Copy(conn, os.Stdin)
	//
	//_ = conn.Close()
	//
	//log.Println("main wait goroutine")
	//message := <-done
	//log.Println(message)

	naturals := make(chan int)
	squares := make(chan int)
	go func() {
		for x := 0; x < 5; x++ {
			naturals <- x
			time.Sleep(1 * time.Second)
		}
		close(naturals)
	}()

	go func() {
		//for {
		//	x, ok := <-naturals
		//	if !ok {
		//		break;
		//	}
		//	squares <- x * x
		//}
		for x := range naturals {
			squares <- x * x
		}
		close(squares)

	}()

	for {
		x, ok := <-squares
		if !ok {
			break
		}
		fmt.Println(x)
	}
	fmt.Println("done")

}
