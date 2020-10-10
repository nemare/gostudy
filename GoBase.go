package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func main() {
	var i = new(int)
	fmt.Println(reflect.TypeOf(i))
	fmt.Println(*i)
	*i = 2
	fmt.Println(*i)

	p := newInt0()
	q := newInt0()
	fmt.Println(p)
	fmt.Println(q)
	fmt.Println(p == q)

	s := "hello, world"
	fmt.Println(len(s))
	fmt.Println(s[:5])
	fmt.Println(s[7:])

	t := s
	s += ", right fo ot"
	fmt.Println(s)
	fmt.Println(t)

	b := []byte(s)
	fmt.Println(b)

	x := 123
	y := fmt.Sprintf("%d", x)
	fmt.Println(y, strconv.Itoa(x))
	fmt.Println(y, strconv.FormatInt(int64(x), 2))

	var nums = [...]int{1: 2, 2: 3, 3: 5, 4: 8, 5: 9}
	slice1 := nums[0:3]
	slice2 := nums[2:4]
	fmt.Println(&slice1[2])
	fmt.Println(slice2)
	//slice3 := slice2[:6]
	//fmt.Println(slice3)
	slice4 := slice2[:3]
	fmt.Println(slice4)

	a := [...]int{1, 2, 3}
	reverse(a[:])
	c := []int{1, 2, 3}
	reverse(c)
	fmt.Println(c)

	o := [3]int{1, 2, 3}
	u := [...]int{1, 2, 3}
	fmt.Println(o == u)
	r := []int{}
	r = make([]int, 3)
	fmt.Println(len(r), r == nil)
	//r = []int(nil)
	r = nil
	reverse(r)
	fmt.Println(r)
}

func newInt() *int {
	return new(int)
}

func newInt0() *[0]int {
	return new([0]int)
}

func reverse(nums []int) {
	for i, j := 0, len(nums)-1; i < j; i, j = i+1, j-1 {
		nums[i], nums[j] = nums[j], nums[i]
	}
}
