package interview

import (
	"fmt"
	"strings"
)

func IsUniqueString(s string) bool {
	if strings.Count(s, "") > 3000 {
		return false
	}

	// for range 可以自动把字符转成ascii的吗?
	for _, v := range s {
		fmt.Println(v)
		if v > 127 {
			return false
		}
		if strings.Count(s, string(v)) > 1 {
			return false
		}
	}
	return true
}

func IsUniqueString2(s string) bool {
	if strings.Count(s, "") > 3000 {
		return false
	}
	for k, v := range s {
		if v > 127 {
			return false
		}
		if strings.Index(s, string(v)) != k {
			return false
		}
	}
	return true
}
