package subslice

import (
	"go-pizza/shape"
	"sort"
)

func (s *SubSlicer) Defragmentate(shapeList []shape.Shape) {
	for _, slice := range s.Slices {
		s.ClearSlice(&slice)
	}
	s.Slices = []Slice{}

	sort.Sort(Sort(shapeList))

	for i := 0; i < s.Height; i++ {
		for j := 0; j < s.Width; j++ {
			for _, shape := range shapeList {
				if s.validateShape(j, i, &shape) {
					// todo set value filled
					break
				}
			}
		}
	}
}

func (s *SubSlicer) validateShape(x, y int, shape *shape.Shape) bool {
	if y + shape.Height > s.Height || x + shape.Width > s.Width {
		return false
	}

	if !s.validateShapeForFill(x, y, *shape) {
		return false
	}

	// todo check assorti

	return true
}

type Sort []shape.Shape

func (s Sort) Len() int {
	return len(s)
}

func (s Sort) Less(i, j int) bool {
	return (s[i].Width * s[i].Height) > (s[j].Width * s[j].Height)
}

func (s Sort) Swap(i, j int) {
	s[j], s[i] = s[i], s[j]
}
