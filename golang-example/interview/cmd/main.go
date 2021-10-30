package main

import (
	"fmt"
	"math/rand"
	"sync"
)

type Student struct {
	Name string
}

func main() {

	fmt.Println(&Student{Name: "menglu"} == &Student{Name: "menglu"})
	fmt.Println(Student{Name: "menglu"} == Student{Name: "menglu"})

	// str1 := []string{"a", "b", "c"}
	// str2 := str1[1:]
	// str2[1] = "new"

	// fmt.Println(str1)

	// str2 = append(str2, "z", "x", "y")
	// fmt.Println(str1, str2)
	// s := make([]int, 5)
	// s = append(s, 1, 2, 3)
	// fmt.Println(s)

	// runtime.GOMAXPROCS(1)
	// wg := sync.WaitGroup{}
	// wg.Add(20)
	// for i := 0; i < 10; i++ {
	// 	go func() {
	// 		fmt.Println("i: ", i)
	// 		wg.Done()
	// 	}()
	// }
	// for i := 0; i < 10; i++ {
	// 	go func(i int) {
	// 		fmt.Println("i: ", i)
	// 		wg.Done()
	// 	}(i)
	// }
	// wg.Wait()
}

func RandNumber() {
	out := make(chan int)
	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		defer wg.Done()
		for i := 0; i < 5; i++ {
			out <- rand.Intn(5)
		}
		close(out)
	}()
	go func() {
		defer wg.Done()
		for i := range out {
			fmt.Println(i)
		}
	}()
	wg.Wait()
}
