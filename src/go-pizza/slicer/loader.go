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

type Loader struct {
	file *os.File
}

func NewLoader(file string) (*Loader, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	return &Loader{file: f}, nil
}

func (l *Loader) Load() (*Slicer, error) {
	buffer := make([]byte, ReadBufferSize, ReadBufferSize)
	n, err := l.file.Read(buffer)
	if err != nil {
		return nil, err
	}
	if n == 0 {
		return nil, errors.New("empty file")
	}

	offset := 0
	height, err := scanDigit(buffer)
	if err != nil {
		return nil, err
	}

	log.Printf("Height: %d", height)

	offset += len(strconv.Itoa(int(height))) + 1

	width, err := scanDigit(buffer[offset:])
	if err != nil {
		return nil, err
	}

	log.Printf("Width: %d", width)

	offset += len(strconv.Itoa(int(width))) + 1

	minIngredienCount, err := scanDigit(buffer[offset:])
	if err != nil {
		return nil, err
	}

	log.Printf("Min ingredient count: %d", minIngredienCount)

	offset += len(strconv.Itoa(int(minIngredienCount))) + 1

	maxSizeSlice, err := scanDigit(buffer[offset:])
	if err != nil {
		return nil, err
	}

	log.Printf("Max slice size: %d", maxSizeSlice)
	offset += len(strconv.Itoa(int(maxSizeSlice)))
	spaces := 0

	for _, b := range buffer[offset:] {
		if b == byte(0x0D) || b == byte(0x0A) {
			spaces++
			continue
		}
		break
	}
	offset += spaces

	slicer := Slicer{}

	slicer.stream = make([]byte, 0, width * height)

	return &slicer, nil
}

func (l *Loader) scanRowPizza(height int64, width int64, buffer []byte) {
	
}

func scanDigit(buffer []byte) (int64, error) {
	size := 0
	for _, b := range buffer {
		if b == byte(0x20) || b == byte(0x0D) {
			break
		}
		if b >= byte('0') && b <= byte('9') {
			size++
		} else {
			return 0, errors.New("invalid character")
		}
	}

	if size == 0 {
		return 0, errors.New("digits not found")
	}

	digit, err := strconv.ParseInt(string(buffer[:size]), 10, 64)
	if err != nil {
		return 0, err
	}
	if digit < 0 {
		return 0, errors.New("invalid digit value")
	}
	return digit, nil
}
