
## Overview

Package os provides a platform-independent interface to operating system functionality. The design is Unix-like, although the error handling is Go-like; failing calls return values of type error rather than error numbers. Often, more information is available within the error. For example, if a call that takes a file name fails, such as Open or Stat, the error will include the failing file name when printed and will be of type *PathError, which may be unpacked for more information

包os提供了一个独立于平台的操作系统功能接口。该设计类似于Unix，尽管错误处理类似于Go；失败的调用返回错误类型的值而不是错误数字。通常，在错误中可以得到更多的信息。例如，如果一个需要文件名的调用失败了，如Open或Stat，错误在打印时将包括失败的文件名，并且是*PathError类型，它可以被解包以获得更多信息。

## Examples


[Examples-url](https://pkg.go.dev/os@go1.17.6#pkg-examples)

## Constants

```
const (
	// Exactly one of O_RDONLY, O_WRONLY, or O_RDWR must be specified.
	O_RDONLY int = syscall.O_RDONLY // open the file read-only.
	O_WRONLY int = syscall.O_WRONLY // open the file write-only.
	O_RDWR   int = syscall.O_RDWR   // open the file read-write.
	// The remaining values may be or'ed in to control behavior.
	O_APPEND int = syscall.O_APPEND // append data to the file when writing.
	O_CREATE int = syscall.O_CREAT  // create a new file if none exists.
	O_EXCL   int = syscall.O_EXCL   // used with O_CREATE, file must not exist.
	O_SYNC   int = syscall.O_SYNC   // open for synchronous I/O.
	O_TRUNC  int = syscall.O_TRUNC  // truncate regular writable file when opened.
)
```

Flags to OpenFile wrapping those of the underlying system. Not all flags may be implemented on a given system.

OpenFile的标志是对底层系统的包装。并非所有的标志都能在一个特定的系统上实现。

------

```

const (
	SEEK_SET int = 0 // seek relative to the origin of the file
	SEEK_CUR int = 1 // seek relative to the current offset
	SEEK_END int = 2 // seek relative to the end
)
```

Seek whence values.

Deprecated: Use io.SeekStart, io.SeekCurrent, and io.SeekEnd.

寻求从哪里来的值。

已经废弃。使用io.SeekStart, io.SeekCurrent, 和io.SeekEnd。


-----

```
const (
	PathSeparator     = '/' // OS-specific path separator
	PathListSeparator = ':' // OS-specific path list separator
)
```

-----

```
const (
	// The single letters are the abbreviations
	// used by the String method's formatting.
	ModeDir        = fs.ModeDir        // d: is a directory
	ModeAppend     = fs.ModeAppend     // a: append-only
	ModeExclusive  = fs.ModeExclusive  // l: exclusive use
	ModeTemporary  = fs.ModeTemporary  // T: temporary file; Plan 9 only
	ModeSymlink    = fs.ModeSymlink    // L: symbolic link
	ModeDevice     = fs.ModeDevice     // D: device file
	ModeNamedPipe  = fs.ModeNamedPipe  // p: named pipe (FIFO)
	ModeSocket     = fs.ModeSocket     // S: Unix domain socket
	ModeSetuid     = fs.ModeSetuid     // u: setuid
	ModeSetgid     = fs.ModeSetgid     // g: setgid
	ModeCharDevice = fs.ModeCharDevice // c: Unix character device, when ModeDevice is set
	ModeSticky     = fs.ModeSticky     // t: sticky
	ModeIrregular  = fs.ModeIrregular  // ?: non-regular file; nothing else is known about this file

	// Mask for the type bits. For regular files, none will be set.
	ModeType = fs.ModeType

	ModePerm = fs.ModePerm // Unix permission bits, 0o777
)
```

The defined file mode bits are the most significant bits of the FileMode. The nine least-significant bits are the standard Unix rwxrwxrwx permissions. The values of these bits should be considered part of the public API and may be used in wire protocols or disk representations: they must not be changed, although new bits might be added.

定义的文件模式位是FileMode中最重要的位。九个最不重要的位是标准的Unix rwxrwxrwx权限。这些位的值应该被认为是公共API的一部分，可以在线程协议或磁盘表示法中使用：它们不能被改变，尽管可能会添加新的位。

------


`const DevNull = "/dev/null"`

DevNull is the name of the operating system's “null device.” On Unix-like systems, it is "/dev/null"; on Windows, "NUL".

DevNull是操作系统的 "空设备 "的名称。在类Unix系统中，它是"/dev/null"；在Windows中，是 "NUL"。



## Variables


```
var (
	// ErrInvalid indicates an invalid argument.
	// Methods on File will return this error when the receiver is nil.
	ErrInvalid = fs.ErrInvalid // "invalid argument"

	ErrPermission = fs.ErrPermission // "permission denied"
	ErrExist      = fs.ErrExist      // "file already exists"
	ErrNotExist   = fs.ErrNotExist   // "file does not exist"
	ErrClosed     = fs.ErrClosed     // "file already closed"

	ErrNoDeadline       = errNoDeadline()       // "file type does not support deadline"
	ErrDeadlineExceeded = errDeadlineExceeded() // "i/o timeout"
)
```

Portable analogs of some common system call errors.

Errors returned from this package may be tested against these errors with errors.Is.

一些常见的系统调用错误的便携类似物。

从该包返回的错误可以用error.Is.测试这些错误。

-----

```
var (
	Stdin  = NewFile(uintptr(syscall.Stdin), "/dev/stdin")
	Stdout = NewFile(uintptr(syscall.Stdout), "/dev/stdout")
	Stderr = NewFile(uintptr(syscall.Stderr), "/dev/stderr")
)
```

Stdin, Stdout, and Stderr are open Files pointing to the standard input, standard output, and standard error file descriptors.

Note that the Go runtime writes to standard error for panics and crashes; closing Stderr may cause those messages to go elsewhere, perhaps to a file opened later.

Stdin、Stdout和Stderr是打开的文件，指向标准输入、标准输出和标准错误文件描述符。

请注意，Go 运行时间会将恐慌和崩溃的信息写入标准错误；关闭 Stderr 可能会导致这些信息流向其他地方，可能是后来打开的文件。

------

`var Args []string`

Args hold the command-line arguments, starting with the program name.

Args存放命令行参数，从程序名称开始。

-----

`var ErrProcessDone = errors.New("os: process already finished")`

ErrProcessDone indicates a Process has finished.

ErrProcessDone表示一个进程已经完成。


## Functions

### func Chdir

`func Chdir(dir string) error`


Chdir changes the current working directory to the named directory. If there is an error, it will be of type *PathError.

Chdir将当前工作目录改为指定的目录。如果有错误，它将是*PathError类型。

### func Chmod

`func Chmod(name string, mode FileMode) error`

Chmod changes the mode of the named file to mode. If the file is a symbolic link, it changes the mode of the link's target. If there is an error, it will be of type *PathError.

A different subset of the mode bits are used, depending on the operating system.

On Unix, the mode's permission bits, ModeSetuid, ModeSetgid, and ModeSticky are used.

On Windows, only the 0200 bit (owner writable) of mode is used; it controls whether the file's read-only attribute is set or cleared. The other bits are currently unused. For compatibility with Go 1.12 and earlier, use a non-zero mode. Use mode 0400 for a read-only file and 0600 for a readable+writable file.

On Plan 9, the mode's permission bits, ModeAppend, ModeExclusive, and ModeTemporary are used.


Chmod将命名文件的模式改为模式。如果该文件是一个符号链接，它将改变该链接的目标模式。如果有一个错误，它将是*PathError类型的。

模式位的不同子集被使用，取决于操作系统。

在Unix上，模式的权限位、ModeSetuid、ModeSetgid和ModeSticky被使用。

在Windows上，只有模式的0200位（所有者可写）被使用；它控制文件的只读属性是否被设置或清除。其他位目前没有使用。为了与Go 1.12和更早的版本兼容，请使用非零的模式。对只读文件使用模式0400，对可读+可写文件使用模式0600。

在计划9中，使用模式的权限位，ModeAppend，ModeExclusive，和ModeTemporary。


### func Chown

`func Chown(name string, uid, gid int) error`

Chown changes the numeric uid and gid of the named file. If the file is a symbolic link, it changes the uid and gid of the link's target. A uid or gid of -1 means to not change that value. If there is an error, it will be of type *PathError.

On Windows or Plan 9, Chown always returns the syscall.EWINDOWS or EPLAN9 error, wrapped in *PathError.

Chown改变命名文件的数字uid和gid。如果该文件是一个符号链接，它会改变该链接目标的uid和gid。uid或gid为-1意味着不改变该值。如果有一个错误，它将是*PathError类型的。

在Windows或Plan 9上，Chown总是返回syscall.EWINDOWS或EPLAN9的错误，被包裹在*PathError中。

### func Chtimes

`func Chtimes(name string, atime time.Time, mtime time.Time) error`

Chtimes changes the access and modification times of the named file, similar to the Unix utime() or utimes() functions.

The underlying filesystem may truncate or round the values to a less precise time unit. If there is an error, it will be of type *PathError.

Chtimes改变指定文件的访问和修改时间，类似于Unix的utime()或utimes()函数。

底层文件系统可能会截断或四舍五入到一个不太精确的时间单位。如果有错误，它将是*PathError类型。

### func Clearenv

`func Clearenv()`

Clearenv deletes all environment variables.

Clearenv删除了所有的环境变量。

### func DirFS

`func DirFS(dir string) fs.FS`

DirFS returns a file system (an fs.FS) for the tree of files rooted at the directory dir.

Note that DirFS("/prefix") only guarantees that the Open calls it makes to the operating system will begin with "/prefix": DirFS("/prefix").Open("file") is the same as os.Open("/prefix/file"). So if /prefix/file is a symbolic link pointing outside the /prefix tree, then using DirFS does not stop the access any more than using os.Open does. DirFS is therefore not a general substitute for a chroot-style security mechanism when the directory tree contains arbitrary content.

DirFS为以目录dir为根的文件树返回一个文件系统（fs.FS）。

注意，DirFS("/prefix")只保证它向操作系统发出的Open调用将以"/prefix "开始。DirFS("/prefix").Open("file")与os.Open("/prefix/file")相同。因此，如果/prefix/file是一个指向/prefix树外的符号链接，那么使用DirFS并不会像使用os.Open那样停止访问。因此，当目录树包含任意内容时，DirFS不能普遍替代chroot式的安全机制。

### func Environ

`func Environ() []string`

Environ returns a copy of strings representing the environment, in the form "key=value".

Environ返回一个代表环境的字符串的副本，形式为 "key=value"。

### func Executable

`func Executable() (string, error)`

Executable returns the path name for the executable that started the current process. There is no guarantee that the path is still pointing to the correct executable. If a symlink was used to start the process, depending on the operating system, the result might be the symlink or the path it pointed to. If a stable result is needed, path/filepath.EvalSymlinks might help.

Executable returns an absolute path unless an error occurred.

The main use case is finding resources located relative to an executable.

Executable返回启动当前进程的可执行文件的路径名称。不能保证该路径仍然指向正确的可执行文件。如果一个符号链接被用来启动进程，根据操作系统的不同，结果可能是符号链接或它所指向的路径。如果需要一个稳定的结果，path/filepath.EvalSymlinks可能有帮助。

Executable返回一个绝对路径，除非发生错误。

主要的用例是寻找相对于可执行文件的资源。

### func Exit

`func Exit(code int)`

Exit causes the current program to exit with the given status code. Conventionally, code zero indicates success, non-zero an error. The program terminates immediately; deferred functions are not run.

For portability, the status code should be in the range [0, 125].


Exit使当前程序以给定的状态代码退出。传统上，代码0表示成功，非0表示错误。程序立即终止；延迟的函数不被运行。

为了便于移植，状态码应该在[0, 125]范围内。


### func Expand

`func Expand(s string, mapping func(string) string) string`

Expand replaces ${var} or $var in the string based on the mapping function. For example, os.ExpandEnv(s) is equivalent to os.Expand(s, os.Getenv).

Expand 会根据映射函数替换字符串中的 ${var} 或 $var。例如，os.ExpandEnv(s)等同于os.Expand(s, os.Getenv)。

### func ExpandEnv

`func ExpandEnv(s string) string`

ExpandEnv replaces ${var} or $var in the string according to the values of the current environment variables. References to undefined variables are replaced by the empty string.

ExpandEnv根据当前环境变量的值，替换字符串中的${var}或$var。对未定义变量的引用会被替换为空字符串。

### func Getegid

`func Getegid() int`

Getegid returns the numeric effective group id of the caller.

On Windows, it returns -1.

Getegid返回调用者的数字有效组ID。

在Windows下，它返回-1。


### func Getenv

`func Getenv(key string) string`

Getenv retrieves the value of the environment variable named by the key. It returns the value, which will be empty if the variable is not present. To distinguish between an empty value and an unset value, use LookupEnv.

Getenv 检索由键命名的环境变量的值。它返回该值，如果该变量不存在，则该值为空。为了区分空值和未设置的值，使用LookupEnv。

### func Geteuid 

`func Geteuid() int`

Geteuid returns the numeric effective user id of the caller.

On Windows, it returns -1.

Geteuid返回调用者的数字有效用户ID。

在Windows下，它返回-1。

### func Getgid 

`func Getgid() int`

Getgid returns the numeric group id of the caller.

On Windows, it returns -1.

Getgid返回调用者的数字组ID。

在Windows下，它返回-1。


### func Getgroups 

`func Getgroups() ([]int, error)`

Getgroups returns a list of the numeric ids of groups that the caller belongs to.

On Windows, it returns syscall.EWINDOWS. See the os/user package for a possible alternative.

Getgroups返回调用者所属的组的数字ID列表。

在Windows上，它返回syscall.EWINDOWS。参见os/user软件包，以获得一个可能的替代方案。

### func Getpagesize 

`func Getpagesize() int`

Getpagesize returns the underlying system's memory page size.

Getpagesize返回底层系统的内存页面大小。

### func Getpid

`func Getpid() int`

Getpid returns the process id of the caller.

Getpid返回调用者的进程ID。

### func Getppid 

`func Getppid() int`

Getppid returns the process id of the caller's parent.

Getppid返回调用者的父进程ID。

### func Getuid

`func Getuid() int`

Getuid returns the numeric user id of the caller.

On Windows, it returns -1.

Getuid返回调用者的数字用户ID。

在Windows下，它返回-1。

### func Getwd

`func Getwd() (dir string, err error)`

Getwd returns a rooted path name corresponding to the current directory. If the current directory can be reached via multiple paths (due to symbolic links), Getwd may return any one of them.

Getwd返回对应于当前目录的有根路径名称。如果当前目录可以通过多个路径到达（由于符号链接），Getwd可以返回其中任何一个。




















