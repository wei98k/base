package interview

import (
	"fmt"
	"testing"
)

func TestReverString(t *testing.T) {
	str := "abcabckkkpd"
	rstr, res := ReverString(str)

	fmt.Println(rstr, res)
}

func TestIsRegroup(t *testing.T) {
	s1 := "abcdefg"
	s2 := "abcdefg"
	res := IsRegroup(s1, s2)
	fmt.Println(res)
}
