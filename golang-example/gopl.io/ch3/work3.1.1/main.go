package main

import (
	"fmt"
	"strconv"
)

// 关于算术运算、逻辑运算和比较运算的二元运算符，它们按照优先级递减的顺序排列
// *      /      %      <<       >>     &       &^
// +      -      |      ^
// ==     !=     <      <=       >      >=
// &&
// ||

// bit位操作运算符
// &      位运算 AND
// |      位运算 OR
// ^      位运算 XOR
// &^     位清空（AND NOT）
// <<     左移
// >>     右移

func main() {
	// a := 30 * 6
	// b := 30 << 3
	// fmt.Println(a, b)
	// %b 十进制转二进制
	// %d 输出十进制整型
	// %o 十进制转八进制
	// %x 十进制转十六进制
	var b, o, x string = "1001", "123", "ae12"
	var d int16 = 123

	// 二进制转十进制
	dec, err := strconv.ParseUint(b, 2, 8)
	if err != nil {
		return
	}
	fmt.Printf("二转十: %s == %d\n", b, dec)
	// 二转八
	fmt.Printf("二转八: %s == %o\n", b, dec)
	// 二转十六
	fmt.Printf("二转十六: %s == %x\n", b, dec)

	// 八转十 octonary
	oct, err := strconv.ParseUint(o, 8, 8)
	if err != nil {
		panic(err)
	}
	fmt.Printf("八转十: %s == %d\n", o, oct)
	// 八转二
	fmt.Printf("八转二: %s == %b\n", o, oct)
	// 八转十六
	fmt.Printf("八转十六: %s == %X\n", o, oct)
	// 十转二
	fmt.Printf("十转二: %d == %b\n", d, d)
	// 十转八
	fmt.Printf("十转八: %d == %o\n", d, d)
	// 十转十六
	fmt.Printf("十转十六: %d == %x\n", d, d)
	// 十六转十
	hex, err := strconv.ParseUint(x, 16, 32)
	if err != nil {
		panic(err)
	}
	fmt.Printf("十六转十: %s == %d\n", x, hex)
	// 十六转二
	fmt.Printf("十六转二: %s == %b\n", x, hex)
	// 十六转八
	fmt.Printf("十六转八: %s == %o\n", x, hex)

	fmt.Println("==========")

	// bit位操作运算符

	// << 左位移符 >> 右位移符 [x<<n]
	fmt.Println("==== << 左位移符 >> 右位移符")
	var a1, n1 = 7, 2
	var a8, n2 = 42, 3
	// var tt = 7
	// 十进制 7 转成二进制 111 左位移 2 位后 11100   28
	// 十进制 42 转成二进制 10 1010 左位移 3 位后 1 0101 0000  336
	// 十进制 42 转成二进制 10 1010 右位移 3 位后 101  5
	// tt <<= n1
	// fmt.Println(tt)

	fmt.Printf("%d << %d 结果: %d\n", a1, n1, a1<<n1)
	fmt.Printf("%d << %d 结果: %d\n", a8, n2, a8<<n2)
	fmt.Printf("%d >> %d 结果: %d\n", a8, n2, a8>>n2)

	// & 与操作 AND
	fmt.Println("==== & 与操作 AND")
	var aa, bb, cc, dd = 1, 0, 1, 0
	a2 := aa & cc
	a3 := aa & bb
	a4 := bb & dd

	// var b1 = 11101
	// b1 &= 101101
	// fmt.Printf("%b\n", b1)

	fmt.Printf("%d & %d 结果: %d\n", aa, cc, a2)
	fmt.Printf("%d & %d 结果: %d\n", aa, bb, a3)
	fmt.Printf("%d & %d 结果: %d\n", bb, dd, a4)

	// | 或操作 OR
	fmt.Println("==== | 或操作 OR AND")
	a5 := aa | cc
	a6 := aa | bb
	a7 := bb | dd
	fmt.Printf("%d | %d 结果: %d\n", aa, cc, a5)
	fmt.Printf("%d | %d 结果: %d\n", aa, bb, a6)
	fmt.Printf("%d | %d 结果: %d\n", bb, dd, a7)
	// ^ 位运算 XOR 异或
	fmt.Println("=== ^ 位运算 XOR 异或")
	a9 := aa ^ cc
	a10 := aa ^ bb
	a11 := bb ^ dd
	fmt.Printf("%d ^ %d 结果: %d\n", aa, cc, a9)
	fmt.Printf("%d ^ %d 结果: %d\n", aa, bb, a10)
	fmt.Printf("%d ^ %d 结果: %d\n", bb, dd, a11)
	// &^     位清空（AND NOT）
	var ee uint8 = 1<<1 | 1<<5
	fmt.Println("=== &^ 位清空 AND NOT")
	// 1 << 1 = 10  十进制  2
	// 1 << 5 = 100000 十进制  32
	//  10 | 100000
	//  10 0000
	//       10
	//  10 0010  = 32
	var ff uint8 = 1<<1 | 1<<2

	fmt.Printf("%08b\n", ee)
	fmt.Printf("%08b\n", ff)
	fmt.Printf("%08b\n", ee&^ff)
}
