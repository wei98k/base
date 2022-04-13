## Overview

Package testing provides support for automated testing of Go packages. It is intended to be used in concert with the "go test" command, which automates execution of any function of the form

包测试提供了对Go包自动测试的支持。它旨在与 "go test "命令一起使用，后者可以自动执行任何形式的函数

```
func TestXxx(*testing.T)
```

where Xxx does not start with a lowercase letter. The function name serves to identify the test routine.

Within these functions, use the Error, Fail or related methods to signal failure.

To write a new test suite, create a file whose name ends _test.go that contains the TestXxx functions as described here. Put the file in the same package as the one being tested. The file will be excluded from regular package builds but will be included when the "go test" command is run. For more detail, run "go help test" and "go help testflag".

A simple test function looks like this:

其中Xxx不是以小写字母开头的。函数名称用于识别测试程序。

在这些函数中，使用Error、Fail或相关的方法来表示失败。

要编写一个新的测试套件，创建一个文件，文件名以_test.go结尾，其中包含这里描述的TestXxx函数。将该文件放在与被测试的包相同的包中。该文件将被排除在常规软件包的构建之外，但在运行 "go test "命令时将被包括在内。更多细节，请运行 "go help test "和 "go help testflag"。

一个简单的测试函数看起来像这样。

```
func TestAbs(t *testing.T) {
    got := Abs(-1)
    if got != 1 {
        t.Errorf("Abs(-1) = %d; want 1", got)
    }
}
```

### Benchmarks

Functions of the form

```
func BenchmarkXxx(*testing.B)
```

are considered benchmarks, and are executed by the "go test" command when its -bench flag is provided. Benchmarks are run sequentially.

For a description of the testing flags, see https://golang.org/cmd/go/#hdr-Testing_flags

A sample benchmark function looks like this:

被视为基准，在提供-bench标志的情况下，由 "go test "命令执行。基准是按顺序运行的。

关于测试标志的描述，见https://golang.org/cmd/go/#hdr-Testing_flags

一个基准函数的样本看起来像这样。

```
func BenchmarkRandInt(b *testing.B) {
    for i := 0; i < b.N; i++ {
        rand.Int()
    }
}
```

The benchmark function must run the target code b.N times. During benchmark execution, b.N is adjusted until the benchmark function lasts long enough to be timed reliably. The output

基准函数必须运行目标代码b.N次。在基准执行过程中，b.N被调整，直到基准函数持续足够长的时间，以便可靠地进行计时。输出

```
BenchmarkRandInt-8   	68453040	        17.8 ns/op
```

means that the loop ran 68453040 times at a speed of 17.8 ns per loop.

If a benchmark needs some expensive setup before running, the timer may be reset:


意味着该循环运行了68453040次，每循环速度为17.8ns。

如果一个基准在运行前需要一些昂贵的设置，则可以重置计时器。

```
func BenchmarkBigLen(b *testing.B) {
    big := NewBig()
    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        big.Len()
    }
}
```

If a benchmark needs to test performance in a parallel setting, it may use the RunParallel helper function; such benchmarks are intended to be used with the go test -cpu flag:

如果一个基准需要在并行环境下测试性能，它可以使用RunParallel辅助函数；这样的基准旨在与go test -cpu标志一起使用。

```
func BenchmarkTemplateParallel(b *testing.B) {
    templ := template.Must(template.New("test").Parse("Hello, {{.}}!"))
    b.RunParallel(func(pb *testing.PB) {
        var buf bytes.Buffer
        for pb.Next() {
            buf.Reset()
            templ.Execute(&buf, "World")
        }
    })
}
```

### Examples

The package also runs and verifies example code. Example functions may include a concluding line comment that begins with "Output:" and is compared with the standard output of the function when the tests are run. (The comparison ignores leading and trailing space.) These are examples of an example:

该软件包还可以运行和验证示例代码。示例函数可能包括一个以 "Output: "开头的结尾行注释，并在测试运行时与该函数的标准输出进行比较。(比较时忽略前导和尾部的空格。)这些是例子的例子。

```
func ExampleHello() {
    fmt.Println("hello")
    // Output: hello
}

func ExampleSalutations() {
    fmt.Println("hello, and")
    fmt.Println("goodbye")
    // Output:
    // hello, and
    // goodbye
}
```

The comment prefix "Unordered output:" is like "Output:", but matches any line order:

注释前缀 "无序输出："与 "输出："类似，但与任何行序相匹配。

```
func ExamplePerm() {
    for _, value := range Perm(5) {
        fmt.Println(value)
    }
    // Unordered output: 4
    // 2
    // 1
    // 3
    // 0
}
```

Example functions without output comments are compiled but not executed.

The naming convention to declare examples for the package, a function F, a type T and method M on type T are:

没有输出注释的示例函数被编译，但不执行。

声明包的例子，一个函数F，一个类型T和类型T上的方法M的命名规则是。

```
func Example() { ... }
func ExampleF() { ... }
func ExampleT() { ... }
func ExampleT_M() { ... }
```

Multiple example functions for a package/type/function/method may be provided by appending a distinct suffix to the name. The suffix must start with a lower-case letter.

一个包/类型/函数/方法的多个示例函数可以通过在名称后面添加一个不同的后缀来提供。后缀必须以小写字母开头。

```
func Example_suffix() { ... }
func ExampleF_suffix() { ... }
func ExampleT_suffix() { ... }
func ExampleT_M_suffix() { ... }
```

The entire test file is presented as the example when it contains a single example function, at least one other function, type, variable, or constant declaration, and no test or benchmark functions.

当整个测试文件包含一个示例函数，至少有一个其他函数、类型、变量或常量声明，并且没有测试或基准函数时，该文件将作为示例呈现。


### Skipping

Tests or benchmarks may be skipped at run time with a call to the Skip method of *T or *B:

在运行时，可以通过调用*T或*B的跳过方法跳过测试或基准。

```
func TestTimeConsuming(t *testing.T) {
    if testing.Short() {
        t.Skip("skipping test in short mode.")
    }
    ...
}
```

### Subtests and Sub-benchmarks (次级测试和次级基准)

The Run methods of T and B allow defining subtests and sub-benchmarks, without having to define separate functions for each. This enables uses like table-driven benchmarks and creating hierarchical tests. It also provides a way to share common setup and tear-down code:

T和B的运行方法允许定义子测试和子基准，而不需要为每一个定义单独的函数。这使得像表格驱动的基准和创建分层测试这样的用途成为可能。它还提供了一种分享共同的设置和拆除代码的方法。

```
func TestFoo(t *testing.T) {
    // <setup code>
    t.Run("A=1", func(t *testing.T) { ... })
    t.Run("A=2", func(t *testing.T) { ... })
    t.Run("B=1", func(t *testing.T) { ... })
    // <tear-down code>
}
```

Each subtest and sub-benchmark has a unique name: the combination of the name of the top-level test and the sequence of names passed to Run, separated by slashes, with an optional trailing sequence number for disambiguation.

The argument to the -run and -bench command-line flags is an unanchored regular expression that matches the test's name. For tests with multiple slash-separated elements, such as subtests, the argument is itself slash-separated, with expressions matching each name element in turn. Because it is unanchored, an empty expression matches any string. For example, using "matching" to mean "whose name contains":

每个子测试和子基准都有一个独特的名字：顶层测试的名字和传递给Run的名字序列的组合，用斜线隔开，后面的序列号是可选的，用于消除歧义。

运行 "和 "基准 "命令行标志的参数是一个非锚定的正则表达式，与测试的名称相匹配。对于有多个斜线分隔元素的测试，比如子测试，参数本身是斜线分隔的，表达式依次匹配每个名称元素。因为它是无锚的，所以空的表达式可以匹配任何字符串。例如，用 "匹配 "来表示 "其名称包含"。

```
go test -run ”      # Run all tests.
go test -run Foo     # Run top-level tests matching "Foo", such as "TestFooBar".
go test -run Foo/A=  # For top-level tests matching "Foo", run subtests matching "A=".
go test -run /A=1    # For all top-level tests, run subtests matching "A=1".
```

Subtests can also be used to control parallelism. A parent test will only complete once all of its subtests complete. In this example, all tests are run in parallel with each other, and only with each other, regardless of other top-level tests that may be defined:

子测试也可以用来控制并行性。一个父测试只有在其所有的子测试完成后才会完成。在这个例子中，所有的测试都是相互平行运行的，而且只相互平行运行，不考虑其他可能定义的顶层测试。

```
func TestGroupedParallel(t *testing.T) {
    for _, tc := range tests {
        tc := tc // capture range variable
        t.Run(tc.Name, func(t *testing.T) {
            t.Parallel()
            ...
        })
    }
}
```

The race detector kills the program if it exceeds 8128 concurrent goroutines, so use care when running parallel tests with the -race flag set.

Run does not return until parallel subtests have completed, providing a way to clean up after a group of parallel tests:

如果程序超过8128个并发的goroutine，竞赛检测器就会杀死它，所以在设置了-race标志的情况下运行并行测试时要小心。

在并行子测试完成之前，运行不会返回，提供了一种在一组并行测试之后进行清理的方法。

```
func TestTeardownParallel(t *testing.T) {
    // This Run will not return until the parallel tests finish.
    t.Run("group", func(t *testing.T) {
        t.Run("Test1", parallelTest1)
        t.Run("Test2", parallelTest2)
        t.Run("Test3", parallelTest3)
    })
    // <tear-down code>
}
```

### Main

It is sometimes necessary for a test or benchmark program to do extra setup or teardown before or after it executes. It is also sometimes necessary to control which code runs on the main thread. To support these and other cases, if a test file contains a function:

有时，测试或基准程序有必要在执行之前或之后进行额外的设置或拆分。有时也需要控制哪些代码在主线程上运行。为了支持这些和其他情况，如果一个测试文件包含一个函数。

```
func TestMain(m *testing.M)
```

then the generated test will call TestMain(m) instead of running the tests or benchmarks directly. TestMain runs in the main goroutine and can do whatever setup and teardown is necessary around a call to m.Run. m.Run will return an exit code that may be passed to os.Exit. If TestMain returns, the test wrapper will pass the result of m.Run to os.Exit itself.

When TestMain is called, flag.Parse has not been run. If TestMain depends on command-line flags, including those of the testing package, it should call flag.Parse explicitly. Command line flags are always parsed by the time test or benchmark functions run.

A simple implementation of TestMain is:


那么生成的测试将调用TestMain(m)而不是直接运行测试或基准。TestMain在主goroutine中运行，可以围绕调用m.Run做任何必要的设置和拆除。m.Run将返回一个退出代码，可以传递给os.Exit。如果TestMain返回，测试包装器将把m.Run的结果传递给os.Exit本身。

当TestMain被调用时，flag.Parse还没有被运行。如果TestMain依赖于命令行标志，包括测试包的标志，它应该明确地调用flag.Parse。命令行标志总是在测试或基准函数运行时被解析。

TestMain的一个简单实现是。

```
func TestMain(m *testing.M) {
	// call flag.Parse() here if TestMain uses flags
	os.Exit(m.Run())
}
```

TestMain is a low-level primitive and should not be necessary for casual testing needs, where ordinary test functions suffice.

TestMain是一个低级的基元，对于随意的测试需求不应该是必要的，普通的测试函数就足够了。

## Functions

### func AllocsPerRun

`func AllocsPerRun(runs int, f func()) (avg float64)`

AllocsPerRun returns the average number of allocations during calls to f. Although the return value has type float64, it will always be an integral value.

To compute the number of allocations, the function will first be run once as a warm-up. The average number of allocations over the specified number of runs will then be measured and returned.

AllocsPerRun sets GOMAXPROCS to 1 during its measurement and will restore it before returning.

AllocsPerRun返回调用f过程中的平均分配次数。虽然返回值的类型是float64，但它总是一个积分值。

为了计算分配次数，该函数将首先被运行一次作为热身。然后将测量并返回指定运行次数的平均分配数。

AllocsPerRun在测量期间将GOMAXPROCS设置为1，并在返回之前将其恢复。

### func CoverMode

`func CoverMode() string`

CoverMode reports what the test coverage mode is set to. The values are "set", "count", or "atomic". The return value will be empty if test coverage is not enabled.

CoverMode报告测试覆盖率模式的设置情况。其值是 "set"、"count "或 "atomic"。如果测试覆盖率没有被启用，返回值将为空。

### func Coverage

`func Coverage() float64`

Coverage reports the current code coverage as a fraction in the range [0, 1]. If coverage is not enabled, Coverage returns 0.

When running a large set of sequential test cases, checking Coverage after each one can be useful for identifying which test cases exercise new code paths. It is not a replacement for the reports generated by 'go test -cover' and 'go tool cover'.

覆盖率以[0, 1]范围内的分数报告当前的代码覆盖率。如果覆盖率没有被启用，Coverage返回0。

当运行一大组连续的测试用例时，在每一个测试用例之后检查Coverage对于识别哪些测试用例行使了新的代码路径非常有用。它不能替代 "go test -cover "和 "go tool cover "所产生的报告。

### func Init

`func Init()`

Init registers testing flags. These flags are automatically registered by the "go test" command before running test functions, so Init is only needed when calling functions such as Benchmark without using "go test".

Init has no effect if it was already called.

Init注册了测试标志。在运行测试函数之前，这些标志会被 "go test "命令自动注册，所以只有在调用Benchmark等函数而不使用 "go test "时才需要Init。

如果Init已经被调用，它就没有任何作用。

### func Main

`func Main(matchString func(pat, str string) (bool, error), tests []InternalTest, benchmarks []InternalBenchmark, examples []InternalExample)`

Main is an internal function, part of the implementation of the "go test" command. It was exported because it is cross-package and predates "internal" packages. It is no longer used by "go test" but preserved, as much as possible, for other systems that simulate "go test" using Main, but Main sometimes cannot be updated as new functionality is added to the testing package. Systems simulating "go test" should be updated to use MainStart.

Main是一个内部函数，是实现 "go test "命令的一部分。它被导出是因为它是跨包的，而且比 "内部 "包更早。它不再被 "go test "使用，但尽可能地为其他使用Main模拟 "go test "的系统所保留，但Main有时不能被更新，因为新的功能被添加到测试包。模拟 "go test "的系统应该被更新为使用MainStart。

### func RegisterCover

`func RegisterCover(c Cover)`

RegisterCover records the coverage data accumulators for the tests. NOTE: This function is internal to the testing infrastructure and may change. It is not covered (yet) by the Go 1 compatibility guidelines.

RegisterCover记录测试的覆盖数据累积器。注意：这个功能是测试基础设施的内部功能，可能会改变。它不包括在Go 1的兼容性指南中（还没有）。

### func RunBenchmarks

`func RunBenchmarks(matchString func(pat, str string) (bool, error), benchmarks []InternalBenchmark)`

RunBenchmarks is an internal function but exported because it is cross-package; it is part of the implementation of the "go test" command.

RunBenchmarks是一个内部函数，但被导出，因为它是跨包的；它是实现 "go test "命令的一部分。

### func RunExamples

`func RunExamples(matchString func(pat, str string) (bool, error), examples []InternalExample) (ok bool)`

RunExamples is an internal function but exported because it is cross-package; it is part of the implementation of the "go test" command.

RunExamples是一个内部函数，但被导出，因为它是跨包的；它是实现 "go test "命令的一部分。

### func RunTests

`func RunTests(matchString func(pat, str string) (bool, error), tests []InternalTest) (ok bool)`

RunTests is an internal function but exported because it is cross-package; it is part of the implementation of the "go test" command.

RunTests是一个内部函数，但被导出，因为它是跨包的；它是实现 "go test "命令的一部分。

### func Short

`func Short() bool`

Short reports whether the -test.short flag is set.

Short报告是否设置了-test.short标志。

### func Verbose

`func Verbose() bool`

Verbose reports whether the -test.v flag is set.

Verbose报告是否设置了-test.v标志。

## type B

```
type B struct {
	N int
	// contains filtered or unexported fields
}
```

B is a type passed to Benchmark functions to manage benchmark timing and to specify the number of iterations to run.

A benchmark ends when its Benchmark function returns or calls any of the methods FailNow, Fatal, Fatalf, SkipNow, Skip, or Skipf. Those methods must be called only from the goroutine running the Benchmark function. The other reporting methods, such as the variations of Log and Error, may be called simultaneously from multiple goroutines.

Like in tests, benchmark logs are accumulated during execution and dumped to standard output when done. Unlike in tests, benchmark logs are always printed, so as not to hide output whose existence may be affecting benchmark results.

B是传递给Benchmark函数的一个类型，用于管理基准的时间，并指定运行的迭代次数。

当Benchmark函数返回或调用任何方法FailNow、Fatal、Fatalf、SkipNow、Skip或Skipf时，基准就会结束。这些方法必须只从运行Benchmark函数的goroutine中调用。其他报告方法，例如Log和Error的变体，可以从多个goroutine同时调用。

和测试一样，基准日志在执行过程中积累，完成后转储到标准输出。与测试不同的是，基准日志总是被打印出来，这样就不会隐藏那些可能影响基准结果的输出。

### func (*B) Cleanup

`func (c *B) Cleanup(f func())`

Cleanup registers a function to be called when the test (or subtest) and all its subtests complete. Cleanup functions will be called in last added, first called order.

Cleanup注册了一个函数，当测试（或子测试）和它的所有子测试完成时，将被调用。清理函数将按照最后添加、首先调用的顺序调用。

### func (*B) Error

`func (c *B) Error(args ...interface{})`

Error is equivalent to Log followed by Fail.

Error等同于Log，后面是Fail。

### func (*B) Errorf

`func (c *B) Errorf(format string, args ...interface{})`

Errorf is equivalent to Logf followed by Fail.

Errorf等同于Logf后面的Fail。

### func (*B) Fail

`func (c *B) Fail()`

Fail marks the function as having failed but continues execution.

Fail标志着该函数已经失败，但继续执行。

### func (*B) FailNow

`func (c *B) FailNow()`

FailNow marks the function as having failed and stops its execution by calling runtime.Goexit (which then runs all deferred calls in the current goroutine). Execution will continue at the next test or benchmark. FailNow must be called from the goroutine running the test or benchmark function, not from other goroutines created during the test. Calling FailNow does not stop those other goroutines.

FailNow将函数标记为失败，并通过调用runtime.Goexit（然后运行当前goroutine中的所有延迟调用）停止其执行。执行将在下一次测试或基准测试时继续进行。FailNow必须从运行测试或基准函数的goroutine中调用，而不是从测试期间创建的其他goroutine中调用。调用FailNow不会停止那些其他的goroutine。

### func (*B) Failed

`func (c *B) Failed() bool`

Failed reports whether the function has failed.

Failed报告该函数是否失败。

### func (*B) Fatal

`func (c *B) Fatal(args ...interface{})`

Fatal is equivalent to Log followed by FailNow.

Fatal等同于Log，后面是FailNow。

### func (*B) Fatalf 

`func (c *B) Fatalf(format string, args ...interface{})`

Fatalf is equivalent to Logf followed by FailNow.

Fatalf相当于Logf后面的FailNow。

### func (*B) Helper 

`func (c *B) Helper()`

Helper marks the calling function as a test helper function. When printing file and line information, that function will be skipped. Helper may be called simultaneously from multiple goroutines.

Helper标志着调用的函数是一个测试辅助函数。当打印文件和行信息时，该函数将被跳过。Helper可以从多个goroutine同时调用。

### func (*B) Log

`func (c *B) Log(args ...interface{})`

Log formats its arguments using default formatting, analogous to Println, and records the text in the error log. For tests, the text will be printed only if the test fails or the -test.v flag is set. For benchmarks, the text is always printed to avoid having performance depend on the value of the -test.v flag.

Log使用默认格式化其参数，类似于Println，并在错误日志中记录文本。对于测试，只有在测试失败或设置了-test.v标志的情况下才会打印文本。对于基准测试，文本总是被打印出来，以避免性能依赖于-test.v标志的值。

### func (*B) Logf 

`func (c *B) Logf(format string, args ...interface{})`

Logf formats its arguments according to the format, analogous to Printf, and records the text in the error log. A final newline is added if not provided. For tests, the text will be printed only if the test fails or the -test.v flag is set. For benchmarks, the text is always printed to avoid having performance depend on the value of the -test.v flag.

Logf根据格式化其参数，类似于Printf，并在错误日志中记录文本。如果没有提供最后的换行符，则会添加一个换行符。对于测试，只有在测试失败或设置了-test.v标志的情况下才会打印文本。对于基准测试，文本总是被打印出来，以避免性能依赖于-test.v标志的值。

### func (*B) Name

`func (c *B) Name() string`

Name returns the name of the running (sub-) test or benchmark.

The name will include the name of the test along with the names of any nested sub-tests. If two sibling sub-tests have the same name, Name will append a suffix to guarantee the returned name is unique.

名称返回正在运行的（子）测试或基准的名称。

该名称将包括测试的名称以及任何嵌套的子测试的名称。如果两个同级别的子测试有相同的名称，Name将附加一个后缀，以保证返回的名称是唯一的。

### func (*B) ReportAllocs

`func (b *B) ReportAllocs()`

ReportAllocs enables malloc statistics for this benchmark. It is equivalent to setting -test.benchmem, but it only affects the benchmark function that calls ReportAllocs.

ReportAllocs可以为这个基准启用malloc统计。它等同于设置-test.benchmem，但它只影响调用ReportAllocs的基准函数。








TDO.............................





















