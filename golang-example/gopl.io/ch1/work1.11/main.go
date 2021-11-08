package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

// 练习 1.11： 在fetchall中尝试使用长一些的参数列表，
// 比如使用在alexa.com的上百万网站里排名靠前的。
// 如果一个网站没有回应，程序将采取怎样的行为？
// （Section8.9 描述了在这种情况下的应对机制）。

func main() {
	start := time.Now()
	ch := make(chan string)
	result := getSomeURL(100)
	for _, url := range result {
		go fetch(url, ch)
	}
	for range result {
		fmt.Println(<-ch)
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func getSomeURL(n int) []string {
	f, err := os.Open("top-1m.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	result := make([]string, 0, n)
	r := csv.NewReader(f)
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		if n == 0 {
			break
		}
		result = append(result, record[1])
		n--
	}
	return result
}

func fetch(url string, ch chan<- string) {
	if !strings.HasPrefix(url, "http://") {
		url = "http://" + url
	}
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}
	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	resp.Body.Close()
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs %7d %s", secs, nbytes, url)
}
