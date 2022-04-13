
## Overview

Package fs defines basic interfaces to a file system. A file system can be provided by the host operating system but also by other packages.


包 fs 定义了文件系统的基本接口。文件系统可以由主机操作系统提供，也可以由其他包提供。


## Examples

## Variables

```
var (
	ErrInvalid    = errInvalid()    // "invalid argument"
	ErrPermission = errPermission() // "permission denied"
	ErrExist      = errExist()      // "file already exists"
	ErrNotExist   = errNotExist()   // "file does not exist"
	ErrClosed     = errClosed()     // "file already closed"
)
```

Generic file system errors. Errors returned by file systems can be tested against these errors using errors.Is.

通用的文件系统错误。文件系统返回的错误可以用error.Is来测试这些错误。

-----

`var SkipDir = errors.New("skip this directory")`

SkipDir is used as a return value from WalkDirFuncs to indicate that the directory named in the call is to be skipped. It is not returned as an error by any function.

SkipDir被用作WalkDirFuncs的返回值，表示要跳过调用中指定的目录。它不会被任何函数作为错误返回。


## Functions

### func Glob 

`func Glob(fsys FS, pattern string) (matches []string, err error)`

Glob returns the names of all files matching pattern or nil if there is no matching file. The syntax of patterns is the same as in path.Match. The pattern may describe hierarchical names such as usr/*/bin/ed.

Glob ignores file system errors such as I/O errors reading directories. The only possible returned error is path.ErrBadPattern, reporting that the pattern is malformed.

If fs implements GlobFS, Glob calls fs.Glob. Otherwise, Glob uses ReadDir to traverse the directory tree and look for matches for the pattern.

Glob返回与模式匹配的所有文件的名称，如果没有匹配的文件，则返回nil。模式的语法与path.Match中相同。模式可以描述层次化的名称，如usr/*/bin/ed。

Glob忽略文件系统错误，例如I/O错误读取目录。唯一可能返回的错误是path.ErrBadPattern，报告模式是错误的。

如果fs实现了GlobFS，Glob会调用fs.Glob。否则，Glob使用ReadDir遍历目录树，寻找与模式匹配的内容。

### func ReadFile

`func ReadFile(fsys FS, name string) ([]byte, error)`

ReadFile reads the named file from the file system fs and returns its contents. A successful call returns a nil error, not io.EOF. (Because ReadFile reads the whole file, the expected EOF from the final Read is not treated as an error to be reported.)

If fs implements ReadFileFS, ReadFile calls fs.ReadFile. Otherwise ReadFile calls fs.Open and uses Read and Close on the returned file.

ReadFile从文件系统fs中读取命名的文件并返回其内容。一个成功的调用会返回一个 nil 错误，而不是 io.EOF。(因为 ReadFile 读取的是整个文件，所以来自最终读取的预期 EOF 不会被视为一个要报告的错误)。

如果 fs 实现了 ReadFileFS，ReadFile 会调用 fs.ReadFile。否则 ReadFile 会调用 fs.Open 并对返回的文件进行读取和关闭。

### func ValidPath

`func ValidPath(name string) bool`

ValidPath reports whether the given path name is valid for use in a call to Open.

Path names passed to open are UTF-8-encoded, unrooted, slash-separated sequences of path elements, like “x/y/z”. Path names must not contain an element that is “.” or “..” or the empty string, except for the special case that the root directory is named “.”. Paths must not start or end with a slash: “/x” and “x/” are invalid.

Note that paths are slash-separated on all systems, even Windows. Paths containing other characters such as backslash and colon are accepted as valid, but those characters must never be interpreted by an FS implementation as path element separators.


ValidPath报告给定的路径名称在调用Open时是否有效。

传递给open的路径名是UTF-8编码的、无根的、斜线分隔的路径元素序列，如 "x/y/z"。路径名不能包含". "或". "或空字符串的元素，除了根目录被命名为". "的特殊情况。路径不能以斜线开始或结束："/x "和 "x/"是无效的。

请注意，在所有的系统上，甚至是Windows，路径都是以斜线分隔的。含有反斜杠和冒号等其他字符的路径可以被接受为有效，但这些字符决不能被FS实现解释为路径元素分隔符。


### func WalkDir 

`func WalkDir(fsys FS, root string, fn WalkDirFunc) error`

WalkDir walks the file tree rooted at root, calling fn for each file or directory in the tree, including root.

All errors that arise visiting files and directories are filtered by fn: see the fs.WalkDirFunc documentation for details.

The files are walked in lexical order, which makes the output deterministic but requires WalkDir to read an entire directory into memory before proceeding to walk that directory.

WalkDir does not follow symbolic links found in directories, but if root itself is a symbolic link, its target will be walked.

WalkDir行走以根为根的文件树，为树上的每个文件或目录（包括根）调用fn。

访问文件和目录时出现的所有错误都被fn过滤：详情请参见fs.WalkDirFunc文档。

文件是按词法顺序走的，这使得输出是确定的，但要求WalkDir在继续走该目录之前将整个目录读入内存。

WalkDir不跟踪在目录中发现的符号链接，但是如果root本身是一个符号链接，它的目标将被行走。

## type DirEntry

```
type DirEntry interface {
	// Name returns the name of the file (or subdirectory) described by the entry.
	// This name is only the final element of the path (the base name), not the entire path.
	// For example, Name would return "hello.go" not "home/gopher/hello.go".
	Name() string

	// IsDir reports whether the entry describes a directory.
	IsDir() bool

	// Type returns the type bits for the entry.
	// The type bits are a subset of the usual FileMode bits, those returned by the FileMode.Type method.
	Type() FileMode

	// Info returns the FileInfo for the file or subdirectory described by the entry.
	// The returned FileInfo may be from the time of the original directory read
	// or from the time of the call to Info. If the file has been removed or renamed
	// since the directory read, Info may return an error satisfying errors.Is(err, ErrNotExist).
	// If the entry denotes a symbolic link, Info reports the information about the link itself,
	// not the link's target.
	Info() (FileInfo, error)
}
```

A DirEntry is an entry read from a directory (using the ReadDir function or a ReadDirFile's ReadDir method).

DirEntry是一个从目录中读取的条目（使用ReadDir函数或ReadDirFile的ReadDir方法）。

### func FileInfoToDirEntry

`func FileInfoToDirEntry(info FileInfo) DirEntry`

FileInfoToDirEntry returns a DirEntry that returns information from info. If info is nil, FileInfoToDirEntry returns nil.

FileInfoToDirEntry 返回一个 DirEntry，它从 info 中返回信息。如果info是nil，FileInfoToDirEntry返回nil。

### func ReadDir 

`func ReadDir(fsys FS, name string) ([]DirEntry, error)`

ReadDir reads the named directory and returns a list of directory entries sorted by filename.

If fs implements ReadDirFS, ReadDir calls fs.ReadDir. Otherwise ReadDir calls fs.Open and uses ReadDir and Close on the returned file.

ReadDir 读取命名的目录并返回一个按文件名排序的目录条目列表。

如果fs实现了ReadDirFS，ReadDir调用fs.ReadDir。否则ReadDir调用fs.Open并对返回的文件使用ReadDir和Close。

## type FS

```
type FS interface {
	// Open opens the named file.
	//
	// When Open returns an error, it should be of type *PathError
	// with the Op field set to "open", the Path field set to name,
	// and the Err field describing the problem.
	//
	// Open should reject attempts to open names that do not satisfy
	// ValidPath(name), returning a *PathError with Err set to
	// ErrInvalid or ErrNotExist.
	Open(name string) (File, error)
}
```

An FS provides access to a hierarchical file system.

The FS interface is the minimum implementation required of the file system. A file system may implement additional interfaces, such as ReadFileFS, to provide additional or optimized functionality.

一个FS提供对分层文件系统的访问。

FS接口是文件系统所需的最小实现。一个文件系统可以实现额外的接口，如ReadFileFS，以提供额外的或优化的功能。

### func Sub

`func Sub(fsys FS, dir string) (FS, error)`

Sub returns an FS corresponding to the subtree rooted at fsys's dir.

If dir is ".", Sub returns fsys unchanged. Otherwise, if fs implements SubFS, Sub returns fsys.Sub(dir). Otherwise, Sub returns a new FS implementation sub that, in effect, implements sub.Open(name) as fsys.Open(path.Join(dir, name)). The implementation also translates calls to ReadDir, ReadFile, and Glob appropriately.

Note that Sub(os.DirFS("/"), "prefix") is equivalent to os.DirFS("/prefix") and that neither of them guarantees to avoid operating system accesses outside "/prefix", because the implementation of os.DirFS does not check for symbolic links inside "/prefix" that point to other directories. That is, os.DirFS is not a general substitute for a chroot-style security mechanism, and Sub does not change that fact.

Sub返回一个对应于以fsys的dir为根的子树的FS。

如果dir是"."，Sub返回fsys而不改变。否则，如果 fs 实现了 SubFS，Sub 返回 fsys.Sub(dir)。否则，Sub返回一个新的FS实现sub，实际上，它将sub.Open(name)实现为fsys.Open(path.Join(dir, name))。该实现还适当地翻译了对 ReadDir、ReadFile 和 Glob 的调用。

注意Sub(os.DirFS("/"), "prefix")等同于os.DirFS("/prefix")，它们都不能保证避免操作系统对"/prefix "之外的访问，因为os.DirFS的实现并不检查"/prefix "内指向其他目录的符号链接。也就是说，os.DirFS并不是chroot式安全机制的一般替代品，Sub并不能改变这一事实。


## type File

```
type File interface {
	Stat() (FileInfo, error)
	Read([]byte) (int, error)
	Close() error
}
```

A File provides access to a single file. The File interface is the minimum implementation required of the file. Directory files should also implement ReadDirFile. A file may implement io.ReaderAt or io.Seeker as optimizations.

一个文件提供对单个文件的访问。File接口是文件所需的最小实现。目录文件也应该实现ReadDirFile。一个文件可以实现io.ReaderAt或io.Seeker作为优化。

## type FileInfo

```
type FileInfo interface {
	Name() string       // base name of the file
	Size() int64        // length in bytes for regular files; system-dependent for others
	Mode() FileMode     // file mode bits
	ModTime() time.Time // modification time
	IsDir() bool        // abbreviation for Mode().IsDir()
	Sys() interface{}   // underlying data source (can return nil)
}
```

A FileInfo describes a file and is returned by Stat.

一个FileInfo描述了一个文件，并由Stat返回。

### func Stat

`func Stat(fsys FS, name string) (FileInfo, error)`

Stat returns a FileInfo describing the named file from the file system.

If fs implements StatFS, Stat calls fs.Stat. Otherwise, Stat opens the file to stat it.

Stat返回一个描述来自文件系统的命名文件的FileInfo。

如果fs实现了StatFS，Stat调用fs.Stat。否则，Stat会打开文件来统计它。

## type FileMode

`type FileMode uint32`

A FileMode represents a file's mode and permission bits. The bits have the same definition on all systems, so that information about files can be moved from one system to another portably. Not all bits apply to all systems. The only required bit is ModeDir for directories.

一个FileMode代表一个文件的模式和权限位。这些位在所有系统上都有相同的定义，因此有关文件的信息可以从一个系统移植到另一个系统。不是所有的位都适用于所有的系统。唯一需要的位是目录的ModeDir。

```
const (
	// The single letters are the abbreviations
	// used by the String method's formatting.
	ModeDir        FileMode = 1 << (32 - 1 - iota) // d: is a directory
	ModeAppend                                     // a: append-only
	ModeExclusive                                  // l: exclusive use
	ModeTemporary                                  // T: temporary file; Plan 9 only
	ModeSymlink                                    // L: symbolic link
	ModeDevice                                     // D: device file
	ModeNamedPipe                                  // p: named pipe (FIFO)
	ModeSocket                                     // S: Unix domain socket
	ModeSetuid                                     // u: setuid
	ModeSetgid                                     // g: setgid
	ModeCharDevice                                 // c: Unix character device, when ModeDevice is set
	ModeSticky                                     // t: sticky
	ModeIrregular                                  // ?: non-regular file; nothing else is known about this file

	// Mask for the type bits. For regular files, none will be set.
	ModeType = ModeDir | ModeSymlink | ModeNamedPipe | ModeSocket | ModeDevice | ModeCharDevice | ModeIrregular

	ModePerm FileMode = 0777 // Unix permission bits
)
```

The defined file mode bits are the most significant bits of the FileMode. The nine least-significant bits are the standard Unix rwxrwxrwx permissions. The values of these bits should be considered part of the public API and may be used in wire protocols or disk representations: they must not be changed, although new bits might be added.

定义的文件模式位是FileMode中最重要的位。九个最不重要的位是标准的Unix rwxrwxrwx权限。这些位的值应该被认为是公共API的一部分，可以在线程协议或磁盘表示法中使用：它们不能被改变，尽管可能会添加新的位。

### func (FileMode) IsDir

`func (m FileMode) IsDir() bool`

IsDir reports whether m describes a directory. That is, it tests for the ModeDir bit being set in m.

IsDir报告m是否描述了一个目录。也就是说，它测试m中的ModeDir位是否被设置。

### func (FileMode) IsRegular

`func (m FileMode) IsRegular() bool`

IsRegular reports whether m describes a regular file. That is, it tests that no mode type bits are set.

IsRegular报告m是否描述了一个常规文件。也就是说，它测试是否没有设置模式类型位。

### func (FileMode) Perm

`func (m FileMode) Perm() FileMode`

Perm returns the Unix permission bits in m (m & ModePerm).

Perm返回m中的Unix权限位（m & ModePerm）。

### func (FileMode) String 

`func (m FileMode) String() string`

### func (FileMode) Type 

`func (m FileMode) Type() FileMode`

Type returns type bits in m (m & ModeType).

类型返回m中的类型位（m & ModeType）。

## type GlobFS

```
type GlobFS interface {
	FS

	// Glob returns the names of all files matching pattern,
	// providing an implementation of the top-level
	// Glob function.
	Glob(pattern string) ([]string, error)
}
```

A GlobFS is a file system with a Glob method.

GlobFS是一个具有Glob方法的文件系统。

## type PathError 

```
type PathError struct {
	Op   string
	Path string
	Err  error
}
```

PathError records an error and the operation and file path that caused it.

PathError记录了一个错误以及导致该错误的操作和文件路径。

### func (*PathError) Error

`func (e *PathError) Error() string`

### func (*PathError) Timeout 

`func (e *PathError) Timeout() bool`

Timeout reports whether this error represents a timeout.

超时报告该错误是否代表超时。

### func (*PathError) Unwrap

`func (e *PathError) Unwrap() error`

## type ReadDirFS

```
type ReadDirFS interface {
	FS

	// ReadDir reads the named directory
	// and returns a list of directory entries sorted by filename.
	ReadDir(name string) ([]DirEntry, error)
}
```

ReadDirFS is the interface implemented by a file system that provides an optimized implementation of ReadDir.

ReadDirFS是由一个文件系统实现的接口，它提供了ReadDir的优化实现。

## type ReadDirFile

```
type ReadDirFile interface {
	File

	// ReadDir reads the contents of the directory and returns
	// a slice of up to n DirEntry values in directory order.
	// Subsequent calls on the same file will yield further DirEntry values.
	//
	// If n > 0, ReadDir returns at most n DirEntry structures.
	// In this case, if ReadDir returns an empty slice, it will return
	// a non-nil error explaining why.
	// At the end of a directory, the error is io.EOF.
	//
	// If n <= 0, ReadDir returns all the DirEntry values from the directory
	// in a single slice. In this case, if ReadDir succeeds (reads all the way
	// to the end of the directory), it returns the slice and a nil error.
	// If it encounters an error before the end of the directory,
	// ReadDir returns the DirEntry list read until that point and a non-nil error.
	ReadDir(n int) ([]DirEntry, error)
}
```

A ReadDirFile is a directory file whose entries can be read with the ReadDir method. Every directory file should implement this interface. (It is permissible for any file to implement this interface, but if so ReadDir should return an error for non-directories.)

一个ReadDirFile是一个目录文件，其条目可以用ReadDir方法读取。每个目录文件都应该实现这个接口。(任何文件都可以实现这个接口，但如果是这样的话，ReadDir应该对非目录文件返回一个错误。)

## type ReadFileFS

```
type ReadFileFS interface {
	FS

	// ReadFile reads the named file and returns its contents.
	// A successful call returns a nil error, not io.EOF.
	// (Because ReadFile reads the whole file, the expected EOF
	// from the final Read is not treated as an error to be reported.)
	//
	// The caller is permitted to modify the returned byte slice.
	// This method should return a copy of the underlying data.
	ReadFile(name string) ([]byte, error)
}
```

ReadFileFS is the interface implemented by a file system that provides an optimized implementation of ReadFile.

ReadFileFS 是由文件系统实现的接口，它提供了 ReadFile 的优化实现。


## type StatFS

```
type StatFS interface {
	FS

	// Stat returns a FileInfo describing the file.
	// If there is an error, it should be of type *PathError.
	Stat(name string) (FileInfo, error)
}
```

A StatFS is a file system with a Stat method.

一个StatFS是一个具有Stat方法的文件系统。

## type SubFS

```
type SubFS interface {
	FS

	// Sub returns an FS corresponding to the subtree rooted at dir.
	Sub(dir string) (FS, error)
}
```

A SubFS is a file system with a Sub method.

SubFS是一个具有Sub方法的文件系统。

## type WalkDirFunc

```
type WalkDirFunc func(path string, d DirEntry, err error) error
```

WalkDirFunc is the type of the function called by WalkDir to visit each file or directory.

WalkDirFunc是WalkDir为访问每个文件或目录而调用的函数的类型。









