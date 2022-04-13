## Overview

Package hash provides interfaces for hash functions.

包hash提供了哈希函数的接口。

## type Hash

```
type Hash interface {
	// Write (via the embedded io.Writer interface) adds more data to the running hash.
	// It never returns an error.
	io.Writer

	// Sum appends the current hash to b and returns the resulting slice.
	// It does not change the underlying hash state.
	Sum(b []byte) []byte

	// Reset resets the Hash to its initial state.
	Reset()

	// Size returns the number of bytes Sum will return.
	Size() int

	// BlockSize returns the hash's underlying block size.
	// The Write method must be able to accept any amount
	// of data, but it may operate more efficiently if all writes
	// are a multiple of the block size.
	BlockSize() int
}
```

Hash is the common interface implemented by all hash functions.

Hash implementations in the standard library (e.g. hash/crc32 and crypto/sha256) implement the encoding.BinaryMarshaler and encoding.BinaryUnmarshaler interfaces. Marshaling a hash implementation allows its internal state to be saved and used for additional processing later, without having to re-write the data previously written to the hash. The hash state may contain portions of the input in its original form, which users are expected to handle for any possible security implications.

Compatibility: Any future changes to hash or crypto packages will endeavor to maintain compatibility with state encoded using previous versions. That is, any released versions of the packages should be able to decode data written with any previously released version, subject to issues such as security fixes. See the Go compatibility document for background:

Hash是所有哈希函数实现的通用接口。

标准库中的哈希实现（如hash/crc32和crypto/sha256）实现了encoding.BinaryMarshaler和encoding.BinaryUnmarshaler接口。对散列的实现允许其内部状态被保存并用于以后的额外处理，而不需要重新写之前写入散列的数据。散列状态可能包含原始形式的输入的一部分，用户应该处理任何可能的安全问题。

兼容性。任何未来对哈希或加密包的修改都将努力保持与使用以前版本编码的状态的兼容性。也就是说，任何发布的包的版本都应该能够解码用以前发布的任何版本编写的数据，但要注意安全修复等问题。请参阅Go兼容性文件了解背景。

## type Hash32

```
type Hash32 interface {
	Hash
	Sum32() uint32
}
```

Hash32 is the common interface implemented by all 32-bit hash functions.

Hash32是所有32位哈希函数实现的通用接口。

## type Hash64

```
type Hash64 interface {
	Hash
	Sum64() uint64
}
```
Hash64 is the common interface implemented by all 64-bit hash functions.

Hash64是所有64位哈希函数实现的通用接口。
