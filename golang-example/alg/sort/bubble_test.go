package mysort

import "testing"

func TestMysort(t *testing.T) {
	a1 := []int{1, 3, 5, 3, 5}
	bubble(a1)
	t.Log(a1)
}
