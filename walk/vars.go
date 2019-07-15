package main

import (
	"fmt"
)

func main() {
	// initialize a variable i an int
	// var i int
	// assing 42 to variable i
	// i=42

	// initialize l a float of type 32 bit
	var l  float32 = 3.142
	// print out the sum of i and l
	fmt.Println(l)

	// Implicit initialization syntax
	firstName := "Darius"
	fmt.Println(firstName)

	// boolean
	b :=false
	fmt.Println(b)

	//
	c := complex(3,4)
	fmt.Println(c)

	r, i := real(c), imag(c)
	fmt.Println(r,i)
}
