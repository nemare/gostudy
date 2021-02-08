package main

import (
	"fmt"
	"time"
)

func main() {
	dd := "02"
	fmt.Println(time.Now().Add(24 * time.Hour).Format(dd))
}
