package structs_and_methods

import "testing"

func TestPerimeter(t *testing.T) {

	t.Run("rectangle", func(t *testing.T) {
		rectangle := Rectangle{10.0, 10.0}
		got := rectangle.Perimeter()
		want := 40.0

		if got != want {
			t.Errorf("Expected '%.2f' got '%.2f'", want, got)
		}
	})

	t.Run("circle", func(t *testing.T) {
		circle := Circle{10.0}
		got := circle.Perimeter()
		want := 20 * 3.141592653589793

		if got != want {
			t.Errorf("Expected %g got %.g", want, got)
		}
	})

}

func TestArea(t *testing.T) {
	//Table driven tests:
	//They are a great fit when you wish to test various implementations of an interface,
	//or if the data being passed in to a function has lots of different requirements that need testing.
	areaTests := []struct {
		name  string
		shape Shape
		want  float64
	}{
		{"Rectangle", Rectangle{12, 6}, 72.0},
		{"Circle", Circle{10}, 314.1592653589793},
		{"Triangle", Triangle{12, 6}, 36.0},
	}

	for _, tt := range areaTests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.shape.Area()
			if got != tt.want {
				t.Errorf("%#v got %g want %g", tt.shape, got, tt.want)
			}
		})

	}

}
