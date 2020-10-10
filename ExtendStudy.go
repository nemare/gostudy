package main

import "fmt"

type Parent struct {
	val int
}

type Child struct {
	Parent
	val int
}

func (p_child *Child) set(val int) {
	p_child.val = val
}

func main() {

	child := Child{Parent{3}, 4}
	fmt.Println(child.Parent.val)
	fmt.Println(child.val)

	var child1 Child
	fmt.Println(child1.val)

	child.set(5)
	fmt.Println(child.val)
}
