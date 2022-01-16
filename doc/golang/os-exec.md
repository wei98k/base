
## Overview

Package exec runs external commands. It wraps os.StartProcess to make it easier to remap stdin and stdout, connect I/O with pipes, and do other adjustments.

Unlike the "system" library call from C and other languages, the os/exec package intentionally does not invoke the system shell and does not expand any glob patterns or handle other expansions, pipelines, or redirections typically done by shells. The package behaves more like C's "exec" family of functions. To expand glob patterns, either call the shell directly, taking care to escape any dangerous input, or use the path/filepath package's Glob function. To expand environment variables, use package os's ExpandEnv.

Note that the examples in this package assume a Unix system. They may not run on Windows, and they do not run in the Go Playground used by golang.org and godoc.org.

包exec运行外部命令。它包装了os.StartProcess，使其更容易重新映射stdin和stdout，用管道连接I/O，并做其他调整。

与 C 和其他语言的 "系统 "库调用不同，os/exec 包有意不调用系统 shell，也不扩展任何 glob 模式或处理其他通常由 shell 完成的扩展、管道或重定向。这个包的行为更像 C 的 "exec" 系列函数。要扩展glob模式，可以直接调用shell，注意转义任何危险的输入，或者使用path/filepath包的Glob函数。要扩展环境变量，可以使用包os的ExpandEnv。

请注意，本软件包中的例子是假设Unix系统的。它们可能无法在 Windows 上运行，也无法在 golang.org 和 godoc.org 使用的 Go Playground 中运行。

## Examples

[Examples-url](https://pkg.go.dev/os/exec@go1.17.6#pkg-examples)

## Variables

```
var ErrNotFound = errors.New("executable file not found in $PATH")
```
ErrNotFound is the error resulting if a path search failed to find an executable file.

ErrNotFound是如果路径搜索未能找到可执行文件而产生的错误。

## Functions


### func LookPath

`func LookPath(file string) (string, error)`

LookPath searches for an executable named file in the directories named by the PATH environment variable. If file contains a slash, it is tried directly and the PATH is not consulted. The result may be an absolute path or a path relative to the current directory.

LookPath在由PATH环境变量命名的目录中搜索一个名为文件的可执行文件。如果文件包含一个斜线，就会直接尝试，而不参考PATH。其结果可能是一个绝对路径或相对于当前目录的路径。


## type Cmd

```
type Cmd struct {
	// Path is the path of the command to run.
	//
	// This is the only field that must be set to a non-zero
	// value. If Path is relative, it is evaluated relative
	// to Dir.
	Path string

	// Args holds command line arguments, including the command as Args[0].
	// If the Args field is empty or nil, Run uses {Path}.
	//
	// In typical use, both Path and Args are set by calling Command.
	Args []string

	// Env specifies the environment of the process.
	// Each entry is of the form "key=value".
	// If Env is nil, the new process uses the current process's
	// environment.
	// If Env contains duplicate environment keys, only the last
	// value in the slice for each duplicate key is used.
	// As a special case on Windows, SYSTEMROOT is always added if
	// missing and not explicitly set to the empty string.
	Env []string

	// Dir specifies the working directory of the command.
	// If Dir is the empty string, Run runs the command in the
	// calling process's current directory.
	Dir string

	// Stdin specifies the process's standard input.
	//
	// If Stdin is nil, the process reads from the null device (os.DevNull).
	//
	// If Stdin is an *os.File, the process's standard input is connected
	// directly to that file.
	//
	// Otherwise, during the execution of the command a separate
	// goroutine reads from Stdin and delivers that data to the command
	// over a pipe. In this case, Wait does not complete until the goroutine
	// stops copying, either because it has reached the end of Stdin
	// (EOF or a read error) or because writing to the pipe returned an error.
	Stdin io.Reader

	// Stdout and Stderr specify the process's standard output and error.
	//
	// If either is nil, Run connects the corresponding file descriptor
	// to the null device (os.DevNull).
	//
	// If either is an *os.File, the corresponding output from the process
	// is connected directly to that file.
	//
	// Otherwise, during the execution of the command a separate goroutine
	// reads from the process over a pipe and delivers that data to the
	// corresponding Writer. In this case, Wait does not complete until the
	// goroutine reaches EOF or encounters an error.
	//
	// If Stdout and Stderr are the same writer, and have a type that can
	// be compared with ==, at most one goroutine at a time will call Write.
	Stdout io.Writer
	Stderr io.Writer

	// ExtraFiles specifies additional open files to be inherited by the
	// new process. It does not include standard input, standard output, or
	// standard error. If non-nil, entry i becomes file descriptor 3+i.
	//
	// ExtraFiles is not supported on Windows.
	ExtraFiles []*os.File

	// SysProcAttr holds optional, operating system-specific attributes.
	// Run passes it to os.StartProcess as the os.ProcAttr's Sys field.
	SysProcAttr *syscall.SysProcAttr

	// Process is the underlying process, once started.
	Process *os.Process

	// ProcessState contains information about an exited process,
	// available after a call to Wait or Run.
	ProcessState *os.ProcessState
	// contains filtered or unexported fields
}
```

Cmd represents an external command being prepared or run.

A Cmd cannot be reused after calling its Run, Output or CombinedOutput methods.

Cmd代表一个正在准备或运行的外部命令。

一个Cmd在调用它的Run、Output或CombinedOutput方法后不能被重复使用。

### func Command

`func Command(name string, arg ...string) *Cmd`

Command returns the Cmd struct to execute the named program with the given arguments.

It sets only the Path and Args in the returned structure.

If name contains no path separators, Command uses LookPath to resolve name to a complete path if possible. Otherwise it uses name directly as Path.

The returned Cmd's Args field is constructed from the command name followed by the elements of arg, so arg should not include the command name itself. For example, Command("echo", "hello"). Args[0] is always name, not the possibly resolved Path.

On Windows, processes receive the whole command line as a single string and do their own parsing. Command combines and quotes Args into a command line string with an algorithm compatible with applications using CommandLineToArgvW (which is the most common way). Notable exceptions are msiexec.exe and cmd.exe (and thus, all batch files), which have a different unquoting algorithm. In these or other similar cases, you can do the quoting yourself and provide the full command line in SysProcAttr.CmdLine, leaving Args empty.

Command返回Cmd结构，用给定的参数执行指定的程序。

它只在返回的结构中设置Path和Args。

如果name不包含路径分隔符，Command使用LookPath将name解析为一个完整的路径。否则它直接使用name作为Path。

返回的Cmd的Args字段是由命令名和arg的元素组成的，所以arg不应该包括命令名本身。例如，Command("echo", "hello")。Args[0]总是名字，而不是可能解决的Path。

在Windows上，进程会将整个命令行作为一个字符串接收，并进行自己的解析。Command将Args合并并引用为一个命令行字符串，其算法与使用CommandLineToArgvW的应用程序兼容（这是最常见的方式）。值得注意的例外是msiexec.exe和cmd.exe（因此，所有批处理文件），它们有不同的解引算法。在这些或其他类似的情况下，你可以自己做引号，并在SysProcAttr.CmdLine中提供完整的命令行，把Args留空。

### func CommandContext

`func CommandContext(ctx context.Context, name string, arg ...string) *Cmd`

CommandContext is like Command but includes a context.

The provided context is used to kill the process (by calling os.Process.Kill) if the context becomes done before the command completes on its own.

CommandContext和Command一样，但包括一个上下文。

如果在命令自己完成之前，上下文已经完成，那么所提供的上下文将被用来杀死进程（通过调用os.Process.Kill）。

### func (*Cmd) CombinedOutput

`func (c *Cmd) CombinedOutput() ([]byte, error)`

CombinedOutput runs the command and returns its combined standard output and standard error.

CombinedOutput运行命令并返回其合并的标准输出和标准错误。

### func (*Cmd) Output

`func (c *Cmd) Output() ([]byte, error)`

Output runs the command and returns its standard output. Any returned error will usually be of type *ExitError. If c.Stderr was nil, Output populates ExitError.Stderr.

输出运行该命令并返回其标准输出。任何返回的错误通常是*ExitError类型的。如果c.Stderr是nil，Output会填充ExitError.Stderr。

### func (*Cmd) Run

`func (c *Cmd) Run() error`

Run starts the specified command and waits for it to complete.

The returned error is nil if the command runs, has no problems copying stdin, stdout, and stderr, and exits with a zero exit status.

If the command starts but does not complete successfully, the error is of type *ExitError. Other error types may be returned for other situations.

If the calling goroutine has locked the operating system thread with runtime.LockOSThread and modified any inheritable OS-level thread state (for example, Linux or Plan 9 name spaces), the new process will inherit the caller's thread state.

RUN启动指定的命令并等待其完成。

如果命令运行，复制stdin、stdout和stderr没有问题，并且以0退出状态退出，返回的错误是nil。

如果命令启动但没有成功完成，错误类型为*ExitError。在其他情况下可能会返回其他错误类型。

如果调用的goroutine用runtime.LockOSThread锁定了操作系统线程，并修改了任何可继承的操作系统级线程状态（例如，Linux或Plan 9名称空间），新进程将继承调用者的线程状态。

### func (*Cmd) Start

`func (c *Cmd) Start() error`

Start starts the specified command but does not wait for it to complete.

If Start returns successfully, the c.Process field will be set.

The Wait method will return the exit code and release associated resources once the command exits

Start启动指定的命令，但不等待它完成。

如果Start成功返回，c.Process字段将被设置。

一旦命令退出，Wait方法将返回退出代码并释放相关资源

### func (*Cmd) StderrPipe 

`func (c *Cmd) StderrPipe() (io.ReadCloser, error)`

StderrPipe returns a pipe that will be connected to the command's standard error when the command starts.

Wait will close the pipe after seeing the command exit, so most callers need not close the pipe themselves. It is thus incorrect to call Wait before all reads from the pipe have completed. For the same reason, it is incorrect to use Run when using StderrPipe. See the StdoutPipe example for idiomatic usage.

StderrPipe返回一个管道，当命令启动时，该管道将被连接到命令的标准错误。

Wait将在看到命令退出后关闭该管道，所以大多数调用者不需要自己关闭该管道。因此，在所有从管道中读取的数据完成之前调用Wait是不正确的。出于同样的原因，在使用StderrPipe时，使用Run是不正确的。请看StdoutPipe的例子以了解习惯性的用法。

### func (*Cmd) StdinPipe 

`func (c *Cmd) StdinPipe() (io.WriteCloser, error)`

StdinPipe returns a pipe that will be connected to the command's standard input when the command starts. The pipe will be closed automatically after Wait sees the command exit. A caller need only call Close to force the pipe to close sooner. For example, if the command being run will not exit until standard input is closed, the caller must close the pipe.

StdinPipe返回一个管道，当命令启动时，该管道将被连接到命令的标准输入。在Wait看到命令退出后，该管道将被自动关闭。调用者只需要调用Close来强制管道提前关闭。例如，如果正在运行的命令在标准输入关闭之前不会退出，调用者必须关闭该管道。

### func (*Cmd) StdoutPipe

`func (c *Cmd) StdoutPipe() (io.ReadCloser, error)`

StdoutPipe returns a pipe that will be connected to the command's standard output when the command starts.

Wait will close the pipe after seeing the command exit, so most callers need not close the pipe themselves. It is thus incorrect to call Wait before all reads from the pipe have completed. For the same reason, it is incorrect to call Run when using StdoutPipe. See the example for idiomatic usage.

StdoutPipe返回一个管道，当命令启动时，该管道将被连接到命令的标准输出。

Wait将在看到命令退出后关闭该管道，所以大多数调用者不需要自己关闭该管道。因此，在所有从管道中读取的数据完成之前调用Wait是不正确的。出于同样的原因，在使用StdoutPipe时调用Run是不正确的。请看示例中的习惯性用法。

### func (*Cmd) String

`func (c *Cmd) String() string`

String returns a human-readable description of c. It is intended only for debugging. In particular, it is not suitable for use as input to a shell. The output of String may vary across Go releases.

字符串返回对c的可读描述，它只用于调试。特别是，它不适合作为shell的输入。String的输出可能在不同的Go版本中有所不同。

### func (*Cmd) Wait

`func (c *Cmd) Wait() error`

Wait waits for the command to exit and waits for any copying to stdin or copying from stdout or stderr to complete.

The command must have been started by Start.

The returned error is nil if the command runs, has no problems copying stdin, stdout, and stderr, and exits with a zero exit status.

If the command fails to run or doesn't complete successfully, the error is of type *ExitError. Other error types may be returned for I/O problems.

If any of c.Stdin, c.Stdout or c.Stderr are not an *os.File, Wait also waits for the respective I/O loop copying to or from the process to complete.

Wait releases any resources associated with the Cmd.

Wait等待命令退出，并等待任何复制到stdin或从stdout或stderr复制的工作完成。

该命令必须是由Start启动的。

如果命令运行，复制stdin、stdout和stderr没有问题，并且以零退出状态退出，返回的错误是nil。

如果命令不能运行或没有成功完成，错误类型为*ExitError。对于I/O问题可能会返回其他错误类型。

如果c.Stdin、c.Stdout或c.Stderr中的任何一个不是*os.File，Wait也会等待各自的I/O循环复制到进程中或从进程中复制出来，以完成。

Wait释放任何与Cmd.Stdin相关的资源。

## type Error

```
type Error struct {
	// Name is the file name for which the error occurred.
	Name string
	// Err is the underlying error.
	Err error
}
```

Error is returned by LookPath when it fails to classify a file as an executable.

当LookPath未能将一个文件归类为可执行文件时，会返回错误。

### func (*Error) Error

`func (e *Error) Error() string`

### func (*Error) Unwrap

`func (e *Error) Unwrap() error`

## type ExitError

```
type ExitError struct {
	*os.ProcessState

	// Stderr holds a subset of the standard error output from the
	// Cmd.Output method if standard error was not otherwise being
	// collected.
	//
	// If the error output is long, Stderr may contain only a prefix
	// and suffix of the output, with the middle replaced with
	// text about the number of omitted bytes.
	//
	// Stderr is provided for debugging, for inclusion in error messages.
	// Users with other needs should redirect Cmd.Stderr as needed.
	Stderr []byte
}
```

An ExitError reports an unsuccessful exit by a command.

ExitError报告了一个命令的不成功退出。

### func (*ExitError) Error 

`func (e *ExitError) Error() string`