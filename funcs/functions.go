package funcs

// Plus returns sum of 3 ints
func Plus(nums ...int) int {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	return sum
}
