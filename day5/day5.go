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

func numJumpsToExitAlternative(jumps []int) uint {
	var position int
	var numJumps uint
	jumpSize := len(jumps)

	for ; position < jumpSize && position >= 0; numJumps++ {
		jump := jumps[position]
		if jump >= 3 {
			jumps[position]--
		} else {
			jumps[position]++
		}

		position += jump
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

	// Copy so they can both manipulate arrays
	alternativeJumps := make([]int, len(jumps))
	copy(alternativeJumps, jumps)

	numJumps := numJumpsToExit(jumps)
	fmt.Println(numJumps)

	numAlternativeJumps := numJumpsToExitAlternative(alternativeJumps)
	fmt.Println(numAlternativeJumps)
}
