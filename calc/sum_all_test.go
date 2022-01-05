package calc

import (
	"reflect"
	"testing"
)

func TestSumAll(t *testing.T) {
	t.Run("given a slice of slices, returns a slice that contains total for each slice", func(t *testing.T) {
		numbers1 := []int{1, 2, 3}
		numbers2 := []int{1, 2}

		got := SumAll(numbers1, numbers2)
		want := []int{6, 3}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})
}
