package main

import "fmt"

// 练习 3.12： 编写一个函数，判断两个字符串是否是相互打乱的，
// 也就是说它们有着相同的字符，但是对应不同的顺序。

// 思路
// 1 判断两个字符串长度一致
// 2 判断两个字符串是否完全相等
// 3 判断两个字符串中的每个字符出现的字符数量是否一致
func main() {
	// a := "abcjjjklll" b := "abcjjjlllk" 返回true
	fmt.Println(isAnagram("abcd", "abdc"))

}

func isAnagram(s1, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}

	if s1 == s2 {
		return false
	}

	m := make(map[rune]int, len(s1))

	for _, v := range s1 {
		m[v]++
	}

	for _, v := range s2 {
		if m[v] == 0 {
			return false
		}
		m[v]--
	}
	return true
}
