package main

import (
	"io"
	"log"
	"net"
	"time"
)

func main() {
	//监听本机TCP的8000端口
	listen, err := net.Listen("tcp", "localhost:8000")
	//如果错误 终止程序 反馈错误
	if err != nil {
		log.Fatal(err)
	}
	// for 死循环当前线程
	for {
		conn, err := listen.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(conn)
	}
	// Accept()
	// 如果错误 反馈错误 跳过本次

	// 以goroutines启动handle
}

func handleConn(c net.Conn) {
	// defer 关闭IO
	defer c.Close()
	// FOR死循环
	for {
		_, err := io.WriteString(c, time.Now().Format("15:04:05\n"))
		if err != nil {
			return
		}
		time.Sleep(1 * time.Second)
	}
	// 把当前时间写入 net.Conn
	// 延迟1秒钟
}
