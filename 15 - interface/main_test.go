package main

import (
	"math"
	"testing"
)

func TestArea(t *testing.T) {
	t.Run("Rectangle", func(t *testing.T) {
		t.Parallel()

		testableRectangles := []Rectangle{
			{1, 2},
			{3, 4},
			{100, 31232},
			{1.23123, 312.3123},
			{43, 31},
		}

		for _, rectangle := range testableRectangles {
			returned := rectangle.Area()
			expected := rectangle.heigth * rectangle.width

			if expected != returned {
				t.Fatalf(
					"\nReturned: %f\nExpected: %f",
					returned,
					expected,
				)
			}
		}
	})

	t.Run("Circle", func(t *testing.T) {
		t.Parallel()

		testableCircles := []Circle{
			{radius: 1},
			{radius: 3},
			{radius: 00},
			{radius: .23123},
			{radius: 43123},
		}

		for _, circle := range testableCircles {
			returned := circle.Area()
			expected := math.Pow(circle.radius, 2) * math.Pi

			if expected != returned {
				t.Fatalf(
					"\nReturned: %f\nExpected: %f",
					returned,
					expected,
				)
			}
		}
	})
}
