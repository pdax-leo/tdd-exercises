package iteration

// Repeat func repeats char
func Repeat(str string, iteration int) string {
	var repeated string

	for i := 0; i < iteration; i++ {
		repeated += str
	}

	return repeated
}
