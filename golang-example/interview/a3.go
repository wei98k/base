package interview

import (
	"strings"
)

func ReverString(s string) (string, bool) {
	str := []rune(s)

	strLen := len(str)
	if strLen > 5000 {
		return s, false
	}

	for i := 0; i < strLen/2; i++ {
		str[i], str[strLen-1-i] = str[strLen-1-i], str[i]
	}
	return string(str), true
}

func IsRegroup(s1, s2 string) bool {
	sl1 := len([]rune(s1))
	sl2 := len([]rune(s2))

	if sl1 > 5000 || sl2 > 5000 || sl1 != sl2 {
		return false
	}
	for k, v := range s1 {
		// fmt.Printf("字符: %v , %v", s1[k], s2[k])
		if strings.Count(s1, string(v)) != strings.Count(s2, string(v)) {
			return false
		}

		if s1[k] != s2[k] {
			return false
		}
	}
	return true
}
