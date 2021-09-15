package main

import "fmt"

func main() {

	s := make([]int, 5)
	s = append(s, 1, 2, 3)
	fmt.Println(s)

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
