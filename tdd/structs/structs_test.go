package structs_test

import (
	"testing"

	"github.com/cpustejovsky/ready-steady-go/tdd/structs"
)

func TestPerimeter(t *testing.T) {
	perimeterTests := []struct {
		shape structs.Shape
		want  float64
	}{
		{&structs.Rectangle{12, 6}, 36},
		{&structs.Circle{10}, 62.83185307179586},
		{&structs.Triangle{12, 6, []float64{1, 2, 3}}, 6},
	}
	for _, tt := range perimeterTests {
		got := tt.shape.Perimeter()
		if got != tt.want {
			t.Errorf("got %g want %g", got, tt.want)
		}
	}
}

func TestArea(t *testing.T) {

	areaTests := []struct {
		name  string
		shape structs.Shape
		want  float64
	}{
		{"Rectangle", &structs.Rectangle{12, 6}, 72.0},
		{"Circle", &structs.Circle{10}, 314.1592653589793},
		{"Triangle", &structs.Triangle{12, 6, []float64{1, 2, 3}}, 36.0},
	}

	for _, tt := range areaTests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.shape.Area()
			if got != tt.want {
				t.Errorf("got %g want %g", got, tt.want)
			}
		})
	}

}
