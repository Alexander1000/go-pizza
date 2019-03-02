package main

import (
	"flag"
	"log"

	"go-pizza/slicer"
)

func main() {
	file := flag.String("file", "", "path to file")
	flag.Parse()
	if *file == "" {
		log.Fatal("File required")
	}

	loader, err := slicer.NewLoader(*file)
	if err != nil {
		log.Fatalf("loader error: %v", err)
	}
	slicer, err := loader.Load()
	if err != nil {
		log.Fatalf("parse error: %v", err)
	}
	if slicer == nil {
		log.Fatal("empty slicer given")
	}

	slicer.Scan()
}
