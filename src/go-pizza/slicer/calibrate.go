package slicer

import (
	"fmt"
	"go-pizza/subslice"
	"math"
)

func (s *Slicer) calibrate() {
	blackPointList := make([]BlackPoint, 0)

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
				blackPointList = append(blackPointList, BlackPoint{X: j, Y: i})
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

	for _, blackPoint := range blackPointList {
		// try optimize/ defragmentation
		startX := math.Max(float64(blackPoint.X) - float64(s.maxSlice) * 1.2, 0)
		startY := math.Max(float64(blackPoint.Y) - float64(s.maxSlice) * 1.2, 0)

		stopX := math.Min(float64(blackPoint.X) + float64(s.maxSlice) * 1.2, float64(s.width))
		stopY := math.Min(float64(blackPoint.Y) + float64(s.maxSlice) * 1.2, float64(s.height))

		sizeSubSlicerWidth := int(stopX - startX)
		sizeSubSlicerHeight := int(stopY - startY)

		subSlicer := subslice.SubSlicer{
			Height: sizeSubSlicerHeight,
			Width: sizeSubSlicerWidth,
			Buffer: make([]byte, sizeSubSlicerHeight * sizeSubSlicerWidth, sizeSubSlicerHeight * sizeSubSlicerWidth),
			Filled: make([]bool, sizeSubSlicerHeight * sizeSubSlicerWidth, sizeSubSlicerHeight * sizeSubSlicerWidth),
		}

		slices := make([]Slice, 0)
		for x := startX; x < stopX; x++ {
			for y := startY; y < stopY; y++ {
				slice := s.findSlice(int64(x), int64(y))
				if slice != nil {
					if slice.X >= int64(startX) &&
						slice.Y >= int64(startY) &&
						slice.X + int64(slice.Shape.Width) <= int64(stopX) &&
						slice.Y + int64(slice.Shape.Height) <= int64(stopY) {
						slices = append(slices, *slice)
					}
				}
			}
		}

		fmt.Printf("Count slices for subslice: %d\n", len(slices))

		for i := 0; i < sizeSubSlicerHeight; i++ {
			offset := s.getOffset(int64(startX) + int64(i), int64(startY))
			for index, filled := range s.filled[offset : offset+int64(sizeSubSlicerWidth)] {
				subSlicer.Filled[i * sizeSubSlicerWidth + index] = filled
			}
			for index, data := range s.stream[offset : offset+int64(sizeSubSlicerWidth)] {
				subSlicer.Buffer[i * sizeSubSlicerWidth + index] = data
			}
		}
	}
}

type BlackPoint struct {
	X int64
	Y int64
}
