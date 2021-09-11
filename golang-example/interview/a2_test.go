package interview

import (
	"fmt"
	"testing"
)

func TestIsUniqueString(t *testing.T) {
	str := "abc123"

	res := IsUniqueString(str)
	fmt.Println("out: ", res)
}
