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
					// make fill
					continue
				}
			}
		}
	}
}

func (s *Slicer) validateShape(x, y int64, shape shape.Shape) bool {
	if !s.validateShapeForFill(x, y, shape) {
		return false
	}
	assorti := make(map[byte]bool)
	stream := s.getStreamForShape(x, y, shape)
	for _, data := range stream {
		assorti[data] = true
	}

	if len(assorti) != s.minSlice {
		return false
	}
	return true
}
