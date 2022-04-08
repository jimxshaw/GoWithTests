package structs

import (
	"fmt"
	"testing"
)

func TestPerimeter(t *testing.T) {
	got := Perimeter(2.0, 4.0)
	expected := 20.0

	if got != expected {
		fmt.Errorf("expected %.2f but got %.2f", expected, got)
	}
}
