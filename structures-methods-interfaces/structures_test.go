package structures_methods_interfaces

import "testing"

func TestPerimeter(t *testing.T) {
	result := Perimeter(Retangle{10.0, 10.0})
	expected := 40.0

	if result != expected {
		t.Errorf("got %.2f want %.2f", result, expected)
	}
}

func TestArea(t *testing.T) {
	checkArea := func(t *testing.T, form Form, expected float64) {
		t.Helper()
		result := form.Area()
		if result != expected {
			t.Errorf("got %.2f want %.2f", result, expected)
		}
	}

	t.Run("Area of retangle", func(t *testing.T) {
		retangle := Retangle{Length: 12.0, Width: 6.0}
		checkArea(t, retangle, 72.0)
	})

	t.Run("Area of circle", func(t *testing.T) {
		circle := Circle{10}
		checkArea(t, circle, 314.1592653589793)
	})
}
