package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	assertCorrectMessage := func(t testing.TB, expected, repeated string) {
		t.Helper()
		if repeated != expected {
			t.Errorf("expected %q but got %q", expected, repeated)
		}
	}

	t.Run("Repeat a specified number of times", func(t *testing.T) {
		repeated := Repeat("y", 3)
		expected := "yyy"
		assertCorrectMessage(t, expected, repeated)
	})

	t.Run("Prevent repeat if invalid input", func(t *testing.T) {
		repeated := Repeat("y", 0)
		expected := "howManyTimes can only be between 1 and 1000"
		assertCorrectMessage(t, expected, repeated)
	})

}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("y", 2)
	}
}

func ExampleRepeat() {
	repeated := Repeat("a", 7)
	fmt.Println(repeated)
	// Output: aaaaaaa
}
