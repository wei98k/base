package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

// 练习 8.12： 使broadcaster能够将arrival事件
// 通知当前所有的客户端。
// 这需要你在clients集合中，
// 以及entering和leaving的channel中记录客户端的名字。

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

type client struct {
	ch chan<- string
	// 客户的名字
	name string
}

var (
	entering = make(chan client)
	leaving  = make(chan client)
	messages = make(chan string)
)

func broadcaster() {
	clients := make(map[client]bool)
	for {
		select {
		case msg := <-messages:
			for cli := range clients {
				cli.ch <- msg
			}
		case cli := <-entering:
			clients[cli] = true
			var names []string
			// 搜索名字
			for c := range clients {
				names = append(names, c.name)
			}
			// 发送给新登录的客户
			cli.ch <- fmt.Sprintf("%d arrival: %v\n", len(names), names)

		case cli := <-leaving:
			delete(clients, cli)
			close(cli.ch)

		}
	}
}

func handleConn(conn net.Conn) {
	ch := make(chan string)
	go clientWriter(conn, ch)

	who := conn.RemoteAddr().String()
	ch <- "You are" + who
	messages <- who + " has arrived"
	entering <- client{ch, who}

	input := bufio.NewScanner(conn)
	for input.Scan() {
		messages <- who + ": " + input.Text()
	}

	leaving <- client{ch, who}
	messages <- who + " has left"
	conn.Close()
}

func clientWriter(conn net.Conn, ch <-chan string) {
	for msg := range ch {
		fmt.Fprintln(conn, msg)
	}
}
