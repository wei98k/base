package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"sync"
	"time"
)

var (
	downloadDestFolder = "/Users/jw/workspace/base/golang-example/ex-concurrent/crawl/down/tmp"
	urlFilePath        = "/Users/jw/workspace/base/golang-example/ex-concurrent/crawl/down/url.txt"
)

func init() {
	log.SetFlags(log.Lshortfile)
}

func main() {

	start := time.Now()

	fi, err := os.Open(urlFilePath)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}
	defer fi.Close()
	br := bufio.NewReader(fi)
	var w sync.WaitGroup
	for {
		line, _, err := br.ReadLine()
		if err != nil {
			log.Println("read url complete")
			break
		}
		// list := strings.Split(string(line), ",")
		w.Add(1)
		fNmae := path.Base(string(line))
		go download(string(line), fNmae, &w)
	}
	w.Wait()

	fmt.Printf("%.2fs run time\n", time.Since(start).Seconds())
}

func download(url string, filename string, w *sync.WaitGroup) {
	res, err := http.Get(url)
	if err != nil {
		log.Printf("http.Get -> %v", err)
		return
	}
	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Printf("ioutil.ReadAll -> %s", err.Error())
		return
	}
	defer res.Body.Close()
	_ = os.MkdirAll(downloadDestFolder, 0777)
	if err = ioutil.WriteFile(downloadDestFolder+string(filepath.Separator)+filename, data, 0777); err != nil {
		log.Println("Error Saving:", filename, err)
	} else {
		log.Println("Saved:", filename)
	}
	w.Done()
}
