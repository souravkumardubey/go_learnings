package main

import "fmt"

// * pass by value
// func changeNum(num int)  {
// 	num = 5
// 	fmt.Println("In changeNum: ", num)
// }

// * pass by reference
func changeNum(num *int) {
	*num = 5
	fmt.Println("Changed value:", *num)
}
func main()  {
	num := 1
	// changeNum(num)	

	fmt.Println("memory Address:", &num)
	changeNum(&num)
	fmt.Println(num)
}