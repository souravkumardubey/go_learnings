package main

import "fmt"

func main() {

	// * Go does not have ternary operator
	
	// age := 10
	// if age >= 18 {
	// 	fmt.Println("Person is an adult")
	// } else if age >= 12 {
	// 	fmt.Println("Person is a teenager")
	// } else {
	// 	fmt.Println("Person is a child")
	// }

	// ! variable can be declared inside the if condition and can be using in if-else block
	if age := 18; age >= 18 {
		fmt.Println("Person is an adult")
	} else if age >= 12 {
		fmt.Println("Person is a teenager")
	} else {
		fmt.Println("Person is a child")
	}

}