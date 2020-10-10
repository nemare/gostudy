package main

import (
	"io/ioutil"
	"os"
)

func main() {
	b, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	err = ioutil.WriteFile("output.txt", b, os.ModeAppend)
	if err != nil {
		panic(err)
	}
}
