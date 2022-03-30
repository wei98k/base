package recu

import "fmt"

// 九九递归
func Mul(i, j int) {
	if i < 10 {
		if j < 10 {
			fmt.Printf("%d * %d = %2d \n", i, j, i*j)
			Mul(i, j+1)
		} else {
			fmt.Println("kk")
			i++
			Mul(i, 1)
		}
	}
}

func Mul2(i int) {
	if i == 1 {
		fmt.Println("1 * 1 = 1")
	} else {
		Mul2(i - 1)
		for j := 1; j <= i; j++ {
			fmt.Printf("%d * %d = %2d ", j, i, j*i)
		}
	}
}

// https://blog.csdn.net/hcsnxdld/article/details/107904655
