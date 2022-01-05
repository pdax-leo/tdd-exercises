package calc

import (
	"reflect"
	"testing"
)

func TestSumTail(t *testing.T) {

	cases := []struct {
		Description string
		Got         []int
		Want        []int
	}{
		{
			"given slice of number slices, returns a slice with the slice tail total",
			SumTail([]int{1, 2, 3}, []int{1, 2}),
			[]int{5, 2},
		},
		{
			"safely sum empty slices",
			SumTail([]int{}, []int{3, 4, 5}),
			[]int{0, 9},
		},
	}

	for _, test := range cases {
		t.Run(test.Description, func(t *testing.T) {
			if !reflect.DeepEqual(test.Got, test.Want) {
				t.Errorf("got %v want %v", test.Got, test.Want)
			}
		})
	}

}
