package main

import (
	"fmt"
	"reflect"
)

type Adad struct {
	s string
}

func main() {

	a1 := Adad{s: "abc"}
	a2 := Adad{s: "abc"}
	if reflect.DeepEqual(a1, a2) {
		fmt.Println(a1, "==", a2)
	}

	b1 := []int{1, 2}
	b2 := []int{1, 2}
	if reflect.DeepEqual(b1, b2) {
		fmt.Println(b1, "==", b2)

	}

	fmt.Println(fmt.Sprintf("%vdada", b2))

	c1 := map[string]int{"a": 1, "b": 2}
	c2 := map[string]int{"a": 1, "b": 2}
	if reflect.DeepEqual(c1, c2) {
		fmt.Println(c1, "==", c2)
	}

	d1 := map[string]interface{}{
		"a": 1, "b": 0,
	}
	d2 := map[string]interface{}{
		"a": 1, "b": 1,
	}
	fmt.Println(reflect.DeepEqual(d1, d2))

}
