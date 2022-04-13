package core

import (
	"io/ioutil"
	"strings"
	"testing"
)

func BenchmarkToShortUrl(b *testing.B) {

	// 如何使用自定义数据集？
	// 将数据从文件中读取然后保存到一个slice中
	for i := 0; i < b.N; i++ {
		url := "http://www.baidu.com/QlLxLlUpMZFzRaXcteNt/hyuoAOdzGlmTDjuhfIKx/QyCpHrjAYcIQpvxDJAiM/"
		toShortUrl(url)
	}
}

func TestDb(t *testing.T) {
	// 把整个文件读取到内存中
	f, _ := ioutil.ReadFile("../url.txt")
	// 把内存的字符串转成数组
	strings.Split(string(f), "\n")
}
