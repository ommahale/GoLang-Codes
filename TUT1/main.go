package main

import (
	"fmt"
	"tutorial-1/helper"
)

type BaseStruct struct {
	name    string
	surname string
	age     int
}

func (b BaseStruct) getFullName() string {
	return b.name + " " + b.surname
}
func main() {
	structure := BaseStruct{
		name:    "Foo",
		surname: "Barr",
		age:     34,
	}
	fmt.Println(structure.getFullName(), helper.MyGlobl)
}
