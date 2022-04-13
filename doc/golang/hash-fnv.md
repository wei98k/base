## Overview

Package fnv implements FNV-1 and FNV-1a, non-cryptographic hash functions created by Glenn Fowler, Landon Curt Noll, and Phong Vo. See https://en.wikipedia.org/wiki/Fowler-Noll-Vo_hash_function.

All the hash.Hash implementations returned by this package also implement encoding.BinaryMarshaler and encoding.BinaryUnmarshaler to marshal and unmarshal the internal state of the hash.

包fnv实现了FNV-1和FNV-1a，由Glenn Fowler、Landon Curt Noll和Phong Vo创建的非加密哈希函数。参见https://en.wikipedia.org/wiki/Fowler-Noll-Vo_hash_function。

该包返回的所有hash.Hash实现也实现了encoding.BinaryMarshaler和encoding.BinaryUnmarshaler，以对哈希的内部状态进行marshal和unmarshal。

## Functions


### func New128

`func New128() hash.Hash`

New128 returns a new 128-bit FNV-1 hash.Hash. Its Sum method will lay the value out in big-endian byte order.

New128返回一个新的128位FNV-1哈希值。它的Sum方法将以big-endian的字节顺序排列数值。

### func New128a

`func New128a() hash.Hash`

New128a returns a new 128-bit FNV-1a hash.Hash. Its Sum method will lay the value out in big-endian byte order.

New128a返回一个新的128位FNV-1a哈希值。它的Sum方法将以big-endian的字节顺序排列数值。

### func New32

`func New32() hash.Hash32`

New32 returns a new 32-bit FNV-1 hash.Hash. Its Sum method will lay the value out in big-endian byte order.

New32返回一个新的32位FNV-1哈希.Hash。它的Sum方法将以big-endian的字节顺序排列数值。

### func New32a

`func New32a() hash.Hash32`

New32a returns a new 32-bit FNV-1a hash.Hash. Its Sum method will lay the value out in big-endian byte order.

New32a返回一个新的32位FNV-1a哈希.Hash。它的Sum方法将以big-endian的字节顺序排列数值。

### func New64

`func New64() hash.Hash64`

New64 returns a new 64-bit FNV-1 hash.Hash. Its Sum method will lay the value out in big-endian byte order.

New64返回一个新的64位FNV-1哈希.Hash。它的Sum方法将以big-endian的字节顺序排列数值。

### func New64a

`func New64a() hash.Hash64`

New64a returns a new 64-bit FNV-1a hash.Hash. Its Sum method will lay the value out in big-endian byte order.

New64a返回一个新的64位FNV-1a hash.Hash。它的Sum方法将以big-endian的字节顺序排列数值。