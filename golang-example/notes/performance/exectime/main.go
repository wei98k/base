package main

import (
	"fmt"
	"time"
)

// 统计一个程序运行的时间
func main() {
	start := time.Now()

	// other more flow code
	time.Sleep(1000 * 3 * time.Millisecond)

	fmt.Printf("%.2fs run time\n", time.Since(start).Seconds())
}
