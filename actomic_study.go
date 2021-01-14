package main

import (
	"encoding/json"
	"fmt"
	"sync/atomic"
)

func test1(val *atomic.Value) {
	var m1 map[string]interface{}
	fmt.Printf("[test1] m1 address before unmarshal: %p\n", m1)
	defer val.Store(m1)

	json.Unmarshal([]byte(`{"a":1}`), &m1)
	fmt.Printf("[test1] m1 address after unmarshal : %p\n", m1)
}

func test2(val *atomic.Value) {
	var m2 map[string]interface{}
	fmt.Printf("[test2] m2 address before unmarshal: %p\n", m2)
	defer func() {
		val.Store(m2)
	}()

	json.Unmarshal([]byte(`{"a":1}`), &m2)
	fmt.Printf("[test2] m2 address after unmarshal : %p\n", m2)
}

func test3(val *atomic.Value) {
	m3 := make(map[string]interface{})
	fmt.Printf("[test3] m3 address before unmarshal: %p\n", m3)
	defer val.Store(m3)

	json.Unmarshal([]byte(`{"a":1}`), &m3)
	fmt.Printf("[test3] m3 address after unmarshal : %p\n", m3)
}

func test8() {
	m := map[string]int{"a": 1}
	defer fmt.Println("[test4]", m)

	m["a"] = 2
}

func test5() {
	m := map[string]int{"a": 1}
	defer fmt.Println("[test5]", m["a"])

	m["a"] = 2
}

func test6() {
	m := map[string]map[string]int{"a": {"b": 1}}
	defer fmt.Println("[test6]", m["a"])

	m["a"]["b"] = 2
}

func test7() {
	a := 1
	defer fmt.Println("[test7]", a)

	a++
}

func main() {
	val1 := atomic.Value{}
	test1(&val1)
	fmt.Printf("[test1] val1: %v,\taddress: %p\n", val1.Load(), val1.Load())

	val2 := atomic.Value{}
	test2(&val2)
	fmt.Printf("[test2] val2: %v,\taddress: %p\n", val2.Load(), val2.Load())

	val3 := atomic.Value{}
	test3(&val3)
	fmt.Printf("[test3] val3: %v,\taddress: %p\n", val3.Load(), val3.Load())

	test8()
	test5()
	test6()
	test7()
}
