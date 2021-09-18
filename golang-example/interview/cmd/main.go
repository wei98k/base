package main

import (
	"fmt"
	"math/rand"
	"sync"
)

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

	myMap := new(Map)

	myMap.Out("keya", "ok123")
}

type Map struct {
	c   map[string]*entry
	rmx *sync.RWMutex
}

type entry struct {
	ch      chan struct{}
	value   interface{}
	isExist bool
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

func (m *Map) Out(key string, val interface{}) {
	m.rmx.Lock()
	defer m.rmx.Unlock()

	//if m.c == nil {
	//	m.c = make(map[string]*entry)
	//}

	item, ok := m.c[key]
	if !ok {
		m.c[key] = &entry{
			value:   val,
			isExist: true,
		}
		return
	}
	item.value = val
	if !item.isExist {
		if item.ch != nil {
			close(item.ch)
			item.ch = nil
		}
	}
	return
}
