package strings

// Testing various functions within the
// standard library's strings package.

import (
	"strings"
	"testing"
)

func assertCorrectMessageInt(t testing.TB, expected, actual int) {
	t.Helper()
	if actual != expected {
		t.Errorf("expected %q but got %q", expected, actual)
	}
}

func assertCorrectMessageStr(t testing.TB, expected, actual string) {
	t.Helper()
	if actual != expected {
		t.Errorf("expected %q but got %q", expected, actual)
	}
}

func TestCompare(t *testing.T) {
	t.Run("Compare: first is 'less than' second", func(t *testing.T) {
		actual := strings.Compare("a", "b")
		expected := -1
		assertCorrectMessageInt(t, expected, actual)
	})

	t.Run("Compare: first is 'greater than' second", func(t *testing.T) {
		actual := strings.Compare("b", "a")
		expected := 1
		assertCorrectMessageInt(t, expected, actual)
	})

	t.Run("Compare: first is 'equal to' second", func(t *testing.T) {
		actual := strings.Compare("b", "b")
		expected := 0
		assertCorrectMessageInt(t, expected, actual)
	})
}
