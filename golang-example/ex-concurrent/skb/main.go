package main

import (
	"crypto/md5"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
)

// type skbUrl struct {
// 	ShortKey string
// 	OrgUrl   string
// }
var skbUrl = make(map[string]string)

//skb-短链接服务
func main() {

	//===== web1.0

	// handel := func(w http.ResponseWriter, req *http.Request) {
	// 	toOrgUrl()
	// }
	http.HandleFunc("/", toShortUrl)
	http.HandleFunc("/05c67da88247ad5968c7dff9de44b184", toOrgUrl)
	log.Fatal(http.ListenAndServe(":1234", nil))

}

func toShortUrl(w http.ResponseWriter, req *http.Request) {
	// 字面创建map
	// skbUrl := map[string]string{}

	// 使用make创建map

	// http://www.majianwei.com/archives/docs/golang/qa/%e9%85%8d%e7%bd%ae%e4%bb%a3%e7%90%86
	url1 := "http://www.majianwei.com/archives/docs/golang/qa/%e9%85%8d%e7%bd%ae%e4%bb%a3%e7%90%86"
	h := md5.New()
	io.WriteString(h, url1)
	m1 := h.Sum(nil)

	// skbUrl[string(m1)] = url1

	skbUrl["05c67da88247ad5968c7dff9de44b184"] = url1

	click := fmt.Sprintf("<a href='http://127.0.0.1:1234/%x'>点击转向原URL</a>", string(m1))

	echo := fmt.Sprintf("shortKey: %x\norgUrl: %s\n%s", string(m1), skbUrl[string(m1)], click)
	// fmt.Println(echo)
	io.WriteString(w, echo)

	// fmt.Printf("md5: %x num:%d\n", m1, utf8.RuneCountInString(string(m1))) // 32 16 * 2 = 32
}

func toOrgUrl(w http.ResponseWriter, req *http.Request) {
	k := strings.TrimLeft(req.RequestURI, "/")
	fmt.Println(string(k))
	fmt.Println(skbUrl)
	http.Redirect(w, req, skbUrl[k], http.StatusFound)
}
