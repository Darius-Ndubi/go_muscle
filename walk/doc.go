// package main

// func main() {
// 	println("yeey pdt")
// }

package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) !=2 {
		println(len(os.Args))
		os.Exit(1)
	}

	fmt.Println("It's over", os.Args[1])
}
