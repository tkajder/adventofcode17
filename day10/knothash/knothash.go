package knothash

import "errors"

type KnotHash struct {
	data     []int
	currPos  int
	skipSize int
}

func New() *KnotHash {
	data := make([]int, 256)
	for i := 0; i < 256; i++ {
		data[i] = i
	}

	return &KnotHash{data: data, currPos: 0, skipSize: 0}
}

func (kh *KnotHash) Knot(length int) error {
	dataLen := len(kh.data)

	if length < 0 {
		return errors.New("Length of knot cannot be less than 0")
	}

	if length > dataLen {
		return errors.New("Length of knot cannot be greater than length of data")
	}

	// Find start and end of length
	lengthStart := kh.currPos
	lengthEnd := lengthStart + length

	// Reverse length with wrap around indexing
	for i := 0; i < length/2; i++ {
		// Calculate swap indices with wrapped length
		swapStart := (lengthStart + i) % dataLen
		swapEnd := (lengthEnd - i - 1) % dataLen

		// Perform swap
		temp := kh.data[swapStart]
		kh.data[swapStart] = kh.data[swapEnd]
		kh.data[swapEnd] = temp
	}

	// Move current position forward by length and skipSize with wrapping
	kh.currPos = (kh.currPos + length + kh.skipSize) % dataLen

	// Increment skipSize
	kh.skipSize++

	return nil
}

func (kh *KnotHash) Hash() int {
	return kh.data[0] * kh.data[1]
}
