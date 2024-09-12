package main

import "fmt"

func sum(nums ...int) int  {
	total := 0
	for _, num := range nums {
		total += num
	}
	return total
}

func main() {
	
	// * function which can take in n number of parameters are called variadic func
	fmt.Println(sum(3, 4,5 ,6))
}