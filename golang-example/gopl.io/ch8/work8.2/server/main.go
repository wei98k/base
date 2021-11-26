package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net"
	"os"
	"strconv"
	"strings"
)

// 练习 8.2： 实现一个并发FTP服务器。
// 服务器应该解析客户端发来的一些命令，
// 比如cd命令来切换目录，ls来列出目录内文件，
// get和send来传输文件，close来关闭连接。
// 你可以用标准的ftp命令来作为客户端，或者也可以自己实现一个。

func main() {
	listener, err := net.Listen("tcp", "localhost:8888")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
			continue
		}
		go handleConn(conn)
	}
}

func handleConn(conn net.Conn) {
	defer conn.Close()
	sc := bufio.NewScanner(conn)
	cwd := "."
	// 这儿是什么意思?
CLOSE:
	for sc.Scan() {
		args := strings.Fields(sc.Text())
		cmd := args[0]
		switch cmd {
		case "close":
			break CLOSE
		case "ls":
			if len(args) < 2 {
				ls(conn, cwd)
			} else {
				path := args[1]
				if err := ls(conn, path); err != nil {
					fmt.Fprint(conn, err)
				}
			}
		case "cd":
			if len(args) < 2 {
				fmt.Fprintln(conn, "not enough argument")
			} else {
				cwd += "/" + args[1]
			}
		case "get":
			if len(args) < 2 {
				fmt.Fprintln(conn, "not enough argument")
			} else {
				filename := args[1]
				data, err := ioutil.ReadFile(filename)
				if err != nil {
					fmt.Fprint(conn, err)
				}
				fmt.Fprintf(conn, "%s\n", data)
			}
		case "send":
			filename := args[1]
			f, err := os.Create(filename)
			if err != nil {
				fmt.Fprint(conn, err)
			}
			defer f.Close()

			c, err := strconv.Atoi(args[2])
			if err != nil {
				fmt.Fprint(conn, err)
			}
			var texts string
			for i := 0; i < c && sc.Scan(); i++ {
				texts += sc.Text() + "\n"
			}
			texts = strings.TrimSuffix(texts, "\n")

			fmt.Fprint(f, texts)
		}
	}
}

func ls(w io.Writer, dir string) error {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return err
	}
	for _, file := range files {
		fmt.Fprintf(w, "%s\n", file.Name())
	}
	return nil
}
