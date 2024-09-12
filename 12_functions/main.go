package main

import "fmt"

// func add(a int, b int) int {
// 	return a + b
// }

// * returing multiple values
// func getLanguages() (string, string, string)  {
// 	return "JS", "C++", "Py"
// }

func processIt(fn func(a int) int) {
	val := fn(1)
	fmt.Println(val)
}

func main()  {

	// ? passing function as parameter
	fn := func (a int) int {
		return a
	}
	processIt(fn)
	// fmt.Println( )
	
}