package funcs

// Plusx returns sum of 3 ints
func Plusx(nums ...int) int {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	return sum
}

// Plus returns sum of 3 ints
func Plus(nums []int) int {
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

//Recursions
//Fact
func Fact(n int) int {
	if n <= 0 {
		return 1
	}
	return n * Fact(n-1)
}
