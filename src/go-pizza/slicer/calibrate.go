package slicer

import "fmt"

func (s *Slicer) calibrate() {
	blackPoints := make([]BlackPoint, 0)

	fmt.Printf("Black points\n")

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
				blackPoints = append(blackPoints, BlackPoint{X: j, Y: i})
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
}

type BlackPoint struct {
	X int64
	Y int64
}