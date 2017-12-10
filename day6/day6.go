package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/tkajder/adventofcode17/fileutils"
)

func balance(unbalanced []uint) []uint {
	balanced := make([]uint, len(unbalanced))
	copy(balanced, unbalanced)

	if len(balanced) == 0 {
		return balanced
	}

	var index int
	var max uint
	for i, val := range balanced {
		if val > max {
			index = i
			max = val
		}
	}

	balanced[index] = 0
	index = (index + 1) % len(balanced)
	for redistribute := max; redistribute > 0; redistribute-- {
		balanced[index]++
		index = (index + 1) % len(balanced)
	}

	return balanced
}

func arrEquals(a []uint, b []uint) bool {
	if len(a) != len(b) {
		return false
	}

	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}

func main() {
	fileNamePtr := flag.String("file", "", "Required: file containing inverse captcha")
	flag.Parse()

	// Check if file provided
	if *fileNamePtr == "" {
		flag.Usage()
		os.Exit(1)
	}

	fulltext, err := fileutils.WholeFile(*fileNamePtr)
	if err != nil {
		log.Fatal(err)
	}

	strs := strings.Split(strings.Trim(fulltext, "\r\n"), "\t")
	memoryBanks := make([]uint, len(strs))
	for i := 0; i < len(strs); i++ {
		val, err := strconv.Atoi(strs[i])
		if err != nil {
			log.Fatal(err)
		}

		memoryBanks[i] = uint(val)
	}

	allEqual := false
	cycles := [][]uint{memoryBanks}
	for !allEqual {
		next := balance(cycles[len(cycles)-1])

		for _, prev := range cycles {
			if arrEquals(prev, next) {
				allEqual = true
				break
			}
		}

		cycles = append(cycles, next)
	}
	fmt.Println(len(cycles) - 1)
}
