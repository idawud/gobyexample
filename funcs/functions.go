package funcs

// Plus returns sum of 3 ints
func Plus(nums ...int) int {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	return sum
}

// Vals get multple returns
func Vals() (int, int) {
	return 3, 7
}

// Closures return another func
func Closures() func() int {
	count := 0
	return func() int {
		count++
		return count
	}
}
