package main

import (
	"encoding/json"
	"fmt"
)

type Parent struct {
	val int
}

type Child struct {
	Parent
	val int
}

type GrandFather struct {
	Val int
}

func (p_grandFather *GrandFather) set(val int) {
	p_grandFather.Val = val
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

	var grandFathers []*GrandFather
	var grandFathersLocal []*GrandFather
	grandFather := &GrandFather{1}
	for i := 0; i < 3; i++ {
		grandFatherLocal := &GrandFather{i}
		grandFathers = append(grandFathers, grandFather)
		grandFathersLocal = append(grandFathersLocal, grandFatherLocal)
	}

	grandFathersStr, err := json.Marshal(grandFathersLocal)
	if err != nil {
		return
	}
	fmt.Println(string(grandFathersStr))
	grandFather.set(2)
	grandFathersStr, err = json.Marshal(grandFathersLocal)
	if err != nil {
		return
	}
	fmt.Println(string(grandFathersStr))
}
