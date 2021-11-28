package word

import "unicode"

// 练习 11.3: TestRandomPalindromes测试函数
// 只测试了回文字符串。编写新的随机测试生成器，
// 用于测试随机生成的非回文字符串。

func IsPalindrome(s string) bool {
	var letters []rune
	for _, r := range s {
		if unicode.IsLetter(r) {
			letters = append(letters, unicode.ToLower(r))
		}
	}
	for i := range letters {
		if letters[i] != letters[len(letters)-1-i] {
			return false
		}
	}
	return true
}
