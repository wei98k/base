
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


### func Hostname

`func Hostname() (name string, err error)`

Hostname returns the host name reported by the kernel.

Hostname返回内核报告的主机名。

### func IsExist

`func IsExist(err error) bool`

IsExist returns a boolean indicating whether the error is known to report that a file or directory already exists. It is satisfied by ErrExist as well as some syscall errors.

This function predates errors.Is. It only supports errors returned by the os package. New code should use errors.Is(err, fs.ErrExist).

IsExist返回一个布尔值，表明该错误是否知道报告一个文件或目录已经存在。它被ErrExist以及一些系统调用错误所满足。

这个函数比 errors.Is 更早。它只支持由os包返回的错误。新代码应该使用 errors.Is(err, fs.ErrExist)。

### func IsNotExist

`func IsNotExist(err error) bool`

IsNotExist returns a boolean indicating whether the error is known to report that a file or directory does not exist. It is satisfied by ErrNotExist as well as some syscall errors.

This function predates errors.Is. It only supports errors returned by the os package. New code should use errors.Is(err,

IsNotExist返回一个布尔值，表示错误是否知道报告文件或目录不存在。它被ErrNotExist以及一些syscall错误满足。

这个函数比 errors.Is 更早。它只支持由os包返回的错误。新代码应该使用 errors.Is(err,

### func IsPathSeparator 

`func IsPathSeparator(c uint8) bool`

IsPathSeparator reports whether c is a directory separator character.

IsPathSeparator报告c是否是一个目录分隔符。

### func IsPermission

`func IsPermission(err error) bool`

IsPermission returns a boolean indicating whether the error is known to report that permission is denied. It is satisfied by ErrPermission as well as some syscall errors.

This function predates errors.Is. It only supports errors returned by the os package. New code should use errors.Is(err, fs.ErrPermission).

IsPermission返回一个布尔值，表明该错误是否知道报告权限被拒绝。它被ErrPermission以及一些系统调用错误所满足。

这个函数比 errors.Is 更早。它只支持由os包返回的错误。新代码应该使用 errors.Is(err, fs.ErrPermission)。

### func IsTimeout 

`func IsTimeout(err error) bool`

IsTimeout returns a boolean indicating whether the error is known to report that a timeout occurred.

This function predates errors.Is, and the notion of whether an error indicates a timeout can be ambiguous. For example, the Unix error EWOULDBLOCK sometimes indicates a timeout and sometimes does not. New code should use errors.Is with a value appropriate to the call returning the error, such as os.ErrDeadlineExceeded.

IsTimeout 返回一个布尔值，表明该错误是否已知报告发生了超时。

这个函数早于 errors.Is，而且一个错误是否表示超时的概念可能是不明确的。例如，Unix的错误EWOULDBLOCK有时表示超时，有时不表示。新的代码应该使用 errors.Is，其值适合于返回错误的调用，例如 os.ErrDeadlineExceeded。

### func Lchown 

`func Lchown(name string, uid, gid int) error`

Lchown changes the numeric uid and gid of the named file. If the file is a symbolic link, it changes the uid and gid of the link itself. If there is an error, it will be of type *PathError.

On Windows, it always returns the syscall.EWINDOWS error, wrapped in *PathError.

Lchown改变命名文件的数字uid和gid。如果该文件是一个符号链接，它将改变该链接本身的uid和gid。如果有一个错误，它将是*PathError类型的。

在Windows上，它总是返回syscall.EWINDOWS错误，被包裹在*PathError中。


### func Link

`func Link(oldname, newname string) error`

Link creates newname as a hard link to the oldname file. If there is an error, it will be of type *LinkError.

链接创建newname作为一个硬链接到oldname文件。如果有错误，它将是*LinkError类型。

### func LookupEnv 

`func LookupEnv(key string) (string, bool)`

LookupEnv retrieves the value of the environment variable named by the key. If the variable is present in the environment the value (which may be empty) is returned and the boolean is true. Otherwise the returned value will be empty and the boolean will be false.

LookupEnv检索由键命名的环境变量的值。如果该变量存在于环境中，则返回值（可能为空），布尔值为真。否则，返回的值将是空的，布尔值将是假的。

### func Mkdir

`func Mkdir(name string, perm FileMode) error`

Mkdir creates a new directory with the specified name and permission bits (before umask). If there is an error, it will be of type *PathError.

Mkdir用指定的名称和权限位（在umask之前）创建一个新目录。如果有错误，它将是*PathError类型。

### func MkdirAll

`func MkdirAll(path string, perm FileMode) error`

MkdirAll creates a directory named path, along with any necessary parents, and returns nil, or else returns an error. The permission bits perm (before umask) are used for all directories that MkdirAll creates. If path is already a directory, MkdirAll does nothing and returns nil.

MkdirAll创建一个名为path的目录，以及任何必要的父目录，并返回nil，否则返回一个错误。权限位perm（在umask之前）被用于MkdirAll创建的所有目录。如果path已经是一个目录，MkdirAll不做任何事情，并返回nil。

### func MkdirTemp

`func MkdirTemp(dir, pattern string) (string, error)`

MkdirTemp creates a new temporary directory in the directory dir and returns the pathname of the new directory. The new directory's name is generated by adding a random string to the end of pattern. If pattern includes a "*", the random string replaces the last "*" instead. If dir is the empty string, MkdirTemp uses the default directory for temporary files, as returned by TempDir. Multiple programs or goroutines calling MkdirTemp simultaneously will not choose the same directory. It is the caller's responsibility to remove the directory when it is no longer needed.

MkdirTemp在目录dir中创建一个新的临时目录并返回新目录的路径名。新目录的名称是通过在pattern的末尾添加一个随机字符串生成的。如果pattern包括一个 "*"，随机字符串将代替最后的 "*"。如果dir是空字符串，MkdirTemp使用TempDir返回的默认目录作为临时文件。多个程序或goroutine同时调用MkdirTemp不会选择相同的目录。当不再需要该目录时，由调用者负责删除该目录。

### func NewSyscallError

`func NewSyscallError(syscall string, err error) error`

NewSyscallError returns, as an error, a new SyscallError with the given system call name and error details. As a convenience, if err is nil, NewSyscallError returns nil.

NewSyscallError作为一个错误返回一个新的SyscallError，并给出系统调用名称和错误细节。为了方便起见，如果err为nil，NewSyscallError返回nil。

### func Pipe

`func Pipe() (r *File, w *File, err error)`

Pipe returns a connected pair of Files; reads from r return bytes written to w. It returns the files and an error, if any.

Pipe返回一对相连的Files；从r读取的数据返回写给w的字节，如果有的话，它返回文件和错误。

### func ReadFile

`func ReadFile(name string) ([]byte, error)`

ReadFile reads the named file and returns the contents. A successful call returns err == nil, not err == EOF. Because ReadFile reads the whole file, it does not treat an EOF from Read as an error to be reported.

ReadFile 读取命名的文件并返回其内容。一个成功的调用返回 err == nil, 而不是 err == EOF。因为ReadFile读取了整个文件，它不把来自Read的EOF当作一个要报告的错误。

### func Readlink

`func Readlink(name string) (string, error)`

Readlink returns the destination of the named symbolic link. If there is an error, it will be of type *PathError.

Readlink返回命名的符号链接的目的地。如果有一个错误，它将是*PathError类型的。

### func Remove

`func Remove(name string) error`

Remove removes the named file or (empty) directory. If there is an error, it will be of type *PathError.

移除指定的文件或（空）目录。如果有一个错误，它将是*PathError类型的。

### func RemoveAll 

`func RemoveAll(path string) error`

RemoveAll removes path and any children it contains. It removes everything it can but returns the first error it encounters. If the path does not exist, RemoveAll returns nil (no error). If there is an error, it will be of type *PathError.

RemoveAll删除路径和它所包含的任何子代。它删除所有可以删除的东西，但会返回它遇到的第一个错误。如果路径不存在，RemoveAll返回nil（没有错误）。如果有一个错误，它将是*PathError类型的。


### func Rename

`func Rename(oldpath, newpath string) error`

Rename renames (moves) oldpath to newpath. If newpath already exists and is not a directory, Rename replaces it. OS-specific restrictions may apply when oldpath and newpath are in different directories. If there is an error, it will be of type *LinkError.

Rename将oldpath重命名（移动）到newpath。如果newpath已经存在，并且不是一个目录，Rename就会替换它。当oldpath和newpath在不同的目录中时，操作系统的特定限制可能适用。如果有一个错误，它将是*LinkError类型的。


### func SameFile

`func SameFile(fi1, fi2 FileInfo) bool`

SameFile reports whether fi1 and fi2 describe the same file. For example, on Unix this means that the device and inode fields of the two underlying structures are identical; on other systems the decision may be based on the path names. SameFile only applies to results returned by this package's Stat. It returns false in other cases.

SameFile报告fi1和fi2是否描述同一个文件。例如，在 Unix 上，这意味着两个底层结构的设备和 inode 字段是相同的；在其他系统上，这个决定可能是基于路径名称的。SameFile只适用于这个包的Stat所返回的结果。在其他情况下，它会返回false。

### func Setenv

`func Setenv(key, value string) error`

Setenv sets the value of the environment variable named by the key. It returns an error, if any.

Setenv设置由键命名的环境变量的值。如果有的话，它会返回一个错误。

### func Symlink

`func Symlink(oldname, newname string) error`

Symlink creates newname as a symbolic link to oldname. On Windows, a symlink to a non-existent oldname creates a file symlink; if oldname is later created as a directory the symlink will not work. If there is an error, it will be of type *LinkError.

Symlink将newname创建为oldname的符号链接。在Windows上，指向不存在的oldname的符号链接会创建一个文件符号链接；如果oldname后来被创建为一个目录，那么符号链接将不起作用。如果有错误，它将是*LinkError类型的。

### func TempDir 

`func TempDir() string`

TempDir returns the default directory to use for temporary files.

On Unix systems, it returns $TMPDIR if non-empty, else /tmp. On Windows, it uses GetTempPath, returning the first non-empty value from %TMP%, %TEMP%, %USERPROFILE%, or the Windows directory. On Plan 9, it returns /tmp.

The directory is neither guaranteed to exist nor have accessible permissions.

TempDir返回用于临时文件的默认目录。

在Unix系统中，如果非空，它返回$TMPDIR，否则返回/tmp。在Windows上，它使用GetTempPath，返回%TMP%、%TEMP%、%USERPROFILE%或Windows目录中第一个非空值。在计划9中，它返回/tmp。

该目录既不保证存在，也没有可访问的权限。

### func Truncate

`func Truncate(name string, size int64) error`

Truncate changes the size of the named file. If the file is a symbolic link, it changes the size of the link's target. If there is an error, it will be of type *PathError.

Truncate改变命名文件的大小。如果该文件是一个符号链接，它将改变该链接的目标大小。如果有一个错误，它将是*PathError类型的。


### func Unsetenv

`func Unsetenv(key string) error`

Unsetenv unsets a single environment variable.

Unsetenv取消一个环境变量。

### func UserCacheDir 

`func UserCacheDir() (string, error)`

UserCacheDir returns the default root directory to use for user-specific cached data. Users should create their own application-specific subdirectory within this one and use that.

UserCacheDir返回默认的根目录，用于用户特定的缓存数据。用户应该在这个目录中创建他们自己的特定应用的子目录并使用它。

### func UserConfigDir

`func UserConfigDir() (string, error)`

UserConfigDir returns the default root directory to use for user-specific configuration data. Users should create their own application-specific subdirectory within this one and use that.

UserConfigDir返回默认的根目录，用于用户特定的配置数据。用户应该在这个目录中创建他们自己的特定应用的子目录并使用它。

### func UserHomeDir 

`func UserHomeDir() (string, error)`

UserHomeDir returns the current user's home directory.

UserHomeDir返回当前用户的主目录。

### func WriteFile

`func WriteFile(name string, data []byte, perm FileMode) error`

WriteFile writes data to the named file, creating it if necessary. If the file does not exist, WriteFile creates it with permissions perm (before umask); otherwise WriteFile truncates it before writing, without changing permissions.

WriteFile 将数据写入被命名的文件，如果有必要的话，将创建它。如果该文件不存在，WriteFile 会以权限 perm (在 umask 之前) 创建它；否则，WriteFile 会在写入前将其截断，而不改变权限。

## type DirEntry 

```
type DirEntry = fs.DirEntry
```

A DirEntry is an entry read from a directory (using the ReadDir function or a File's ReadDir method).

DirEntry是一个从目录中读取的条目（使用ReadDir函数或File的ReadDir方法）。

### func ReadDir

`func ReadDir(name string) ([]DirEntry, error)`

ReadDir reads the named directory, returning all its directory entries sorted by filename. If an error occurs reading the directory, ReadDir returns the entries it was able to read before the error, along with the error.

ReadDir 读取命名的目录，返回其所有按文件名排序的目录条目。如果在读取目录时发生错误，ReadDir会返回它在出错前能够读取的条目，以及错误信息。


## type File

```
type File struct {
	// contains filtered or unexported fields
}
```

File represents an open file descriptor.

File代表一个开放的文件描述符。

### func Create

`func Create(name string) (*File, error)`

Create creates or truncates the named file. If the file already exists, it is truncated. If the file does not exist, it is created with mode 0666 (before umask). If successful, methods on the returned File can be used for I/O; the associated file descriptor has mode O_RDWR. If there is an error, it will be of type *PathError.

创建创建或截断命名的文件。如果该文件已经存在，它将被截断。如果文件不存在，将以0666模式（在umask之前）创建。如果成功，返回的文件上的方法可以用于I/O；相关的文件描述符具有O_RDWR模式。如果有一个错误，它将是*PathError类型的。

### func CreateTemp

`func CreateTemp(dir, pattern string) (*File, error)`

CreateTemp creates a new temporary file in the directory dir, opens the file for reading and writing, and returns the resulting file. The filename is generated by taking pattern and adding a random string to the end. If pattern includes a "*", the random string replaces the last "*". If dir is the empty string, CreateTemp uses the default directory for temporary files, as returned by TempDir. Multiple programs or goroutines calling CreateTemp simultaneously will not choose the same file. The caller can use the file's Name method to find the pathname of the file. It is the caller's responsibility to remove the file when it is no longer needed.

CreateTemp在目录dir中创建一个新的临时文件，打开该文件进行读写，并返回生成的文件。文件名是通过提取pattern并在末尾添加一个随机字符串生成的。如果pattern包括一个 "*"，随机字符串将取代最后的 "*"。如果dir是空字符串，CreateTemp使用临时文件的默认目录，如TempDir返回。多个程序或goroutine同时调用CreateTemp不会选择同一个文件。调用者可以使用文件的Name方法来查找文件的路径名。当不再需要该文件时，调用者有责任将其删除。

### func NewFile

`func NewFile(fd uintptr, name string) *File`

NewFile returns a new File with the given file descriptor and name. The returned value will be nil if fd is not a valid file descriptor. On Unix systems, if the file descriptor is in non-blocking mode, NewFile will attempt to return a pollable File (one for which the SetDeadline methods work).

After passing it to NewFile, fd may become invalid under the same conditions described in the comments of the Fd method, and the same constraints apply.

NewFile 返回一个具有给定文件描述符和名称的新文件。如果fd不是一个有效的文件描述符，返回值将是nil。在 Unix 系统中, 如果文件描述符是在非阻塞模式中, NewFile 会尝试返回一个可轮询的文件 (SetDeadline 方法对其有效).

在把它传递给 NewFile 之后，fd 可能会在 Fd 方法的注释中所描述的相同条件下变得无效，并且同样的限制条件也适用。

### func Open

`func Open(name string) (*File, error)`

Open opens the named file for reading. If successful, methods on the returned file can be used for reading; the associated file descriptor has mode O_RDONLY. If there is an error, it will be of type *PathError.

Open打开命名的文件以供阅读。如果成功，返回的文件上的方法可以用来阅读；相关的文件描述符有O_RDONLY模式。如果有一个错误，它将是*PathError类型。

### func OpenFile

`func OpenFile(name string, flag int, perm FileMode) (*File, error)`

OpenFile is the generalized open call; most users will use Open or Create instead. It opens the named file with specified flag (O_RDONLY etc.). If the file does not exist, and the O_CREATE flag is passed, it is created with mode perm (before umask). If successful, methods on the returned File can be used for I/O. If there is an error, it will be of type *PathError.

OpenFile是通用的打开调用；大多数用户会使用Open或Create代替。它用指定的标志（O_RDONLY等）打开命名的文件。如果文件不存在，并且传递了O_CREATE标志，它将以perm模式（在umask之前）创建。如果成功，返回的文件上的方法可以被用于I/O。如果有一个错误，它将是*PathError类型。

### func (*File) Chdir

`func (f *File) Chdir() error`

Chdir changes the current working directory to the file, which must be a directory. If there is an error, it will be of type *PathError.

Chdir将当前工作目录改为文件，该文件必须是一个目录。如果有一个错误，它将是*PathError类型。

### func (*File) Chmod

`func (f *File) Chmod(mode FileMode) error`

Chmod changes the mode of the file to mode. If there is an error, it will be of type *PathError.

Chmod将文件的模式改为模式。如果有一个错误，它将是*PathError类型。

### func (*File) Chown

`func (f *File) Chown(uid, gid int) error`

Chown changes the numeric uid and gid of the named file. If there is an error, it will be of type *PathError.

On Windows, it always returns the syscall.EWINDOWS error, wrapped in *PathError.

Chown改变命名文件的数字uid和gid。如果有一个错误，它将是*PathError类型的。

在Windows上，它总是返回syscall.EWINDOWS错误，被包裹在*PathError中。

### func (*File) Close 

`func (f *File) Close() error`

Close closes the File, rendering it unusable for I/O. On files that support SetDeadline, any pending I/O operations will be canceled and return immediately with an error. Close will return an error if it has already been called.

关闭文件，使其不能用于I/O。在支持SetDeadline的文件上，任何未决的I/O操作将被取消，并立即返回一个错误。如果已经被调用，Close将返回一个错误。

### func (*File) Fd

`func (f *File) Fd() uintptr`

Fd returns the integer Unix file descriptor referencing the open file. If f is closed, the file descriptor becomes invalid. If f is garbage collected, a finalizer may close the file descriptor, making it invalid; see runtime.SetFinalizer for more information on when a finalizer might be run. On Unix systems this will cause the SetDeadline methods to stop working. Because file descriptors can be reused, the returned file descriptor may only be closed through the Close method of f, or by its finalizer during garbage collection. Otherwise, during garbage collection the finalizer may close an unrelated file descriptor with the same (reused) number.

Fd返回整数的Unix文件描述符，引用打开的文件。如果f被关闭，文件描述符就会失效。如果f被垃圾回收，终结者可能会关闭文件描述符，使其失效；关于终结者何时可能被运行的更多信息，见runtime.SetFinalizer。在Unix系统上，这将导致SetDeadline方法停止工作。因为文件描述符可以被重复使用，返回的文件描述符只能通过f的Close方法关闭，或者在垃圾回收期间由其最终化器关闭。否则，在垃圾回收过程中，终结者可能会关闭一个具有相同（重复使用）编号的无关的文件描述符。

### func (*File) Name 

`func (f *File) Name() string`

Name returns the name of the file as presented to Open.

Name返回给Open的文件名。

### func (*File) Read

`func (f *File) Read(b []byte) (n int, err error)`

Read reads up to len(b) bytes from the File. It returns the number of bytes read and any error encountered. At end of file, Read returns 0, io.EOF.

读取从文件中最多读取len(b)个字节。它返回读取的字节数和遇到的任何错误。在文件结束时，Read 返回 0, io.EOF。

### func (*File) ReadAt

`func (f *File) ReadAt(b []byte, off int64) (n int, err error)`

ReadAt reads len(b) bytes from the File starting at byte offset off. It returns the number of bytes read and the error, if any. ReadAt always returns a non-nil error when n < len(b). At end of file, that error is io.EOF.

ReadAt从文件中读取len(b)字节，从字节偏移量off开始。它返回读取的字节数和错误，如果有的话。当n < len(b)时，ReadAt总是返回一个非零的错误。在文件结束时，这个错误是io.EOF。

### func (*File) ReadDir 

`func (f *File) ReadDir(n int) ([]DirEntry, error)`

ReadDir reads the contents of the directory associated with the file f and returns a slice of DirEntry values in directory order. Subsequent calls on the same file will yield later DirEntry records in the directory.

If n > 0, ReadDir returns at most n DirEntry records. In this case, if ReadDir returns an empty slice, it will return an error explaining why. At the end of a directory, the error is io.EOF.

If n <= 0, ReadDir returns all the DirEntry records remaining in the directory. When it succeeds, it returns a nil error (not io.EOF).

ReadDir读取与文件f相关的目录的内容，并按目录顺序返回DirEntry值的片断。对同一文件的后续调用将产生该目录中后来的DirEntry记录。

如果n>0，ReadDir最多返回n条DirEntry记录。在这种情况下，如果ReadDir返回一个空片，它将返回一个错误，解释为什么。在一个目录的末端，错误是io.EOF。

如果n<=0，ReadDir返回目录中剩余的所有DirEntry记录。当它成功时，它返回一个nil错误（不是io.EOF）

### func (*File) ReadFrom 

`func (f *File) ReadFrom(r io.Reader) (n int64, err error)`

ReadFrom implements io.ReaderFrom.

ReadFrom 实现了 io.ReaderFrom。


=====

TODO https://pkg.go.dev/os@go1.17.6#File.Readdir














