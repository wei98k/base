package array

//
func duplicate(numbers [7]int) (duplicate int, succ bool) {

	for i := 0; i < len(numbers); i++ {
		for numbers[i] != i {
			if numbers[i] == numbers[numbers[i]] {

				return numbers[i], true
			}
			numbers[i], numbers[numbers[i]] = numbers[numbers[i]], numbers[i]
		}
	}
	return 0, false
}
