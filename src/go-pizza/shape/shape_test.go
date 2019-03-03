package shape

import (
	"testing"
)

func TestGenerate_MinMax_ShapeList(t *testing.T) {
	shapeList := Generate(5, 9)
	if len(shapeList) != 15 {
		t.Fatalf("unexpected length of list, given: %d", len(shapeList))
	}

	dataSource := []Shape{
		{Height: 1, Width: 5},
		{Height: 1, Width: 6},
		{Height: 1, Width: 7},
		{Height: 1, Width: 8},
		{Height: 1, Width: 9},
		{Height: 2, Width: 3},
		{Height: 2, Width: 4},
		{Height: 3, Width: 2},
		{Height: 3, Width: 3},
		{Height: 4, Width: 2},
		{Height: 5, Width: 1},
		{Height: 6, Width: 1},
		{Height: 7, Width: 1},
		{Height: 8, Width: 1},
		{Height: 9, Width: 1},
	}

	for _, expectedShape := range dataSource {
		found := false
		for _, actualShape := range shapeList {
			if actualShape.Height == expectedShape.Height && actualShape.Width == expectedShape.Width {
				found = true
				break
			}
		}
		if !found {
			t.Fatalf("Not found shape(height: %d; width: %d)", expectedShape.Height, expectedShape.Width)
		}
	}
}
