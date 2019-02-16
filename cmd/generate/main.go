package main

import (
	"flag"
	"bufio"
	"os"
	"fmt"
)

func main() {
	size := flag.String("s", "size", "usage size")
	flag.Parse()

	if size == nil {
		for size == nil {
			fmt.Println("Input size of pizza:")
			s := bufio.NewScanner(os.Stdin)
			for s.Scan() {
				line := s.Text()
				if len(line) > 0 {
					size = &line
				}
			}
		}
	}
}
