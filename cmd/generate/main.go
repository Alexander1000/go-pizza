package main

import (
	"flag"
	"bufio"
	"os"
	"fmt"
	"regexp"
	"strings"
	"strconv"

	"go-pizza/generator"
)

func main() {
	size := flag.String("size", "", "size of pizza")
	minCount := flag.Int("min", 1, "min count each ingredient")
	maxCount := flag.Int("max", 6, "max number of ceil")
	file := flag.String("file", "", "path to file")
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

	fmt.Printf("Size: %dx%d\r\n", height, width)

	fmt.Println("Conditions:")
	fmt.Printf("Min count each ingredient: %d\r\n", *minCount)
	fmt.Printf("Max size: %d\r\n", *maxCount)

	for *file == "" {
		fmt.Println("Input output file for:")
		s := bufio.NewScanner(os.Stdin)
		for s.Scan() {
			line := s.Text()
			if len(line) > 0 {
				file = &line
				break
			}
		}
	}

	fmt.Printf("File: %s\r\n", *file)

	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
	generator.GenerateMap(out, int(height), int(width))
}
