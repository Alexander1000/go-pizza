package subslice

type SubSlicer struct {
	Width int
	Height int
	Buffer []byte
	Filled []bool
	Slices []Slice
}

func (s *SubSlicer) CountEmpty() int {
	count := 0
	for _, filled := range s.Filled {
		if !filled {
			count++
		}
	}
	return count
}
