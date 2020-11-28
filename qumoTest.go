package main

import (
	"fmt"
	"time"
)

func main() {
	a := 299068487572983 % 1024
	fmt.Println(a)
	fmt.Println(GetYesterday20(time.Now()))
	config := 618
	var b float64 = float64(config / 100)
	fmt.Println(b)
}

func GetDay20(d time.Time) time.Time {
	return time.Date(d.Year(), d.Month(), d.Day(), 20, 0, 0, 0, d.Location())
}

func GetYesterday20(d time.Time) time.Time {
	yesDay := d.AddDate(0, 0, -1)
	return GetDay20(yesDay)
}
