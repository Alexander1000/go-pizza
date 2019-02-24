package slicer

import (
	"os"
	"errors"
	"strconv"
	"log"
)

const (
	ReadBufferSize = 1024
)

func Load(file string) (*Slicer, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}

	buffer := make([]byte, ReadBufferSize, ReadBufferSize)
	n, err := f.Read(buffer)
	if err != nil {
		return nil, err
	}
	if n == 0 {
		return nil, errors.New("empty file")
	}

	heightSize := 0
	for _, b := range buffer {
		if b == byte(0x20) {
			break
		}
		if b >= byte('0') && b <= byte('9') {
			heightSize++
		} else {
			err = errors.New("invalid character")
			break
		}
	}

	if heightSize == 0 {
		return nil, errors.New("empty height")
	}
	if err != nil {
		return nil, err
	}

	height, err := strconv.ParseInt(string(buffer[0:heightSize]), 10, 64)
	if err != nil {
		return nil, err
	}
	if height < 0 {
		return nil, errors.New("invalid height size")
	}

	log.Printf("Height: %d", height)

	widthSize := 0
	for _, b := range buffer[heightSize+1:] {
		if b == byte(0x20) {
			break
		}
		if b >= byte('0') && b <= byte('9') {
			widthSize++
		} else {
			err = errors.New("invalid character")
			break
		}
	}

	if widthSize == 0 {
		return nil, errors.New("empty width")
	}
	if err != nil {
		return nil, err
	}

	width, err := strconv.ParseInt(string(buffer[heightSize+1:heightSize + 1 + widthSize]), 10, 64)
	if err != nil {
		return nil, err
	}
	if width < 0 {
		return nil, errors.New("invalid width size")
	}

	log.Printf("Width: %d", width)

	return nil, nil
}
