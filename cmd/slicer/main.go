package main

import (
	"flag"
	"log"
)

func main() {
	file := flag.String("file", "", "path to file")
	flag.Parse()
	if *file == "" {
		log.Fatal("File required")
	}
}
