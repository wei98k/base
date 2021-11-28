package word

import "unicode"

// 我们现在的任务就是修复这些错误。
// 简要分析后发现第一个BUG的原因是我们采用了 byte而不是rune序列，
// 所以像“été”中的é等非ASCII字符不能正确处理。
// 第二个BUG是因为没有忽略空格和字母的大小写导致的。

// 针对上述两个BUG，我们仔细重写了函数：

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
