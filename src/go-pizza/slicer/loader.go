package slicer

import (
	"os"
	"errors"
	"strconv"
)

const (
	ReadBufferSize = 1024
)

func Load(file string) (*Slicer, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}

	buffer := make([]byte, 0, ReadBufferSize)
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

	return nil, nil
}
