
## Overview

Package sort provides primitives for sorting slices and user-defined collections.

包sort提供了对切片和用户定义的集合进行排序的基元。


## Examples

[Examples-url](https://pkg.go.dev/sort@go1.17.6)

## Functions

### func Float64s

`func Float64s(x []float64)`

Float64s sorts a slice of float64s in increasing order. Not-a-number (NaN) values are ordered before other values.

Float64s以递增的顺序对浮点数的片断进行排序。非数字（NaN）值在其他值之前排序。

### func Float64sAreSorted

`func Float64sAreSorted(x []float64) bool`

Float64sAreSorted reports whether the slice x is sorted in increasing order, with not-a-number (NaN) values before any other values.

Float64sAreSorted报告分片x是否按递增顺序排序，非数字（NaN）值在任何其他值之前。

### func Ints

`func Ints(x []int)`

Ints sorts a slice of ints in increasing order.

Ints以递增的顺序对一个整数片进行排序。

### func IntsAreSorted 

`func IntsAreSorted(x []int) bool`

IntsAreSorted reports whether the slice x is sorted in increasing order.

IntsAreSorted报告片断x是否按递增顺序排序。


### func IsSorted

`func IsSorted(data Interface) bool`

IsSorted reports whether data is sorted.

IsSorted报告数据是否被排序。

### func Search

`func Search(n int, f func(int) bool) int`

Search uses binary search to find and return the smallest index i in [0, n) at which f(i) is true, assuming that on the range [0, n), f(i) == true implies f(i+1) == true. That is, Search requires that f is false for some (possibly empty) prefix of the input range [0, n) and then true for the (possibly empty) remainder; Search returns the first true index. If there is no such index, Search returns n. (Note that the "not found" return value is not -1 as in, for instance, strings.Index.) Search calls f(i) only for i in the range [0, n).

A common use of Search is to find the index i for a value x in a sorted, indexable data structure such as an array or slice. In this case, the argument f, typically a closure, captures the value to be searched for, and how the data structure is indexed and ordered.

For instance, given a slice data sorted in ascending order, the call Search(len(data), func(i int) bool { return data[i] >= 23 }) returns the smallest index i such that data[i] >= 23. If the caller wants to find whether 23 is in the slice, it must test data[i] == 23 separately.

Searching data sorted in descending order would use the <= operator instead of the >= operator.

To complete the example above, the following code tries to find the value x in an integer slice data sorted in ascending order:


Search使用二进制搜索来寻找并返回f(i)为真的最小索引i，假设在[0, n]范围内，f(i) == true意味着f(i+1) == true。也就是说，Search要求f在输入范围[0, n]的某个（可能是空的）前缀是假的，然后在（可能是空的）剩余部分是真的；Search返回第一个真索引。如果没有这样的索引，Search返回n。（注意，"没有找到 "的返回值不是像strings.Index那样的-1。） Search只对范围[0, n]内的i调用f(i)。

Search的一个常见用途是为一个排序的、可索引的数据结构（如数组或片断）中的一个值x寻找索引i。在这种情况下，参数f，通常是一个闭包，捕获要搜索的值，以及数据结构如何被索引和排序。

例如，给定一个按升序排序的片断数据，调用Search(len(data), func(i int) bool { return data[i] >= 23 })返回最小的索引i，使data[i] >= 23。如果调用者想知道23是否在片断中，它必须单独测试data[i] == 23。

搜索按降序排序的数据将使用<=操作符而不是>=操作符。

为了完成上面的例子，下面的代码试图在一个以升序排序的整数片数据中找到值x。


### func SearchFloat64s

`func SearchFloat64s(a []float64, x float64) int`

SearchFloat64s searches for x in a sorted slice of float64s and returns the index as specified by Search. The return value is the index to insert x if x is not present (it could be len(a)). The slice must be sorted in ascending order.

SearchFloat64s在一个排序的float64s片断中搜索x，并返回Search指定的索引。如果x不存在，返回值是插入x的索引（可以是len(a)）。片断必须以升序排序。

### func SearchInts

`func SearchInts(a []int, x int) int`

SearchInts searches for x in a sorted slice of ints and returns the index as specified by Search. The return value is the index to insert x if x is not present (it could be len(a)). The slice must be sorted in ascending order.

SearchInts在一个排序的ints片断中搜索x，并返回由Search指定的索引。如果x不存在，返回值是插入x的索引（可以是len(a)）。片断必须以升序排序。

### func SearchStrings

`func SearchStrings(a []string, x string) int`

SearchStrings searches for x in a sorted slice of strings and returns the index as specified by Search. The return value is the index to insert x if x is not present (it could be len(a)). The slice must be sorted in ascending order.

SearchStrings在一个排序的字符串片断中搜索x，并返回Search所指定的索引。如果x不存在，返回值是插入x的索引（可以是len(a)）。片断必须以升序排序。

### func Slice

`func Slice(x interface{}, less func(i, j int) bool)`

Slice sorts the slice x given the provided less function. It panics if x is not a slice.

The sort is not guaranteed to be stable: equal elements may be reversed from their original order. For a stable sort, use SliceStable.

The less function must satisfy the same requirements as the Interface type's Less method.

Slice对给定的less函数的片断x进行排序。如果x不是一个片断，它就会恐慌。

该排序不保证是稳定的：相等的元素可能会从它们的原始顺序中颠倒过来。对于一个稳定的排序，请使用SliceStable。

less函数必须满足与接口类型的Less方法相同的要求。


### func SliceIsSorted 

`func SliceIsSorted(x interface{}, less func(i, j int) bool) bool`

SliceIsSorted reports whether the slice x is sorted according to the provided less function. It panics if x is not a slice.

SliceIsSorted报告片断x是否按照提供的less函数进行排序。如果x不是一个片断，它就会恐慌。

### func SliceStable

`func SliceStable(x interface{}, less func(i, j int) bool)`

SliceStable sorts the slice x using the provided less function, keeping equal elements in their original order. It panics if x is not a slice.

The less function must satisfy the same requirements as the Interface type's Less method.

SliceStable使用提供的less函数对slice x进行排序，保持相等元素的原始顺序。如果x不是一个片断，它就会恐慌。

less函数必须满足与接口类型的Less方法相同的要求。

### func Sort

`func Sort(data Interface)`

Sort sorts data. It makes one call to data.Len to determine n and O(n*log(n)) calls to data.Less and data.Swap. The sort is not guaranteed to be stable.

排序对数据进行排序。它对data.Len进行一次调用以确定n，并对data.Less和data.Swap进行O(n*log(n))调用。该排序不保证是稳定的。

### func Stable

`func Stable(data Interface)`

Stable sorts data while keeping the original order of equal elements.

It makes one call to data.Len to determine n, O(n*log(n)) calls to data.Less and O(n*log(n)*log(n)) calls to data.Swap.

稳定地对数据进行排序，同时保持相等元素的原始顺序。

它对data.Len进行一次调用以确定n，对data.Less进行O(n*log(n))调用，对data.Swap进行O(n*log(n))调用。

### func Strings

`func Strings(x []string)`

Strings sorts a slice of strings in increasing order.

弦的排序是按照递增的顺序对弦的片断进行排序。

### func StringsAreSorted

`func StringsAreSorted(x []string) bool`

StringsAreSorted reports whether the slice x is sorted in increasing order.

StringsAreSorted报告片断x是否按递增顺序排序。


## type Float64Slice

`type Float64Slice []float64`

Float64Slice implements Interface for a []float64, sorting in increasing order, with not-a-number (NaN) values ordered before other values.

Float64Slice实现了[]float64的接口，以递增的顺序排序，非数字（NaN）值在其他值之前排序。

### func (Float64Slice) Len

`func (x Float64Slice) Len() int`


### func (Float64Slice) Less

`func (x Float64Slice) Less(i, j int) bool`

Less reports whether x[i] should be ordered before x[j], as required by the sort Interface. Note that floating-point comparison by itself is not a transitive relation: it does not report a consistent ordering for not-a-number (NaN) values. This implementation of Less places NaN values before any others, by using:

Less报告x[i]是否应该排序在x[j]之前，这是排序界面的要求。请注意，浮点比较本身并不是一种传递性关系：它并不报告非数字（NaN）值的一致排序。这个Less的实现通过使用以下方法将NaN值放在任何其他值之前。


### func (Float64Slice) Search 

`func (p Float64Slice) Search(x float64) int`

Search returns the result of applying SearchFloat64s to the receiver and x.

搜索返回对接收器和x应用SearchFloat64s的结果。

### func (Float64Slice) Sort

`func (x Float64Slice) Sort()`

Sort is a convenience method: x.Sort() calls Sort(x).

Sort是一个方便的方法：x.Sort()调用Sort(x)。


### func (Float64Slice) Swap

`func (x Float64Slice) Swap(i, j int)`

## type IntSlice

`type IntSlice []int`

IntSlice attaches the methods of Interface to []int, sorting in increasing order.

IntSlice将Interface的方法附加到[]int上，按照递增的顺序进行排序。

### func (IntSlice) Len

`func (x IntSlice) Len() int`

### func (IntSlice) Less

`func (x IntSlice) Less(i, j int) bool`

### func (IntSlice) Search 

`func (p IntSlice) Search(x int) int`

Search returns the result of applying SearchInts to the receiver and x.

搜索返回对接收器和x应用SearchInts的结果。

### func (IntSlice) Sort 

`func (x IntSlice) Sort()`

Sort is a convenience method: x.Sort() calls Sort(x).

Sort是一个方便的方法：x.Sort()调用Sort(x)。

### func (IntSlice) Swap 

`func (x IntSlice) Swap(i, j int)`

## type Interface

```
type Interface interface {
	// Len is the number of elements in the collection.
	Len() int

	// Less reports whether the element with index i
	// must sort before the element with index j.
	//
	// If both Less(i, j) and Less(j, i) are false,
	// then the elements at index i and j are considered equal.
	// Sort may place equal elements in any order in the final result,
	// while Stable preserves the original input order of equal elements.
	//
	// Less must describe a transitive ordering:
	//  - if both Less(i, j) and Less(j, k) are true, then Less(i, k) must be true as well.
	//  - if both Less(i, j) and Less(j, k) are false, then Less(i, k) must be false as well.
	//
	// Note that floating-point comparison (the < operator on float32 or float64 values)
	// is not a transitive ordering when not-a-number (NaN) values are involved.
	// See Float64Slice.Less for a correct implementation for floating-point values.
	Less(i, j int) bool

	// Swap swaps the elements with indexes i and j.
	Swap(i, j int)
}
```

An implementation of Interface can be sorted by the routines in this package. The methods refer to elements of the underlying collection by integer index.

Interface的一个实现可以通过这个包中的例程进行排序。这些方法通过整数索引来引用底层集合的元素。

### func Reverse 

`func Reverse(data Interface) Interface`

Reverse returns the reverse order for data.

反向返回数据的反向顺序。


## type StringSlice

`type StringSlice []string`

StringSlice attaches the methods of Interface to []string, sorting in increasing order.

StringSlice将Interface的方法附加到[]string上，以递增的顺序进行排序。

### func (StringSlice) Len

`func (x StringSlice) Len() int`

### func (StringSlice) Less 

`func (x StringSlice) Less(i, j int) bool`

### func (StringSlice) Search

`func (p StringSlice) Search(x string) int`

Search returns the result of applying SearchStrings to the receiver and x.

搜索返回对接收器和x应用SearchStrings的结果。

### func (StringSlice) Sort

`func (x StringSlice) Sort()`

Sort is a convenience method: x.Sort() calls Sort(x).

Sort是一个方便的方法：x.Sort()调用Sort(x)。

### func (StringSlice) Swap 

`func (x StringSlice) Swap(i, j int)`
