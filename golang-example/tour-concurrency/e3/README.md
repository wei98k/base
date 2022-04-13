
## 原文

Buffered Channels
Channels can be buffered. Provide the buffer length as the second argument to make to initialize a buffered channel:

ch := make(chan int, 100)
Sends to a buffered channel block only when the buffer is full. Receives block when the buffer is empty.

Modify the example to overfill the buffer and see what happens.

## 译文-By DeeL

缓冲通道
通道可以是缓冲的。提供缓冲区长度作为make的第二个参数，以初始化一个缓冲通道。

ch := make(chan int, 100)
只在缓冲区满时向缓冲通道发送块。当缓冲区为空时接收块。

修改这个例子，使缓冲区过满，看看会发生什么。

会提示 `deadlock`

```
fatal error: all goroutines are asleep - deadlock!

goroutine 1 [chan send]:
main.main()
```

## 相关资料

[官方-Buffered Channels](https://go.dev/tour/concurrency/3)