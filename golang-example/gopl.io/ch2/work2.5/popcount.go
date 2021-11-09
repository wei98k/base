package popcount

var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

func PopCount(x uint64) int {
	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}

func PopCountClean(x uint64) int {
	n := 0
	for x != 0 {
		n++
		// 先运算然后取地址 赋值
		x &= x - 1
		// 试试把&去掉? 试试把测试单元总的 x 变量值改成 0xAA? 执行观察变化
	}
	return n
}
