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
	areaTests := []struct {
		Shape    Shape
		Expected float64
	}{
		{Shape: Rectangle{Width: 2.0, Height: 4.0}, Expected: 8.0},
		{Shape: Circle{Radius: 10.0}, Expected: 314.1592653589793},
		{Shape: Triangle{Base: 6, Height: 6}, Expected: 18.0},
	}

	// This is an example of Table Driven Testing:
	// https://github.com/golang/go/wiki/TableDrivenTests
	for _, tableTest := range areaTests {
		got := tableTest.Shape.Area()

		if got != tableTest.Expected {
			// The g format string will print more
			// precise decimal number in the message.
			t.Errorf("expected %g but got %g", tableTest.Expected, got)
		}
	}
}
