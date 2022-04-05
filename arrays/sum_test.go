package arrays

import "testing"

func TestSum(t *testing.T) {
	assertCorrectMessage := func(t testing.TB, got, expected int, given [5]int) {
		t.Helper()
		if got != expected {
			t.Errorf("expected %d but got %d given %v", expected, got, given)
		}
	}

	t.Run("Sum all ints in an array", func(t *testing.T) {
		// Arrays have a fixed capacity that is
		// defined when they are declared.
		numbers := [5]int{2, 4, 6, 8, 10}

		got := Sum(numbers)
		expected := 30

		assertCorrectMessage(t, got, expected, numbers)
	})

}
