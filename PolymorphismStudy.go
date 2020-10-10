package main

import "fmt"

type act interface {
	write()
}

type studenta struct {
	name string
}

type studentb struct {
	name string
}

func (p_studenta *studenta) set(name string) {
	p_studenta.name = name
}

func (p_studenta *studenta) write() {
	fmt.Println(p_studenta.name)
}

func (p_studentb *studentb) set(name string) {
	p_studentb.name = name
}

func (p_studentb *studentb) write() {
	fmt.Println(p_studentb.name)
}

func main() {

	var w act
	a := studenta{"piao"}

	b := studentb{"zhikang"}

	w = &a
	w.write()

	w = &b
	w.write()

}
