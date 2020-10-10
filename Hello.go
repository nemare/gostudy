package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

type Vertex struct {
	Lat, Long float64
}

var m map[string]Vertex

var ptr *int

const (
	a = 1
	c = unsafe.Sizeof(a)
)

// 自定义类型 原因是必须要创建枚举再调用
type StereoType int

const (
	red StereoType = 1 << iota
	_
	_
	blue
	black
)

func main() {

	fmt.Println("Hello World!")
	pow := make([]int, 10)
	for i := range pow {
		pow[i] = 1 << uint(i) // == 2**i
	}
	for _, value := range pow {
		fmt.Printf("%d\n", value)
	}
	m = make(map[string]Vertex, 2)
	m["Bell Labs"] = Vertex{
		40.68433, -74.39967,
	}

	var e = Vertex{1.2, 3.2}
	var f = e
	fmt.Println(&e, &f)

	fmt.Println(m["Bell Labs"])
	fmt.Println(reflect.ValueOf(m))

	var a string = "1"
	var c string = "a"
	var d = a
	fmt.Println(&a)
	fmt.Println(&c)
	fmt.Println(&d)
	fmt.Println(a, c)
	fmt.Println(red, blue, black)

	var z int = 1
	fmt.Println(^z)

	var z1 string
	fmt.Println(z1)

	var x interface{}
	switch i := x.(type) {
	case nil:
		fmt.Println("i:", i)
	case int:
		fmt.Println("x是int32")
	default:

	}

	var ptr1 *int
	fmt.Println("ptr1= ", ptr1)

	var num int
	fmt.Println(num)

	as := []int{1, 2, 3}
	var as1 [3]int
	for i := 0; i < 3; i++ {
		as1[i] = i
	}
	fmt.Println(as1)
	for i, v := range as {
		fmt.Println(i, v)
	}

	var ptrs [10]*int

	for i := 0; i < 3; i++ {
		ptrs[i] = &as[i]
		fmt.Print(*ptrs[i])
	}
	fmt.Println("\n")

	var wos = [...]int{1, 2, 3, 4}
	fmt.Println(len(wos))

	twoDemension := [3][]int{
		{3, 4, 5, 6},
		{2, 3, 5, 6},
		{1, 2, 3, 4},
	}

	fmt.Println(twoDemension[0])

	result := getAverage(twoDemension[0], len(twoDemension))
	fmt.Println(result)
}

func getAverage(arr []int, size int) float32 {
	var result float32
	sum := 0
	for i := 0; i < len(arr); i++ {
		sum += arr[i]
	}

	result = float32(sum / size)
	return result
}
