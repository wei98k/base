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
