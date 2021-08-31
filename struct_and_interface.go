/*
 * Copyright (c) 2021.
 */

package main

import "math"

type Shape interface {
	Area() float64
	Perimeter() float64
}

type Rectangle struct {
	Width  float64
	Height float64
}

// This function calculate the Area of a Rectangle
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// This function calculate the Perimeter of a Rectangle
func (r Rectangle) Perimeter() float64 {
	return (r.Width + r.Height) * 2
}

type Circle struct {
	Radius float64
}

// This function calculate the Area of a Circle
func (c Circle) Area() float64 {
	return math.Pow(c.Radius, 2) * math.Pi
}

// This function calculate the Perimeter of a Circle
func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

type Triangle struct {
	Base   float64
	Height float64
	Sides  []float64
}

// This function calculate the Area of a Triangle
func (t Triangle) Area() float64 {
	return (t.Base * t.Height) * 0.5
}

// This function calculate the Perimeter of a Triangle
func (t Triangle) Perimeter() float64 {
	return AddFloat(t.Sides)
}

func AddFloat(integers []float64) (sum float64) {
	for _, v := range integers {
		sum += v
	}
	return sum
}
