package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

//func funname (a interface{}) string {
//	return string(a)
//}
type InterfaceStd struct {
}

type interfaceA interface {
	a()
}

type interfaceB interface {
	b()
}

type interfaceC interface {
	c()
}


type interfaceD interface {
	interfaceA
	interfaceB
}

func main() {
	//var a interface{}
	// 会产生panic 避免
	//fmt.Println("woshi1", a.(string))

	//value, ok := a.(string)
	//if !ok {
	//	fmt.Println("It's not ok for type string")
	//	return
	//}
	//fmt.Println("The value is ", value)

	//var b interface{} = nil
	//fmt.Println(b == nil)
	//
	//value, ok := b.(int)
	//if !ok {
	//	fmt.Println("not ok")
	//	return
	//}
	//fmt.Println(value)
	//
	//var d *InterfaceStd = nil
	//var c interface{} = d
	//fmt.Println(c == nil)

	var w io.Writer
	fmt.Println(w == nil)

	w = os.Stdin
	fmt.Println(w == nil)
	w = new(bytes.Buffer)
	fmt.Println(w == nil)
	w = nil
	fmt.Println(w == nil)

	var x interface{} = []int{1,2,3}
	fmt.Println(x == x)


}
