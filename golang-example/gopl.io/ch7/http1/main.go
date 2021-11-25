package main

import (
	"fmt"
	"log"
	"net/http"
)

// 下面这个程序可能是能想到的最简单的实现了。
// 它将库存清单模型化为一个命名为database的map类型，
// 我们给这个类型一个ServeHttp方法，
// 这样它可以满足http.Handler接口。
// 这个handler会遍历整个map并输出物品信息。

func main() {
	// db := database{"shoes": 50, "socks": 5}
	m := myok(1)
	log.Fatal(http.ListenAndServe("localhost:8000", m))
}

// 定义一个dollars类型
type dollars float32

//
func (d dollars) String() string { return fmt.Sprintf("$%.2f", d) }

// 定义模拟数据库的map
// 只要是这个类型实现了这个方法
// 那么这个类型就实现了这个接口?
type database map[string]dollars

//实现Handler接口中的ServerHTTP方法
func (db database) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	for item, price := range db {
		fmt.Fprintf(w, "%s: %s\n", item, price)
	}
}

type myok int16

func (m myok) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "%d", m)
}
