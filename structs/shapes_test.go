package structs

import "testing"

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{2.0, 4.0}
	got := Perimeter(rectangle)
	expected := 12.0

	if got != expected {
		// The f format string is for
		// float64 to 2 decimal places.
		t.Errorf("expected %.2f but got %.2f", expected, got)
	}
}

func TestArea(t *testing.T) {
	checkArea := func(t testing.TB, shape Shape, expected float64) {
		t.Helper()
		got := shape.Area()

		if got != expected {
			// The g format string will print more
			// precise decimal number in the message.
			t.Errorf("expected %g but got %g", expected, got)
		}
	}

	t.Run("Rectangles", func(t *testing.T) {
		rectangle := Rectangle{2.0, 4.0}
		expected := 8.0
		checkArea(t, rectangle, expected)
	})

	t.Run("Circles", func(t *testing.T) {
		circle := Circle{10.0}
		expected := 314.1592653589793
		checkArea(t, circle, expected)
	})

}
