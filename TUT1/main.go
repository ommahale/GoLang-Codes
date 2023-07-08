package main

import (
	"fmt"
	"strings"
	"tutorial-1/helper" //external module for utility funtions
)

var packageLevel = "This is a package level var"

/*
packageLevel := "This is a package level var" will not work outside functional scope
*/

func main() {
	/*


	 Variable declaration in go and array and slices


	*/
	fmt.Println(packageLevel)
	var arr []string //variable declaration. You must decalre data type for late value allocation
	arr = append(arr, "This", "is", "a", "Loop")
	isGreaterThan3 := len(arr) > 3
	fmt.Printf("Size of arr: %v\n Is greater than 3: %v\n", len(arr), isGreaterThan3)
	/*


	 Simple loop in go


	*/
	for index, ele := range arr {
		fmt.Printf("Index %v : %v\n", index, ele)
	}
	//when we want to ignore a value we simply use _ (eg. for index , _ or for _, ele)
	end := false //dynamic declaration

	/*


	 I / O in go


	*/
	var name string
	fmt.Scanln(&name) // &name gives the reference to the variable name and acts as a ponter to the vvariable
	fmt.Printf("The name with reference: %v\nAddress of the variable name: %v\n", name, &name)
	/*


	 conditonals and flow control


	*/
	for _, ele := range arr {
		if ele == "a" {
			end = true
		}
		fmt.Print(ele, " ")
		if end {
			break
		} else {
			fmt.Println("This is a part of else block")
		}
	}
	/*


	 String methods


	*/
	statement := "This is a statement that can be split using Fields method of string package"
	words := strings.Fields(statement) //methord that splits string to form array based on spaces
	for _, word := range words {
		fmt.Print(word, " ")
	}
	end = false
	helper.SwitchCase(end) //function call
	fmt.Println()
	fmt.Println(statement, "\n", words)

}
