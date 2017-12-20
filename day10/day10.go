package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/tkajder/adventofcode17/day10/knothash"

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

	kh := knothash.New()

	// Knot the knothash with each length
	for _, lengthStr := range strings.Split(strings.Trim(fulltext, " \t\r\n"), ",") {
		length, err := strconv.Atoi(lengthStr)
		if err != nil {
			log.Fatal(err)
		}

		kh.Knot(length)
	}

	fmt.Println(kh.Hash())
}
