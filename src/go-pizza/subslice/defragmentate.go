package subslice

func (s *SubSlicer) Defragmentate() {
	for _, slice := range s.Slices {
		s.ClearSlice(&slice)
	}
	s.Slices = []Slice{}
}
