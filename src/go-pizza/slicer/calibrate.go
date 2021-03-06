package slicer

import (
	"fmt"
	"go-pizza/subslice"
	"math"
	"go-pizza/coord"
	"sort"
)

func (s *Slicer) calibrate() {
	blackPointList := make([]*coord.Point, 0)

	for i := int64(0); i < s.height; i++ {
		for j := int64(0); j < s.width; j++ {
			if !s.filled[i * s.width + j] {
				blackPointList = append(blackPointList, &coord.Point{X: j, Y: i})
			}
		}
	}

	// todo optimize with clusterisation

	subSliceList := make([]*subSliceStatistic, 0, len(blackPointList))

	for _, blackPoint := range blackPointList {
		count := s.countEmptyInSector(s.getRectangleAroundPoint(blackPoint))
		subSliceList = append(subSliceList, &subSliceStatistic{Point: blackPoint, CountEmpty: count})
	}

	sort.Sort(sortSubSliceStatistic(subSliceList))

	for _, sbSlice := range subSliceList {
		fmt.Printf("Empty fields: %d\n", sbSlice.CountEmpty)
		subSlice := s.importToSubSlice(s.getRectangleAroundPoint(sbSlice.Point))
		subSlice.Defragmentate(s.shapeList)
		// todo iterate while not get perfect result
		// todo import best result in main matrix
		// todo recheck
		break
	}
}

func (s *Slicer) getRectangleAroundPoint(point *coord.Point) (*coord.Point, *coord.Point) {
	startX := math.Max(float64(point.X) - float64(s.maxSlice) * 1.5, 0)
	startY := math.Max(float64(point.Y) - float64(s.maxSlice) * 1.5, 0)
	stopX := math.Min(float64(point.X) + float64(s.maxSlice) * 1.5, float64(s.width))
	stopY := math.Min(float64(point.Y) + float64(s.maxSlice) * 1.5, float64(s.height))
	return &coord.Point{X: int64(startX), Y: int64(startY)}, &coord.Point{X: int64(stopX), Y: int64(stopY)}
}

func (s *Slicer) countEmptyInSector(start *coord.Point, stop *coord.Point) int {
	sizeSubSlicerWidth := stop.X - start.X
	sizeSubSlicerHeight := stop.Y - start.Y
	count := 0
	for i := int64(0); i < sizeSubSlicerHeight; i++ {
		offset := s.getOffset(start.X, start.Y + int64(i))
		for j := int64(0); j < sizeSubSlicerWidth; j++ {
			if !s.filled[offset + j] {
				count++
			}
		}
	}
	return count
}

func (s *Slicer) importToSubSlice(start *coord.Point, stop *coord.Point) *subslice.SubSlicer {
	sizeSubSlicerWidth := int(stop.X - start.X)
	sizeSubSlicerHeight := int(stop.Y - start.Y)

	subSlicer := subslice.SubSlicer{
		Height: sizeSubSlicerHeight,
		Width: sizeSubSlicerWidth,
		Buffer: make([]byte, sizeSubSlicerHeight * sizeSubSlicerWidth, sizeSubSlicerHeight * sizeSubSlicerWidth),
		Filled: make([]bool, sizeSubSlicerHeight * sizeSubSlicerWidth, sizeSubSlicerHeight * sizeSubSlicerWidth),
	}

	// copy cross slices
	slices := make([]Slice, 0)
	for x := start.X; x < stop.X; x++ {
		for y := start.Y; y < stop.Y; y++ {
			slice := s.findSlice(int64(x), int64(y))
			if slice != nil {
				if slice.X >= start.X &&
					slice.Y >= start.Y &&
					slice.X + int64(slice.Shape.Width) <= stop.X &&
					slice.Y + int64(slice.Shape.Height) <= stop.Y {
					found := false
					for _, tSlice := range slices {
						if tSlice.X == slice.X && tSlice.Y == slice.Y {
							found = true
							break
						}
					}
					if !found {
						slices = append(slices, *slice)
					}
				}
			}
		}
	}

	for i := 0; i < sizeSubSlicerHeight; i++ {
		offset := s.getOffset(start.X, start.Y + int64(i))
		for j := 0; j < sizeSubSlicerWidth; j++ {
			subSlicer.Filled[i * sizeSubSlicerWidth + j] = s.filled[offset + int64(j)]
			subSlicer.Buffer[i * sizeSubSlicerWidth + j] = s.stream[offset + int64(j)]
		}
	}

	subSlicer.Slices = make([]subslice.Slice, 0, len(slices))

	for _, slice := range slices {
		nSlice := subslice.Slice{
			X: slice.X - start.X,
			Y: slice.Y - start.Y,
			Shape: slice.Shape,
		}
		subSlicer.Slices = append(subSlicer.Slices, nSlice)
	}

	return &subSlicer
}

type subSliceStatistic struct {
	Point *coord.Point
	CountEmpty int
}

type sortSubSliceStatistic []*subSliceStatistic

func (s sortSubSliceStatistic) Len() int {
	return len(s)
}

func (s sortSubSliceStatistic) Less(i, j int) bool {
	return s[i].CountEmpty > s[j].CountEmpty
}

func (s sortSubSliceStatistic) Swap(i, j int) {
	s[j], s[i] = s[i], s[j]
}
