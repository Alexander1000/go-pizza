package slicer

import (
	"go-pizza/shape"
)

type Slicer struct {
	height int64
	width int64
	stream []byte
	filled []bool
	minSlice int
	maxSlice int
	slices []Slice
}

func (s *Slicer) getOffset(x, y int64) int64 {
	return y * s.width + x
}

func (s *Slicer) getStreamForShape(x, y int64, shape shape.Shape) []byte {
	stream := make([]byte, 0, shape.Width * shape.Height)
	for i := 0; i < shape.Height; i++ {
		offset := s.getOffset(x + int64(i), y)
		for j := 0; j < shape.Width; j++ {
			stream = append(stream, s.stream[offset + int64(j)])
		}
	}
	return stream
}

func (s *Slicer) validateShapeForFill(x, y int64, shape shape.Shape) bool {
	for i := 0; i < shape.Height; i++ {
		offset := s.getOffset(x + int64(i), y)
		for j := 0; j < shape.Width; j++ {
			if s.filled[offset + int64(j)] {
				return false
			}
		}
	}

	return true
}

func (s *Slicer) fill(x, y int64, shape shape.Shape) {
	s.slices = append(s.slices, Slice{X: x, Y: y, Shape: shape})
	for i := 0; i < shape.Height; i++ {
		offset := s.getOffset(x + int64(i), y)
		for j := 0; j < shape.Width; j++ {
			s.filled[offset + int64(j)] = true
		}
	}
}
