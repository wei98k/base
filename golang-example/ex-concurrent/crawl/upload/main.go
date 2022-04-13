package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"math/rand"
	"net/http"
	"os"
	"time"
)

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func init() {
	// rand.Seed(time.Now().UnixNano())
	rand.Seed(time.Now().UnixNano())
}

func uploadHandler(w http.ResponseWriter, r *http.Request) {
	reader, err := r.MultipartReader()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	for {
		part, err := reader.NextPart()
		if err == io.EOF {
			break
		}

		fmt.Printf("FileName=[%s], FormName=[%s]\n", part.FileName(), part.FormName())

		if part.FileName() == "" { //formdata
			data, _ := ioutil.ReadAll(part)
			fmt.Printf("FormData=[%s]\n", string(data))
		} else {
			file, _ := os.Create("./" + part.FileName())
			defer file.Close()
			io.Copy(file, part)
		}
	}
}

func main() {
	http.HandleFunc("/upload", uploadHandler)
	http.ListenAndServe(":8080", nil)
}

func createUrl() {
	// http://www.sdfsdf.com/sdfsdfds/dsfsdfj/kjwrewjrlkjcxcvd

	baseUrl := "http://www.baidu.com/"

	file, _ := os.Create("./url.txt")

	for i := 0; i < 999; i++ {
		var path string
		for j := 0; j < 3; j++ {
			p := RandStringRunes(20)
			path += p + "/"
		}
		f := baseUrl + path + "\n"
		file.WriteString(f)

	}

}

func RandStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	r := string(b)
	return r
}
