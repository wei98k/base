package ping

// 练习 9.5: 写一个有两个goroutine的程序，
// 两个goroutine会向两个无buffer channel
// 反复地发送ping-pong消息。
// 这样的程序每秒可以支持多少次通信？

import "testing"

func BenchmarkPing(b *testing.B) {
	ch1 := make(chan string)
	ch2 := make(chan string)

	done := make(chan struct{})
	go func() {
		for i := 0; i < b.N; i++ {
			ch1 <- "ping"
			<-ch2
		}
		done <- struct{}{}
	}()

	go func() {
		for i := 0; i < b.N; i++ {
			<-ch1
			ch2 <- "pong"
		}
		done <- struct{}{}
	}()
	<-done
	<-done

	close(ch1)
	close(ch2)
	close(done)
}

// cpu: Intel(R) Core(TM) i5-5350U CPU @ 1.80GHz
// BenchmarkPing-4          2000244              605.1 ns/op
// PASS
// ok      example/gopl.io/ch9/work9.5    1.828s
