package shapes

import "testing"

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{5.5, 2.5}
	got := Perimeter(rectangle)
	want := 16.0

	if got != want {
		t.Errorf("got %g want %g", got, want)
	}
}

func TestArea(t *testing.T) {
	tests := []struct {
		name  string
		shape Shape
		want  float64
	}{
		{name: "rectangle", shape: Rectangle{Width: 5.5, Height: 2.5}, want: 13.75},
		{name: "circle", shape: Circle{Radius: 10}, want: 314.1592653589793},
		{name: "triangle", shape: Triangle{Base: 10, Height: 5}, want: 25},
	}

	for _, v := range tests {
		t.Run(v.name, func(t *testing.T) {
			got := v.shape.Area()
			if got != v.want {
				t.Errorf("%#v got %g want %g", v.shape, got, v.want)
			}
		})
	}
}
