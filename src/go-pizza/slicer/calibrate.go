package slicer

import (
	"fmt"
	"go-pizza/subslice"
	"math"
	"go-pizza/coord"
)

func (s *Slicer) calibrate() {
	blackPointList := make([]coord.Point, 0)

	fmt.Println("Black points")

	for i := int64(0); i < s.height; i++ {
		if i == 0 {
			fmt.Print("=")
			for j := int64(0); j < s.width; j++ {
				fmt.Print("=")
			}
			fmt.Println("=")
		}

		fmt.Print("|")

		for j := int64(0); j < s.width; j++ {
			if !s.filled[i * s.width + j] {
				blackPointList = append(blackPointList, coord.Point{X: j, Y: i})
				fmt.Print("O")
			} else {
				fmt.Print("*")
			}
		}

		fmt.Println("|")
	}

	fmt.Print("=")
	for j := int64(0); j < s.width; j++ {
		fmt.Print("=")
	}
	fmt.Println("=")

	// todo optimize with clusterisation

	subSliceList := make([]*subslice.SubSlicer, 0, len(blackPointList))

	for _, blackPoint := range blackPointList {
		// try optimize/ defragmentation
		startX := math.Max(float64(blackPoint.X) - float64(s.maxSlice) * 1.5, 0)
		startY := math.Max(float64(blackPoint.Y) - float64(s.maxSlice) * 1.5, 0)

		stopX := math.Min(float64(blackPoint.X) + float64(s.maxSlice) * 1.5, float64(s.width))
		stopY := math.Min(float64(blackPoint.Y) + float64(s.maxSlice) * 1.5, float64(s.height))

		subSlice := s.importToSubSlice(int64(startX), int64(startY), int64(stopX), int64(stopY))
		if subSlice == nil {
			fmt.Print("Fiasko\n")
		}

		subSliceList = append(subSliceList, subSlice)

		fmt.Printf("Empty fields: %d\n", subSlice.CountEmpty())
	}

	// todo выбрать subslice with maximum count empties и оптимизировать их
}

func (s *Slicer) importToSubSlice(startX, startY, stopX, stopY int64) *subslice.SubSlicer {
	sizeSubSlicerWidth := int(stopX - startX)
	sizeSubSlicerHeight := int(stopY - startY)

	subSlicer := subslice.SubSlicer{
		Height: sizeSubSlicerHeight,
		Width: sizeSubSlicerWidth,
		Buffer: make([]byte, sizeSubSlicerHeight * sizeSubSlicerWidth, sizeSubSlicerHeight * sizeSubSlicerWidth),
		Filled: make([]bool, sizeSubSlicerHeight * sizeSubSlicerWidth, sizeSubSlicerHeight * sizeSubSlicerWidth),
	}

	// copy cross slices
	slices := make([]Slice, 0)
	for x := startX; x < stopX; x++ {
		for y := startY; y < stopY; y++ {
			slice := s.findSlice(int64(x), int64(y))
			if slice != nil {
				if slice.X >= int64(startX) &&
					slice.Y >= int64(startY) &&
					slice.X + int64(slice.Shape.Width) <= int64(stopX) &&
					slice.Y + int64(slice.Shape.Height) <= int64(stopY) {
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
		offset := s.getOffset(int64(startX), int64(startY) + int64(i))
		for j := 0; j < sizeSubSlicerWidth; j++ {
			subSlicer.Filled[i * sizeSubSlicerWidth + j] = s.filled[offset + int64(j)]
			subSlicer.Buffer[i * sizeSubSlicerWidth + j] = s.stream[offset + int64(j)]
		}
	}

	subSlicer.Slices = make([]subslice.Slice, 0, len(slices))

	for _, slice := range slices {
		nSlice := subslice.Slice{
			X: slice.X - startX,
			Y: slice.Y - startY,
			Shape: slice.Shape,
		}
		subSlicer.Slices = append(subSlicer.Slices, nSlice)
	}

	return &subSlicer
}
