package main

import "log"

func main() {
    log.Println("welcome my world")
    a()
}


func a() (i int){
    defer func() {
	log.Println(i)
	i++
    }()
    return 1
}
