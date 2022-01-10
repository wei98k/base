## Overview

Package bytes implements functions for the manipulation of byte slices. It is analogous to the facilities of the strings package.

包byte实现了操作字节片的功能。它类似于字符串包的设施。

## Examples

[Examples](https://pkg.go.dev/bytes@go1.17.6#pkg-examples)

## Constants

`const MinRead = 512`

MinRead is the minimum slice size passed to a Read call by Buffer.ReadFrom. As long as the Buffer has at least MinRead bytes beyond what is required to hold the contents of r, ReadFrom will not grow the underlying buffer.

MinRead是Buffer.ReadFrom传递给Read调用的最小分片大小。只要Buffer至少有MinRead字节超过容纳r内容的要求，ReadFrom就不会增加底层的缓冲区。


## Variables

```
var ErrTooLarge = errors.New("bytes.Buffer: too large")
```

ErrTooLarge is passed to panic if memory cannot be allocated to store data in a buffer.

如果不能分配内存来存储缓冲区中的数据，ErrTooLarge将被传递给恐慌。

## Functions

### func Compare

`func Compare(a, b []byte) int`

Compare returns an integer comparing two byte slices lexicographically. The result will be 0 if a==b, -1 if a < b, and +1 if a > b. A nil argument is equivalent to an empty slice.

比较返回一个整数，按字母顺序比较两个字节片。如果a==b，结果是0；如果a<b，结果是-1；如果a>b，结果是+1。

### func Contains

`func Contains(b, subslice []byte) bool`

Contains reports whether subslice is within b.

包含报告子片是否在b内。

### func ContainsAny

`func ContainsAny(b []byte, chars string) bool`

ContainsAny reports whether any of the UTF-8-encoded code points in chars are within b.

ContainsAny报告chars中是否有UTF-8编码的代码点在b内。

### func ContainsRune

`func ContainsRune(b []byte, r rune) bool`

ContainsRune reports whether the rune is contained in the UTF-8-encoded byte slice b.

ContainsRune报告符文是否包含在UTF-8编码的字节片b中。

### func Count

`func Count(s, sep []byte) int`

Count counts the number of non-overlapping instances of sep in s. If sep is an empty slice, Count returns 1 + the number of UTF-8-encoded code points in s.

如果sep是一个空片，Count返回1 + s中UTF-8编码的码位数。

### func Equal 

`func Equal(a, b []byte) bool`

Equal reports whether a and b are the same length and contain the same bytes. A nil argument is equivalent to an empty slice.

Equal报告a和b是否相同的长度和包含相同的字节。一个nil参数等同于一个空的片断。

### func EqualFold

`func EqualFold(s, t []byte) bool`

EqualFold reports whether s and t, interpreted as UTF-8 strings, are equal under Unicode case-folding, which is a more general form of case-insensitivity.

EqualFold报告s和t，解释为UTF-8字符串，在Unicode大小写折叠下是否相等，这是一种更普遍的大小写不敏感的形式。


### func Fields

`func Fields(s []byte) [][]byte`

Fields interprets s as a sequence of UTF-8-encoded code points. It splits the slice s around each instance of one or more consecutive white space characters, as defined by unicode.IsSpace, returning a slice of subslices of s or an empty slice if s contains only white space.

Fields将s解释为一串UTF-8编码的代码点。它在一个或多个连续的空白字符的每个实例周围分割s片，正如unicode.IsSpace所定义的那样，返回s片的子片，如果s片只包含空白，则返回一个空片。


### func FieldsFunc

`func FieldsFunc(s []byte, f func(rune) bool) [][]byte`

FieldsFunc interprets s as a sequence of UTF-8-encoded code points. It splits the slice s at each run of code points c satisfying f(c) and returns a slice of subslices of s. If all code points in s satisfy f(c), or len(s) == 0, an empty slice is returned.

FieldsFunc makes no guarantees about the order in which it calls f(c) and assumes that f always returns the same value for a given c.

FieldsFunc将s解释为一串UTF-8编码的代码点。如果s中的所有代码点都满足f(c)，或者len(s)==0，就会返回一个空片。

FieldsFunc不保证它调用f(c)的顺序，并假定f对给定的c总是返回相同的值。


### func HasPrefix

`func HasPrefix(s, prefix []byte) bool`

HasPrefix tests whether the byte slice s begins with prefix.

HasPrefix测试字节片s是否以前缀开始。

### func HasSuffix

`func HasSuffix(s, suffix []byte) bool`

HasSuffix tests whether the byte slice s ends with suffix.

HasSuffix测试字节片s是否以后缀结束。


### func Index

`func Index(s, sep []byte) int`

Index returns the index of the first instance of sep in s, or -1 if sep is not present in s.

Index返回s中sep的第一个实例的索引，如果s中没有sep，则返回-1。

### func IndexAny

`func IndexAny(s []byte, chars string) int`

IndexAny interprets s as a sequence of UTF-8-encoded Unicode code points. It returns the byte index of the first occurrence in s of any of the Unicode code points in chars. It returns -1 if chars is empty or if there is no code point in common.

IndexAny将s解释为一串UTF-8编码的Unicode代码点。它返回s中任何一个Unicode代码点在chars中第一次出现的字节索引。如果chars为空或者没有共同的码位，则返回-1。


### func IndexByte

`func IndexByte(b []byte, c byte) int`

IndexByte returns the index of the first instance of c in b, or -1 if c is not present in b.

IndexByte返回b中c的第一个实例的索引，如果c在b中不存在，则返回-1。

### func IndexFunc 

`func IndexFunc(s []byte, f func(r rune) bool) int`

IndexFunc interprets s as a sequence of UTF-8-encoded code points. It returns the byte index in s of the first Unicode code point satisfying f(c), or -1 if none do.

IndexFunc将s解释为一串UTF-8编码的代码点。它返回s中第一个满足f(c)的Unicode编码点的字节索引，如果没有，则返回-1。

### func IndexRune

`func IndexRune(s []byte, r rune) int`

IndexRune interprets s as a sequence of UTF-8-encoded code points. It returns the byte index of the first occurrence in s of the given rune. It returns -1 if rune is not present in s. If r is utf8.RuneError, it returns the first instance of any invalid UTF-8 byte sequence.

IndexRune将s解释为一串UTF-8编码的代码点。它返回s中第一次出现的给定符文的字节索引。如果r是utf8.RuneError，它将返回任何无效的UTF-8字节序列的第一个实例。

### func Join 

`func Join(s [][]byte, sep []byte) []byte`

Join concatenates the elements of s to create a new byte slice. The separator sep is placed between elements in the resulting slice.


### func LastIndex

`func LastIndex(s, sep []byte) int`

LastIndex returns the index of the last instance of sep in s, or -1 if sep is not present in s.

LastIndex返回s中最后一个sep实例的索引，如果s中不存在sep，则返回-1。

### func LastIndexAny

`func LastIndexAny(s []byte, chars string) int`

LastIndexAny interprets s as a sequence of UTF-8-encoded Unicode code points. It returns the byte index of the last occurrence in s of any of the Unicode code points in chars. It returns -1 if chars is empty or if there is no code point in common.

LastIndexAny将s解释为一串UTF-8编码的Unicode代码点。它返回s中任何一个Unicode代码点在chars中最后出现的字节索引。如果chars为空或者没有共同的码位，则返回-1。

### func LastIndexByte

`func LastIndexByte(s []byte, c byte) int`

LastIndexByte returns the index of the last instance of c in s, or -1 if c is not present in s.

LastIndexByte返回s中c的最后一个实例的索引，如果c在s中不存在，则返回-1。

### func LastIndexFunc

`func LastIndexFunc(s []byte, f func(r rune) bool) int`

LastIndexFunc interprets s as a sequence of UTF-8-encoded code points. It returns the byte index in s of the last Unicode code point satisfying f(c), or -1 if none do.

LastIndexFunc将s解释为一串UTF-8编码的代码点。它返回s中满足f(c)的最后一个Unicode码位的字节索引，如果没有，则返回-1。

### func Map

`func Map(mapping func(r rune) rune, s []byte) []byte`

Map returns a copy of the byte slice s with all its characters modified according to the mapping function. If mapping returns a negative value, the character is dropped from the byte slice with no replacement. The characters in s and the output are interpreted as UTF-8-encoded code points.

Map返回一个字节片s的副本，并根据映射函数修改其所有字符。如果映射返回一个负值，则该字符将从字节片中删除，不作替换。s中的字符和输出被解释为UTF-8编码的代码点。

### func Repeat

`func Repeat(b []byte, count int) []byte`

Repeat returns a new byte slice consisting of count copies of b.

It panics if count is negative or if the result of (len(b) * count) overflows.

Repeat返回一个新的字节片，该字节片由b的count副本组成。

如果count为负数或者(len(b) * count)的结果溢出，它就会慌乱。

### func ReplaceAll

`func ReplaceAll(s, old, new []byte) []byte`

ReplaceAll returns a copy of the slice s with all non-overlapping instances of old replaced by new. If old is empty, it matches at the beginning of the slice and after each UTF-8 sequence, yielding up to k+1 replacements for a k-rune slice.

ReplaceAll返回一个片断s的副本，其中所有不重叠的旧实例都被新实例替换。如果old是空的，它将在片段的开头和每个UTF-8序列之后进行匹配，对于一个k-rune的片段，最多产生k+1个替换。

### func Runes

`func Runes(s []byte) []rune`

Runes interprets s as a sequence of UTF-8-encoded code points. It returns a slice of runes (Unicode code points) equivalent to s.

Runes将s解释为一串UTF-8编码的代码点。它返回相当于s的符文（Unicode代码点）的片断。

### func Split

`func Split(s, sep []byte) [][]byte`

Split slices s into all subslices separated by sep and returns a slice of the subslices between those separators. If sep is empty, Split splits after each UTF-8 sequence. It is equivalent to SplitN with a count of -1.

Split将s切成由sep分隔的所有子片，并返回这些分隔符之间的子片的一个片断。如果sep为空，Split会在每个UTF-8序列之后进行分割。它等同于SplitN，计数为-1。

### func SplitAfter

`func SplitAfter(s, sep []byte) [][]byte`

SplitAfter slices s into all subslices after each instance of sep and returns a slice of those subslices. If sep is empty, SplitAfter splits after each UTF-8 sequence. It is equivalent to SplitAfterN with a count of -1.

SplitAfter在每个sep实例之后将s切成所有子片，并返回这些子片的一个片断。如果sep是空的，SplitAfter在每个UTF-8序列之后进行分割。它等同于SplitAfterN，计数为-1。

### func SplitAfterN

`func SplitAfterN(s, sep []byte, n int) [][]byte`

SplitAfterN slices s into subslices after each instance of sep and returns a slice of those subslices. If sep is empty, SplitAfterN splits after each UTF-8 sequence. The count determines the number of subslices to return:

SplitAfterN在每个sep的实例之后将s切成子片，并返回这些子片的一个片断。如果sep是空的，SplitAfterN在每个UTF-8序列之后进行分割。Count决定了要返回的子片的数量。


### func SplitN

`func SplitN(s, sep []byte, n int) [][]byte`

SplitN slices s into subslices separated by sep and returns a slice of the subslices between those separators. If sep is empty, SplitN splits after each UTF-8 sequence. The count determines the number of subslices to return:

SplitN将s切成由sep分隔的子片，并返回这些分隔符之间的子片的片断。如果sep为空，SplitN会在每个UTF-8序列之后进行分割。Count决定了要返回的子片的数量。

### func Title 

`func Title(s []byte) []byte`

Title treats s as UTF-8-encoded bytes and returns a copy with all Unicode letters that begin words mapped to their title case.


Title将s处理为UTF-8编码的字节，并返回一个带有所有Unicode字母的副本，这些字母是单词的开头，被映射到它们的标题大小写。

BUG(rsc)。Title用于单词边界的规则不能正确处理Unicode标点符号。

### func ToLower

`func ToLower(s []byte) []byte`

ToLower returns a copy of the byte slice s with all Unicode letters mapped to their lower case.

ToLower返回所有Unicode字母映射为小写字母的字节片s的副本。

### func ToLowerSpecial

`func ToLowerSpecial(c unicode.SpecialCase, s []byte) []byte`

ToLowerSpecial treats s as UTF-8-encoded bytes and returns a copy with all the Unicode letters mapped to their lower case, giving priority to the special casing rules.

ToLowerSpecial将s视为UTF-8编码的字节，并返回一个所有Unicode字母被映射为小写字母的副本，优先考虑特殊大小写规则。

### func ToTitle

`func ToTitle(s []byte) []byte`

ToTitle treats s as UTF-8-encoded bytes and returns a copy with all the Unicode letters mapped to their title case.

ToTitle将s视为UTF-8编码的字节，并返回一个带有所有Unicode字母映射到标题大小写的副本。

### func ToTitleSpecial 

`func ToTitleSpecial(c unicode.SpecialCase, s []byte) []byte`

ToTitleSpecial treats s as UTF-8-encoded bytes and returns a copy with all the Unicode letters mapped to their title case, giving priority to the special casing rules.

ToTitleSpecial将s处理为UTF-8编码的字节，并返回一个带有所有Unicode字母映射到其标题大小写的副本，优先考虑特殊大小写规则。

### func ToUpper 

`func ToUpper(s []byte) []byte`

ToUpper returns a copy of the byte slice s with all Unicode letters mapped to their upper case.

ToUpper返回所有Unicode字母映射为大写字母的字节片s的副本。

### func ToUpperSpecial

`func ToUpperSpecial(c unicode.SpecialCase, s []byte) []byte`

ToUpperSpecial treats s as UTF-8-encoded bytes and returns a copy with all the Unicode letters mapped to their upper case, giving priority to the special casing rules.

ToUpperSpecial将s视为UTF-8编码的字节，并返回一个所有Unicode字母被映射为大写字母的副本，优先考虑特殊大小写规则。

### func ToValidUTF8

`func ToValidUTF8(s, replacement []byte) []byte`

ToValidUTF8 treats s as UTF-8-encoded bytes and returns a copy with each run of bytes representing invalid UTF-8 replaced with the bytes in replacement, which may be empty.

ToValidUTF8将s视为UTF-8编码的字节，并返回一个副本，其中每一个运行的字节代表无效的UTF-8，用替换的字节替换，该字节可能为空。

### func Trim 

`func Trim(s []byte, cutset string) []byte`

Trim returns a subslice of s by slicing off all leading and trailing UTF-8-encoded code points contained in cutset.

Trim通过切掉cutet中包含的所有前面和后面的UTF-8编码的代码点，返回s的子片。

### func TrimFunc

`func TrimFunc(s []byte, f func(r rune) bool) []byte`

TrimFunc returns a subslice of s by slicing off all leading and trailing UTF-8-encoded code points c that satisfy f(c).

TrimFunc通过切掉所有满足f(c)的前面和后面的UTF-8编码的码位c，返回s的子片。

### func TrimLeft 

`func TrimLeft(s []byte, cutset string) []byte`

TrimLeft returns a subslice of s by slicing off all leading UTF-8-encoded code points contained in cutset.

TrimLeft通过切掉cutet中包含的所有领先的UTF-8编码的代码点，返回s的子片。

### func TrimLeftFunc 

`func TrimLeftFunc(s []byte, f func(r rune) bool) []byte`

TrimLeftFunc treats s as UTF-8-encoded bytes and returns a subslice of s by slicing off all leading UTF-8-encoded code points c that satisfy f(c).

TrimLeftFunc将s视为UTF-8编码的字节，并通过切掉所有满足f(c)的领先UTF-8编码的码位c，返回s的子片。

### func TrimPrefix

`func TrimPrefix(s, prefix []byte) []byte`

TrimPrefix returns s without the provided leading prefix string. If s doesn't start with prefix, s is returned unchanged.

TrimPrefix返回没有提供前缀字符串的s。如果s不以前缀开始，s将被原样返回。

### func TrimRight

`func TrimRight(s []byte, cutset string) []byte`

TrimRight returns a subslice of s by slicing off all trailing UTF-8-encoded code points that are contained in cutset.

TrimRight通过切掉cutet中包含的所有UTF-8编码的尾部代码点，返回s的一个子片。

### func TrimRightFunc

`func TrimRightFunc(s []byte, f func(r rune) bool) []byte`

TrimRightFunc returns a subslice of s by slicing off all trailing UTF-8-encoded code points c that satisfy f(c).

TrimRightFunc通过切掉所有满足f(c)的尾部UTF-8编码的码位c，返回s的子片。

### func TrimSpace

`func TrimSpace(s []byte) []byte`

TrimSpace returns a subslice of s by slicing off all leading and trailing white space, as defined by Unicode.

根据Unicode的定义，TrimSpace通过切掉所有前导和尾部的白色空间，返回s的子片。

### func TrimSuffix

`func TrimSuffix(s, suffix []byte) []byte`

TrimSuffix returns s without the provided trailing suffix string. If s doesn't end with suffix, s is returned unchanged.

TrimSuffix返回s，不包括提供的后缀字符串。如果s不以后缀结尾，s将被原样返回。

## type Buffer

```
type Buffer struct {
	// contains filtered or unexported fields
}
```

A Buffer is a variable-sized buffer of bytes with Read and Write methods. The zero value for Buffer is an empty buffer ready to use.

Buffer是一个可变大小的字节缓冲区，有读和写的方法。Buffer的零值是一个准备使用的空缓冲区。

### func NewBuffer

`func NewBuffer(buf []byte) *Buffer`

NewBuffer creates and initializes a new Buffer using buf as its initial contents. The new Buffer takes ownership of buf, and the caller should not use buf after this call. NewBuffer is intended to prepare a Buffer to read existing data. It can also be used to set the initial size of the internal buffer for writing. To do that, buf should have the desired capacity but a length of zero.

In most cases, new(Buffer) (or just declaring a Buffer variable) is sufficient to initialize a Buffer.

NewBuffer创建并初始化一个新的Buffer，使用buf作为其初始内容。新的Buffer拥有buf的所有权，调用者在这次调用后不应该再使用buf。NewBuffer是用来准备一个Buffer来读取现有的数据。它也可以用来设置内部缓冲区的初始大小，以便写入。要做到这一点，buf应该有理想的容量，但长度为0。

### func NewBufferString

`func NewBufferString(s string) *Buffer`

NewBufferString creates and initializes a new Buffer using string s as its initial contents. It is intended to prepare a buffer to read an existing string.

In most cases, new(Buffer) (or just declaring a Buffer variable) is sufficient to initialize a Buffer.

NewBufferString使用字符串s作为其初始内容创建并初始化一个新的Buffer。它的目的是准备一个缓冲区来读取一个现有的字符串。

在大多数情况下，new(Buffer)（或者只是声明一个Buffer变量）就足以初始化一个Buffer。

### func (*Buffer) Bytes

`func (b *Buffer) Bytes() []byte`

Bytes returns a slice of length b.Len() holding the unread portion of the buffer. The slice is valid for use only until the next buffer modification (that is, only until the next call to a method like Read, Write, Reset, or Truncate). The slice aliases the buffer content at least until the next buffer modification, so immediate changes to the slice will affect the result of future reads.

Bytes返回一个长度为b.Len()的片断，持有缓冲区的未读部分。这个片断只在下次修改缓冲区之前有效（也就是说，只在下次调用读、写、重置或截断等方法之前有效）。分片至少在下一次修改缓冲区之前，是缓冲区内容的别名，所以立即改变分片会影响未来的读取结果。

### func (*Buffer) Cap

`func (b *Buffer) Cap() int`

Cap returns the capacity of the buffer's underlying byte slice, that is, the total space allocated for the buffer's data.

Cap返回缓冲区底层字节片的容量，也就是说，分配给缓冲区数据的总空间。

### func (*Buffer) Grow

`func (b *Buffer) Grow(n int)`

Grow grows the buffer's capacity, if necessary, to guarantee space for another n bytes. After Grow(n), at least n bytes can be written to the buffer without another allocation. If n is negative, Grow will panic. If the buffer can't grow it will panic with ErrTooLarge.

如果有必要，Grow会增加缓冲区的容量，以保证另外n个字节的空间。在Grow(n)之后，至少有n个字节可以被写入缓冲区而不需要再次分配。如果n是负数，Grow会出现恐慌。如果缓冲区不能增长，就会出现ErrTooLarge的恐慌。

### func (*Buffer) Len

`func (b *Buffer) Len() int`

Len returns the number of bytes of the unread portion of the buffer; b.Len() == len(b.Bytes()).

Len返回缓冲区未读部分的字节数；b.Len() == len(b.Bytes() )。

### func (*Buffer) Next 

`func (b *Buffer) Next(n int) []byte`

Next returns a slice containing the next n bytes from the buffer, advancing the buffer as if the bytes had been returned by Read. If there are fewer than n bytes in the buffer, Next returns the entire buffer. The slice is only valid until the next call to a read or write method.

Next返回一个包含缓冲区下一个n个字节的片断，将缓冲区向前推进，就像这些字节是由Read返回的一样。如果缓冲区内少于n个字节，Next会返回整个缓冲区。这个片断只在下次调用读或写方法之前有效。

### func (*Buffer) Read

`func (b *Buffer) Read(p []byte) (n int, err error)`

Read reads the next len(p) bytes from the buffer or until the buffer is drained. The return value n is the number of bytes read. If the buffer has no data to return, err is io.EOF (unless len(p) is zero); otherwise it is nil.

Read从缓冲区中读取下一个len(p)字节或直到缓冲区被耗尽。返回值n是读取的字节数。如果缓冲区没有数据返回，err为io.EOF（除非len(p)为零）；否则为nil。

### func (*Buffer) ReadByte

`func (b *Buffer) ReadByte() (byte, error)`

ReadByte reads and returns the next byte from the buffer. If no byte is available, it returns error io.EOF.

ReadByte读取并返回缓冲区的下一个字节。如果没有可用的字节，它会返回错误io.EOF。

### func (*Buffer) ReadBytes

`func (b *Buffer) ReadBytes(delim byte) (line []byte, err error)`

ReadBytes reads until the first occurrence of delim in the input, returning a slice containing the data up to and including the delimiter. If ReadBytes encounters an error before finding a delimiter, it returns the data read before the error and the error itself (often io.EOF). ReadBytes returns err != nil if and only if the returned data does not end in delim.

ReadBytes读到输入中第一次出现delim为止，返回一个包含数据的片断，直到并包括分界符。如果ReadBytes在找到定界符之前遇到了错误，它将返回在错误之前读取的数据和错误本身（通常是io.EOF）。如果且仅当返回的数据不以定界符结束时，ReadBytes返回err !=nil。

### func (*Buffer) ReadFrom

`func (b *Buffer) ReadFrom(r io.Reader) (n int64, err error)`

ReadFrom reads data from r until EOF and appends it to the buffer, growing the buffer as needed. The return value n is the number of bytes read. Any error except io.EOF encountered during the read is also returned. If the buffer becomes too large, ReadFrom will panic with ErrTooLarge.

ReadFrom从r中读取数据，直到EOF，并将其附加到缓冲区中，根据需要增加缓冲区。返回值n是读取的字节数。在读取过程中遇到的任何错误（除了io.EOF）也会被返回。如果缓冲区变得太大，ReadFrom将以ErrTooLarge惊慌失措。

### func (*Buffer) ReadRune 

`func (b *Buffer) ReadRune() (r rune, size int, err error)`

ReadRune reads and returns the next UTF-8-encoded Unicode code point from the buffer. If no bytes are available, the error returned is io.EOF. If the bytes are an erroneous UTF-8 encoding, it consumes one byte and returns U+FFFD, 1.

ReadRune从缓冲区中读取并返回下一个UTF-8编码的Unicode代码点。如果没有可用的字节，返回的错误是io.EOF。如果字节是错误的UTF-8编码，它将消耗一个字节并返回U+FFFD，1。

### func (*Buffer) ReadString

`func (b *Buffer) ReadString(delim byte) (line string, err error)`

ReadString reads until the first occurrence of delim in the input, returning a string containing the data up to and including the delimiter. If ReadString encounters an error before finding a delimiter, it returns the data read before the error and the error itself (often io.EOF). ReadString returns err != nil if and only if the returned data does not end in delim.

ReadString 读取到输入中第一次出现的 delim，返回一个包含数据的字符串，直到并包括分界符。如果ReadString在找到定界符之前遇到错误，它会返回在错误之前读取的数据和错误本身（通常是io.EOF）。如果且仅当返回的数据不以定界符结束时，ReadString返回err !=nil。

### func (*Buffer) Reset

`func (b *Buffer) Reset()`

Reset resets the buffer to be empty, but it retains the underlying storage for use by future writes. Reset is the same as Truncate(0).

Reset将缓冲区重置为空，但是它保留了底层的存储空间，供将来的写操作使用。Reset与Truncate(0)相同。

### func (*Buffer) String

`func (b *Buffer) String() string`

String returns the contents of the unread portion of the buffer as a string. If the Buffer is a nil pointer, it returns "<nil>".

To build strings more efficiently, see the strings.Builder type.

String将缓冲区未读部分的内容作为一个字符串返回。如果Buffer是一个nil指针，它返回"<nil>"。

要想更有效地构建字符串，请参阅strings.Builder类型。

### func (*Buffer) Truncate

`func (b *Buffer) Truncate(n int)`

Truncate discards all but the first n unread bytes from the buffer but continues to use the same allocated storage. It panics if n is negative or greater than the length of the buffer.

Truncate丢弃了缓冲区中除前n个未读字节以外的所有字节，但继续使用相同的分配存储空间。如果n是负数或大于缓冲区的长度，它就会陷入恐慌。

### func (*Buffer) UnreadByte

`func (b *Buffer) UnreadByte() error`

UnreadByte unreads the last byte returned by the most recent successful read operation that read at least one byte. If a write has happened since the last read, if the last read returned an error, or if the read read zero bytes, UnreadByte returns an error.

UnreadByte解读最近一次成功读取至少一个字节的读取操作所返回的最后一个字节。如果在最后一次读取后发生了写操作，如果最后一次读取返回了一个错误，或者如果读取的字节数为零，UnreadByte返回一个错误。

### func (*Buffer) UnreadRune 

`func (b *Buffer) UnreadRune() error`

UnreadRune unreads the last rune returned by ReadRune. If the most recent read or write operation on the buffer was not a successful ReadRune, UnreadRune returns an error. (In this regard it is stricter than UnreadByte, which will unread the last byte from any read operation.)

UnreadRune解读由ReadRune返回的最后一个符文。如果最近对缓冲区的读或写操作不是一个成功的ReadRune，UnreadRune会返回一个错误。(在这方面，它比UnreadByte更严格，UnreadByte将取消任何读操作的最后一个字节。)

### func (*Buffer) Write

`func (b *Buffer) Write(p []byte) (n int, err error)`

Write appends the contents of p to the buffer, growing the buffer as needed. The return value n is the length of p; err is always nil. If the buffer becomes too large, Write will panic with ErrTooLarge.

Write将p的内容追加到缓冲区，根据需要增加缓冲区。返回值n是p的长度；err总是nil。如果缓冲区变得太大，Write将以ErrTooLarge惊慌失措。

### func (*Buffer) WriteByte

`func (b *Buffer) WriteByte(c byte) error`

WriteByte appends the byte c to the buffer, growing the buffer as needed. The returned error is always nil, but is included to match bufio.Writer's WriteByte. If the buffer becomes too large, WriteByte will panic with ErrTooLarge.

WriteByte将字节c追加到缓冲区，根据需要增加缓冲区。返回的错误总是nil，但被包括在内以匹配bufio.Writer的WriteByte。如果缓冲区变得太大，WriteByte将以ErrTooLarge惊慌失措。

### func (*Buffer) WriteRune

`func (b *Buffer) WriteRune(r rune) (n int, err error)`

WriteRune appends the UTF-8 encoding of Unicode code point r to the buffer, returning its length and an error, which is always nil but is included to match bufio.Writer's WriteRune. The buffer is grown as needed; if it becomes too large, WriteRune will panic with ErrTooLarge.

WriteRune将Unicode代码点r的UTF-8编码附加到缓冲区，并返回其长度和一个错误，这个错误总是为零，但被包括在内以匹配bufio.Writer的WriteRune。缓冲区根据需要增长；如果它变得太大，WriteRune将以ErrTooLarge惊慌失措。

### func (*Buffer) WriteString

`func (b *Buffer) WriteString(s string) (n int, err error)`

WriteString appends the contents of s to the buffer, growing the buffer as needed. The return value n is the length of s; err is always nil. If the buffer becomes too large, WriteString will panic with ErrTooLarge.

WriteString将s的内容追加到缓冲区，根据需要增加缓冲区。返回值n是s的长度；err总是nil。如果缓冲区变得太大，WriteString将以ErrTooLarge惊慌。

### func (*Buffer) WriteTo

`func (b *Buffer) WriteTo(w io.Writer) (n int64, err error)`

WriteTo writes data to w until the buffer is drained or an error occurs. The return value n is the number of bytes written; it always fits into an int, but it is int64 to match the io.WriterTo interface. Any error encountered during the write is also returned.

WriteTo向w写入数据，直到缓冲区被耗尽或发生错误。返回值n是写入的字节数；它总是适合于一个int，但它是int64以匹配io.WriterTo接口。写入过程中遇到的任何错误也会被返回。

## type Reader

```
type Reader struct {
	// contains filtered or unexported fields
}
```

A Reader implements the io.Reader, io.ReaderAt, io.WriterTo, io.Seeker, io.ByteScanner, and io.RuneScanner interfaces by reading from a byte slice. Unlike a Buffer, a Reader is read-only and supports seeking. The zero value for Reader operates like a Reader of an empty slice.

读取器通过读取字节片实现io.Reader、io.ReaderAt、io.WriterTo、io.Seeker、io.ByteScanner和io.RuneScanner接口。与Buffer不同，Reader是只读的，并且支持寻道。读取器的零值与空片的读取器操作类似。

### func NewReader 

`func NewReader(b []byte) *Reader`

NewReader returns a new Reader reading from b.

NewReader返回一个新的从b读取的阅读器。

### func (*Reader) Len 

`func (r *Reader) Len() int`

Len returns the number of bytes of the unread portion of the slice.

Len返回分片中未读部分的字节数。

### func (*Reader) Read

`func (r *Reader) Read(b []byte) (n int, err error)`

Read implements the io.Reader interface.

读取实现了io.Reader接口。

### func (*Reader) ReadAt

`func (r *Reader) ReadAt(b []byte, off int64) (n int, err error)`

ReadAt implements the io.ReaderAt interface.

ReadAt实现了io.ReaderAt接口。

### func (*Reader) ReadByte

`func (r *Reader) ReadByte() (byte, error)`

ReadByte implements the io.ByteReader interface.

ReadByte实现了io.ByteReader接口。

### func (*Reader) ReadRune

`func (r *Reader) ReadRune() (ch rune, size int, err error)`

ReadRune implements the io.RuneReader interface.

ReadRune实现了io.RuneReader接口。

### func (*Reader) Reset

`func (r *Reader) Reset(b []byte)`

Reset resets the Reader to be reading from b.

重置将阅读器重置为从b读取。

### func (*Reader) Seek

`func (r *Reader) Seek(offset int64, whence int) (int64, error)`

Seek implements the io.Seeker interface.

Seek实现了io.Seeker接口。

### func (*Reader) Size

`func (r *Reader) Size() int64`

Size returns the original length of the underlying byte slice. Size is the number of bytes available for reading via ReadAt. The returned value is always the same and is not affected by calls to any other method.

Size返回底层字节片的原始长度。Size是可用于通过ReadAt读取的字节数。返回的值总是相同的，并且不受调用任何其他方法的影响。

### func (*Reader) UnreadByte

`func (r *Reader) UnreadByte() error`

UnreadByte complements ReadByte in implementing the io.ByteScanner interface.

UnreadByte在实现io.ByteScanner接口时补充了ReadByte。

### func (*Reader) UnreadRune

`func (r *Reader) UnreadRune() error`

UnreadRune complements ReadRune in implementing the io.RuneScanner interface.

UnreadRune在实现io.RuneScanner接口方面是对ReadRune的补充。

### func (*Reader) WriteTo

`func (r *Reader) WriteTo(w io.Writer) (n int64, err error)`

WriteTo implements the io.WriterTo interface.

WriteTo实现了io.WriterTo接口。


