package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/tkajder/adventofcode17/day11/hexagonal"
	"github.com/tkajder/adventofcode17/fileutils"
)

func main() {
	fileNamePtr := flag.String("file", "", "Required: file containing inverse captcha")
	flag.Parse()

	// Check if file provided
	if *fileNamePtr == "" {
		flag.Usage()
		os.Exit(1)
	}

	fileContents, err := fileutils.WholeFile(*fileNamePtr)
	if err != nil {
		log.Fatal(err)
	}

	origin, err := hexagonal.NewCoordinate(0, 0, 0)
	if err != nil {
		log.Fatal(err)
	}
	coordinate, err := hexagonal.NewCoordinate(0, 0, 0)
	if err != nil {
		log.Fatal(err)
	}
	maxDistance := 0

	for _, s := range strings.Split(strings.Trim(fileContents, "\r\n"), ",") {
		direction, err := hexagonal.ParseDirection(s)
		if err != nil {
			log.Fatal(err)
		}
		coordinate.Move(direction)

		// Hold maximum distance for part 2
		distance := coordinate.Distance(origin)
		if distance > maxDistance {
			maxDistance = distance
		}
	}

	fmt.Println(coordinate.Distance(origin))
	fmt.Println(maxDistance)
}
