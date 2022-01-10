
## Overview

Package io provides basic interfaces to I/O primitives. Its primary job is to wrap existing implementations of such primitives, such as those in package os, into shared public interfaces that abstract the functionality, plus some other related primitives.

Because these interfaces and primitives wrap lower-level operations with various implementations, unless otherwise informed clients should not assume they are safe for parallel execution.

包io提供了I/O原语的基本接口。它的主要工作是将这些基元的现有实现（比如包os中的那些）打包成共享的公共接口，以抽象出这些功能以及其他一些相关的基元。

因为这些接口和基元用不同的实现包装了低级别的操作，除非另行通知，否则客户不应该认为它们对并行执行是安全的。


## Examples

[Examples](https://pkg.go.dev/io@go1.17.6#pkg-examples)

## Constants

```
const (
	SeekStart   = 0 // seek relative to the origin of the file
	SeekCurrent = 1 // seek relative to the current offset
	SeekEnd     = 2 // seek relative to the end
)
```

Seek whence values.


## Variables

`var EOF = errors.New("EOF")`

EOF is the error returned by Read when no more input is available. (Read must return EOF itself, not an error wrapping EOF, because callers will test for EOF using ==.) Functions should return EOF only to signal a graceful end of input. If the EOF occurs unexpectedly in a structured data stream, the appropriate error is either ErrUnexpectedEOF or some other error giving more detail.

EOF是Read在没有更多输入时返回的错误。(Read必须返回EOF本身，而不是一个包裹EOF的错误，因为调用者会用==来测试EOF)。函数应该只返回EOF，以示对输入的优雅结束。如果EOF在结构化数据流中意外发生，适当的错误是ErrUnexpectedEOF或者其他给出更多细节的错误。

------

`var ErrClosedPipe = errors.New("io: read/write on closed pipe")`

ErrClosedPipe is the error used for read or write operations on a closed pipe.

ErrClosedPipe是用于对封闭管道进行读或写操作的错误。


-----

`var ErrNoProgress = errors.New("multiple Read calls return no data or error")`

ErrNoProgress is returned by some clients of an Reader when many calls to Read have failed to return any data or error, usually the sign of a broken Reader implementation.

ErrNoProgress是由Reader的一些客户端返回的，当许多对Read的调用都没有返回任何数据或错误时，通常是Reader实现失败的标志。

-----

`var ErrShortBuffer = errors.New("short buffer")`

ErrShortBuffer means that a read required a longer buffer than was provided.

ErrShortBuffer指的是读取时需要的缓冲区比提供的长。

-----

`var ErrShortWrite = errors.New("short write")`

ErrShortWrite means that a write accepted fewer bytes than requested but failed to return an explicit error.

ErrShortWrite意味着一个写操作接受的字节数比要求的少，但未能返回一个明确的错误。


-----

`var ErrUnexpectedEOF = errors.New("unexpected EOF")`

ErrUnexpectedEOF means that EOF was encountered in the middle of reading a fixed-size block or data structure.

ErrUnexpectedEOF意味着在读取一个固定大小的块或数据结构的过程中遇到了EOF。

## Functions

### func Copy

`func Copy(dst Writer, src Reader) (written int64, err error)`

Copy copies from src to dst until either EOF is reached on src or an error occurs. It returns the number of bytes copied and the first error encountered while copying, if any.

A successful Copy returns err == nil, not err == EOF. Because Copy is defined to read from src until EOF, it does not treat an EOF from Read as an error to be reported.

If src implements the WriterTo interface, the copy is implemented by calling src.WriteTo(dst). Otherwise, if dst implements the ReaderFrom interface, the copy is implemented by calling dst.ReadFrom(src).

拷贝从src到dst，直到src上达到EOF或发生错误。它返回复制的字节数和复制时遇到的第一个错误（如果有）。

一个成功的Copy返回err == nil，而不是err == EOF。因为Copy被定义为从src读到EOF，所以它不把Read的EOF当作要报告的错误。

如果src实现了WriterTo接口，则通过调用src.WriteTo(dst)实现拷贝。否则，如果dst实现了ReaderFrom接口，那么就通过调用dst.ReadFrom(src)来实现拷贝。

### func CopyBuffer

`func CopyBuffer(dst Writer, src Reader, buf []byte) (written int64, err error)`

CopyBuffer is identical to Copy except that it stages through the provided buffer (if one is required) rather than allocating a temporary one. If buf is nil, one is allocated; otherwise if it has zero length, CopyBuffer panics.

If either src implements WriterTo or dst implements ReaderFrom, buf will not be used to perform the copy.

CopyBuffer与Copy相同，只是它通过提供的缓冲区（如果需要的话），而不是分配一个临时缓冲区。如果buf为nil，就会分配一个；否则，如果它的长度为零，CopyBuffer就会陷入困境。

如果src实现了WriterTo或者dst实现了ReaderFrom，buf将不会被用于执行拷贝。

### func CopyN

`func CopyN(dst Writer, src Reader, n int64) (written int64, err error)`

CopyN copies n bytes (or until an error) from src to dst. It returns the number of bytes copied and the earliest error encountered while copying. On return, written == n if and only if err == nil.

If dst implements the ReaderFrom interface, the copy is implemented using it.

CopyN从src向dst拷贝n个字节（或直到出现错误）。它返回复制的字节数和复制时遇到的最早的错误。返回时，如果且仅当err == nil时，write == n。

如果dst实现了ReaderFrom接口，则使用该接口实现拷贝。

### func Pipe 

`func Pipe() (*PipeReader, *PipeWriter)`

Pipe creates a synchronous in-memory pipe. It can be used to connect code expecting an io.Reader with code expecting an io.Writer.

Reads and Writes on the pipe are matched one to one except when multiple Reads are needed to consume a single Write. That is, each Write to the PipeWriter blocks until it has satisfied one or more Reads from the PipeReader that fully consume the written data. The data is copied directly from the Write to the corresponding Read (or Reads); there is no internal buffering.

It is safe to call Read and Write in parallel with each other or with Close. Parallel calls to Read and parallel calls to Write are also safe: the individual calls will be gated sequentially.

Pipe创建一个同步的内存管道。它可以用来连接期望使用io.Reader的代码和期望使用io.Writer的代码。

管道上的读和写是一对一的，除非需要多个读来消耗一个写。也就是说，PipeWriter的每个Write都会被阻断，直到它满足了PipeReader的一个或多个完全消耗所写数据的Read。数据被直接从Write复制到相应的Read（或Reads）；没有内部缓冲。

互相平行调用Read和Write是安全的，也可以和Close一起调用。平行调用Read和平行调用Write也是安全的：各个调用将按顺序被门控。

### func ReadAll

`func ReadAll(r Reader) ([]byte, error)`

ReadAll reads from r until an error or EOF and returns the data it read. A successful call returns err == nil, not err == EOF. Because ReadAll is defined to read from src until EOF, it does not treat an EOF from Read as an error to be reported.

ReadAll从r读取数据，直到出现错误或EOF，并返回它所读取的数据。一个成功的调用返回err == nil，而不是err == EOF。因为ReadAll被定义为从src读到EOF为止，所以它不把Read的EOF当作一个需要报告的错误。

### func ReadAtLeast

`func ReadAtLeast(r Reader, buf []byte, min int) (n int, err error)`

ReadAtLeast reads from r into buf until it has read at least min bytes. It returns the number of bytes copied and an error if fewer bytes were read. The error is EOF only if no bytes were read. If an EOF happens after reading fewer than min bytes, ReadAtLeast returns ErrUnexpectedEOF. If min is greater than the length of buf, ReadAtLeast returns ErrShortBuffer. On return, n >= min if and only if err == nil. If r returns an error having read at least min bytes, the error is dropped.

ReadAtLeast从r读到buf，直到它至少读了min字节。它返回复制的字节数，如果读取的字节数较少，则返回错误。只有在没有读取任何字节的情况下，错误才是EOF。如果在读取少于min字节后发生EOF，ReadAtLeast返回ErrUnexpectedEOF。如果 min 大于 buf 的长度，ReadAtLeast 返回 ErrShortBuffer。在返回时，n >= min，当且仅当err == nil。如果r在至少读完min字节后返回一个错误，那么这个错误将被放弃。

### func ReadFull

`func ReadFull(r Reader, buf []byte) (n int, err error)`

ReadFull reads exactly len(buf) bytes from r into buf. It returns the number of bytes copied and an error if fewer bytes were read. The error is EOF only if no bytes were read. If an EOF happens after reading some but not all the bytes, ReadFull returns ErrUnexpectedEOF. On return, n == len(buf) if and only if err == nil. If r returns an error having read at least len(buf) bytes, the error is dropped.

ReadFull从r中准确地读取len(buf)字节到buf中。它返回复制的字节数，如果读取的字节数更少，则返回错误。只有在没有读取任何字节的情况下，错误才是EOF。如果EOF发生在读取了一些但不是全部的字节之后，ReadFull返回ErrUnexpectedEOF。返回时，如果且仅当err == nil时，n == len(buf)。如果r在读取了至少len(buf)字节后返回一个错误，那么这个错误将被放弃。

### func WriteString

`func WriteString(w Writer, s string) (n int, err error)`

WriteString writes the contents of the string s to w, which accepts a slice of bytes. If w implements StringWriter, its WriteString method is invoked directly. Otherwise, w.Write is called exactly once.

WriteString将字符串s的内容写到w，w接受一个字节片。如果w实现了StringWriter，它的WriteString方法被直接调用。否则，w.Write正好被调用一次。

## type ByteReader 

```
type ByteReader interface {
	ReadByte() (byte, error)
}
```

ByteReader is the interface that wraps the ReadByte method.

ReadByte reads and returns the next byte from the input or any error encountered. If ReadByte returns an error, no input byte was consumed, and the returned byte value is undefined.

ReadByte provides an efficient interface for byte-at-time processing. A Reader that does not implement ByteReader can be wrapped using bufio.NewReader to add this method.

ByteReader是包装ReadByte方法的接口。

ReadByte 读取并返回输入的下一个字节或遇到的任何错误。如果ReadByte返回错误，说明没有消耗任何输入字节，返回的字节值是未定义的。

ReadByte为字节时间处理提供了一个有效的接口。一个没有实现ByteReader的Reader可以使用bufio.NewReader包装来添加这个方法。

## type ByteScanner

```
type ByteScanner interface {
	ByteReader
	UnreadByte() error
}
```

ByteScanner is the interface that adds the UnreadByte method to the basic ReadByte method.

UnreadByte causes the next call to ReadByte to return the same byte as the previous call to ReadByte. It may be an error to call UnreadByte twice without an intervening call to ReadByte.

ByteScanner是在基本的ReadByte方法上增加UnreadByte方法的接口。

UnreadByte使下一次对ReadByte的调用返回与之前对ReadByte的调用相同的字节。在没有调用ReadByte的情况下，两次调用UnreadByte可能是一个错误。


## type ByteWriter

```
type ByteWriter interface {
	WriteByte(c byte) error
}
```

ByteWriter is the interface that wraps the WriteByte method.

ByteWriter是包装WriteByte方法的接口。

## type Closer

```
type Closer interface {
	Close() error
}
```

Closer is the interface that wraps the basic Close method.

The behavior of Close after the first call is undefined. Specific implementations may document their own behavior.

Closer是包装基本Close方法的接口。

第一次调用后，Close的行为是未定义的。具体的实现可以记录他们自己的行为。

## type LimitedReader

```
type LimitedReader struct {
	R Reader // underlying reader
	N int64  // max bytes remaining
}
```

A LimitedReader reads from R but limits the amount of data returned to just N bytes. Each call to Read updates N to reflect the new amount remaining. Read returns EOF when N <= 0 or when the underlying R returns EOF.

一个有限读取器从R中读取数据，但将返回的数据量限制在N个字节。每次对Read的调用都会更新N，以反映新的剩余量。当N<=0或底层R返回EOF时，Read返回EOF。

### func (*LimitedReader) Read

`func (l *LimitedReader) Read(p []byte) (n int, err error)`

## type PipeReader

```
type PipeReader struct {
	// contains filtered or unexported fields
}
```

A PipeReader is the read half of a pipe.

PipeReader是一个管道的读取部分。

### func (*PipeReader) Close

`func (r *PipeReader) Close() error`

Close closes the reader; subsequent writes to the write half of the pipe will return the error ErrClosedPipe.

Close关闭了阅读器；随后对管道的写半部分的写入将返回错误ErrClosedPipe。

### func (*PipeReader) CloseWithError

`func (r *PipeReader) CloseWithError(err error) error`

CloseWithError closes the reader; subsequent writes to the write half of the pipe will return the error err.

CloseWithError never overwrites the previous error if it exists and always returns nil.

CloseWithError关闭阅读器；随后向管道的写半部分的写操作将返回错误err。

如果存在错误，CloseWithError不会覆盖之前的错误，并且总是返回nil。


### func (*PipeReader) Read

`func (r *PipeReader) Read(data []byte) (n int, err error)`

Read implements the standard Read interface: it reads data from the pipe, blocking until a writer arrives or the write end is closed. If the write end is closed with an error, that error is returned as err; otherwise err is EOF.

读取实现了标准的读取接口：它从管道中读取数据，阻塞直到有写者到达或写端被关闭。如果写端关闭时出现了错误，该错误将作为err返回；否则err就是EOF。

## type PipeWriter 

```
type PipeWriter struct {
	// contains filtered or unexported fields
}
```

A PipeWriter is the write half of a pipe.

PipeWriter是一个管道的写入部分。

### func (*PipeWriter) Close

`func (w *PipeWriter) Close() error`

Close closes the writer; subsequent reads from the read half of the pipe will return no bytes and EOF.

Close关闭写入器；随后从管道的读半部分读出的数据将不返回字节和EOF。

### func (*PipeWriter) CloseWithError

`func (w *PipeWriter) CloseWithError(err error) error`

CloseWithError closes the writer; subsequent reads from the read half of the pipe will return no bytes and the error err, or EOF if err is nil.

CloseWithError never overwrites the previous error if it exists and always returns nil.

CloseWithError关闭写入器；随后从管道的读半部分读出的数据将不返回字节和错误err，如果err为nil，则返回EOF。

如果存在错误，CloseWithError不会覆盖之前的错误，并且总是返回nil。


### func (*PipeWriter) Write

`func (w *PipeWriter) Write(data []byte) (n int, err error)`

Write implements the standard Write interface: it writes data to the pipe, blocking until one or more readers have consumed all the data or the read end is closed. If the read end is closed with an error, that err is returned as err; otherwise err is ErrClosedPipe.

Write实现了标准的Write接口：它向管道写入数据，阻塞直到一个或多个读者消耗了所有的数据或读取端被关闭。如果读取端关闭时出现了错误，则返回err；否则err是ErrClosedPipe。

## type ReadCloser

```
type ReadCloser interface {
	Reader
	Closer
}
```

ReadCloser is the interface that groups the basic Read and Close methods.

ReadCloser是一个接口，它将基本的读取和关闭方法分组。

### func NopCloser

`func NopCloser(r Reader) ReadCloser`

NopCloser returns a ReadCloser with a no-op Close method wrapping the provided Reader r.

NopCloser返回一个ReadCloser，它有一个无操作的关闭方法，包裹着提供的Reader r。


## type ReadSeekCloser

```
type ReadSeekCloser interface {
	Reader
	Seeker
	Closer
}
```

ReadSeekCloser is the interface that groups the basic Read, Seek and Close methods.

ReadSeekCloser是将基本的读取、寻找和关闭方法分组的接口。

## type ReadSeeker

```
type ReadSeeker interface {
	Reader
	Seeker
}
```

ReadSeeker is the interface that groups the basic Read and Seek methods.

ReadSeeker是一个接口，它将基本的Read和Seek方法分组。


## type ReadWriteCloser

```
type ReadWriteCloser interface {
	Reader
	Writer
	Closer
}
```

ReadWriteCloser is the interface that groups the basic Read, Write and Close methods.

ReadWriteCloser是一个接口，它将基本的读、写和关闭方法分组。

## type ReadWriteSeeker 

```
type ReadWriteSeeker interface {
	Reader
	Writer
	Seeker
}
```

ReadWriteSeeker is the interface that groups the basic Read, Write and Seek methods.

ReadWriteSeeker是将基本的Read、Write和Seek方法分组的接口。

## type ReadWriter

```
type ReadWriter interface {
	Reader
	Writer
}
```

ReadWriter is the interface that groups the basic Read and Write methods.

ReadWriter是一个接口，它将基本的读和写方法分组。

## type Reader

```
type Reader interface {
	Read(p []byte) (n int, err error)
}
```

Reader is the interface that wraps the basic Read method.

Read reads up to len(p) bytes into p. It returns the number of bytes read (0 <= n <= len(p)) and any error encountered. Even if Read returns n < len(p), it may use all of p as scratch space during the call. If some data is available but not len(p) bytes, Read conventionally returns what is available instead of waiting for more.

When Read encounters an error or end-of-file condition after successfully reading n > 0 bytes, it returns the number of bytes read. It may return the (non-nil) error from the same call or return the error (and n == 0) from a subsequent call. An instance of this general case is that a Reader returning a non-zero number of bytes at the end of the input stream may return either err == EOF or err == nil. The next Read should return 0, EOF.

Callers should always process the n > 0 bytes returned before considering the error err. Doing so correctly handles I/O errors that happen after reading some bytes and also both of the allowed EOF behaviors.

Implementations of Read are discouraged from returning a zero byte count with a nil error, except when len(p) == 0. Callers should treat a return of 0 and nil as indicating that nothing happened; in particular it does not indicate EOF.

Implementations must not retain p.

读取器是包装基本读取方法的接口。

它返回读取的字节数（0 <= n <= len(p)）和遇到的任何错误。即使Read返回n < len(p)，它也可能在调用过程中使用所有的p作为抓取空间。如果有些数据是可用的，但不是len(p)字节，Read通常会返回可用的数据，而不是等待更多。

当Read在成功读取n>0字节后遇到错误或文件结束的情况时，它会返回所读取的字节数。它可以在同一个调用中返回（非零）错误，或者在后续调用中返回错误（和n == 0）。这个一般情况的一个例子是，在输入流结束时返回非零字节数的Reader可能返回err == EOF或者err == nil。下一个Read应该返回0，即EOF。

调用者应该总是在考虑错误err之前处理返回的n > 0字节。这样做可以正确处理读取一些字节后发生的I/O错误，也可以处理两种允许的EOF行为。

不鼓励Read的实现在返回0字节数时出现nil错误，除非len(p)==0。调用者应该把返回0和nil视为表示什么都没有发生；特别是它不表示EOF。

实现不得保留p。

### func LimitReader

`func LimitReader(r Reader, n int64) Reader`

LimitReader returns a Reader that reads from r but stops with EOF after n bytes. The underlying implementation is a *LimitedReader.


LimitReader返回一个从r中读取的阅读器，但在n个字节后以EOF停止。底层实现是一个*LimitedReader。

### func MultiReader

`func MultiReader(readers ...Reader) Reader`

MultiReader returns a Reader that's the logical concatenation of the provided input readers. They're read sequentially. Once all inputs have returned EOF, Read will return EOF. If any of the readers return a non-nil, non-EOF error, Read will return that error.

MultiReader返回一个阅读器，它是所提供的输入阅读器的逻辑串联。它们被依次读取。一旦所有的输入都返回EOF，Read将返回EOF。如果任何一个读取器返回一个非零、非EOF的错误，Read将返回该错误。

### func TeeReader

`func TeeReader(r Reader, w Writer) Reader`

TeeReader returns a Reader that writes to w what it reads from r. All reads from r performed through it are matched with corresponding writes to w. There is no internal buffering - the write must complete before the read completes. Any error encountered while writing is reported as a read error.

TeeReader返回一个读取器，该读取器将其从r中读取的内容写入w中，所有通过它从r中进行的读取都与写入w的内容相匹配。写入时遇到的任何错误都被报告为读取错误。

## type ReaderAt

```
type ReaderAt interface {
	ReadAt(p []byte, off int64) (n int, err error)
}
```

ReaderAt is the interface that wraps the basic ReadAt method.

ReadAt reads len(p) bytes into p starting at offset off in the underlying input source. It returns the number of bytes read (0 <= n <= len(p)) and any error encountered.

When ReadAt returns n < len(p), it returns a non-nil error explaining why more bytes were not returned. In this respect, ReadAt is stricter than Read.

Even if ReadAt returns n < len(p), it may use all of p as scratch space during the call. If some data is available but not len(p) bytes, ReadAt blocks until either all the data is available or an error occurs. In this respect ReadAt is different from Read.

If the n = len(p) bytes returned by ReadAt are at the end of the input source, ReadAt may return either err == EOF or err == nil.

If ReadAt is reading from an input source with a seek offset, ReadAt should not affect nor be affected by the underlying seek offset.

Clients of ReadAt can execute parallel ReadAt calls on the same input source.

Implementations must not retain p.

ReaderAt是包装基本ReadAt方法的接口。

ReadAt从底层输入源的偏移量开始，将len(p)字节读入p。它返回读取的字节数（0 <= n <= len(p)）和遇到的任何错误。

当 ReadAt 返回 n < len(p)时，它会返回一个非零的错误，解释为什么没有返回更多的字节。在这一点上，ReadAt比Read更严格。

即使 ReadAt 返回 n < len(p)，它也可能在调用过程中使用所有 p 作为抓取空间。如果有些数据是可用的，但不是 len(p) 字节，ReadAt 就会阻塞，直到所有的数据都可用或者发生错误。在这一点上，ReadAt与Read不同。

如果ReadAt返回的n = len(p)字节在输入源的末端，ReadAt可能返回err == EOF或err == nil。

如果ReadAt是从一个有寻址偏移的输入源读取的，ReadAt不应该影响也不应该受到底层寻址偏移的影响。

ReadAt的客户端可以在同一个输入源上执行并行的ReadAt调用。

实现不得保留p。

## type ReaderFrom

```
type ReaderFrom interface {
	ReadFrom(r Reader) (n int64, err error)
}
```

ReaderFrom is the interface that wraps the ReadFrom method.

ReadFrom reads data from r until EOF or error. The return value n is the number of bytes read. Any error except EOF encountered during the read is also returned.

The Copy function uses ReaderFrom if available.

ReaderFrom是包装ReadFrom方法的接口。

ReadFrom从r读取数据，直到EOF或错误。返回值n是读取的字节数。读取过程中遇到的除EOF外的任何错误也会被返回。

如果有的话，Copy函数使用ReaderFrom。

## type RuneReader

```
type RuneReader interface {
	ReadRune() (r rune, size int, err error)
}
```

RuneReader is the interface that wraps the ReadRune method.

ReadRune reads a single UTF-8 encoded Unicode character and returns the rune and its size in bytes. If no character is available, err will be set.

RuneReader是包装ReadRune方法的接口。

ReadRune读取一个UTF-8编码的Unicode字符，并返回符文和它的大小（字节）。如果没有可用的字符，将设置err。

## type RuneScanner

```
type RuneScanner interface {
	RuneReader
	UnreadRune() error
}
```

RuneScanner is the interface that adds the UnreadRune method to the basic ReadRune method.

UnreadRune causes the next call to ReadRune to return the same rune as the previous call to ReadRune. It may be an error to call UnreadRune twice without an intervening call to ReadRune.

RuneScanner是在基本的ReadRune方法上增加UnreadRune方法的接口。

UnreadRune使下一次对ReadRune的调用返回与之前对ReadRune的调用相同的符文。在没有调用ReadRune的情况下两次调用UnreadRune可能是一个错误。

## type SectionReader

```
type SectionReader struct {
	// contains filtered or unexported fields
}
```

SectionReader implements Read, Seek, and ReadAt on a section of an underlying ReaderAt.

SectionReader实现了对底层ReaderAt的一个部分的读取、寻找和ReadAt。

### func NewSectionReader 

`func NewSectionReader(r ReaderAt, off int64, n int64) *SectionReader`

NewSectionReader returns a SectionReader that reads from r starting at offset off and stops with EOF after n bytes.

NewSectionReader返回一个SectionReader，从偏移量off开始读取r，并在n个字节后以EOF停止。

### func (*SectionReader) Read

`func (s *SectionReader) Read(p []byte) (n int, err error)`

### func (*SectionReader) ReadAt

`func (s *SectionReader) ReadAt(p []byte, off int64) (n int, err error)`

### func (*SectionReader) Seek

`func (s *SectionReader) Seek(offset int64, whence int) (int64, error)`

### func (*SectionReader) Size

`func (s *SectionReader) Size() int64`

Size returns the size of the section in bytes.

Size返回该部分的大小，单位为字节。

## type Seeker

```
type Seeker interface {
	Seek(offset int64, whence int) (int64, error)
}
```

Seeker is the interface that wraps the basic Seek method.

Seek sets the offset for the next Read or Write to offset, interpreted according to whence: SeekStart means relative to the start of the file, SeekCurrent means relative to the current offset, and SeekEnd means relative to the end. Seek returns the new offset relative to the start of the file and an error, if any.

Seeking to an offset before the start of the file is an error. Seeking to any positive offset is legal, but the behavior of subsequent I/O operations on the underlying object is implementation-dependent.

Seeker是包装基本Seek方法的接口。

Seek将下一次读或写的偏移量设置为偏移量，根据whence进行解释。SeekStart表示相对于文件的开始，SeekCurrent表示相对于当前的偏移量，而SeekEnd表示相对于结束。Seek返回相对于文件开始的新偏移量和一个错误，如果有的话。

在文件开始之前寻找一个偏移量是一个错误。寻求任何正的偏移量都是合法的，但是对底层对象的后续I/O操作的行为取决于实现。

## type StringWriter

```
type StringWriter interface {
	WriteString(s string) (n int, err error)
}
```

StringWriter is the interface that wraps the WriteString method.

StringWriter是包装WriteString方法的接口。


## type WriteCloser

```
type WriteCloser interface {
	Writer
	Closer
}
```

WriteCloser is the interface that groups the basic Write and Close methods.

WriteCloser是一个接口，它将基本的Write和Close方法分组。


## type WriteSeeker 

```
type WriteSeeker interface {
	Writer
	Seeker
}
```

WriteSeeker is the interface that groups the basic Write and Seek methods.

WriteSeeker是一个接口，将基本的Write和Seek方法分组。

## type Writer

```
type Writer interface {
	Write(p []byte) (n int, err error)
}
```

Writer is the interface that wraps the basic Write method.

Write writes len(p) bytes from p to the underlying data stream. It returns the number of bytes written from p (0 <= n <= len(p)) and any error encountered that caused the write to stop early. Write must return a non-nil error if it returns n < len(p). Write must not modify the slice data, even temporarily.

Implementations must not retain p.

Writer是包装基本Write方法的接口。

Write将len(p)字节从p写到底层数据流中。它返回从p写入的字节数（0 <= n <= len(p)），以及遇到的导致写入提前停止的任何错误。如果Write返回n < len(p)，它必须返回一个非零的错误。写入时不能修改分片数据，即使是暂时的。

实现不得保留p。

### func MultiWriter

`func MultiWriter(writers ...Writer) Writer`

MultiWriter creates a writer that duplicates its writes to all the provided writers, similar to the Unix tee(1) command.

Each write is written to each listed writer, one at a time. If a listed writer returns an error, that overall write operation stops and returns the error; it does not continue down the list.

MultiWriter创建了一个写入器，它向所有提供的写入器重复写入，类似于Unix的tee(1)命令。

每个写操作都会被写到每个列出的写器上，一次一个。如果一个列出的写入器返回一个错误，那么整个写入操作就会停止并返回错误；它不会继续往下写。

## type WriterAt

```
type WriterAt interface {
	WriteAt(p []byte, off int64) (n int, err error)
}
```

WriterAt is the interface that wraps the basic WriteAt method.

WriteAt writes len(p) bytes from p to the underlying data stream at offset off. It returns the number of bytes written from p (0 <= n <= len(p)) and any error encountered that caused the write to stop early. WriteAt must return a non-nil error if it returns n < len(p).

If WriteAt is writing to a destination with a seek offset, WriteAt should not affect nor be affected by the underlying seek offset.

Clients of WriteAt can execute parallel WriteAt calls on the same destination if the ranges do not overlap.

Implementations must not retain p.

WriterAt是包装基本WriteAt方法的接口。

WriteAt将len(p)字节从p写到底层数据流的偏移量处。它返回从p写入的字节数（0 <= n <= len(p)），以及遇到的导致写入提前停止的任何错误。如果WriteAt返回n < len(p)，它必须返回一个非零的错误。

如果WriteAt正在向有寻址偏移的目的地写入，WriteAt不应该影响也不应该受到基础寻址偏移的影响。

如果范围不重叠，WriteAt 的客户端可以在同一个目标上执行平行的 WriteAt 调用。

实现不得保留p。

## type WriterTo

```
type WriterTo interface {
	WriteTo(w Writer) (n int64, err error)
}
```

WriterTo is the interface that wraps the WriteTo method.

WriteTo writes data to w until there's no more data to write or when an error occurs. The return value n is the number of bytes written. Any error encountered during the write is also returned.

The Copy function uses WriterTo if available.

WriterTo是包装WriteTo方法的接口。

WriteTo向w写数据，直到没有数据可写或发生错误时。返回值n是写入的字节数。写入过程中遇到的任何错误也会被返回。

如果有的话，Copy函数使用WriterTo。



















