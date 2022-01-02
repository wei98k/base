## Overview 

Package bufio implements buffered I/O. It wraps an io.Reader or io.Writer object, creating another object (Reader or Writer) that also implements the interface but provides buffering and some help for textual I/O.

包bufio实现了缓冲的I/O。它包装了一个io.Reader或io.Writer对象，创建了另一个对象（Reader或Writer），它也实现了该接口，但提供了缓冲和一些文本I/O的帮助。

## Examples

[Scanner (Custom)](https://pkg.go.dev/bufio@go1.17.5#example-Scanner-Custom)

[Scanner (EmptyFinalToken)](https://pkg.go.dev/bufio@go1.17.5#example-Scanner-EmptyFinalToken)

[Scanner (Lines)](https://pkg.go.dev/bufio@go1.17.5#example-Scanner-Lines)

[Scanner (Words)](https://pkg.go.dev/bufio@go1.17.5#example-Scanner-Words)

[Scanner.Bytes](https://pkg.go.dev/bufio@go1.17.5#example-Scanner.Bytes)

[Writer](https://pkg.go.dev/bufio@go1.17.5#example-Writer)

## Constants

```
const (
	// MaxScanTokenSize is the maximum size used to buffer a token
	// unless the user provides an explicit buffer with Scanner.Buffer.
	// The actual maximum token size may be smaller as the buffer
	// may need to include, for instance, a newline.
	MaxScanTokenSize = 64 * 1024
)
```

## Variables

```
var (
	ErrInvalidUnreadByte = errors.New("bufio: invalid use of UnreadByte")
	ErrInvalidUnreadRune = errors.New("bufio: invalid use of UnreadRune")
	ErrBufferFull        = errors.New("bufio: buffer full")
	ErrNegativeCount     = errors.New("bufio: negative count")
)
```

```
var (
	ErrTooLong         = errors.New("bufio.Scanner: token too long")
	ErrNegativeAdvance = errors.New("bufio.Scanner: SplitFunc returns negative advance count")
	ErrAdvanceTooFar   = errors.New("bufio.Scanner: SplitFunc returns advance count beyond input")
	ErrBadReadCount    = errors.New("bufio.Scanner: Read returned impossible count")
)
```

Errors returned by Scanner.

```
var ErrFinalToken = errors.New("final token")
```

## Functions

### func ScanBytes

`func ScanBytes(data []byte, atEOF bool) (advance int, token []byte, err error)`

ScanBytes is a split function for a Scanner that returns each byte as a token.

ScanBytes是Scanner的一个分割函数，它将每个字节作为一个标记返回。

### func ScanLines

`func ScanLines(data []byte, atEOF bool) (advance int, token []byte, err error)`

ScanLines is a split function for a Scanner that returns each line of text, stripped of any trailing end-of-line marker. The returned line may be empty. The end-of-line marker is one optional carriage return followed by one mandatory newline. In regular expression notation, it is `\r?\n`. The last non-empty line of input will be returned even if it has no newline.

ScanLines是Scanner的一个分割函数，它返回每一行的文本，去掉任何尾部的行尾标记。返回的行可能是空的。行末标记是一个可选的回车键，后面是一个强制性的换行键。在正则表达式中，它是`r?\n`。最后一行非空的输入将被返回，即使它没有换行。


### func ScanRunes

`func ScanRunes(data []byte, atEOF bool) (advance int, token []byte, err error)`

ScanRunes is a split function for a Scanner that returns each UTF-8-encoded rune as a token. The sequence of runes returned is equivalent to that from a range loop over the input as a string, which means that erroneous UTF-8 encodings translate to U+FFFD = "\xef\xbf\xbd". Because of the Scan interface, this makes it impossible for the client to distinguish correctly encoded replacement runes from encoding errors.

ScanRunes是Scanner的一个分割函数，它将每个UTF-8编码的符文作为一个标记返回。返回的符文序列等同于从输入的范围循环中得到的字符串，这意味着错误的UTF-8编码会转化为U+FFFD = "\xef\xbf\xbd"。由于扫描接口的存在，这使得客户端无法区分正确编码的替换符文和编码错误。

### func ScanWords

`func ScanWords(data []byte, atEOF bool) (advance int, token []byte, err error)`

ScanWords is a split function for a Scanner that returns each space-separated word of text, with surrounding spaces deleted. It will never return an empty string. The definition of space is set by unicode.IsSpace.

ScanWords是Scanner的一个分割函数，它返回每个以空格分隔的文本字，并删除周围的空格。它永远不会返回一个空字符串。空间的定义是由unicode.IsSpace设置的。


## Types

## type ReadWriter

```
type ReadWriter struct {
	*Reader
	*Writer
}
```

ReadWriter stores pointers to a Reader and a Writer. It implements io.ReadWriter.

ReadWriter存储指向一个Reader和一个Writer的指针。它实现了io.ReadWriter

### func NewReadWriter

`func NewReadWriter(r *Reader, w *Writer) *ReadWriter`

NewReadWriter allocates a new ReadWriter that dispatches to r and w.

NewReadWriter分配了一个新的ReadWriter，向r和w派发。

## type Reader

```
type Reader struct {
	// contains filtered or unexported fields
}
```

Reader implements buffering for an io.Reader object.

读者实现了io.Reader对象的缓冲。

### func NewReader

`func NewReader(rd io.Reader) *Reader`

NewReader returns a new Reader whose buffer has the default size.

NewReader返回一个新的阅读器，其缓冲区具有默认大小。

### func NewReaderSize

`func NewReaderSize(rd io.Reader, size int) *Reader`

NewReaderSize returns a new Reader whose buffer has at least the specified size. If the argument io.Reader is already a Reader with large enough size, it returns the underlying Reader.

NewReaderSize返回一个新的阅读器，其缓冲区至少有指定的大小。如果参数io.Reader已经是一个具有足够大尺寸的Reader，它将返回底层的Reader。

### func (*Reader) Buffered

`func (b *Reader) Buffered() int`

Buffered returns the number of bytes that can be read from the current buffer.

Buffered返回可从当前缓冲区读取的字节数。


### func (*Reader) Discard

`func (b *Reader) Discard(n int) (discarded int, err error)`

Discard skips the next n bytes, returning the number of bytes discarded.
If Discard skips fewer than n bytes, it also returns an error. If 0 <= n <= b.Buffered(), Discard is guaranteed to succeed without reading from the underlying io.Reader.

Discard跳过接下来的n个字节，返回被丢弃的字节数。
如果Discard跳过的字节数少于n，它也会返回一个错误。如果0 <= n <= b.Buffered()，Discard保证成功，不需要从底层的io.Reader中读取。

### func (*Reader) Peek

`func (b *Reader) Peek(n int) ([]byte, error)`

Peek returns the next n bytes without advancing the reader. The bytes stop being valid at the next read call. If Peek returns fewer than n bytes, it also returns an error explaining why the read is short. The error is ErrBufferFull if n is larger than b's buffer size.

Calling Peek prevents a UnreadByte or UnreadRune call from succeeding until the next read operation.

Peek返回下一个n个字节，而不推进阅读器。这些字节在下一次读取调用时不再有效。如果Peek返回的字节数少于n，它也会返回一个错误，解释为什么读的时间短。如果n大于b的缓冲区大小，这个错误就是ErrBufferFull。

调用Peek可以防止UnreadByte或UnreadRune调用成功，直到下一次读操作。

### func (*Reader) Read

`func (b *Reader) Read(p []byte) (n int, err error)`

Read reads data into p. It returns the number of bytes read into p. The bytes are taken from at most one Read on the underlying Reader, hence n may be less than len(p). To read exactly len(p) bytes, use io.ReadFull(b, p). At EOF, the count will be zero and err will be io.EOF.

读取数据到p中。它返回读到p中的字节数。这些字节最多取自底层阅读器上的一个Read，因此n可能小于len(p)。要准确地读取len(p)字节，请使用io.ReadFull(b, p)。在EOF时，计数将为零，err将为io.EOF。

### func (*Reader) ReadByte

`func (b *Reader) ReadByte() (byte, error)`

ReadByte reads and returns a single byte. If no byte is available, returns an error.

ReadByte 读取并返回一个单字节。如果没有可用的字节，则返回一个错误。

### func (*Reader) ReadBytes

`func (b *Reader) ReadBytes(delim byte) ([]byte, error)`

ReadBytes reads until the first occurrence of delim in the input, returning a slice containing the data up to and including the delimiter. If ReadBytes encounters an error before finding a delimiter, it returns the data read before the error and the error itself (often io.EOF). ReadBytes returns err != nil if and only if the returned data does not end in delim. For simple uses, a Scanner may be more convenient.


ReadBytes读到输入中第一次出现delim为止，返回一个包含数据的片断，直到并包括分界符。如果ReadBytes在找到定界符之前遇到了错误，它将返回在错误之前读取的数据和错误本身（通常是io.EOF）。如果且仅当返回的数据不以定界符结束时，ReadBytes返回err != nil。对于简单的使用，一个扫描器可能更方便。

### func (*Reader) ReadLine

`func (b *Reader) ReadLine() (line []byte, isPrefix bool, err error)`


ReadLine is a low-level line-reading primitive. Most callers should use ReadBytes('\n') or ReadString('\n') instead or use a Scanner.

ReadLine tries to return a single line, not including the end-of-line bytes. If the line was too long for the buffer then isPrefix is set and the beginning of the line is returned. The rest of the line will be returned from future calls. isPrefix will be false when returning the last fragment of the line. The returned buffer is only valid until the next call to ReadLine. ReadLine either returns a non-nil line or it returns an error, never both.

The text returned from ReadLine does not include the line end ("\r\n" or "\n"). No indication or error is given if the input ends without a final line end. Calling UnreadByte after ReadLine will always unread the last byte read (possibly a character belonging to the line end) even if that byte is not part of the line returned by ReadLine.

ReadLine是一个低级的读行原语。大多数调用者应该使用ReadBytes('\n')或ReadString('\n')代替，或者使用Scanner。

ReadLine试图返回一个单行，不包括行尾的字节。如果该行对缓冲区来说太长，那么isPrefix被设置，并返回该行的开头。当返回该行的最后一个片段时，isPrefix将为假。返回的缓冲区只在下次调用ReadLine之前有效。ReadLine要么返回一个非零的行，要么返回一个错误，绝不会同时返回。

从ReadLine返回的文本不包括行尾（"\r\n "或"\n"）。如果输入结束时没有最后的行尾，则不会给出任何指示或错误。在ReadLine之后调用UnreadByte总是会取消读取的最后一个字节（可能是属于行尾的字符），即使该字节不是ReadLine返回的行的一部分。


### func (*Reader) ReadRune

`func (b *Reader) ReadRune() (r rune, size int, err error)`

ReadRune reads a single UTF-8 encoded Unicode character and returns the rune and its size in bytes. If the encoded rune is invalid, it consumes one byte and returns unicode.ReplacementChar (U+FFFD) with a size of 1.

ReadRune读取一个UTF-8编码的Unicode字符，并返回符文和它的大小（字节）。如果编码的符文是无效的，它将消耗一个字节并返回unicode.ReplacementChar（U+FFFD），其大小为1。

### func (*Reader) ReadSlice

`func (b *Reader) ReadSlice(delim byte) (line []byte, err error)`

ReadSlice reads until the first occurrence of delim in the input, returning a slice pointing at the bytes in the buffer. The bytes stop being valid at the next read. If ReadSlice encounters an error before finding a delimiter, it returns all the data in the buffer and the error itself (often io.EOF). ReadSlice fails with error ErrBufferFull if the buffer fills without a delim. Because the data returned from ReadSlice will be overwritten by the next I/O operation, most clients should use ReadBytes or ReadString instead. ReadSlice returns err != nil if and only if line does not end in delim.


ReadSlice读到输入中第一次出现delim为止，返回一个指向缓冲区中的字节的切片。这些字节在下次读取时不再有效。如果ReadSlice在找到定界符之前遇到错误，它会返回缓冲区中的所有数据和错误本身（通常是io.EOF）。如果缓冲区在没有分界符的情况下被填满，ReadSlice就会失败，错误为ErrBufferFull。因为从ReadSlice返回的数据将被下一个I/O操作所覆盖，大多数客户端应该使用ReadBytes或ReadString来代替。如果且仅当一行没有以分界线结束时，ReadSlice返回err !=nil。

### func (*Reader) ReadString

`func (b *Reader) ReadString(delim byte) (string, error)`

ReadString reads until the first occurrence of delim in the input, returning a string containing the data up to and including the delimiter. If ReadString encounters an error before finding a delimiter, it returns the data read before the error and the error itself (often io.EOF). ReadString returns err != nil if and only if the returned data does not end in delim. For simple uses, a Scanner may be more convenient.

ReadString 读取到输入中第一次出现的 delim，返回一个包含数据的字符串，直到并包括分界符。如果ReadString在找到定界符之前遇到错误，它会返回在错误之前读取的数据和错误本身（通常是io.EOF）。如果且仅当返回的数据不以定界符结束时，ReadString返回err != nil。对于简单的使用，Scanner可能更方便。


### func (*Reader) Reset

`func (b *Reader) Reset(r io.Reader)`

Reset discards any buffered data, resets all state, and switches the buffered reader to read from r.

重置丢弃任何缓冲数据，重置所有状态，并将缓冲的读卡器切换到从r读取。

### func (*Reader) Size

`func (b *Reader) Size() int`

Size returns the size of the underlying buffer in bytes.

Size返回底层缓冲区的大小，单位是字节。

### func (*Reader) UnreadByte

`func (b *Reader) UnreadByte() error`

UnreadByte unreads the last byte. Only the most recently read byte can be unread.

UnreadByte returns an error if the most recent method called on the Reader was not a read operation. Notably, Peek is not considered a read operation.

UnreadByte解读最后一个字节。只有最近读取的字节可以被解读。

如果最近在阅读器上调用的方法不是读操作，UnreadByte会返回一个错误。值得注意的是，Peek不被认为是一个读操作。

### func (*Reader) UnreadRune

`func (b *Reader) UnreadRune() error`

UnreadRune unreads the last rune. If the most recent method called on the Reader was not a ReadRune, UnreadRune returns an error. (In this regard it is stricter than UnreadByte, which will unread the last byte from any read operation.)

UnreadRune解除对最后一个符文的读取。如果最近在阅读器上调用的方法不是ReadRune，UnreadRune会返回一个错误。(在这方面，它比UnreadByte更严格，UnreadByte会从任何读操作中取消最后一个字节。)

### func (*Reader) WriteTo

`func (b *Reader) WriteTo(w io.Writer) (n int64, err error)`

WriteTo implements io.WriterTo. This may make multiple calls to the Read method of the underlying Reader. If the underlying reader supports the WriteTo method, this calls the underlying WriteTo without buffering.

WriteTo实现了io.WriterTo。这可能会多次调用底层阅读器的读取方法。如果底层阅读器支持WriteTo方法，这就会调用底层WriteTo而不需要缓冲。


## type Scanner

```
type Scanner struct {
	// contains filtered or unexported fields
}
```

Scanner provides a convenient interface for reading data such as a file of newline-delimited lines of text. Successive calls to the Scan method will step through the 'tokens' of a file, skipping the bytes between the tokens. The specification of a token is defined by a split function of type SplitFunc; the default split function breaks the input into lines with line termination stripped. Split functions are defined in this package for scanning a file into lines, bytes, UTF-8-encoded runes, and space-delimited words. The client may instead provide a custom split function.

Scanning stops unrecoverably at EOF, the first I/O error, or a token too large to fit in the buffer. When a scan stops, the reader may have advanced arbitrarily far past the last token. Programs that need more control over error handling or large tokens, or must run sequential scans on a reader, should use bufio.Reader instead.

扫描器提供了一个方便的接口来读取数据，比如一个以换行符分隔的文本文件。对扫描方法的连续调用将逐步通过文件的 "标记"，跳过标记之间的字节。标记的规格由SplitFunc类型的分割函数定义；默认的分割函数将输入分成几行，并剥离行终止。本软件包中定义了分割函数，用于将文件扫描成行、字节、UTF-8编码的符文和以空格分隔的单词。客户端可以提供一个自定义的分割函数。

扫描在EOF、第一个I/O错误、或一个大到无法装入缓冲区的符号时不可恢复地停止。当扫描停止时，阅读器可能已经任意地推进到了最后一个标记。需要对错误处理或大标记进行更多控制的程序，或者必须在阅读器上运行连续的扫描，应该使用bufio.Reader来代替。


### func NewScanner

`func NewScanner(r io.Reader) *Scanner`


NewScanner returns a new Scanner to read from r. The split function defaults to ScanLines.

NewScanner返回一个新的Scanner以从r中读取，分割函数默认为ScanLines。

### func (*Scanner) Buffer

`func (s *Scanner) Buffer(buf []byte, max int)`

Buffer sets the initial buffer to use when scanning and the maximum size of buffer that may be allocated during scanning. The maximum token size is the larger of max and cap(buf). If max <= cap(buf), Scan will use this buffer only and do no allocation.

By default, Scan uses an internal buffer and sets the maximum token size to MaxScanTokenSize.

Buffer panics if it is called after scanning has started.

Buffer设置扫描时使用的初始缓冲区，以及扫描时可能分配的最大缓冲区大小。最大token大小是max和cap(buf)中较大的一个。如果max <= cap(buf)，Scan将只使用这个缓冲区而不进行分配。

默认情况下，Scan使用一个内部缓冲区，并将最大令牌大小设置为MaxScanTokenSize。

如果在扫描开始后被调用，Buffer就会慌乱。

### func (*Scanner) Bytes

`func (s *Scanner) Bytes() []byte`

Bytes returns the most recent token generated by a call to Scan. The underlying array may point to data that will be overwritten by a subsequent call to Scan. It does no allocation.

Bytes返回由调用Scan产生的最近的token。底层数组可能指向将被后续调用Scan覆盖的数据。它不做任何分配。

### func (*Scanner) Err

`func (s *Scanner) Err() error`

Err returns the first non-EOF error that was encountered by the Scanner.

Err返回扫描器遇到的第一个非EOF错误。

### func (*Scanner) Scan

`func (s *Scanner) Scan() bool`

Scan advances the Scanner to the next token, which will then be available through the Bytes or Text method. It returns false when the scan stops, either by reaching the end of the input or an error. After Scan returns false, the Err method will return any error that occurred during scanning, except that if it was io.EOF, Err will return nil. Scan panics if the split function returns too many empty tokens without advancing the input. This is a common error mode for scanners.

扫描将扫描器推进到下一个标记，然后可以通过Bytes或Text方法获得该标记。当扫描停止时，它将返回false，因为它达到了输入的终点或出现了错误。在Scan返回false后，Err方法将返回扫描过程中发生的任何错误，除了如果是io.EOF，Err将返回nil。如果split函数返回了太多的空符号而没有推进输入，Scan就会惊慌失措。这是扫描器的一种常见错误模式。

### func (*Scanner) Split

`func (s *Scanner) Split(split SplitFunc)`

Split sets the split function for the Scanner. The default split function is ScanLines.

Split panics if it is called after scanning has started.

分割设置扫描仪的分割功能。默认的分割功能是ScanLines。

如果在扫描开始后调用Split，就会出现恐慌。

### func (*Scanner) Text

`func (s *Scanner) Text() string`

Text returns the most recent token generated by a call to Scan as a newly allocated string holding its bytes.

Text将调用Scan所产生的最新令牌作为一个新分配的字符串返回，持有其字节数。

## type SplitFunc

```
type SplitFunc func(data []byte, atEOF bool) (advance int, token []byte, err error)
```

SplitFunc is the signature of the split function used to tokenize the input. The arguments are an initial substring of the remaining unprocessed data and a flag, atEOF, that reports whether the Reader has no more data to give. The return values are the number of bytes to advance the input and the next token to return to the user, if any, plus an error, if any.

Scanning stops if the function returns an error, in which case some of the input may be discarded. If that error is ErrFinalToken, scanning stops with no error.

Otherwise, the Scanner advances the input. If the token is not nil, the Scanner returns it to the user. If the token is nil, the Scanner reads more data and continues scanning; if there is no more data--if atEOF was true--the Scanner returns. If the data does not yet hold a complete token, for instance if it has no newline while scanning lines, a SplitFunc can return (0, nil, nil) to signal the Scanner to read more data into the slice and try again with a longer slice starting at the same point in the input.

The function is never called with an empty data slice unless atEOF is true. If atEOF is true, however, data may be non-empty and, as always, holds unprocessed text.

SplitFunc是用于对输入进行标记的分割函数的签名。参数是剩余未处理数据的初始子串和一个标志atEOF，它报告阅读器是否没有更多的数据可以提供。返回值是推进输入的字节数和返回给用户的下一个标记，如果有的话，再加上一个错误，如果有的话。

如果该函数返回一个错误，扫描就会停止，在这种情况下，一些输入可能被丢弃。如果该错误是ErrFinalToken，则扫描停止，没有错误。

否则，扫描器会推进输入。如果token不是nil，Scanner会将其返回给用户。如果令牌为零，Scanner会读取更多的数据并继续扫描；如果没有更多的数据--如果atEOF为真--Scanner会返回。如果数据还没有包含一个完整的标记，例如在扫描行时没有换行，SplitFunc可以返回（0, nil, nil），以示Scanner向分片中读取更多的数据，并从输入的同一点开始用更长的分片重新尝试。

除非atEOF为真，否则该函数不会以空的数据片被调用。然而，如果atEOF为真，数据可能是非空的，并且像往常一样，持有未处理的文本。

## type Writer

```
type Writer struct {
	// contains filtered or unexported fields
}
```

Writer implements buffering for an io.Writer object. If an error occurs writing to a Writer, no more data will be accepted and all subsequent writes, and Flush, will return the error. After all data has been written, the client should call the Flush method to guarantee all data has been forwarded to the underlying io.Writer.

Writer 实现了 io.Writer 对象的缓冲。如果向Writer写入时发生错误，将不再接受更多的数据，所有后续写入和Flush都将返回错误。在所有数据被写入后，客户端应该调用Flush方法以保证所有数据都被转发到底层的io.Writer。

### func NewWriter

`func NewWriter(w io.Writer) *Writer`

NewWriter returns a new Writer whose buffer has the default size.

NewWriter返回一个新的Writer，其缓冲区具有默认大小。


### func NewWriterSize

`func NewWriterSize(w io.Writer, size int) *Writer`

NewWriterSize returns a new Writer whose buffer has at least the specified size. If the argument io.Writer is already a Writer with large enough size, it returns the underlying Writer.

NewWriterSize返回一个新的Writer，其缓冲区至少有指定的大小。如果参数io.Writer已经是一个具有足够大尺寸的Writer，它将返回底层Writer。

### func (*Writer) Available

`func (b *Writer) Available() int`

Available returns how many bytes are unused in the buffer.

Available返回缓冲区内有多少字节未使用。


### func (*Writer) Buffered 

`func (b *Writer) Buffered() int`

Buffered returns the number of bytes that have been written into the current buffer.

Buffered返回已经写进当前缓冲区的字节数。

### func (*Writer) Flush 

`func (b *Writer) Flush() error`

Flush writes any buffered data to the underlying io.Writer.

Flush将任何缓冲的数据写入底层的io.Writer。

### func (*Writer) ReadFrom

`func (b *Writer) ReadFrom(r io.Reader) (n int64, err error)`

ReadFrom implements io.ReaderFrom. If the underlying writer supports the ReadFrom method, and b has no buffered data yet, this calls the underlying ReadFrom without buffering.

ReadFrom实现了io.ReaderFrom。如果底层写入器支持ReadFrom方法，并且b还没有缓冲的数据，这就会调用底层的ReadFrom而不进行缓冲。

### func (*Writer) Reset

`func (b *Writer) Reset(w io.Writer)`

Reset discards any unflushed buffered data, clears any error, and resets b to write its output to w.

重置会丢弃任何未刷新的缓冲数据，清除任何错误，并重置b，将其输出写入w中。

### func (*Writer) Size

`func (b *Writer) Size() int`

Size returns the size of the underlying buffer in bytes.

Size返回底层缓冲区的大小，单位是字节。

### func (*Writer) Write 

`func (b *Writer) Write(p []byte) (nn int, err error)`

Write writes the contents of p into the buffer. It returns the number of bytes written. If nn < len(p), it also returns an error explaining why the write is short.

Write将p的内容写进缓冲区。它返回写入的字节数。如果nn < len(p)，它也会返回一个错误，解释为什么写得很短。

### func (*Writer) WriteByte

`func (b *Writer) WriteByte(c byte) error`

WriteByte writes a single byte.

WriteByte写一个单字节。

### func (*Writer) WriteRune

`func (b *Writer) WriteRune(r rune) (size int, err error)`

WriteRune writes a single Unicode code point, returning the number of bytes written and any error.

WriteRune写入一个单一的Unicode代码点，返回写入的字节数和任何错误。

### func (*Writer) WriteString 

`func (b *Writer) WriteString(s string) (int, error)`

WriteString writes a string. It returns the number of bytes written. If the count is less than len(s), it also returns an error explaining why the write is short.

WriteString 写入一个字符串。它返回写入的字节数。如果计数小于len(s)，它也会返回一个错误，解释为什么写得很短。


