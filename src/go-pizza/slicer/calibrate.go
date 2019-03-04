package slicer

func (s *Slicer) calibrate() {
	blackPoints := make([]BlackPoint, 0)

	for i := int64(0); i < s.height; i++ {
		for j := int64(0); j < s.width; j++ {
			if !s.filled[i * s.width + j] {
				blackPoints = append(blackPoints, BlackPoint{X: j, Y: i})
			}
		}
	}
}

type BlackPoint struct {
	X int64
	Y int64
}