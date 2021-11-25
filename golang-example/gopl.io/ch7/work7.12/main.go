package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

// 练习 7.12：
// 修改/list的handler让它把输出打印成一个HTML的表格
// 而不是文本。html/template包（§4.6）可能会对你有帮助。

func main() {
	db := database{"shoes": 50, "socks": 5}
	http.HandleFunc("/list", db.list)
	http.HandleFunc("/price", db.price)
	log.Fatal(http.ListenAndServe("localhost:8888", nil))
}

type dollars float32

func (d dollars) String() string { return fmt.Sprintf("$%.2f", d) }

type database map[string]dollars

func (db database) list(w http.ResponseWriter, req *http.Request) {
	tmpl := template.Must(template.ParseFiles("list.html"))
	if err := tmpl.Execute(w, db); err != nil {
		log.Println(err)
	}
}

func (db database) price(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	if price, ok := db[item]; ok {
		fmt.Fprintf(w, "%s\n", price)
	} else {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "no such item: %q\n", item)
	}
}
