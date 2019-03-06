package subslice

func (s *SubSlicer) ClearSlice(slice *Slice) {
	for i := 0; i < slice.Shape.Height; i++ {
		for j := 0; j < slice.Shape.Width; j++ {
			s.Filled[(slice.Y + int64(i)) * int64(s.Width) + slice.X + int64(j)] = false
		}
	}
}
