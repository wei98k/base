package main

import (
	"io"
	"log"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}

	TCPConn := conn.(*net.TCPConn)

	done := make(chan struct{})

	go func() {
		io.Copy(os.Stdout, conn)
		log.Println("done")
		done <- struct{}{}
	}()
	mustCopy(conn, os.Stdin)

	TCPConn.CloseWrite()
	// 只关闭网络连接中写的部分，
	// 这样的话后台goroutine可以
	// 在标准输入被关闭后继续打印从reverb1服务器传回的数据。
	<-done

	TCPConn.CloseRead()
}

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}
