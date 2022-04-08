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
	t.Run("Rectangles", func(t *testing.T) {
		rectangle := Rectangle{2.0, 4.0}
		got := rectangle.Area()
		expected := 8.0

		if got != expected {
			t.Errorf("expected %.2f but got %.2f", expected, got)
		}
	})

	t.Run("Circles", func(t *testing.T) {
		circle := Circle{10.0}
		got := circle.Area()
		expected := 314.1592653589793

		if got != expected {
			// The g format string will print more
			// precise decimal number in the message.
			t.Errorf("expected %g but got %g", expected, got)
		}
	})

}
