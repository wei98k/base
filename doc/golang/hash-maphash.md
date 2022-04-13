## Overview

Package maphash provides hash functions on byte sequences. These hash functions are intended to be used to implement hash tables or other data structures that need to map arbitrary strings or byte sequences to a uniform distribution on unsigned 64-bit integers. Each different instance of a hash table or data structure should use its own Seed.

The hash functions are not cryptographically secure. (See crypto/sha256 and crypto/sha512 for cryptographic use.)

包maphash提供了字节序列的哈希函数。这些哈希函数旨在用于实现哈希表或其他数据结构，需要将任意字符串或字节序列映射到无符号64位整数上的统一分布。散列表或数据结构的每个不同实例都应该使用它自己的Seed。

哈希函数在密码学上是不安全的。(参见 crypto/sha256 和 crypto/sha512 的加密使用）。


## type Hash

```
type Hash struct {
	// contains filtered or unexported fields
}
```

A Hash computes a seeded hash of a byte sequence.

The zero Hash is a valid Hash ready to use. A zero Hash chooses a random seed for itself during the first call to a Reset, Write, Seed, or Sum64 method. For control over the seed, use SetSeed.

The computed hash values depend only on the initial seed and the sequence of bytes provided to the Hash object, not on the way in which the bytes are provided. For example, the three sequences

Hash计算一个字节序列的种子哈希值。

零Hash是一个有效的Hash，可以随时使用。在第一次调用Reset、Write、Seed或Sum64方法时，零Hash为自己选择了一个随机的种子。对于种子的控制，可以使用SetSeed。

计算的哈希值只取决于初始种子和提供给哈希对象的字节序列，而不取决于提供字节的方式。例如，三个序列

### func (*Hash) BlockSize

`func (h *Hash) BlockSize() int`

BlockSize returns h's block size.

BlockSize返回h的块大小。

### func (*Hash) Reset

`func (h *Hash) Reset()`

Reset discards all bytes added to h. (The seed remains the same.)

重置丢弃所有加入h的字节（种子保持不变）。

### func (*Hash) Seed

`func (h *Hash) Seed() Seed`

Seed returns h's seed value.

种子返回h的种子值。

### func (*Hash) SetSeed

`func (h *Hash) SetSeed(seed Seed)`

SetSeed sets h to use seed, which must have been returned by MakeSeed or by another Hash's Seed method. Two Hash objects with the same seed behave identically. Two Hash objects with different seeds will very likely behave differently. Any bytes added to h before this call will be discarded.

SetSeed将h设置为使用种子，种子必须是由MakeSeed或其他哈希的种子方法返回的。两个具有相同种子的哈希对象的行为是相同的。两个具有不同种子的哈希对象很可能会有不同的行为。在这个调用之前添加到h中的任何字节都将被丢弃。

### func (*Hash) Size

`func (h *Hash) Size() int`

Size returns h's hash value size, 8 bytes.

Size返回h的哈希值大小，8字节。

### func (*Hash) Sum

`func (h *Hash) Sum(b []byte) []byte`

Sum appends the hash's current 64-bit value to b. It exists for implementing hash.Hash. For direct calls, it is more efficient to use Sum64.

Sum将哈希的当前64位值追加到b中，它的存在是为了实现hash.Hash。对于直接调用，使用Sum64更有效率。


### func (*Hash) Sum64 

`func (h *Hash) Sum64() uint64`

Sum64 returns h's current 64-bit value, which depends on h's seed and the sequence of bytes added to h since the last call to Reset or SetSeed.

All bits of the Sum64 result are close to uniformly and independently distributed, so it can be safely reduced by using bit masking, shifting, or modular arithmetic.

Sum64返回h的当前64位值，这取决于h的种子和自上次调用Reset或SetSeed后添加到h的字节序列。

Sum64结果的所有位都接近于均匀和独立分布，所以它可以通过使用位屏蔽、移位或模块化运算安全地减少。

### func (*Hash) Write

`func (h *Hash) Write(b []byte) (int, error)`

Write adds b to the sequence of bytes hashed by h. It always writes all of b and never fails; the count and error result are for implementing io.Writer.

Write将b添加到由h散列的字节序列中。它总是写下所有的b，并且从不失败；计数和错误结果是为了实现io.Writer。

### func (*Hash) WriteByte

`func (h *Hash) WriteByte(b byte) error`

WriteByte adds b to the sequence of bytes hashed by h. It never fails; the error result is for implementing io.ByteWriter.

WriteByte将b添加到由h散列的字节序列中。它从未失败；错误结果是实现io.ByteWriter。

### func (*Hash) WriteString 

`func (h *Hash) WriteString(s string) (int, error)`

WriteString adds the bytes of s to the sequence of bytes hashed by h. It always writes all of s and never fails; the count and error result are for implementing io.StringWriter.

WriteString将s的字节添加到h的散列字节序列中。它总是写出所有的s，而且从不失败；计数和错误结果是为实现io.StringWriter而设。

## type Seed

```
type Seed struct {
	// contains filtered or unexported fields
}
```

A Seed is a random value that selects the specific hash function computed by a Hash. If two Hashes use the same Seeds, they will compute the same hash values for any given input. If two Hashes use different Seeds, they are very likely to compute distinct hash values for any given input.

A Seed must be initialized by calling MakeSeed. The zero seed is uninitialized and not valid for use with Hash's SetSeed method.

Each Seed value is local to a single process and cannot be serialized or otherwise recreated in a different process.


种子是一个随机值，用于选择哈希计算的特定哈希函数。如果两个哈希算法使用相同的种子，它们将为任何给定的输入计算相同的哈希值。如果两个哈希使用不同的种子，它们很可能对任何给定的输入计算出不同的哈希值。

一个种子必须通过调用MakeSeed来初始化。零的种子是未初始化的，不能用于Hash的SetSeed方法。

每个种子值都是单个进程的本地值，不能被序列化或在不同的进程中重新创建。


### func MakeSeed

`func MakeSeed() Seed`

MakeSeed returns a new random seed.

MakeSeed返回一个新的随机种子。