package main

import ("fmt")

func main() {

	s := []int{0, 1, 2, 3, 4, 5}

	for k, v := range s {
		fmt.Printf("k: %v v: %v\n", k, v)
	}

}

