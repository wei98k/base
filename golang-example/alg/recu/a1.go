package recu

import "fmt"

func Stepa(n int) int {

	fmt.Println(n)

	if n == 1 {
		return 1
	}

	if n == 2 {
		return 2
	}

	return Stepa(n-1) + Stepa(n-2)
}
