package main

import (
	"flag"
	"bufio"
	"os"
	"fmt"
	"regexp"
)

func main() {
	size := flag.String("s", "", "usage size")
	flag.Parse()

	sizeRegexp := regexp.MustCompile(`^\d+\s?[xX]\s?\d+$`)

	for *size == "" {
		fmt.Println("Input size of pizza [height x weight]:")
		s := bufio.NewScanner(os.Stdin)
		for s.Scan() {
			line := s.Text()
			if len(line) > 0 {
				found := sizeRegexp.FindString(line)
				if found != "" {
					size = &found
					break
				}
			}
		}
	}

	fmt.Printf("size: %s", *size)
}
