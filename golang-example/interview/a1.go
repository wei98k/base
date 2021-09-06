package interview

import (
	"fmt"
	"strings"
	"sync"
)

// 交替打印数字和字母
func CrossPrintNumberAndLetter() {

	// 定义两个通道
	letter, number := make(chan bool), make(chan bool)

	// 定义阻塞组
	wait := sync.WaitGroup{}

	go func() {
		i := 1
		for {
			select {
			case <-number:
				fmt.Print(i)
				i++
				fmt.Print(i)
				i++
				// 通知字母协程运行
				letter <- true
				break
			default:
				break
			}
		}
	}()
	// 协程计数器
	wait.Add(1)
	go func(wait *sync.WaitGroup) {
		str := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
		i := 0
		for {
			select {
			case <-letter:
				// 当i的值大于或等于字符的长度时 停止全部协程运行
				if i >= strings.Count(str, "")-1 {
					wait.Done()
					return
				}
				fmt.Print(str[i : i+1])
				i++
				if i >= strings.Count(str, "") {
					i = 0
				}
				fmt.Print(str[i : i+1])
				i++
				number <- true
				break
			default:
				break
			}
		}
	}(&wait)

	// 通过这个number通道告诉 number协程可以运行了
	number <- true

	// 挂载全部协程，等待全部程序结束。
	wait.Wait()
}
