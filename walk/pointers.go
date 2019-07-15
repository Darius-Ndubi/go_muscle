package main

import (
	"fmt"
)
func main() {
	// var firstName *string = new(string)

	// *firstName = "rio"

	// fmt.Println(*firstName)

	firstName := "Dario"

	pinter := &firstName
	fmt.Println(pinter, *pinter)

}