package main

import(
	"fmt"
)

// Q: 如何判断变量是引用还是不是引用(是指针还是非指针)
// A: 如果是引用传值那执行的内存地址就是一致的， 如果不一致那就是值传递

// Q: 如何把值传递改成引用传递呢？

// 引用和值传递有什么区别？

func main() {
	a := []int{1,2,3}

	c := append(a, 9)
	fmt.Printf("c-p: %p c-v: %v\n", c, c)
	b := "hello"
	fmt.Printf("a-p: %p b-p: %p \n", a, &b)
	a1(a, c, &b)

}

func a1(a, c []int, b *string) {

	pv1 := reflect.ValueOf(a)

	

	d := append(a, 8)
	fmt.Printf("a-p: %p b-p: %p c-p: %p d-p: %p \n", a, b, c, d)
}