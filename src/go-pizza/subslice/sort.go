package subslice

type Sort []*SubSlicer

func (s Sort) Len() int {
	return len(s)
}

func (s Sort) Less(i, j int) bool {
	return s[i].CountEmpty() > s[j].CountEmpty()
}

func (s Sort) Swap(i, j int) {
	s[j], s[i] = s[i], s[j]
}
