package arrays

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("Collection of size 5", func(t *testing.T) {
		// Arrays have a fixed capacity that is
		// defined when they are declared.
		numbers := []int{2, 4, 6, 8, 10}

		got := Sum(numbers)
		expected := 30

		if got != expected {
			t.Errorf("expected %d but got %d given %v", expected, got, expected)
		}
	})

	t.Run("Collection of any size", func(t *testing.T) {
		numbers := []int{1, 3, 5}

		got := Sum(numbers)
		expected := 9

		if got != expected {
			t.Errorf("expected %d but got %d given %v", expected, got, expected)
		}
	})
}

func TestSumAll(t *testing.T) {
	got := SumAll([]int{1, 2}, []int{7, 3})
	expected := []int{3, 10}

	// DeepEqual (not type safe) checks if any two
	// variables are equal.
	if !reflect.DeepEqual(got, expected) {
		t.Errorf("expected %d but got %d", expected, got)
	}
}

func TestSumAllTails(t *testing.T) {
	checkSums := func(t testing.TB, got, expected []int) {
		t.Helper()
		if !reflect.DeepEqual(got, expected) {
			t.Errorf("expected %d but got %d", expected, got)
		}
	}

	t.Run("Sum some slices", func(t *testing.T) {
		got := SumAllTails([]int{1, 2}, []int{7, 3})
		expected := []int{2, 3}

		checkSums(t, got, expected)
	})

	t.Run("Safely sum empty slices", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{7, 3})
		expected := []int{0, 3}

		checkSums(t, got, expected)
	})

}
