package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"log"
	"math"
	"os"
	"strings"

	"github.com/tkajder/adventofcode17/sliceutils"
)

func calculateFileChecksum(filename string) (uint64, error) {
	file, err := os.Open(filename)
	if err != nil {
		return 0, err
	}
	defer file.Close()

	var checksum uint64

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		difference, err := computeDifference(scanner.Text())
		if err != nil {
			return 0, err
		}
		checksum += difference
	}

	if err := scanner.Err(); err != nil {
		return 0, err
	}

	return checksum, nil
}

func computeDifference(line string) (uint64, error) {
	strNums := strings.Split(line, "\t")
	nums, err := sliceutils.Atoui64(strNums)
	if err != nil {
		return 0, err
	}

	difference, err := minMaxDifference(nums)
	if err != nil {
		return 0, err
	}

	return difference, nil
}

func minMaxDifference(nums []uint64) (uint64, error) {
	if len(nums) == 0 {
		return 0, errors.New("Cannot calculate difference on empty slice")
	}

	// "Zero" value is 0
	var max uint64
	var min uint64 = math.MaxUint64
	for _, num := range nums {
		if num < min {
			min = num
		}
		if num > max {
			max = num
		}
	}

	return max - min, nil
}

func main() {
	fileNamePtr := flag.String("file", "", "Required: file containing inverse captcha")
	flag.Parse()

	// Check if file provided
	if *fileNamePtr == "" {
		flag.Usage()
		os.Exit(1)
	}

	checksum, err := calculateFileChecksum(*fileNamePtr)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(checksum)
}
