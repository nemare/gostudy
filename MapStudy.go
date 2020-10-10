package main

import "fmt"

func main() {

	age := map[string]int{
		"park": 23,
		"jin":  24,
	}
	for _, v := range age {
		fmt.Println(v)
	}

	fmt.Println(age["xu"])
	if v, ok := age["par"]; !ok {
		fmt.Println("not in", v)
	}

	names := make([]string, 0, len(age))
	for k := range age {
		names = append(names, k)
	}
	fmt.Println(names)

	var ages map[string]int
	ages = make(map[string]int, 0)
	fmt.Println(len(ages))
	ages["w"] = 21
	fmt.Println(ages)
}
