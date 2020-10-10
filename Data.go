package main

import "fmt"

type Data struct {
	val int
}

func (p_data *Data) set(val int) {
	p_data.val = val
}

func (data Data) set1(val int) {
	data.val = val
}

func (p_data *Data) show() {
	fmt.Println(p_data.val)
}

func main() {

	a := &Data{2}
	a.set(4)
	a.show()

	b := Data{3}
	b.set(4)
	b.show()

	c := Data{3}
	c.set1(4)
	c.show()
}
