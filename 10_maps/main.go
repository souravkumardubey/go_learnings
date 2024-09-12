package main

import (
	"fmt"
)

func main() {
	mp := make(map[string]int)
	mp["sourav"] = 1
	// fmt.Println(mp["asd"])
	// * delete element
	// delete(mp, "sourav")
	// * clear map
	// clear(mp)
	// fmt.Println((mp))

	// ! value, bool 
	k, ok := mp["sourav"] // ? ok: an idiom just like flag in other lang
	if ok {
		fmt.Println(k)
	}
	// fmt.Println(maps.Equal(mp1, mp2))
}