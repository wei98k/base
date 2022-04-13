## Overview

Package crc64 implements the 64-bit cyclic redundancy check, or CRC-64, checksum. See https://en.wikipedia.org/wiki/Cyclic_redundancy_check for information.

包crc64实现了64位循环冗余校验，即CRC-64，校验。参见https://en.wikipedia.org/wiki/Cyclic_redundancy_check，了解相关信息。

## Constants

```
const (
	// The ISO polynomial, defined in ISO 3309 and used in HDLC.
	ISO = 0xD800000000000000

	// The ECMA polynomial, defined in ECMA 182.
	ECMA = 0xC96C5795D7870F42
)
```

Predefined polynomials.

预定义的多项式。

-----

```
const Size = 8
```

The size of a CRC-64 checksum in bytes.

CRC-64校验和的大小，以字节为单位。


## Functions

### func Checksum

`func Checksum(data []byte, tab *Table) uint64`

Checksum returns the CRC-64 checksum of data using the polynomial represented by the Table.

Checksum使用表所代表的多项式返回数据的CRC-64校验和。


### func New

`func New(tab *Table) hash.Hash64`

New creates a new hash.Hash64 computing the CRC-64 checksum using the polynomial represented by the Table. Its Sum method will lay the value out in big-endian byte order. The returned Hash64 also implements encoding.BinaryMarshaler and encoding.BinaryUnmarshaler to marshal and unmarshal the internal state of the hash.

Hash64使用表所代表的多项式计算CRC-64校验和。它的Sum方法将以big-endian的字节顺序排列数值。返回的Hash64还实现了encoding.BinaryMarshaler和encoding.BinaryUnmarshaler，以对哈希的内部状态进行处理和解除处理。

### func Update

`func Update(crc uint64, tab *Table, p []byte) uint64`

Update returns the result of adding the bytes in p to the crc.

更新返回将p中的字节添加到crc的结果。

## type Table 

```
type Table [256]uint64
```

Table is a 256-word table representing the polynomial for efficient processing.

表是一个256字的表，代表多项式的有效处理。

### func MakeTable

`func MakeTable(poly uint64) *Table`

MakeTable returns a Table constructed from the specified polynomial. The contents of this Table must not be modified.

MakeTable返回一个由指定多项式构建的表。这个表的内容不能被修改。