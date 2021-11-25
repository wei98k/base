package main

import (
	"bufio"
	"bytes"
	"fmt"
)

func main() {
	s := "Hello"

	var wc WordCounter
	fmt.Fprintf(&wc, s)
	fmt.Println(wc)

	var lc LineCounter
	fmt.Fprintf(&lc, s)
	fmt.Println(lc)
}

type WordCounter int

func (c *WordCounter) Write(p []byte) (int, error) {
	// &取地址 *引用指针
	scanner := bufio.NewScanner(bytes.NewReader(p))
	scanner.Split(bufio.ScanBytes)

	fmt.Println(len(p)) //out: 5

	for scanner.Scan() {
		//*c++
	}
	//?? 这儿怎么会是len(p)
	//A1: io.write接口返回的就应该是写入的字节数
	//A2: 统计单词个数，在调用完writer后使用
	//A3: 看签名定义
	//Q1: *c 和 p 之间是有关系的吗
	//A4: 没有关系
	// *c = *c + 1
	fmt.Println(c) // 当*c++注释后是0否则呢则是5
	return len(p), nil
}

type LineCounter int

func (c *LineCounter) Write(p []byte) (int, error) {
	scanner := bufio.NewScanner(bytes.NewReader(p))
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		*c++
	}
	return len(p), nil
}
