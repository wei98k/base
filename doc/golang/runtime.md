https://pkg.go.dev/runtime@go1.17.5

- update 2021年12月30日07:57:47
- update 2021年12月29日18:51:55

## Overview

## Constants

## Variables

## Functions

### func BlockProfile

`func BlockProfile(p []BlockProfileRecord) (n int, ok bool)`

BlockProfile返回n，即当前阻断配置文件中的记录数。如果len(p)>= n，BlockProfile将配置文件复制到p中，并返回n，true。如果len(p)< n，BlockProfile不改变p，并返回n, false。

大多数客户端应该使用runtime/pprof包或测试包的-test.blockprofile标志，而不是直接调用BlockProfile。

BlockProfile returns n, the number of records in the current blocking profile. If len(p) >= n, BlockProfile copies the profile into p and returns n, true. If len(p) < n, BlockProfile does not change p and returns n, false.

Most clients should use the runtime/pprof package or the testing package's -test.blockprofile flag instead of calling BlockProfile directly.

### func Breakpoint

`func Breakpoint()`

断点执行一个断点陷阱。

Breakpoint executes a breakpoint trap.

### func Caller

`func Caller(skip int) (pc uintptr, file string, line int, ok bool)`

Caller报告调用goroutine的堆栈上的函数调用的文件和行号信息。参数skip是要上升的堆栈帧数，其中0表示Caller的调用者。(由于历史原因，skip的含义在Caller和Callers之间有所不同）。返回值报告程序计数器、文件名和相应调用文件中的行号。如果不可能恢复信息，则布尔值ok为false。

Caller reports file and line number information about function invocations on the calling goroutine's stack. The argument skip is the number of stack frames to ascend, with 0 identifying the caller of Caller. (For historical reasons the meaning of skip differs between Caller and Callers.) The return values report the program counter, file name, and line number within the file of the corresponding call. The boolean ok is false if it was not possible to recover the information.

### func Callers

`func Callers(skip int, pc []uintptr) int`


Callers用调用goroutine的堆栈上的函数调用的返回程序计数器来填充片断pc。参数skip是在pc中记录之前要跳过的堆栈帧数，其中0标识Callers本身的帧，1标识Callers的调用者。它返回写入pc中的条目数。

要把这些PC翻译成符号信息，如函数名和行号，请使用CallersFrames。CallersFrames考虑到了内联函数，并将返回程序计数器调整为调用程序计数器。不鼓励直接在返回的PC片上进行迭代，也不鼓励在任何返回的PC上使用FuncForPC，因为这些不能考虑内联或返回程序计数器的调整。

Callers fills the slice pc with the return program counters of function invocations on the calling goroutine's stack. The argument skip is the number of stack frames to skip before recording in pc, with 0 identifying the frame for Callers itself and 1 identifying the caller of Callers. It returns the number of entries written to pc.

To translate these PCs into symbolic information such as function names and line numbers, use CallersFrames. CallersFrames accounts for inlined functions and adjusts the return program counters into call program counters. Iterating over the returned slice of PCs directly is discouraged, as is using FuncForPC on any of the returned PCs, since these cannot account for inlining or return program counter adjustment.

### func GC

`func GC()`

GC运行一个垃圾回收，并阻塞调用者，直到垃圾回收完成。它也可能阻塞整个程序。

GC runs a garbage collection and blocks the caller until the garbage collection is complete. It may also block the entire program.

### func GOMAXPROCS

`func GOMAXPROCS(n int) int`

GOMAXPROCS 设置可同时使用执行的最大CPU数，并返回先前的设置。 若 n < 1，它就不会更改当前设置。本地机器的逻辑CPU数可通过 NumCPU 查询。 当调度器改进后，此调用将会消失。

```
func main() {
	n1 := runtime.GOMAXPROCS(1)
	fmt.Println(n1) // 这个时候返回的是默认设置参数也就是CPU核实
	n2 := runtime.GOMAXPROCS(2)
	fmt.Println(n2) // 这个时候返回的是上次设置的CPU核数1
}
```

### func GOROOT

`func GOROOT() string`

GOROOT 返回 Go 树的根。它使用 GOROOT 环境变量（如果在进程开始时设置），否则就使用 Go 构建时使用的根。

GOROOT returns the root of the Go tree. It uses the GOROOT environment variable, if set at process start, or else the root used during the Go build.


### func Goexit

`func Goexit()`

Goexit会终止调用它的goroutine。其他goroutine不受影响。Goexit在终止goroutine之前运行所有延迟调用。因为Goexit不是一个恐慌，在这些延迟函数中的任何恢复调用将返回nil。

从主goroutine中调用Goexit会终止该goroutine，而func main不会返回。由于func main没有返回，程序继续执行其他goroutine。如果所有其他的goroutine都退出，程序就会崩溃。

Goexit terminates the goroutine that calls it. No other goroutine is affected. Goexit runs all deferred calls before terminating the goroutine. Because Goexit is not a panic, any recover calls in those deferred functions will return nil.

Calling Goexit from the main goroutine terminates that goroutine without func main returning. Since func main has not returned, the program continues execution of other goroutines. If all other goroutines exit, the program crashes.

### func GoroutineProfile

`func GoroutineProfile(p []StackRecord) (n int, ok bool)`

GoroutineProfile返回n，即活动的goroutine堆栈配置文件中的记录数量。如果len(p)>= n，GoroutineProfile将配置文件复制到p，并返回n，true。如果len(p)< n，GoroutineProfile不改变p并返回n，false。

大多数客户端应该使用runtime/pprof包而不是直接调用GoroutineProfile。

GoroutineProfile returns n, the number of records in the active goroutine stack profile. If len(p) >= n, GoroutineProfile copies the profile into p and returns n, true. If len(p) < n, GoroutineProfile does not change p and returns n, false.

Most clients should use the runtime/pprof package instead of calling GoroutineProfile directly.

### func Gosched

`func Gosched()`

Gosched让出处理器，允许其他goroutine运行。它不会暂停当前的goroutine，所以执行会自动恢复。

Gosched yields the processor, allowing other goroutines to run. It does not suspend the current goroutine, so execution resumes automatically.

### func KeepAlive

`func KeepAlive(x interface{})`

KeepAlive将其参数标记为当前可达。这确保了在程序中调用KeepAlive之前，该对象不会被释放，其终结者也不会被运行。

KeepAlive marks its argument as currently reachable. This ensures that the object is not freed, and its finalizer is not run, before the point in the program where KeepAlive is called.

### func LockOSThread

`func LockOSThread()`

LockOSThread将调用的goroutine连接到其当前的操作系统线程。调用的goroutine将一直在该线程中执行，没有其他goroutine会在该线程中执行，直到调用的goroutine对UnlockOSThread的调用次数与对LockOSThread的调用次数相同。如果调用的goroutine在没有解锁线程的情况下退出，该线程将被终止。

所有的初始函数都在启动线程上运行。从一个init函数中调用LockOSThread将导致主函数在该线程上被调用。

在调用操作系统服务或依赖每线程状态的非围棋库函数之前，一个goroutine应该调用LockOSThread。

LockOSThread wires the calling goroutine to its current operating system thread. The calling goroutine will always execute in that thread, and no other goroutine will execute in it, until the calling goroutine has made as many calls to UnlockOSThread as to LockOSThread. If the calling goroutine exits without unlocking the thread, the thread will be terminated.

All init functions are run on the startup thread. Calling LockOSThread from an init function will cause the main function to be invoked on that thread.

A goroutine should call LockOSThread before calling OS services or non-Go library functions that depend on per-thread state.

### func MemProfile

`func MemProfile(p []MemProfileRecord, inuseZero bool) (n int, ok bool)`

MemProfile返回每个分配点的内存分配和释放的概况。

MemProfile返回n，即当前内存配置文件中的记录数。如果len(p)>= n，MemProfile将配置文件复制到p中，并返回n，true。如果len(p) < n，MemProfile不改变p，并返回n, false。

如果inuseZero为真，配置文件包括r.AllocBytes>0但r.AllocBytes==r.FreeBytes的分配记录。这些地方的内存被分配了，但是已经全部释放给了运行时。

返回的配置文件最多可以是两个垃圾收集周期的时间。这是为了避免配置文件偏向分配；因为分配是实时发生的，而释放是延迟的，直到垃圾收集器执行清扫，配置文件只考虑有机会被垃圾收集器释放的分配。

大多数客户应该使用runtime/pprof包或测试包的-test.memprofile标志，而不是直接调用MemProfile。

MemProfile returns a profile of memory allocated and freed per allocation site.

MemProfile returns n, the number of records in the current memory profile. If len(p) >= n, MemProfile copies the profile into p and returns n, true. If len(p) < n, MemProfile does not change p and returns n, false.

If inuseZero is true, the profile includes allocation records where r.AllocBytes > 0 but r.AllocBytes == r.FreeBytes. These are sites where memory was allocated, but it has all been released back to the runtime.

The returned profile may be up to two garbage collection cycles old. This is to avoid skewing the profile toward allocations; because allocations happen in real time but frees are delayed until the garbage collector performs sweeping, the profile only accounts for allocations that have had a chance to be freed by the garbage collector.

Most clients should use the runtime/pprof package or the testing package's -test.memprofile flag instead of calling MemProfile directly.

### func MutexProfile

`func MutexProfile(p []BlockProfileRecord) (n int, ok bool)`

MutexProfile返回n，即当前mutex profile中的记录数。如果len(p)>= n，MutexProfile将配置文件复制到p中，并返回n，true。否则，MutexProfile不改变p，并返回n，false。

大多数客户应该使用runtime/pprof包而不是直接调用MutexProfile

MutexProfile returns n, the number of records in the current mutex profile. If len(p) >= n, MutexProfile copies the profile into p and returns n, true. Otherwise, MutexProfile does not change p, and returns n, false.

Most clients should use the runtime/pprof package instead of calling MutexProfile directly

### func NumCPU

`func NumCPU() int`

*Tip: 在mac系统中返回的是`sysctl -n machdep.cpu.thread_count`线程数(逻辑CPU)*

NumCPU返回当前进程可用的逻辑CPU的数量。

可用CPU的集合是在进程启动时通过查询操作系统来检查的。进程启动后对操作系统CPU分配的改变不会被反映。

NumCPU returns the number of logical CPUs usable by the current process.

The set of available CPUs is checked by querying the operating system at process startup. Changes to operating system CPU allocation after process startup are not reflected.

### func NumCgoCall

`func NumCgoCall() int64`

NumCgoCall返回当前进程进行的cgo调用的数量。

NumCgoCall returns the number of cgo calls made by the current process.

### func NumGoroutine

`func NumGoroutine() int`

NumGoroutine返回当前存在的goroutine的数量。

NumGoroutine returns the number of goroutines that currently exist.

### func ReadMemStats

`func ReadMemStats(m *MemStats)`

ReadMemStats用内存分配器的统计数据来填充m。

返回的内存分配器统计数据是调用ReadMemStats时的最新情况。这与堆配置文件不同，后者是最近完成的垃圾收集周期的快照。

ReadMemStats populates m with memory allocator statistics.

The returned memory allocator statistics are up to date as of the call to ReadMemStats. This is in contrast with a heap profile, which is a snapshot as of the most recently completed garbage collection cycle.

### func ReadTrace

`func ReadTrace() []byte`

ReadTrace 返回下一个二进制跟踪数据块，阻塞直到数据可用。如果追踪功能被关闭，并且在追踪功能开启时积累的所有数据都已返回，ReadTrace 将返回 nil。调用者必须在再次调用 ReadTrace 之前复制返回的数据。ReadTrace必须一次从一个goroutine中调用。

ReadTrace returns the next chunk of binary tracing data, blocking until data is available. If tracing is turned off and all the data accumulated while it was on has been returned, ReadTrace returns nil. The caller must copy the returned data before calling ReadTrace again. ReadTrace must be called from one goroutine at a time.

### func SetBlockProfileRate

`func SetBlockProfileRate(rate int)`

SetBlockProfileRate控制在阻塞配置文件中报告的goroutine阻塞事件的比例。剖析器的目标是对每一个阻塞率纳秒的阻塞事件进行平均采样。

要在剖析中包括每个阻塞事件，通过率=1。要完全关闭剖析，通过率<=0。

SetBlockProfileRate controls the fraction of goroutine blocking events that are reported in the blocking profile. The profiler aims to sample an average of one blocking event per rate nanoseconds spent blocked.

To include every blocking event in the profile, pass rate = 1. To turn off profiling entirely, pass rate <= 0.

### func SetCPUProfileRate

`func SetCPUProfileRate(hz int)`

SetCPUProfileRate设置CPU剖析率为每秒hz个样本。如果 hz <= 0，SetCPUProfileRate 会关闭剖析功能。如果剖析器处于开启状态，不先关闭它就不能改变速率。

大多数客户端应该使用 runtime/pprof 包或测试包的 -test.cpuprofile 标志，而不是直接调用 SetCPUProfileRate。

SetCPUProfileRate sets the CPU profiling rate to hz samples per second. If hz <= 0, SetCPUProfileRate turns off profiling. If the profiler is on, the rate cannot be changed without first turning it off.

Most clients should use the runtime/pprof package or the testing package's -test.cpuprofile flag instead of calling SetCPUProfileRate directly.

### func SetCgoTraceback

`func SetCgoTraceback(version int, traceback, context, symbolizer unsafe.Pointer)`


SetCgoTraceback记录了三个C函数，用来从C代码中收集回溯信息，并将回溯信息转换成符号信息。这些函数在打印使用cgo的程序的堆栈跟踪时使用。

回溯和上下文函数可以从信号处理程序中调用，因此必须只使用异步信号安全函数。符号化函数可能会在程序崩溃时被调用，因此必须谨慎地使用内存。所有的函数都不能回调到Go中。

SetCgoTraceback records three C functions to use to gather traceback information from C code and to convert that traceback information into symbolic information. These are used when printing stack traces for a program that uses cgo.

The traceback and context functions may be called from a signal handler, and must therefore use only async-signal safe functions. The symbolizer function may be called while the program is crashing, and so must be cautious about using memory. None of the functions may call back into Go.

### func SetFinalizer

`func SetFinalizer(obj interface{}, finalizer interface{})`

SetFinalizer将与obj相关的finalizer设置为提供的finalizer函数。当垃圾收集器发现一个有相关终结者的不可达区块时，它会清除这个关联并在一个单独的goroutine中运行finalizer(obj)。这使得obj再次可达，但现在没有相关的终结者。假设SetFinalizer没有被再次调用，那么下次垃圾收集器看到obj是不可到达的，它将释放obj。 官方文档还有好多内容描述。。。详见官网

SetFinalizer sets the finalizer associated with obj to the provided finalizer function. When the garbage collector finds an unreachable block with an associated finalizer, it clears the association and runs finalizer(obj) in a separate goroutine. This makes obj reachable again, but now without an associated finalizer. Assuming that SetFinalizer is not called again, the next time the garbage collector sees that obj is unreachable, it will free obj.

### func SetMutexProfileFraction

`func SetMutexProfileFraction(rate int) int`

SetMutexProfileFraction控制在mutex profile中报告的mutex争夺事件的比例。平均来说，报告的事件是1/率。之前的比率会被返回。

要完全关闭剖析，请传递速率0。要想只读取当前的速率，请通过速率<0。（对于n>1，采样的细节可能会改变。）

SetMutexProfileFraction controls the fraction of mutex contention events that are reported in the mutex profile. On average 1/rate events are reported. The previous rate is returned.

To turn off profiling entirely, pass rate 0. To just read the current rate, pass rate < 0. (For n>1 the details of sampling may change.)

### func Stack

`func Stack(buf []byte, all bool) int`

Stack将调用的goroutine的堆栈记录格式化为buf，并返回写入buf的字节数。如果全部为真，Stack将所有其他goroutine的堆栈记录格式化到当前goroutine的记录之后的buf中。

Stack formats a stack trace of the calling goroutine into buf and returns the number of bytes written to buf. If all is true, Stack formats stack traces of all other goroutines into buf after the trace for the current goroutine.

### func StartTrace

`func StartTrace() error`

StartTrace 启用对当前进程的跟踪。在跟踪过程中，数据将被缓冲并通过ReadTrace获得。如果追踪功能已经启用，StartTrace会返回一个错误。大多数客户应该使用 runtime/trace 包或测试包的 -test.trace 标志，而不是直接调用 StartTrace。

StartTrace enables tracing for the current process. While tracing, the data will be buffered and available via ReadTrace. StartTrace returns an error if tracing is already enabled. Most clients should use the runtime/trace package or the testing package's -test.trace flag instead of calling StartTrace directly.

### func StartTrace

`func StartTrace() error`

StartTrace 启用对当前进程的跟踪。在跟踪过程中，数据将被缓冲并通过ReadTrace获得。如果追踪功能已经启用，StartTrace会返回一个错误。大多数客户应该使用 runtime/trace 包或测试包的 -test.trace 标志，而不是直接调用 StartTrace。

StartTrace enables tracing for the current process. While tracing, the data will be buffered and available via ReadTrace. StartTrace returns an error if tracing is already enabled. Most clients should use the runtime/trace package or the testing package's -test.trace flag instead of calling StartTrace directly.

### func StopTrace

`func StopTrace()`

StopTrace停止追踪，如果它之前被启用。StopTrace只有在追踪的所有读数完成后才返回。

StopTrace stops tracing, if it was previously enabled. StopTrace only returns after all the reads for the trace have completed.

### func ThreadCreateProfile

`func ThreadCreateProfile(p []StackRecord) (n int, ok bool)`

ThreadCreateProfile返回n，即线程创建配置文件中的记录数量。如果len(p)>= n，ThreadCreateProfile将配置文件复制到p，并返回n，true。如果len(p) < n，ThreadCreateProfile不会改变p，并返回n，false。

大多数客户应该使用runtime/pprof包，而不是直接调用ThreadCreateProfile

ThreadCreateProfile returns n, the number of records in the thread creation profile. If len(p) >= n, ThreadCreateProfile copies the profile into p and returns n, true. If len(p) < n, ThreadCreateProfile does not change p and returns n, false.

Most clients should use the runtime/pprof package instead of calling ThreadCreateProfile directly.

### func UnlockOSThread

`func UnlockOSThread()`

UnlockOSThread会撤销先前对LockOSThread的调用。如果这使得调用的goroutine上的LockOSThread调用数量下降到0，那么它就会将调用的goroutine从其固定的操作系统线程中解开。如果没有活动的LockOSThread调用，这就是一个无用功。

在调用UnlockOSThread之前，调用者必须确保操作系统线程适合运行其他goroutine。如果调用者对线程的状态做了任何永久性的改变，会影响到其他的goroutine，就不应该调用这个函数，从而让goroutine锁定在OS线程上，直到goroutine（也就是线程）退出。

UnlockOSThread undoes an earlier call to LockOSThread. If this drops the number of active LockOSThread calls on the calling goroutine to zero, it unwires the calling goroutine from its fixed operating system thread. If there are no active LockOSThread calls, this is a no-op.

Before calling UnlockOSThread, the caller must ensure that the OS thread is suitable for running other goroutines. If the caller made any permanent changes to the state of the thread that would affect other goroutines, it should not call this function and thus leave the goroutine locked to the OS thread until the goroutine (and hence the thread) exits.

### func Version

`func Version() string`

Version returns the Go tree's version string. It is either the commit hash and date at the time of the build or, when possible, a release tag like "go1.3".

## Type


### type BlockProfileRecord
```
type BlockProfileRecord struct {
	Count  int64
	Cycles int64
	StackRecord
}
```
BlockProfileRecord描述了起源于特定调用序列的阻塞事件（堆栈跟踪）。

BlockProfileRecord describes blocking events originated at a particular call sequence (stack trace).

### type Error

```
type Error interface {
	error

	// RuntimeError is a no-op function but
	// serves to distinguish types that are run time
	// errors from ordinary errors: a type is a
	// run time error if it has a RuntimeError method.
	RuntimeError()
}
```

The Error interface identifies a run time error.

### type Frame

```
type Frame struct {
	// PC is the program counter for the location in this frame.
	// For a frame that calls another frame, this will be the
	// program counter of a call instruction. Because of inlining,
	// multiple frames may have the same PC value, but different
	// symbolic information.
	PC uintptr

	// Func is the Func value of this call frame. This may be nil
	// for non-Go code or fully inlined functions.
	Func *Func

	// Function is the package path-qualified function name of
	// this call frame. If non-empty, this string uniquely
	// identifies a single function in the program.
	// This may be the empty string if not known.
	// If Func is not nil then Function == Func.Name().
	Function string

	// File and Line are the file name and line number of the
	// location in this frame. For non-leaf frames, this will be
	// the location of a call. These may be the empty string and
	// zero, respectively, if not known.
	File string
	Line int

	// Entry point program counter for the function; may be zero
	// if not known. If Func is not nil then Entry ==
	// Func.Entry().
	Entry uintptr
	// contains filtered or unexported fields
}
```

Frame是Frames为每个调用帧返回的信息。

Frame is the information returned by Frames for each call frame.


### type Frames

```
type Frames struct {
	// contains filtered or unexported fields
}
```

框架可用于获得由调用者返回的PC值的片断的功能/文件/行信息。

Frames may be used to get function/file/line information for a slice of PC values returned by Callers.

#### func CallersFrames

`func CallersFrames(callers []uintptr) *Frames`

CallersFrames从Callers返回的PC值中抽取一个片断，准备返回函数/文件/行信息。在你完成对Frames的处理之前，不要改变这个片断。

CallersFrames takes a slice of PC values returned by Callers and prepares to return function/file/line information. Do not change the slice until you are done with the Frames.

#### func (*Frames) Next

`func (ci *Frames) Next() (frame Frame, more bool)`

Next returns a Frame representing the next call frame in the slice of PC values. If it has already returned all call frames, Next returns a zero Frame.
The more result indicates whether the next call to Next will return a valid Frame. It does not necessarily indicate whether this call returned one.
See the Frames example for idiomatic usage.

Next返回一个Frame，代表PC值片断中的下一个调用帧。如果它已经返回了所有的调用帧，Next返回一个零的Frame。
更多的结果表明对Next的下一次调用是否会返回一个有效的Frame。它不一定表明这次调用是否返回一个。
习惯性的用法请参见Frames的例子。


### type Func

```
type Func struct {
	// contains filtered or unexported fields
}
```

A Func represents a Go function in the running binary.

一个Func代表运行中的二进制中的一个Go函数。

#### func FuncForPC

`func FuncForPC(pc uintptr) *Func`

FuncForPC returns a *Func describing the function that contains the given program counter address, or else nil.

If pc represents multiple functions because of inlining, it returns the *Func describing the innermost function, but with an entry of the outermost function.

FuncForPC返回一个描述包含给定程序计数器地址的函数的*Func，否则为零。
如果pc因为内联而代表了多个函数，它返回描述最内层函数的*Func，但有一个最外层函数的条目

#### func (*Func) Entry

`func (f *Func) Entry() uintptr`

Entry returns the entry address of the function.

Entry返回函数的入口地址。

#### func (*Func) FileLine

`func (f *Func) FileLine(pc uintptr) (file string, line int)`

FileLine returns the file name and line number of the source code corresponding to the program counter pc. The result will not be accurate if pc is not a program counter within f.

FileLine返回与程序计数器pc对应的源代码的文件名和行号。如果pc不是f中的程序计数器，结果将不准确。


#### func (*Func) Name

`func (f *Func) Name() string`

Name returns the name of the function.

Name返回函数的名称。


## type MemProfileRecord

```
type MemProfileRecord struct {
	AllocBytes, FreeBytes     int64       // number of bytes allocated, freed
	AllocObjects, FreeObjects int64       // number of objects allocated, freed
	Stack0                    [32]uintptr // stack trace for this record; ends at first 0 entry
}
```

A MemProfileRecord describes the live objects allocated by a particular call sequence (stack trace).

一个MemProfileRecord描述了由一个特定的调用序列（堆栈跟踪）分配的活对象。

### func (*MemProfileRecord) InUseBytes

`func (r *MemProfileRecord) InUseBytes() int64`

InUseBytes returns the number of bytes in use (AllocBytes - FreeBytes).

InUseBytes返回使用中的字节数（AllocBytes - FreeBytes）。

### func (*MemProfileRecord) InUseObjects

`func (r *MemProfileRecord) InUseObjects() int64`

InUseObjects returns the number of objects in use (AllocObjects - FreeObjects).

InUseObjects返回使用中的对象的数量（AllocObjects - FreeObjects）。

### func (*MemProfileRecord) Stack

`func (r *MemProfileRecord) Stack() []uintptr`

Stack returns the stack trace associated with the record, a prefix of r.Stack0.

Stack返回与该记录相关的堆栈跟踪，是r.Stack0的前缀。

## type MemStats

A MemStats records statistics about the memory allocator.

一个MemStats记录了关于内存分配器的统计数据。

## type StackRecord

```
type StackRecord struct {
	Stack0 [32]uintptr // stack trace for this record; ends at first 0 entry
}
```

A StackRecord describes a single execution stack.

一个StackRecord描述了一个单一的执行堆栈。

### func (*StackRecord) Stack

`func (r *StackRecord) Stack() []uintptr`

Stack returns the stack trace associated with the record, a prefix of r.Stack0.

Stack返回与该记录相关的堆栈跟踪，是r.Stack0的前缀。

## type TypeAssertionError

type TypeAssertionError struct {
	// contains filtered or unexported fields
}

A TypeAssertionError explains a failed type assertion.

一个TypeAssertionError解释了一个失败的类型断言。

### func (*TypeAssertionError) Error

`func (e *TypeAssertionError) Error() string`

### func (*TypeAssertionError) RuntimeError

`func (*TypeAssertionError) RuntimeError()`
