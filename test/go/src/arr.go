// @author: mjw <majianwei168@outlook.com>
// @date: 2021/4/23
// @note: 
package main

import "fmt"

func main() {
	//var q [3]int = [3]int{1, 2, 3}
	//var r [3]int = [3]int{1, 2}
	//fmt.Println(r[1])

	a := [2]int{1, 2}
	b := [...]int{1, 3}
	//d := [3]int{1, 2}

	fmt.Println(a == b)
	//fmt.Println(a == d)
}