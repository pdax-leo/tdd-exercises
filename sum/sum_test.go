package sum

import "testing"

func TestSum(t *testing.T) {
	t.Run("accepts an array of numbers and return the total", func(t *testing.T) {
		numbers := []int{1, 2, 3}
		got := Sum(numbers)
		want := 6

		if got != want {
			t.Errorf("got %d, want %d given, %v", got, want, numbers)
		}
	})

}
