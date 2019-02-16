package main

import (
	"flag"
	"bufio"
	"os"
	"fmt"
	"regexp"
	"strings"
	"strconv"

	"github.com/Alexander1000/go-pizza/go-pizza/generator"
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

	result := strings.Split(*size, `x`)
	if len(result) == 1 {
		result = strings.Split(*size, `X`)
	}

	height, _ := strconv.ParseInt(result[0], 10, 32)
	width, _ := strconv.ParseInt(result[1], 10, 32)

	fmt.Printf("size: %dx%d", height, width)

	out := bufio.NewWriter(os.Stdout)
	generator.GenerateMap(out, int(height), int(width))
}
