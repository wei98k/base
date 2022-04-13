## Overview

Package adler32 implements the Adler-32 checksum.

It is defined in RFC 1950:

包adler32实现了Adler-32校验和。

它被定义在RFC 1950中。

## Constants

```
const Size = 4
```

## Functions

### func Checksum

`func Checksum(data []byte) uint32`

Checksum returns the Adler-32 checksum of data.

Checksum返回数据的Adler-32校验和。

### func New

`func New() hash.Hash32`

New returns a new hash.Hash32 computing the Adler-32 checksum. Its Sum method will lay the value out in big-endian byte order. The returned Hash32 also implements encoding.BinaryMarshaler and encoding.BinaryUnmarshaler to marshal and unmarshal the internal state of the hash.

New返回一个新的哈希值。Hash32计算Adler-32校验和。它的Sum方法将以big-endian的字节顺序排列数值。返回的Hash32还实现了encoding.BinaryMarshaler和encoding.BinaryUnmarshaler，以对哈希的内部状态进行marshal和unmarshal。