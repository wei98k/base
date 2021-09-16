package interview

import (
	"fmt"
	"testing"
)

func TestDistance(t *testing.T) {
	p := Point{1, 2}
	q := Point{4, 6}

	fmt.Println(Distance(p, q)) // 函数调用
	fmt.Println(p.Distance(q))  // 方法调用
}
