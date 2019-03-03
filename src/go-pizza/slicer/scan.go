package slicer

import (
	"go-pizza/shape"
)

func (s *Slicer) Scan() {
	shapeList := shape.Generate(s.minSlice, s.maxSlice)

	for i := int64(0); i < s.height; i++ {
		for j := int64(0); j < s.width; j++ {
			for _, shape := range shapeList {
				if s.validateShape(i, j, shape) {
					continue
				}
			}
		}
	}
}

func (s *Slicer) validateShape(x, y int64, shape shape.Shape) bool {
	return false
}
