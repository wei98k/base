package pipeline

// 练习 9.4: 创建一个流水线程序，
// 支持用channel连接任意数量的goroutine，
// 在跑爆内存之前，可以创建多少流水线阶段？
// 一个变量通过整个流水线需要用多久？（这个练习题翻译不是很确定）

func pipeline(stages int) (chan<- interface{}, <-chan interface{}) {
	if stages < 1 {
		return nil, nil
	}

	in := make(chan interface{})
	out := in

	for i := 0; i < stages; i++ {
		prev := out
		next := make(chan interface{})
		go func() {
			for v := range prev {
				next <- v
			}
			close(next)
		}()
		out = next
	}
	return in, out
}
