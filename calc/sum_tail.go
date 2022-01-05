package calc

// SumTail receives slice of number slices and return a slice with total of each slice
func SumTail(numbersToSum ...[]int) []int {
	var result []int

	for _, numbers := range numbersToSum {
		if len(numbers) == 0 {
			result = append(result, 0)
		} else {
			num := numbers[1:]
			result = append(result, Sum(num))
		}
	}

	return result
}
