package interview

import (
	"fmt"
	"testing"
)

func TestReplaceBlank(t *testing.T) {
	str := "abc sdfsfl"
	s, res := ReplaceBlank(str)

	fmt.Println(s, res)
}

func TestMove(t *testing.T) {
	x, y, z := Move("R2(LF)", 0, 0, 1)
	fmt.Println(x, y, z)
}
