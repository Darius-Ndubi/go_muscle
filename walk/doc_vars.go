package main

import (
	"fmt"
)

// func main() {
// 	// var power int = 9000
// 	power := 9000
// 	fmt.Printf("It's over %d\n", power)
// }

// func main() {
// 	power := getPower()
// 	fmt.Println(power)
// }
// func getPower() int {
// 	return 9001
// }


func main() {
	power := 1000
	fmt.Printf("default power is %d\n", power)
	name, power := "Goku", 9000
	fmt.Printf("%s's power is over %d\n", name, power)
}