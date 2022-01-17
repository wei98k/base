## Overview

Package strings implements simple functions to manipulate UTF-8 encoded strings.

For information about UTF-8 strings in Go, see https://blog.golang.org/strings.

Package strings实现了简单的函数来操作UTF-8编码的字符串。

关于Go中UTF-8字符串的信息，请参见https://blog.golang.org/strings。

## Examples

[Examples-url](https://pkg.go.dev/strings@go1.17.6#pkg-examples)

## Functions

### func Compare

`func Compare(a, b string) int`

Compare returns an integer comparing two strings lexicographically. The result will be 0 if a==b, -1 if a < b, and +1 if a > b.

Compare is included only for symmetry with package bytes. It is usually clearer and always faster to use the built-in string comparison operators ==, <, >, and so on.

比较返回一个整数，按字母顺序比较两个字符串。如果a==b，结果是0；如果a<b，结果是-1；如果a>b，结果是+1。

包含Compare只是为了与package bytes对称。通常使用内置的字符串比较运算符==、<、>等会更清晰、更快速。

### func Contains

`func Contains(s, substr string) bool`

Contains reports whether substr is within s.

包含报告substr是否在s内。

### func ContainsAny 

`func ContainsAny(s, chars string) bool`

ContainsAny reports whether any Unicode code points in chars are within s.

ContainsAny报告字符中是否有Unicode代码点在s内。

### func ContainsRune

`func ContainsRune(s string, r rune) bool`

ContainsRune reports whether the Unicode code point r is within s.

ContainsRune报告Unicode代码点r是否在s内。

### func Count

`func Count(s, substr string) int`

Count counts the number of non-overlapping instances of substr in s. If substr is an empty string, Count returns 1 + the number of Unicode code points in s.

Count计算s中substr的非重叠实例的数量。如果substr是一个空字符串，Count返回1+s中Unicode代码点的数量。

### func EqualFold 

`func EqualFold(s, t string) bool`

EqualFold reports whether s and t, interpreted as UTF-8 strings, are equal under Unicode case-folding, which is a more general form of case-insensitivity.

EqualFold报告s和t，解释为UTF-8字符串，在Unicode大小写折叠下是否相等，这是一种更普遍的大小写不敏感的形式。

### func Fields

`func Fields(s string) []string`

Fields splits the string s around each instance of one or more consecutive white space characters, as defined by unicode.IsSpace, returning a slice of substrings of s or an empty slice if s contains only white space.

Fields将字符串s围绕一个或多个连续的空白字符的实例进行分割，正如unicode.IsSpace所定义的那样，返回s的子字符串的片断，如果s只包含空白，则返回一个空片断。

### func FieldsFunc 

`func FieldsFunc(s string, f func(rune) bool) []string`

FieldsFunc splits the string s at each run of Unicode code points c satisfying f(c) and returns an array of slices of s. If all code points in s satisfy f(c) or the string is empty, an empty slice is returned.

FieldsFunc makes no guarantees about the order in which it calls f(c) and assumes that f always returns the same value for a given c.

FieldsFunc在每一个满足f(c)的Unicode代码点c的运行处对字符串s进行分割，并返回s的一个数组切片。

FieldsFunc不保证它调用f(c)的顺序，并假设f对给定的c总是返回相同的值。

### func HasPrefix

`func HasPrefix(s, prefix string) bool`

HasPrefix tests whether the string s begins with prefix.

HasPrefix测试字符串s是否以前缀开始。

### func HasSuffix 

`func HasSuffix(s, suffix string) bool`

HasSuffix tests whether the string s ends with suffix.

HasSuffix测试字符串s是否以后缀结束。

### func Index

`func Index(s, substr string) int`

Index returns the index of the first instance of substr in s, or -1 if substr is not present in s.

Index返回s中substr的第一个实例的索引，如果s中不存在substr，则返回-1。

### func IndexAny

`func IndexAny(s, chars string) int`

IndexAny returns the index of the first instance of any Unicode code point from chars in s, or -1 if no Unicode code point from chars is present in s.

IndexAny返回s中chars的任何Unicode码位的第一个实例的索引，如果s中没有chars的Unicode码位，则返回-1。


### func IndexByte

`func IndexByte(s string, c byte) int`

IndexByte returns the index of the first instance of c in s, or -1 if c is not present in s.

IndexByte返回s中c的第一个实例的索引，如果c在s中不存在，则返回-1。

### func IndexFunc

`func IndexFunc(s string, f func(rune) bool) int`

IndexFunc returns the index into s of the first Unicode code point satisfying f(c), or -1 if none do.

IndexFunc返回满足f(c)的第一个Unicode码位在s中的索引，如果没有，则返回-1。

### func IndexRune

`func IndexRune(s string, r rune) int`

IndexRune returns the index of the first instance of the Unicode code point r, or -1 if rune is not present in s. If r is utf8.RuneError, it returns the first instance of any invalid UTF-8 byte sequence.

IndexRune返回Unicode代码点r的第一个实例的索引，如果s中不存在符文，则返回-1。如果r是utf8.RuneError，则返回任何无效的UTF-8字节序列的第一个实例。

### func Join

`func Join(elems []string, sep string) string`

Join concatenates the elements of its first argument to create a single string. The separator string sep is placed between elements in the resulting string.

Join将其第一个参数中的元素连接起来，创建一个单一的字符串。分隔符sep被放在结果字符串的元素之间。

### func LastIndex

`func LastIndex(s, substr string) int`

LastIndex returns the index of the last instance of substr in s, or -1 if substr is not present in s.

LastIndex返回s中substr的最后一个实例的索引，如果s中不存在substr，则返回-1。

### func LastIndexAny

`func LastIndexAny(s, chars string) int`

LastIndexAny returns the index of the last instance of any Unicode code point from chars in s, or -1 if no Unicode code point from chars is present in s.

LastIndexAny返回s中chars的任何Unicode码位的最后一个实例的索引，如果s中没有chars的Unicode码位，则返回-1。

### func LastIndexByte

`func LastIndexByte(s string, c byte) int`

LastIndexByte returns the index of the last instance of c in s, or -1 if c is not present in s.

LastIndexByte返回s中c的最后一个实例的索引，如果c在s中不存在，则返回-1。

### func LastIndexFunc

`func LastIndexFunc(s string, f func(rune) bool) int`

LastIndexFunc returns the index into s of the last Unicode code point satisfying f(c), or -1 if none do.

LastIndexFunc返回满足f(c)的最后一个Unicode码位在s中的索引，如果没有，则返回-1。

### func Map

`func Map(mapping func(rune) rune, s string) string`

Map returns a copy of the string s with all its characters modified according to the mapping function. If mapping returns a negative value, the character is dropped from the string with no replacement.

Map返回一个字符串s的副本，并根据映射函数修改其所有字符。如果映射返回一个负值，则该字符将从字符串中删除，不作替换。

### func Repeat

`func Repeat(s string, count int) string`

Repeat returns a new string consisting of count copies of the string s.

It panics if count is negative or if the result of (len(s) * count) overflows.

Repeat返回一个新的字符串，由字符串s的count副本组成。

如果count为负数，或者(len(s) * count)的结果溢出，它就会出现恐慌。

### func Replace

`func Replace(s, old, new string, n int) string`

Replace returns a copy of the string s with the first n non-overlapping instances of old replaced by new. If old is empty, it matches at the beginning of the string and after each UTF-8 sequence, yielding up to k+1 replacements for a k-rune string. If n < 0, there is no limit on the number of replacements.

替换返回一个字符串s的副本，其中前n个不重叠的old实例被new替换。如果old是空的，它将在字符串的开头和每个UTF-8序列之后进行匹配，对于一个k-rune字符串，最多产生k+1次替换。如果n<0，则对替换的数量没有限制。

### func ReplaceAll 

`func ReplaceAll(s, old, new string) string`

ReplaceAll returns a copy of the string s with all non-overlapping instances of old replaced by new. If old is empty, it matches at the beginning of the string and after each UTF-8 sequence, yielding up to k+1 replacements for a k-rune string.

ReplaceAll返回一个字符串s的副本，其中所有不重叠的old实例被new替换。如果old是空的，它将在字符串的开头和每个UTF-8序列之后进行匹配，对于一个k-rune字符串，最多产生k+1次替换。

### func Split

`func Split(s, sep string) []string`

Split slices s into all substrings separated by sep and returns a slice of the substrings between those separators.

If s does not contain sep and sep is not empty, Split returns a slice of length 1 whose only element is s.

If sep is empty, Split splits after each UTF-8 sequence. If both s and sep are empty, Split returns an empty slice.

It is equivalent to SplitN with a count of -1.

Split将s切成由sep分隔的所有子字符串，并返回这些分隔符之间的子字符串的一个片断。

如果s不包含sep并且sep不是空的，Split返回一个长度为1的片断，其唯一元素是s。

如果sep是空的，Split在每个UTF-8序列之后进行分割。如果s和sep都是空的，Split返回一个空的片断。

它等同于SplitN，计数为-1。

### func SplitAfter

`func SplitAfter(s, sep string) []string`

SplitAfter slices s into all substrings after each instance of sep and returns a slice of those substrings.

If s does not contain sep and sep is not empty, SplitAfter returns a slice of length 1 whose only element is s.

If sep is empty, SplitAfter splits after each UTF-8 sequence. If both s and sep are empty, SplitAfter returns an empty slice.

It is equivalent to SplitAfterN with a count of -1.

SplitAfter将s切成每个sep实例之后的所有子字符串，并返回这些子字符串的一个片断。

如果s不包含sep并且sep不是空的，SplitAfter返回一个长度为1的片断，其唯一元素是s。

如果sep是空的，SplitAfter在每个UTF-8序列之后进行分割。如果s和sep都是空的，SplitAfter返回一个空的片断。

它等同于SplitAfterN，计数为-1。

### func SplitAfterN

`func SplitAfterN(s, sep string, n int) []string`

SplitAfterN slices s into substrings after each instance of sep and returns a slice of those substrings.

The count determines the number of substrings to return:

SplitAfterN在每个sep的实例之后将s切成子串，并返回这些子串的一个片断。

Count决定了要返回的子串的数量。

### func SplitN

`func SplitN(s, sep string, n int) []string`

SplitN slices s into substrings separated by sep and returns a slice of the substrings between those separators.

The count determines the number of substrings to return:

SplitN将s切成由sep分隔的子串，并返回这些分隔符之间的子串的一个片断。

Count决定了要返回的子串的数量。


### func Title

`func Title(s string) string`

Title returns a copy of the string s with all Unicode letters that begin words mapped to their Unicode title case.

BUG(rsc): The rule Title uses for word boundaries does not handle Unicode punctuation properly.

标题返回一个字符串s的副本，其中包括所有以单词开头的Unicode字母，并映射到其Unicode标题大小写。

BUG(rsc)。Title用于单词边界的规则不能正确处理Unicode标点符号。

### func ToLower

`func ToLower(s string) string`

ToLower returns s with all Unicode letters mapped to their lower case.

ToLower返回s，所有Unicode字母都被映射为小写。

### func ToLowerSpecial

`func ToLowerSpecial(c unicode.SpecialCase, s string) string`

ToLowerSpecial returns a copy of the string s with all Unicode letters mapped to their lower case using the case mapping specified by c.

ToLowerSpecial返回一个字符串s的副本，所有Unicode字母都使用c指定的大小写映射被映射为小写。

### func ToTitle

`func ToTitle(s string) string`

ToTitle returns a copy of the string s with all Unicode letters mapped to their Unicode title case.

ToTitle返回一个字符串s的副本，所有Unicode字母都被映射到Unicode标题大小写。

### func ToTitleSpecial

`func ToTitleSpecial(c unicode.SpecialCase, s string) string`

ToTitleSpecial returns a copy of the string s with all Unicode letters mapped to their Unicode title case, giving priority to the special casing rules.

ToTitleSpecial返回一个字符串s的副本，所有的Unicode字母都被映射到它们的Unicode标题大小写，优先考虑特殊大小写规则。

### func ToUpper

`func ToUpper(s string) string`

ToUpper returns s with all Unicode letters mapped to their upper case.

ToUpper返回s，所有Unicode字母都被映射为大写。

### func ToUpperSpecial

`func ToUpperSpecial(c unicode.SpecialCase, s string) string`

ToUpperSpecial returns a copy of the string s with all Unicode letters mapped to their upper case using the case mapping specified by c.

ToUpperSpecial返回一个字符串s的副本，所有Unicode字母使用c指定的大小写映射被映射为大写。

### func ToValidUTF8

`func ToValidUTF8(s, replacement string) string`

ToValidUTF8 returns a copy of the string s with each run of invalid UTF-8 byte sequences replaced by the replacement string, which may be empty.

ToValidUTF8返回一个字符串s的副本，每个运行中的无效UTF-8字节序列被替换的字符串所取代，该字符串可能是空的。

### func Trim

`func Trim(s, cutset string) string`

Trim returns a slice of the string s with all leading and trailing Unicode code points contained in cutset removed.

Trim返回字符串s的片断，删除cutet中包含的所有前面和后面的Unicode代码点。

### func TrimFunc 

`func TrimFunc(s string, f func(rune) bool) string`

TrimFunc returns a slice of the string s with all leading and trailing Unicode code points c satisfying f(c) removed.

TrimFunc返回一个字符串s的片断，其中删除了所有领先和落后的Unicode代码点c，满足f(c)。

### func TrimLeft

`func TrimLeft(s, cutset string) string`

TrimLeft returns a slice of the string s with all leading Unicode code points contained in cutset removed.

To remove a prefix, use TrimPrefix instead.

TrimLeft返回字符串s的一个片断，删除cutet中包含的所有领先的Unicode代码点。

要删除前缀，请使用TrimPrefix。

### func TrimLeftFunc 

`func TrimLeftFunc(s string, f func(rune) bool) string`

TrimLeftFunc returns a slice of the string s with all leading Unicode code points c satisfying f(c) removed.

TrimLeftFunc返回字符串s的一个片断，所有领先的Unicode代码点c满足f(c)被移除。

### func TrimPrefix

`func TrimPrefix(s, prefix string) string`

TrimPrefix returns s without the provided leading prefix string. If s doesn't start with prefix, s is returned unchanged.

TrimPrefix返回没有提供前缀字符串的s。如果s不以前缀开始，s将被原样返回。

### func TrimRight

`func TrimRight(s, cutset string) string`

TrimRight returns a slice of the string s, with all trailing Unicode code points contained in cutset removed.

To remove a suffix, use TrimSuffix instead.

TrimRight返回字符串s的一个片断，删除cutet中包含的所有尾部Unicode代码点。

要删除后缀，请使用TrimSuffix。

### func TrimRightFunc

`func TrimRightFunc(s string, f func(rune) bool) string`

TrimRightFunc returns a slice of the string s with all trailing Unicode code points c satisfying f(c) removed.

TrimRightFunc返回字符串s的一个片断，所有尾部的Unicode代码点c满足f(c)的删除。

### func TrimSpace

`func TrimSpace(s string) string`

TrimSpace returns a slice of the string s, with all leading and trailing white space removed, as defined by Unicode.

TrimSpace返回字符串s的一个片断，按照Unicode的定义，去除所有前导和尾部的白色空间。

### func TrimSuffix

`func TrimSuffix(s, suffix string) string`

TrimSuffix returns s without the provided trailing suffix string. If s doesn't end with suffix, s is returned unchanged.

TrimSuffix返回s，不包括提供的后缀字符串。如果s不以后缀结尾，s将被原样返回。

## type Builder

```
type Builder struct {
	// contains filtered or unexported fields
}
```

A Builder is used to efficiently build a string using Write methods. It minimizes memory copying. The zero value is ready to use. Do not copy a non-zero Builder.

构建器用于使用Write方法有效地构建一个字符串。它最大限度地减少了内存的复制。零值是可以使用的。不要复制一个非零的生成器。




























