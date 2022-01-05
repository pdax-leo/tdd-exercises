package calc

// SumAll receives slice of number slices and return a slice with total of each slice
func SumAll(numbersToSum ...[]int) []int {
	result := []int{}

	for _, numbers := range numbersToSum {
		result = append(result, Sum(numbers))
	}

	return result
}
