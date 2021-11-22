package main

import (
	"log"
	"time"
)

func main() {
	bigSlowOperation()
}

func bigSlowOperation() {
	// 因为trace返回的是一个函数
	// defer trace("bigSlowOperation")()
	endTrace := trace("bigSlowOperation")
	defer endTrace()

	time.Sleep(10 * time.Second)
}

func trace(msg string) func() {
	start := time.Now()
	log.Printf("enter %s", msg)
	return func() { log.Printf("exit %s (%s)", msg, time.Since(start)) }
}
