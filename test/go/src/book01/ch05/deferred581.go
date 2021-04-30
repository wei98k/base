// @author: mjw <majianwei168@outlook.com>
// @date: 2021/4/25
// @note: 
package main

import "log"

func main() {
	a()
}

func a() int {
	defer b()
	log.Println("this is a action")
	return 1
}

func b() string {
	log.Println("welcome my world")
	return "welcome my world"
}
