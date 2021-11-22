package main

import (
	"errors"
	"fmt"
	"log"
)

// 练习5.15： 编写类似sum的可变参数函数max和min。
// 考虑不传参时，max和min该如何处理，
// 再编写至少接收1个参数的版本。

func main() {
	input := []int{0, 1, 2, 3, 4}
	first := input[0]
	left := input[1:]
	// 1. max
	m, err := max(input...)
	if err != nil {
		log.Println(err)
	}
	fmt.Printf("max: %d\n", m)

	// 2. min
	m, err = min(input...)
	if err != nil {
		log.Println(err)
	}
	fmt.Printf("min: %d\n", m)

	// 3. mustMax
	m = mustMax(first, left...)
	fmt.Printf("mustMax: %d\n", m)

	// 4. mustMin
	m = mustMin(first, left...)
	fmt.Printf("mustMin: %d\n", m)
}

func max(vals ...int) (int, error) {
	if len(vals) == 0 {
		return 0, errors.New("must have at least one argument!")
	}
	max := vals[0]
	for _, v := range vals {
		if v > max {
			max = v
		}
	}
	return max, nil
}

func min(vals ...int) (int, error) {
	if len(vals) == 0 {
		return 0, errors.New("must have at least one argument!")
	}
	min := vals[0]
	for _, v := range vals {
		if v < min {
			min = v
		}
	}
	return min, nil
}

func mustMax(first int, vals ...int) int {
	max := first
	for _, v := range vals {
		if v > max {
			max = v
		}
	}
	return max
}

func mustMin(first int, vals ...int) int {
	min := first
	for _, v := range vals {
		if v < min {
			min = v
		}
	}
	return min
}
