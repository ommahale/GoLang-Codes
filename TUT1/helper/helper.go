package helper

import "fmt"

/*
Functions and arguments in Go
*/

//capitalize the function or variable name to export it/ make it global (eg. switchCase will be package level function)
var MyGlobl = "This is a global variable"

func SwitchCase(end bool) (bool, string) {
	var answer string
	for !end {
		fmt.Println("This is infinite loop")
		fmt.Print("Do you want to end loop(Y for yes N for no):")
		fmt.Scan(&answer)
		switch answer {
		case "Y", "y":
			end = true
		case "N", "n":
			continue
		default:
			continue
		}
	}
	return end, answer
}
