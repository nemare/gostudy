package main

import "fmt"

type NotifyPoint int

const (
	NPEmpty NotifyPoint = 0
)

func main() {
	fmt.Println(NPEmpty == NotifyPoint(6))
}
