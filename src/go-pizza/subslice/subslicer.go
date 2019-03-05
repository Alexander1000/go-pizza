package subslice

type SubSlicer struct {
	Width int
	Height int
	Buffer []byte
	Filled []bool
	Slices []Slice
	countEmpty *int
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
