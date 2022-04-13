package main

import (
	"example/ex-concurrent/skb/v1/core"
	"io"
	"log"
	"net/http"
	"sync"
)

// var skbUrl = make(map[string]string) //模拟数据库-保存URL

type SMap struct {
	sync.RWMutex
	skbUrl map[string]string
}

var baseShortDomain = "http://s.ck/"

// 短链接服务-v1
func main() {

	var mMap *SMap

	mMap = &SMap{
		skbUrl: make(map[string]string),
	}

	// 转换请求
	helloHandler := func(w http.ResponseWriter, req *http.Request) {

		longUrl := req.FormValue("k")
		// fmt.Println(longUrl)
		u := core.ToShortUrl(longUrl)
		io.WriteString(w, baseShortDomain+u)
		// 保存长链接和短连接的关系
		// skbUrl[u] = longUrl
		mMap.writeMap(u, longUrl)
		// fmt.Println(skbUrl)
	}
	//TODO::还原跳转
	http.HandleFunc("/all", func(w http.ResponseWriter, req *http.Request) {
		//fmt.Println(skbUrl)
	})

	http.HandleFunc("/", helloHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))

	// fatal error: concurrent map writes

	// http://tny.im/9xpMH
	// wrk -t5 -c10 http://127.0.0.1:8080?k=http://www.baidu.com/QlLxLlUpMZFzRaXcteNt/hyuoAOdzGlmTDjuhfIKx/QyCpHrjAYcIQpvxDJAiM/
	// wrk -t2 -c2 -s url.lua http://127.0.0.1:8080
}

func (l *SMap) readMap(key string) (string, bool) {
	l.RLock()
	value, ok := l.skbUrl[key]
	l.RUnlock()
	return value, ok
}

func (l *SMap) writeMap(key string, value string) {
	l.Lock()
	l.skbUrl[key] = value
	l.Unlock()
}
