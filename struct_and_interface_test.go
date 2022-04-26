/*
 * Copyright (c) 2021.
 */

package main

import (
	"reflect"
	"testing"
)

func TestArea(t *testing.T) {

	areaTests := []struct {
		name      string
		shape     Shape
		area      float64
		perimeter float64
	}{
		{
			name:      "Rectangle",
			shape:     Rectangle{Width: 12, Height: 6},
			area:      72.0,
			perimeter: 36,
		},
		{
			name:      "Circle",
			shape:     Circle{Radius: 10},
			area:      314.1592653589793,
			perimeter: 62.83185307179586,
		},
		{
			name:      "Triangle",
			shape:     Triangle{Base: 12, Height: 6, Sides: []float64{3.0, 3.0, 3.0}},
			area:      36.0,
			perimeter: 9,
		},
	}

	for _, tt := range areaTests {
		t.Run(tt.name+" calculating area", func(t *testing.T) {
			got := tt.shape.Area()
			if !reflect.DeepEqual(got, tt.area) {
				t.Errorf("%#v got %g want %g", tt.shape, got, tt.area)
			}
		})

		t.Run(tt.name+" calculating perimeter", func(t *testing.T) {
			got := tt.shape.Perimeter()
			if !reflect.DeepEqual(got, tt.perimeter) {
				t.Errorf("%#v got %g want %g", tt.shape, got, tt.perimeter)
			}
		})

	}

}
