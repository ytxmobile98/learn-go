package main_test

import (
	"fmt"
	"reflect"
	"slices"
	"testing"
)

// Tip: Run `go test -v` to see output

type Point interface {
	String() string
}

type Point2D struct {
	X, Y float64
}

func (p Point2D) String() string {
	return fmt.Sprintf("(%f, %f)", p.X, p.Y)
}

type Point3D struct {
	X, Y, Z float64
}

func (p Point3D) String() string {
	return fmt.Sprintf("(%f, %f, %f)", p.X, p.Y, p.Z)
}

func TestRangeOverStruct(t *testing.T) {
	t.Run("Testing slice modification", func(t *testing.T) {
		point2Ddata := []Point2D{
			{1, 2}, {3, 4}, {5, 6},
		}

		t.Run("Iterating over a slice of structs, item only, not changing original data",
			func(t *testing.T) {
				points := slices.Clone(point2Ddata)
				for _, point := range points {
					point.X, point.Y = 10, 20
				}
				// Since points are copied over the original on each iteration,
				// the original data is not modified.
				t.Logf("new data: %v, original data: %v", points, point2Ddata)
				if !reflect.DeepEqual(points, point2Ddata) {
					t.Errorf("Expected %v, got %v", point2Ddata, points)
				}
			})

		t.Run("Iterating over a slice of structs, item and index, changing original data",
			func(t *testing.T) {
				points := slices.Clone(point2Ddata)
				for i, point := range points {
					points[i].X, points[i].Y = float64(point.X)*2, float64(point.Y)*3
				}
				expectedNewData := []Point2D{
					{2, 6}, {6, 12}, {10, 18},
				}
				// We have used `points[i]` to modify original data, so `points`
				// should have all items updated
				t.Logf("new data: %v, original data: %v", points, point2Ddata)
				if !reflect.DeepEqual(points, expectedNewData) {
					t.Errorf("Expected %v, got %v", point2Ddata, points)
				}
			})
	})
}
