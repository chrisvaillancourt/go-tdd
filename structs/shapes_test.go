package structs

import "testing"

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{10.0, 10.0}
	got := Perimeter(rectangle)
	want := 40.0

	if got != want {
		// f for float64 and .2 for 2 decimal places
		t.Errorf("got %.2f want %.2f", got, want)
	}
}

func TestArea(t *testing.T) {
	rectangle := Rectangle{10.0, 5.0}
	got := Area(rectangle)
	want := 50.0
	if got != want {
		// f for float64 and .2 for 2 decimal places
		t.Errorf("got %.2f want %.2f", got, want)
	}
}
