package subslice

import (
	"go-pizza/shape"
)

type SubSlicer struct {
	Width int
	Height int
	Buffer []byte
	Filled []bool
	Slices []Slice
	countEmpty *int
}

func (s *SubSlicer) getOffset(x, y int) int {
	return y * s.Width + x
}

func (s *SubSlicer) CountEmpty() int {
	if s.countEmpty != nil {
		return *s.countEmpty
	}

	count := 0
	for _, filled := range s.Filled {
		if !filled {
			count++
		}
	}
	s.countEmpty = &count
	return count
}

func (s *SubSlicer) validateShapeForFill(x, y int, shape shape.Shape) bool {
	for i := 0; i < shape.Height; i++ {
		offset := s.getOffset(x, y + i)
		for j := 0; j < shape.Width; j++ {
			if s.Filled[offset + j] {
				return false
			}
		}
	}
	return true
}

func (s *SubSlicer) getStreamForShape(x, y int, shape shape.Shape) []byte {
	stream := make([]byte, 0, shape.Width * shape.Height)
	for i := 0; i < shape.Height; i++ {
		offset := s.getOffset(x, y + i)
		for j := 0; j < shape.Width; j++ {
			stream = append(stream, s.Buffer[offset + j])
		}
	}
	return stream
}