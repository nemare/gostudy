package main

import "fmt"

func test4() {
	s := []int{0, 1, 2, 3}
	m := make(map[int]*int)

	for index, value := range s {
		valueCopy := value
		m[index] = &valueCopy
	}

	printMap(m)
}

func printMap(m map[int]*int) {
	for key, value := range m {
		fmt.Printf("map[%v]=%v\n", key, *value)
	}
}

func main() {
	test4()
}
