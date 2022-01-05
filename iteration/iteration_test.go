package iteration

import "testing"

func TestRepeat(t *testing.T) {
	t.Run("repeats an inputed character 5 times", func(t *testing.T) {
		got := Repeat("a", 5)
		want := "aaaaa"

		if got != want {
			t.Errorf("got %s, want %s", got, want)
		}
	})

	t.Run("repeats an inputed character 5 times", func(t *testing.T) {
		got := Repeat("e", 5)
		want := "eeeee"

		if got != want {
			t.Errorf("got %s, want %s", got, want)
		}
	})

	t.Run("repeats an inputed character based on a specified number of iteration", func(t *testing.T) {
		got := Repeat("e", 2)
		want := "ee"

		if got != want {
			t.Errorf("got %s, want %s", got, want)
		}
	})
}
