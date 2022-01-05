package calc

import (
	"reflect"
	"testing"
)

func TestSumTail(t *testing.T) {

	checkSums := func(t testing.TB, got, want []int) {
		t.Helper()
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	}

	t.Run("given slice of number slices, returns a slice with the slice tail total", func(t *testing.T) {
		numbers1 := []int{1, 2, 3}
		numbers2 := []int{1, 2}

		got := SumTail(numbers1, numbers2)
		want := []int{5, 2}
		checkSums(t, got, want)
	})

	t.Run("safely sum empty slices", func(t *testing.T) {
		got := SumTail([]int{}, []int{3, 4, 5})
		want := []int{0, 9}
		checkSums(t, got, want)
	})
}
