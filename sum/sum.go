package sum

// Sum accepts a slice of numbers and return the total
func Sum(numbers []int) int {
	var sum int

	for _, number := range numbers {
		sum += number
	}

	return sum
}
