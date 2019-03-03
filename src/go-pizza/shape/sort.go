package shape

type Sort []Shape

func (s Sort) Len() int {
	return len(s)
}

func (s Sort) Less(i, j int) bool {
	return (s[i].Width * s[i].Height) < (s[j].Width * s[j].Height)
}

func (s Sort) Swap(i, j int) {
	s[j], s[i] = s[i], s[j]
}
