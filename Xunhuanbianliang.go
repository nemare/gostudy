package main

import "fmt"

var slice []func()

func main() {
	sli := []int{1, 2, 3, 4, 5}
	for _, v := range sli {
		// 必须的
		v := v
		fmt.Println(&v)
		slice = append(slice, func() {
			fmt.Println(v * v) // 直接打印结果
		})
	}

	for _, val := range slice {
		val()
	}
	fmt.Printf("%f", 0.1234567)
}
