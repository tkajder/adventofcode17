package main

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"math"
	"os"
	"strings"

	"github.com/tkajder/adventofcode17/fileutils"
	"github.com/tkajder/adventofcode17/sliceutils"
)

func largestDifference(nums []uint64) (uint64, error) {
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

func divisibleNumsSum(nums []uint64) (uint64, error) {
	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums); j++ {
			if nums[i] < nums[j] && nums[j]%nums[i] == 0 {
				return nums[j] / nums[i], nil
			}
		}
	}

	return 0, errors.New("No divisible numbers found")
}

func main() {
	fileNamePtr := flag.String("file", "", "Required: file containing inverse captcha")
	flag.Parse()

	// Check if file provided
	if *fileNamePtr == "" {
		flag.Usage()
		os.Exit(1)
	}

	lines, errc := fileutils.ByLine(*fileNamePtr)

	var (
		diffChecksum uint64
		divChecksum  uint64
	)

	for line := range lines {
		nums, err := sliceutils.Atoui64(strings.Split(line, "\t"))

		diff, err := largestDifference(nums)
		if err != nil {
			log.Fatal(err)
		}
		diffChecksum += diff

		div, err := divisibleNumsSum(nums)
		if err != nil {
			log.Fatal(err)
		}
		divChecksum += div
	}
	if err := <-errc; err != nil {
		log.Fatal(err)
	}

	fmt.Println(diffChecksum, divChecksum)
}
