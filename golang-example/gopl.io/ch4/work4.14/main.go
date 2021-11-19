package main

import (
	"example/gopl.io/ch4/work4.14/issue"
	"fmt"
	"log"
	"net/http"
	"text/template"
)

// 练习 4.14： 创建一个web服务器，查询一次GitHub，
// 然后生成BUG报告、里程碑和对应的用户信息。

func main() {
	http.HandleFunc("/", Handler)
	log.Fatal(http.ListenAndServe("localhost:8888", nil))
}

func Handler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		log.Println(err)
	}
	q := r.FormValue("key")
	fmt.Println(q)
	result, err := issue.SearchIssues(q)
	fmt.Println(result)
	if err != nil {
		log.Println(err)
	}
	tmpl := template.Must(template.ParseFiles("index.html"))
	if err := tmpl.Execute(w, result); err != nil {
		log.Println(err)
	}
}
