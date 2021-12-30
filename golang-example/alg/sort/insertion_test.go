package mysort

import "testing"

func TestInsertion(t *testing.T) {
	arr := []int{1, 9, 2, 4, 4, 6, 2}
	insertion(arr)
	t.Log(arr)
}
