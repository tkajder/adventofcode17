package disk

import (
	"encoding/hex"
	"strconv"

	"github.com/tkajder/adventofcode17/day10/knothash"
)

type Disk struct {
	data [128][128]bool
}

func New(seed string) *Disk {
	d := Disk{}

	for row := 0; row < 128; row++ {
		rowSeed := seed + "-" + strconv.Itoa(row)

		// Assumes that knothash returns 128 bit
		hash := knothash.Hash([]byte(rowSeed))
		hashBytes, err := hex.DecodeString(hash)
		if err != nil {
			panic(err)
		}

		for i, hashByte := range hashBytes {
			for bit := 0; bit < 8; bit++ {
				mask := byte(1 << uint(7-bit))
				d.data[row][i*8+bit] = hashByte&mask > 0
			}
		}
	}

	return &d
}

func (d *Disk) Used() int {
	used := 0

	for _, row := range d.data {
		for _, bit := range row {
			if bit {
				used++
			}
		}
	}

	return used
}
