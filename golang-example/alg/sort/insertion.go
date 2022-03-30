package mysort

func insertion(arr []int) {
	// 从下标1开始-遍历一次数组
	// 取出下标是1的值
	// 定义第二次循环 下标从0开始
	l := len(arr)
	for i := 1; i < l; i++ {
		val := arr[i]
		j := i - 1
		for ; j >= 0; j-- {
			if arr[j] > val {
				arr[j+1] = arr[j]
			} else {
				break
			}
		}
		arr[j+1] = val
	}
}
