package subslice

type SubSlicer struct {
	Width int
	Height int
	Buffer []byte
	Filled []bool
}
