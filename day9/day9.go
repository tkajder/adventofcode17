package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/tkajder/adventofcode17/day9/parser"
)

func main() {
	fileNamePtr := flag.String("file", "", "Required: file containing inverse captcha")
	flag.Parse()

	// Check if file provided
	if *fileNamePtr == "" {
		flag.Usage()
		os.Exit(1)
	}

	r, err := os.Open(*fileNamePtr)
	if err != nil {
		log.Fatal(err)
	}

	groups := []parser.Group{}

	p := parser.NewParser(r)
	groupc, errc := p.Parse()
	for group := range groupc {
		groups = append(groups, group)
	}
	if err := <-errc; err != nil {
		log.Fatal(err)
	}

	var score uint
	for _, group := range groups {
		score += group.Score()
	}
	fmt.Println(score)

	var garbage uint
	for _, group := range groups {
		garbage += group.Garbage()
	}
	fmt.Println(garbage)
}
