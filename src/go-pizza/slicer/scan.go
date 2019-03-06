package slicer

import (
	"go-pizza/shape"
	"fmt"
	"sort"
)

func (s *Slicer) Scan() {
	s.slices = make([]Slice, 0)
	s.shapeList = shape.Generate(s.minSlice, s.maxSlice)
	sort.Sort(shape.Sort(s.shapeList))

	// первичная разметка
	for i := int64(0); i < s.height; i++ {
		for j := int64(0); j < s.width; j++ {
			for _, shape := range s.shapeList {
				if s.validateShape(j, i, shape) {
					s.fill(j, i, shape)
					break
				}
			}
		}
	}

	fmt.Printf("count slices: %d\n", len(s.slices))
	countEmpty := 0
	for _, filled := range s.filled {
		if !filled {
			countEmpty++
		}
	}
	percentEmpty := float64(countEmpty) * 100 / float64(len(s.filled))
	fmt.Printf("Total cells: %d\n", len(s.filled))
	fmt.Printf("Empty fields: %d\n", countEmpty)
	fmt.Printf("Percentage of empties: %0.02f%%\n", percentEmpty)

	if countEmpty > 0 && percentEmpty > 0.5 {
		s.calibrate()
	}
}

func (s *Slicer) validateShape(x, y int64, shape shape.Shape) bool {
	if y + int64(shape.Height) > s.height || x + int64(shape.Width) > s.width {
		return false
	}

	if !s.validateShapeForFill(x, y, shape) {
		return false
	}
	assorti := make(map[byte]bool, 2)
	stream := s.getStreamForShape(x, y, shape)
	for _, data := range stream {
		assorti[data] = true
	}

	if len(assorti) != s.minSlice {
		return false
	}
	return true
}
