package slicer

type Slicer struct {
	height int64
	width int64
	stream []byte
	filled []bool
	minSlice int
	maxSlice int
}

func (s *Slicer) getOffset(x, y int64) int64 {
	return x * s.width + y
}
