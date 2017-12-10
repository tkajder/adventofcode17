package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/tkajder/adventofcode17/fileutils"
)

func numJumpsToExit(jumps []int) uint {
	var position int
	var numJumps uint
	jumpSize := len(jumps)

	for ; position < jumpSize && position >= 0; numJumps++ {
		jumps[position]++
		position += jumps[position] - 1
	}

	return numJumps
}

func main() {
	fileNamePtr := flag.String("file", "", "Required: file containing inverse captcha")
	flag.Parse()

	// Check if file provided
	if *fileNamePtr == "" {
		flag.Usage()
		os.Exit(1)
	}

	jumps := make([]int, 0)

	lines, errc := fileutils.ByLine(*fileNamePtr)
	for line := range lines {
		i, err := strconv.Atoi(line)
		if err != nil {
			log.Fatal(err)
		}

		jumps = append(jumps, i)
	}
	if err := <-errc; err != nil {
		log.Fatal(err)
	}

	numJumps := numJumpsToExit(jumps)
	fmt.Println(numJumps)
}
