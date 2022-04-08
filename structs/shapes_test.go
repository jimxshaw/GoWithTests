package structs

import "testing"

func TestPerimeter(t *testing.T) {
	got := Perimeter(2.0, 4.0)
	expected := 12.0

	if got != expected {
		// The new f format string is for
		// float64 to 2 decimal places.
		t.Errorf("expected %.2f but got %.2f", expected, got)
	}
}
