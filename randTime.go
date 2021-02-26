package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	delayTime := float64(rand.Int63n(60))
	fmt.Println(delayTime)
}
