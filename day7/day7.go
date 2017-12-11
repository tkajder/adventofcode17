package main

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"sync"

	"github.com/tkajder/adventofcode17/day7/disctree"
	"github.com/tkajder/adventofcode17/fileutils"
	"github.com/tkajder/adventofcode17/list"
)

type Edge struct {
	parent *disctree.DiscTree
	child  string
}

func parseProgramInfo(info string) (string, uint, []string, error) {
	sections := strings.Split(info, " ")

	if len(sections) < 2 {
		return "", 0, nil, errors.New("Info does not contain name and weight")
	}

	name := sections[0]
	weight, err := strconv.Atoi(strings.Trim(sections[1], "()"))
	if err != nil {
		return name, 0, nil, err
	}

	var children []string
	if len(sections) > 3 {
		children = sections[3:]
	}

	for i := 0; i < len(children); i++ {
		children[i] = strings.Trim(children[i], ",")
	}
	return name, uint(weight), children, nil
}

func main() {
	fileNamePtr := flag.String("file", "", "Required: file containing inverse captcha")
	flag.Parse()

	// Check if file provided
	if *fileNamePtr == "" {
		flag.Usage()
		os.Exit(1)
	}

	unfinishedTrees := list.New()
	unaddedEdges := list.New()

	var wg sync.WaitGroup
	lines, errc := fileutils.ByLine(*fileNamePtr)
	for line := range lines {
		name, weight, children, err := parseProgramInfo(line)
		if err != nil {
			log.Fatal(err)
		}

		wg.Add(1)
		go func() {
			defer wg.Done()
			tree := disctree.New(name, weight)
			unfinishedTrees.Push(tree)
			for _, child := range children {
				unaddedEdges.Push(&Edge{tree, child})
			}
		}()
	}
	if err := <-errc; err != nil {
		log.Fatal(err)
	}

	wg.Wait()
	for edge := unaddedEdges.Pop(); edge != nil; edge = unaddedEdges.Pop() {
		child := unfinishedTrees.Find(func(value interface{}) bool {
			tree := value.(*disctree.DiscTree).Find(edge.(*Edge).child)
			return tree != nil
		})
		if !unfinishedTrees.Remove(child) {
			log.Panic("Invalid input")
		}

		edge.(*Edge).parent.Insert(child.(*disctree.DiscTree))
	}

	tree := unfinishedTrees.Pop().(*disctree.DiscTree)
	fmt.Println(tree.Name)
}
