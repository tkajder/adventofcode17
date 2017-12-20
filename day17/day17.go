package main

import (
	"container/ring"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"

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

	fulltext, err := fileutils.WholeFile(*fileNamePtr)
	if err != nil {
		log.Fatal(err)
	}

	steps, err := strconv.Atoi(fulltext)
	if err != nil {
		log.Fatal(err)
	}

	// Set up ring
	r := ring.New(1)
	r.Value = 0

	for round := 1; round <= 2017; round++ {
		// Spin baby spin
		for step := 0; step < steps; step++ {
			r = r.Next()
		}

		// Setup the new value
		next := ring.New(1)
		next.Value = round

		// Insert the value and center on it
		r.Link(next)
		r = next
	}

	fmt.Println(r.Next().Value)
}
