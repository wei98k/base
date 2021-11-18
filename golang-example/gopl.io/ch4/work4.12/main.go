package main

import (
	"encoding/json"
	"example/gopl.io/ch4/work4.12/xkcd"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

// 练习 4.12： 流行的web漫画服务xkcd也提供了JSON接口。
// 例如，一个 https://xkcd.com/571/info.0.json
// 请求将返回一个很多人喜爱的571编号的详细描述。
// 下载每个链接（只下载一次）然后创建一个离线索引。
// 编写一个xkcd工具，使用这些离线索引，
// 打印和命令行输入的检索词相匹配的漫画的URL。

// 思路:
// 1. 获取URL响应数据
// 2. 缓存数据
// 3. 搜索的时候 将缓存JSON数据转成结构体 然后通过for range遍历找到的数据放到MAP中然后返回全部找到的数据

// fetch: go run main.go -f -n=100 > in.json
// search: cat in.json | go run main.go keywords #报告错误、待修复 2021/11/18 15:40:51 invalid character 'h' after top-level value
var (
	f = flag.Bool("f", false, "")
	n = flag.Int("n", 100, "")
)

func main() {
	flag.Parse()
	if *f {
		if *n > xkcd.MaxNum {
			log.Fatalf("%d can't bigger than %d", *n, xkcd.MaxNum)
		}
		fetch(*n)
	} else {
		search(flag.Args())
	}
}

func fetch(n int) {
	index := xkcd.New()
	for num := xkcd.MinNum; num < n; num++ {
		c, err := xkcd.Get(num)
		if err != nil {
			log.Fatal(err)
		}
		index.Comics = append(index.Comics, c)
	}
	out, err := json.MarshalIndent(index, "", "    ")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", out)
}

func search(keywords []string) {
	in, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}
	index := xkcd.New()
	if err := json.Unmarshal(in, &index); err != nil {
		log.Fatal(err)
	}
	result := xkcd.Search(index, keywords)
	for _, c := range result {
		fmt.Println(c)
	}
}
