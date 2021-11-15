package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// months := [...]string{1: "January", 12: "December"}
	// fmt.Printf("slice值: %q 容量: %d 长度: %d \n", months, cap(months), len(months))
	a := [...]int{0, 1, 2, 3, 4, 5}
	// a[:] 引用整个数组
	reverse(a[:])
	fmt.Println(a)

	s := []int{0, 1, 2, 3, 4, 5}
	// 参数 0, 1 反转 1 0
	reverse(s[:2]) // 1 0 2 3 4 5
	// 参数 2, 3, 4, 5 反转 5 4 3 2
	reverse(s[2:]) // 1 0 5 4 3 2
	reverse(s)
	fmt.Println(s) // 2 3 4 5 0 1

	input := bufio.NewScanner(os.Stdin)
outer:
	for input.Scan() {
		var ints []int
		for _, s := range strings.Fields(input.Text()) {
			x, err := strconv.ParseInt(s, 10, 64)
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
				continue outer
			}
			ints = append(ints, int(x))
		}
		reverse(ints)
		fmt.Printf("%v\n", ints)
	}

}

func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}
