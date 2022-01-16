## Overview

Package signal implements access to incoming signals.

Signals are primarily used on Unix-like systems. For the use of this package on Windows and Plan 9, see below.

Types of signals ¶
The signals SIGKILL and SIGSTOP may not be caught by a program, and therefore cannot be affected by this package.

Synchronous signals are signals triggered by errors in program execution: SIGBUS, SIGFPE, and SIGSEGV. These are only considered synchronous when caused by program execution, not when sent using os.Process.Kill or the kill program or some similar mechanism. In general, except as discussed below, Go programs will convert a synchronous signal into a run-time panic.

The remaining signals are asynchronous signals. They are not triggered by program errors, but are instead sent from the kernel or from some other program.

Of the asynchronous signals, the SIGHUP signal is sent when a program loses its controlling terminal. The SIGINT signal is sent when the user at the controlling terminal presses the interrupt character, which by default is ^C (Control-C). The SIGQUIT signal is sent when the user at the controlling terminal presses the quit character, which by default is ^\ (Control-Backslash). In general you can cause a program to simply exit by pressing ^C, and you can cause it to exit with a stack dump by pressing ^\.

Default behavior of signals in Go programs ¶
By default, a synchronous signal is converted into a run-time panic. A SIGHUP, SIGINT, or SIGTERM signal causes the program to exit. A SIGQUIT, SIGILL, SIGTRAP, SIGABRT, SIGSTKFLT, SIGEMT, or SIGSYS signal causes the program to exit with a stack dump. A SIGTSTP, SIGTTIN, or SIGTTOU signal gets the system default behavior (these signals are used by the shell for job control). The SIGPROF signal is handled directly by the Go runtime to implement runtime.CPUProfile. Other signals will be caught but no action will be taken.

If the Go program is started with either SIGHUP or SIGINT ignored (signal handler set to SIG_IGN), they will remain ignored.

If the Go program is started with a non-empty signal mask, that will generally be honored. However, some signals are explicitly unblocked: the synchronous signals, SIGILL, SIGTRAP, SIGSTKFLT, SIGCHLD, SIGPROF, and, on Linux, signals 32 (SIGCANCEL) and 33 (SIGSETXID) (SIGCANCEL and SIGSETXID are used internally by glibc). Subprocesses started by os.Exec, or by the os/exec package, will inherit the modified signal mask.

Changing the behavior of signals in Go programs ¶
The functions in this package allow a program to change the way Go programs handle signals.

Notify disables the default behavior for a given set of asynchronous signals and instead delivers them over one or more registered channels. Specifically, it applies to the signals SIGHUP, SIGINT, SIGQUIT, SIGABRT, and SIGTERM. It also applies to the job control signals SIGTSTP, SIGTTIN, and SIGTTOU, in which case the system default behavior does not occur. It also applies to some signals that otherwise cause no action: SIGUSR1, SIGUSR2, SIGPIPE, SIGALRM, SIGCHLD, SIGCONT, SIGURG, SIGXCPU, SIGXFSZ, SIGVTALRM, SIGWINCH, SIGIO, SIGPWR, SIGSYS, SIGINFO, SIGTHR, SIGWAITING, SIGLWP, SIGFREEZE, SIGTHAW, SIGLOST, SIGXRES, SIGJVM1, SIGJVM2, and any real time signals used on the system. Note that not all of these signals are available on all systems.


If the program was started with SIGHUP or SIGINT ignored, and Notify is called for either signal, a signal handler will be installed for that signal and it will no longer be ignored. If, later, Reset or Ignore is called for that signal, or Stop is called on all channels passed to Notify for that signal, the signal will once again be ignored. Reset will restore the system default behavior for the signal, while Ignore will cause the system to ignore the signal entirely.

If the program is started with a non-empty signal mask, some signals will be explicitly unblocked as described above. If Notify is called for a blocked signal, it will be unblocked. If, later, Reset is called for that signal, or Stop is called on all channels passed to Notify for that signal, the signal will once again be blocked.

SIGPIPE ¶
When a Go program writes to a broken pipe, the kernel will raise a SIGPIPE signal.

If the program has not called Notify to receive SIGPIPE signals, then the behavior depends on the file descriptor number. A write to a broken pipe on file descriptors 1 or 2 (standard output or standard error) will cause the program to exit with a SIGPIPE signal. A write to a broken pipe on some other file descriptor will take no action on the SIGPIPE signal, and the write will fail with an EPIPE error.

If the program has called Notify to receive SIGPIPE signals, the file descriptor number does not matter. The SIGPIPE signal will be delivered to the Notify channel, and the write will fail with an EPIPE error.

This means that, by default, command line programs will behave like typical Unix command line programs, while other programs will not crash with SIGPIPE when writing to a closed network connection.

Go programs that use cgo or SWIG ¶
In a Go program that includes non-Go code, typically C/C++ code accessed using cgo or SWIG, Go's startup code normally runs first. It configures the signal handlers as expected by the Go runtime, before the non-Go startup code runs. If the non-Go startup code wishes to install its own signal handlers, it must take certain steps to keep Go working well. This section documents those steps and the overall effect changes to signal handler settings by the non-Go code can have on Go programs. In rare cases, the non-Go code may run before the Go code, in which case the next section also applies.

If the non-Go code called by the Go program does not change any signal handlers or masks, then the behavior is the same as for a pure Go program.

If the non-Go code installs any signal handlers, it must use the SA_ONSTACK flag with sigaction. Failing to do so is likely to cause the program to crash if the signal is received. Go programs routinely run with a limited stack, and therefore set up an alternate signal stack.

If the non-Go code installs a signal handler for any of the synchronous signals (SIGBUS, SIGFPE, SIGSEGV), then it should record the existing Go signal handler. If those signals occur while executing Go code, it should invoke the Go signal handler (whether the signal occurs while executing Go code can be determined by looking at the PC passed to the signal handler). Otherwise some Go run-time panics will not occur as expected.

If the non-Go code installs a signal handler for any of the asynchronous signals, it may invoke the Go signal handler or not as it chooses. Naturally, if it does not invoke the Go signal handler, the Go behavior described above will not occur. This can be an issue with the SIGPROF signal in particular.

The non-Go code should not change the signal mask on any threads created by the Go runtime. If the non-Go code starts new threads of its own, it may set the signal mask as it pleases.

If the non-Go code starts a new thread, changes the signal mask, and then invokes a Go function in that thread, the Go runtime will automatically unblock certain signals: the synchronous signals, SIGILL, SIGTRAP, SIGSTKFLT, SIGCHLD, SIGPROF, SIGCANCEL, and SIGSETXID. When the Go function returns, the non-Go signal mask will be restored.

If the Go signal handler is invoked on a non-Go thread not running Go code, the handler generally forwards the signal to the non-Go code, as follows. If the signal is SIGPROF, the Go handler does nothing. Otherwise, the Go handler removes itself, unblocks the signal, and raises it again, to invoke any non-Go handler or default system handler. If the program does not exit, the Go handler then reinstalls itself and continues execution of the program.

Non-Go programs that call Go code ¶
When Go code is built with options like -buildmode=c-shared, it will be run as part of an existing non-Go program. The non-Go code may have already installed signal handlers when the Go code starts (that may also happen in unusual cases when using cgo or SWIG; in that case, the discussion here applies). For -buildmode=c-archive the Go runtime will initialize signals at global constructor time. For -buildmode=c-shared the Go runtime will initialize signals when the shared library is loaded.

If the Go runtime sees an existing signal handler for the SIGCANCEL or SIGSETXID signals (which are used only on Linux), it will turn on the SA_ONSTACK flag and otherwise keep the signal handler.

For the synchronous signals and SIGPIPE, the Go runtime will install a signal handler. It will save any existing signal handler. If a synchronous signal arrives while executing non-Go code, the Go runtime will invoke the existing signal handler instead of the Go signal handler.

Go code built with -buildmode=c-archive or -buildmode=c-shared will not install any other signal handlers by default. If there is an existing signal handler, the Go runtime will turn on the SA_ONSTACK flag and otherwise keep the signal handler. If Notify is called for an asynchronous signal, a Go signal handler will be installed for that signal. If, later, Reset is called for that signal, the original handling for that signal will be reinstalled, restoring the non-Go signal handler if any.

Go code built without -buildmode=c-archive or -buildmode=c-shared will install a signal handler for the asynchronous signals listed above, and save any existing signal handler. If a signal is delivered to a non-Go thread, it will act as described above, except that if there is an existing non-Go signal handler, that handler will be installed before raising the signal.

Windows ¶
On Windows a ^C (Control-C) or ^BREAK (Control-Break) normally cause the program to exit. If Notify is called for os.Interrupt, ^C or ^BREAK will cause os.Interrupt to be sent on the channel, and the program will not exit. If Reset is called, or Stop is called on all channels passed to Notify, then the default behavior will be restored.

Additionally, if Notify is called, and Windows sends CTRL_CLOSE_EVENT, CTRL_LOGOFF_EVENT or CTRL_SHUTDOWN_EVENT to the process, Notify will return syscall.SIGTERM. Unlike Control-C and Control-Break, Notify does not change process behavior when either CTRL_CLOSE_EVENT, CTRL_LOGOFF_EVENT or CTRL_SHUTDOWN_EVENT is received - the process will still get terminated unless it exits. But receiving syscall.SIGTERM will give the process an opportunity to clean up before termination.

Plan 9 ¶
On Plan 9, signals have type syscall.Note, which is a string. Calling Notify with a syscall.Note will cause that value to be sent on the channel when that string is posted as a note.

包signal实现了对传入信号的访问。

信号主要用在类Unix系统上。关于这个包在Windows和Plan 9上的使用，见下文。

信号的类型
信号SIGKILL和SIGSTOP可能不会被程序捕获，因此不能被这个包影响。

同步信号是由程序执行中的错误触发的信号。SIGBUS, SIGFPE, 和SIGSEGV。这些信号只有在由程序执行引起时才被认为是同步的，而不是在使用os.Process.Kill或kill程序或一些类似机制发送时。一般来说，除了下面讨论的情况，Go程序会将同步信号转化为运行时恐慌。

其余的信号是异步信号。它们不是由程序错误触发的，而是由内核或其他一些程序发送的。

在异步信号中，SIGHUP信号是在程序失去其控制终端时发送的。SIGINT信号是在控制终端的用户按下中断字符时发出的，默认是^C（Control-C）。SIGQUIT信号是在控制终端的用户按下退出字符时发出的，默认情况下是^C（Control-Backslash）。一般来说，你可以通过按^C使程序简单地退出，也可以通过按^C使其退出并进行堆栈转储。

Go程序中信号的默认行为
默认情况下，同步信号会被转换为运行时恐慌。一个SIGHUP、SIGINT或SIGTERM信号会导致程序退出。SIGQUIT、SIGILL、SIGTRAP、SIGABRT、SIGSTKFLT、SIGEMT或SIGSYS信号会导致程序退出并进行堆栈转储。一个SIGTSTP、SIGTTIN或SIGTTOU信号会得到系统的默认行为（这些信号被shell用于工作控制）。SIGPROF信号由Go运行时直接处理，以实现runtime.CPUProfile。其他信号将被捕获，但不会采取任何行动。

如果Go程序在启动时忽略了SIGHUP或SIGINT（信号处理程序设置为SIG_IGN），它们将继续被忽略。

如果Go程序是以非空的信号掩码启动的，一般会被遵守。然而，有些信号是明确解除屏蔽的：同步信号、SIGILL、SIGTRAP、SIGSTKFLT、SIGCHLD、SIGPROF，以及在Linux上，信号32（SIGCANCEL）和33（SIGSETXID）（SIGCANCEL和SIGSETXID由glibc内部使用）。由os.Exec或os/exec包启动的子进程将继承修改后的信号屏蔽。

改变Go程序中的信号行为 ¶.
本软件包中的函数允许程序改变Go程序处理信号的方式。

Notify 禁用一组给定的异步信号的默认行为，而是通过一个或多个注册通道来传递这些信号。具体来说，它适用于信号SIGHUP、SIGINT、SIGQUIT、SIGABRT和SIGTERM。它也适用于工作控制信号SIGTSTP、SIGTTIN和SIGTTOU，在这种情况下，系统的默认行为不会发生。它也适用于一些信号，否则不会引起任何行动。SIGUSR1, SIGUSR2, SIGPIPE, SIGALRM, SIGCHLD, SIGCONT, SIGURG, SIGXCPU, SIGXFSZ, SIGVTALRM, SIGWINCH, SIGIO, SIGPWR, SIGSYS, SIGINFO, SIGTHR, SIGWAITING, SIGLWP, SIGFREEZE, SIGTHAW, SIGLOST, SIGXRES, SIGJVM1, SIGJVM2, 和任何系统上的实时信号。注意，并不是所有的系统都有这些信号。


如果程序在启动时忽略了SIGHUP或SIGINT，并且为这两个信号调用了Notify，那么将为该信号安装一个信号处理程序，它将不再被忽略。如果后来为该信号调用了Reset或Ignore，或者为该信号调用了所有传递给Notify的通道的Stop，该信号将再次被忽略。重置将恢复该信号的系统默认行为，而忽略将导致系统完全忽略该信号。

如果程序是以一个非空的信号掩码启动的，一些信号将被显式地解封，如上所述。如果对一个被封锁的信号调用Notify，它将被解除封锁。如果后来为该信号调用了Reset，或者为该信号调用了传递给Notify的所有通道的Stop，该信号将再次被封锁。

SIGPIPE ¶
当Go程序向一个断裂的管道写入时，内核会发出SIGPIPE信号。

如果程序没有调用Notify来接收SIGPIPE信号，那么行为就取决于文件描述符的编号。在文件描述符1或2（标准输出或标准错误）上向破损的管道写入文件将导致程序以SIGPIPE信号退出。写入其他文件描述符上的断管将不会对SIGPIPE信号采取任何行动，而写入将以EPIPE错误失败。

如果程序已经调用Notify来接收SIGPIPE信号，文件描述符的编号并不重要。SIGPIPE信号将被传递到Notify通道，而写操作将以EPIPE错误失败。

这意味着，在默认情况下，命令行程序的行为与典型的Unix命令行程序一样，而其他程序在向封闭的网络连接写入时不会因SIGPIPE而崩溃。

使用cgo或SWIG的Go程序¶。
在包含非Go代码的Go程序中，通常是使用cgo或SWIG访问的C/C++代码，Go的启动代码通常首先运行。在非围棋启动代码运行之前，它按照围棋运行时间的预期配置信号处理程序。如果非 Go 启动代码希望安装自己的信号处理程序，它必须采取某些步骤以保持 Go 的良好运行。本节记录了这些步骤以及非围棋代码对信号处理程序设置的改变对围棋程序的总体影响。在少数情况下，非围棋代码可能在围棋代码之前运行，在这种情况下，下一节也适用。

如果围棋程序调用的非围棋代码没有改变任何信号处理程序或掩码，那么其行为就与纯围棋程序相同。

如果非围棋代码安装了任何信号处理程序，它必须使用SA_ONSTACK标志和sigaction。如果不这样做，在收到信号时可能会导致程序崩溃。Go程序通常在有限的堆栈中运行，因此设置了一个备用的信号堆栈。

如果非围棋代码为任何同步信号（SIGBUS、SIGFPE、SIGSEGV）安装了信号处理程序，那么它应该记录现有的围棋信号处理程序。如果这些信号在执行Go代码时发生，它应该调用Go信号处理程序（信号是否在执行Go代码时发生，可以通过查看传递给信号处理程序的PC来确定）。否则一些Go运行时的慌乱将不会如期发生。

如果非围棋代码为任何一个异步信号安装了信号处理程序，它可以根据自己的选择调用围棋信号处理程序或不调用。当然，如果它不调用围棋信号处理程序，上述的围棋行为就不会发生。尤其是SIGPROF信号，这可能是一个问题。

非围棋代码不应改变围棋运行时创建的任何线程的信号屏蔽。如果非围棋代码启动了自己的新线程，它可以随意设置信号掩码。

如果非围棋代码启动了一个新的线程，改变了信号掩码，然后在该线程中调用了一个围棋函数，围棋运行时将自动解除对某些信号的封锁：同步信号、SIGILL、SIGTRAP、SIGSTKFLT、SIGCHLD、SIGPROF、SIGCANCEL 和 SIGSETXID。当Go函数返回时，非Go信号屏蔽将被恢复。

如果围棋信号处理程序在未运行围棋代码的非围棋线程上被调用，处理程序一般会将信号转发给非围棋代码，具体如下。如果该信号是SIGPROF，Go处理程序不做任何事情。否则，围棋处理程序会删除自己，解除对信号的封锁，并再次发出信号，以调用任何非围棋处理程序或默认系统处理程序。如果程序没有退出，Go处理程序就会重新安装自己，继续执行程序。

调用围棋代码的非围棋程序¶。
当使用-buildmode=c-shared等选项构建Go代码时，它将作为现有非Go程序的一部分运行。当围棋代码启动时，非围棋代码可能已经安装了信号处理程序（这也可能发生在使用cgo或SWIG的不寻常情况下；在这种情况下，这里的讨论也适用）。对于 -buildmode=c-archive，Go 运行时将在全局构造函数时初始化信号。对于 -buildmode=c-shared，Go 运行时将在加载共享库时初始化信号。

如果Go运行时看到SIGCANCEL或SIGSETXID信号（仅在Linux上使用）的现有信号处理程序，它将打开SA_ONSTACK标志，否则将保留该信号处理程序。

对于同步信号和SIGPIPE，Go运行时将安装一个信号处理器。它将保存任何现有的信号处理程序。如果在执行非 Go 代码时有同步信号出现，Go 运行时将调用现有的信号处理程序而不是 Go 信号处理程序。

使用 -buildmode=c-archive 或 -buildmode=c-shared 构建的 Go 代码默认不会安装任何其他信号处理程序。如果有一个现有的信号处理程序，Go运行时将打开SA_ONSTACK标志，否则将保留该信号处理程序。如果为一个异步信号调用Notify，将为该信号安装Go信号处理器。如果后来为该信号调用了Reset，该信号的原始处理将被重新安装，如果有的话，将恢复非Go信号处理。

没有使用 -buildmode=c-archive 或 -buildmode=c-shared 构建的 Go 代码将为上述的异步信号安装一个信号处理程序，并保存任何现有的信号处理程序。如果一个信号被传递给一个非围棋线程，它将如上所述行事，只是如果有一个现有的非围棋信号处理程序，该处理程序将在提升信号之前被安装。

Windows ¶
在Windows上，^C（Control-C）或^BREAK（Control-Break）通常导致程序退出。如果Notify被调用到os.Interrupt，^C或^BREAK将导致os.Interrupt在通道上被发送，而程序将不会退出。如果 Reset 被调用，或者在所有传递给 Notify 的通道上调用 Stop，那么将恢复默认行为。

此外，如果Notify被调用，并且Windows向进程发送了CTRL_CLOSE_EVENT、CTRL_LOGOFF_EVENT或CTRL_SHUTDOWN_EVENT，Notify将返回syscall.SIGTERM。与Control-C和Control-Break不同，Notify在收到CTRL_CLOSE_EVENT、CTRL_LOGOFF_EVENT或CTRL_SHUTDOWN_EVENT时不会改变进程的行为--进程仍然会被终止，除非它退出。但是收到syscall.SIGTERM将给进程一个机会在终止前进行清理。

计划9 ¶
在Plan 9中，信号的类型是syscall.Note，它是一个字符串。用syscall.Note调用Notify将导致该值在通道上被发送，当该字符串被发布为一个注释时。


## Examples

[Examples-url](https://pkg.go.dev/os/signal@go1.17.6#pkg-examples)

## Functions

### func Ignore

`func Ignore(sig ...os.Signal)`

Ignore causes the provided signals to be ignored. If they are received by the program, nothing will happen. Ignore undoes the effect of any prior calls to Notify for the provided signals. If no signals are provided, all incoming signals will be ignored.

忽略导致所提供的信号被忽略。如果程序收到了这些信号，什么也不会发生。忽略会撤销之前对所提供信号的任何调用Notify的效果。如果没有提供信号，所有传入的信号将被忽略。

### func Ignored

`func Ignored(sig os.Signal) bool`

Ignored reports whether sig is currently ignored.

Ignored报告sig当前是否被忽略。

### func Notify

`func Notify(c chan<- os.Signal, sig ...os.Signal)`

Notify causes package signal to relay incoming signals to c. If no signals are provided, all incoming signals will be relayed to c. Otherwise, just the provided signals will.

Package signal will not block sending to c: the caller must ensure that c has sufficient buffer space to keep up with the expected signal rate. For a channel used for notification of just one signal value, a buffer of size 1 is sufficient.

It is allowed to call Notify multiple times with the same channel: each call expands the set of signals sent to that channel. The only way to remove signals from the set is to call Stop.

It is allowed to call Notify multiple times with different channels and the same signals: each channel receives copies of incoming signals independently.

如果没有提供信号，所有传入的信号都会被转发到c，否则，只有提供的信号会被转发。

包信号不会阻止向c发送信号：调用者必须确保c有足够的缓冲空间来跟上预期的信号速率。对于一个只用于通知一个信号值的通道，一个大小为1的缓冲区就足够了。

允许用同一个通道多次调用Notify：每次调用都会扩大发送到该通道的信号集。从该集合中删除信号的唯一方法是调用Stop。

允许在不同的通道和相同的信号下多次调用Notify：每个通道独立地接收传入信号的副本。

### func NotifyContext 

`func NotifyContext(parent context.Context, signals ...os.Signal) (ctx context.Context, stop context.CancelFunc)`

NotifyContext returns a copy of the parent context that is marked done (its Done channel is closed) when one of the listed signals arrives, when the returned stop function is called, or when the parent context's Done channel is closed, whichever happens first.

The stop function unregisters the signal behavior, which, like signal.Reset, may restore the default behavior for a given signal. For example, the default behavior of a Go program receiving os.Interrupt is to exit. Calling NotifyContext(parent, os.Interrupt) will change the behavior to cancel the returned context. Future interrupts received will not trigger the default (exit) behavior until the returned stop function is called.

The stop function releases resources associated with it, so code should call stop as soon as the operations running in this Context complete and signals no longer need to be diverted to the context.

NotifyContext返回一个父级上下文的副本，当其中一个列出的信号到达时，当返回的stop函数被调用时，或者当父级上下文的Done通道被关闭时（以先发生者为准），该副本被标记为已完成（其Done通道被关闭）。

停止函数取消了对信号行为的注册，与signal.Reset一样，它可以恢复一个给定信号的默认行为。例如，Go程序收到os.Interrupt时的默认行为是退出。调用 NotifyContext(parent, os.Interrupt) 将改变行为，取消返回的上下文。未来收到的中断将不会触发默认（退出）行为，直到调用返回的停止函数。

stop函数释放了与之相关的资源，所以一旦在此上下文中运行的操作完成，信号不再需要被转移到上下文，代码就应该调用stop。

### func Reset 

`func Reset(sig ...os.Signal)`

Reset undoes the effect of any prior calls to Notify for the provided signals. If no signals are provided, all signal handlers will be reset.

重置会撤销之前对所提供信号的Notify的任何调用的效果。如果没有提供信号，所有的信号处理程序将被重置。

### func Stop

`func Stop(c chan<- os.Signal)`

Stop causes package signal to stop relaying incoming signals to c. It undoes the effect of all prior calls to Notify using c. When Stop returns, it is guaranteed that c will receive no more signals.

Stop导致包信号停止向c转发传入的信号。它撤销了之前所有使用c调用Notify的效果。