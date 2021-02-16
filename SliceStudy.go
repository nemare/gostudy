package main

import "fmt"

func main() {
	arr := []int{1, 2, 3}
	s := arr[:]
	fmt.Println(s)

	s1 := make([]int, 3, 4)
	fmt.Println(s1)
	s1[0] = 2
	fmt.Println(s1)
	s1 = append(s1, 1, 2)
	fmt.Println(s1)

	var s2 = make([]int, 4)
	copied := copy(s2, s1)

	fmt.Println(copied, s1)

	s22 := "12"
	fmt.Println(s22[:0])

	var bytes []byte
	fmt.Println(len(string(bytes)))

	filteredSubFactorInfos := make([]int, 0)
	for i := 0; i < 29; i++ {

		filteredSubFactorInfos = append(filteredSubFactorInfos, i)
	}

	fmt.Println(len(filteredSubFactorInfos))
	fmt.Println(filteredSubFactorInfos[:32])
}
