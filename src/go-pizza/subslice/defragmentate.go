package subslice

import (
	"go-pizza/shape"
)

func (s *SubSlicer) Defragmentate(shapeList []shape.Shape) {
	for _, slice := range s.Slices {
		s.ClearSlice(&slice)
	}
	s.Slices = []Slice{}


}
