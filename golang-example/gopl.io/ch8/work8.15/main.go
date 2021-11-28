package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"time"
)

// 练习 8.15：
// 如果一个客户端没有及时地读取数据
// 可能会导致所有的客户端被阻塞。
// 修改broadcaster来跳过一条消息，
// 而不是等待这个客户端一直到其准备好读写。
// 或者为每一个客户端的消息发送channel建立缓冲区，
// 这样大部分的消息便不会被丢掉；
// broadcaster应该用一个非阻塞的send向这个channel中发消息。

type client chan<- string

var (
	entering = make(chan client)
	leaving  = make(chan client)
	messages = make(chan string)
)

func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}

	go broadcaster()
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(conn)
	}
}

func broadcaster() {
	clients := make(map[client]bool)
	for {
		select {
		case msg := <-messages:
			for cli := range clients {
				select {
				case cli <- msg:
				case <-time.After(10 * time.Second):
				}
			}
		case cli := <-entering:
			clients[cli] = true
		case cli := <-leaving:
			delete(clients, cli)
			close(cli)
		}
	}
}

func handleConn(conn net.Conn) {
	ch := make(chan string)
	go clientWriter(conn, ch)

	who := conn.RemoteAddr().String()
	ch <- "You are " + who
	messages <- who + " has arrived"
	entering <- ch

	input := bufio.NewScanner(conn)
	for input.Scan() {
		messages <- who + ": " + input.Text()
	}

	leaving <- ch
	messages <- who + " has left"
	conn.Close()
}

func clientWriter(conn net.Conn, ch <-chan string) {
	for msg := range ch {
		fmt.Fprintln(conn, msg)
	}
}
