package main

import (
	"fmt"
)

func main() {

	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	// a := 9
	// b := 8

	// for a > b {
	// 	fmt.Println("a > b")
	// }

	var key string
	var val interface{}  // element type of m is assignable to val
	m := map[string]int{"mon":0, "tue":1, "wed":2, "thu":3, "fri":4, "sat":5, "sun":6}
	
	for key, val = range m {
		fmt.Println(key)
		fmt.Println(val)
	}


}