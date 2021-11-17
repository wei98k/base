package main

import (
	"example/gopl.io/ch4/github"
	"fmt"
	"log"
	"os"
	"time"
)

// 练习 4.10： 修改issues程序，根据问题的时间进行分类，
// 比如不到一个月的、不到一年的、超过一年。

type class string

const (
	LTOM class = "less than one month"
	MTOM class = "more than one month"
	LTOY class = "less than one year"
	MTOY class = "more than one year"
)

func main() {
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}

	issueClass := make(map[class][]github.Issue)
	for _, item := range result.Items {
		item := *item
		y, m, _ := item.CreatedAt.Date()
		cy, cm, _ := time.Now().Date()
		switch {
		case cm-m <= time.Month(1):
			issueClass[LTOM] = append(issueClass[LTOM], item)
		case cm-m > time.Month(1):
			issueClass[MTOM] = append(issueClass[MTOM], item)
		case cy-y <= 1:
			issueClass[LTOY] = append(issueClass[LTOY], item)
		case cy-y > 1:
			issueClass[MTOY] = append(issueClass[MTOY], item)
		}
	}

	for class, issues := range issueClass {
		fmt.Printf("class: %s, issues: %v\n", class, issues)
	}
}
