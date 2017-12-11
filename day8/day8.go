package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/tkajder/adventofcode17/day8/parser"
	"github.com/tkajder/adventofcode17/day8/registers"
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

	rg := registers.New()
	var runningRegisterMax int

	p := parser.NewParser(r)
	instructions, errc := p.Parse()
	for instruction := range instructions {
		condition := instruction.Condition
		ok := condition.Comparison(rg.Get(condition.RegisterName))
		if ok {
			directive := instruction.Directive
			value := directive.Directive(rg.Get(directive.RegisterName))
			rg.Set(directive.RegisterName, value)

			if value > runningRegisterMax {
				runningRegisterMax = value
			}
		}

	}
	if err := <-errc; err != nil {
		log.Fatal(err)
	}

	fmt.Println(rg.MaxRegisterValue())
	fmt.Println(runningRegisterMax)
}
