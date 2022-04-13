## Overview

Package crc32 implements the 32-bit cyclic redundancy check, or CRC-32, checksum. See https://en.wikipedia.org/wiki/Cyclic_redundancy_check for information.

Polynomials are represented in LSB-first form also known as reversed representation.

See https://en.wikipedia.org/wiki/Mathematics_of_cyclic_redundancy_checks#Reversed_representations_and_reciprocal_polynomials for information.

包crc32实现了32位循环冗余校验，即CRC-32，校验和。参见https://en.wikipedia.org/wiki/Cyclic_redundancy_check，了解相关信息。

多项式以LSB优先的形式表示，也被称为反转表示。

有关信息，请参见https://en.wikipedia.org/wiki/Mathematics_of_cyclic_redundancy_checks#Reversed_representations_and_reciprocal_polynomials。

## Constants

```
const (
	// IEEE is by far and away the most common CRC-32 polynomial.
	// Used by ethernet (IEEE 802.3), v.42, fddi, gzip, zip, png, ...
	IEEE = 0xedb88320

	// Castagnoli's polynomial, used in iSCSI.
	// Has better error detection characteristics than IEEE.
	// https://dx.doi.org/10.1109/26.231911
	Castagnoli = 0x82f63b78

	// Koopman's polynomial.
	// Also has better error detection characteristics than IEEE.
	// https://dx.doi.org/10.1109/DSN.2002.1028931
	Koopman = 0xeb31d82e
)
```

Predefined polynomials.

预定义的多项式。

`const Size = 4`

The size of a CRC-32 checksum in bytes.

CRC-32校验和的大小，以字节为单位。

## Variables

`var IEEETable = simpleMakeTable(IEEE)`

IEEETable is the table for the IEEE polynomial.

IEEETable是IEEE多项式的表格。

## Functions

### func Checksum 

`func Checksum(data []byte, tab *Table) uint32`

Checksum returns the CRC-32 checksum of data using the polynomial represented by the Table.

Checksum使用表所代表的多项式返回数据的CRC-32校验和。

### func ChecksumIEEE

`func ChecksumIEEE(data []byte) uint32`

ChecksumIEEE returns the CRC-32 checksum of data using the IEEE polynomial.

### func New 

`func New(tab *Table) hash.Hash32`

New creates a new hash.Hash32 computing the CRC-32 checksum using the polynomial represented by the Table. Its Sum method will lay the value out in big-endian byte order. The returned Hash32 also implements encoding.BinaryMarshaler and encoding.BinaryUnmarshaler to marshal and unmarshal the internal state of the hash.

Hash32使用表所代表的多项式计算CRC-32校验和。它的Sum方法将以big-endian的字节顺序排列数值。返回的Hash32还实现了encoding.BinaryMarshaler和encoding.BinaryUnmarshaler，以对哈希的内部状态进行处理和解除处理。

### func NewIEEE

`func NewIEEE() hash.Hash32` 

NewIEEE creates a new hash.Hash32 computing the CRC-32 checksum using the IEEE polynomial. Its Sum method will lay the value out in big-endian byte order. The returned Hash32 also implements encoding.BinaryMarshaler and encoding.BinaryUnmarshaler to marshal and unmarshal the internal state of the hash.

NewIEEE创建一个新的Hash.Hash32，使用IEEE多项式计算CRC-32校验和。它的Sum方法将以big-endian的字节顺序排列数值。返回的Hash32还实现了encoding.BinaryMarshaler和encoding.BinaryUnmarshaler，以对哈希的内部状态进行处理和解除处理。

### func Update

`func Update(crc uint32, tab *Table, p []byte) uint32`

Update returns the result of adding the bytes in p to the crc.

更新返回将p中的字节添加到crc的结果

## type Table

```
type Table [256]uint32
```

Table is a 256-word table representing the polynomial for efficient processing.

表是一个256字的表，代表多项式的有效处理。

### func MakeTable

`func MakeTable(poly uint32) *Table`

MakeTable returns a Table constructed from the specified polynomial. The contents of this Table must not be modified.

MakeTable返回一个由指定多项式构建的表。这个表的内容不能被修改。



























