package main

import "fmt"

// ! for -> only construct in go for looping
func main() {

	// ? while implementation with for 
	// i := 1
	// for i <= 10 {
	// 	fmt.Println(i)
	// 	i += 1
	// }

	// ? classic for loop in Go
	// for i := 0; i < 3; i++ {
	// 	fmt.Print(i, " ")
	// }

	// ? for using range
	for i := range(3) { // -> 0 to n - 1
		fmt.Print(i, " ")
	}

}