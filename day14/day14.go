package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/tkajder/adventofcode17/day14/disk"
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

	d := disk.New(fulltext)
	fmt.Println(d.Used())
}
