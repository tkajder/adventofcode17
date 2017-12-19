package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/tkajder/adventofcode17/fileutils"
	"github.com/twmb/algoimpl/go/graph"
)

func parseLine(line string) (srcID int, dstIDs []int, err error) {
	edgeDef := strings.Split(line, "<->")
	if len(edgeDef) != 2 {
		log.Fatal("Invalid line: " + line)
	}

	srcID, err = strconv.Atoi(strings.Trim(edgeDef[0], " "))
	if err != nil {
		return 0, nil, err
	}

	for _, dstIDStr := range strings.Split(edgeDef[1], ",") {
		dstID, err := strconv.Atoi(strings.Trim(dstIDStr, " "))
		if err != nil {
			return 0, nil, err
		}

		dstIDs = append(dstIDs, dstID)
	}

	return
}

func main() {
	fileNamePtr := flag.String("file", "", "Required: file containing inverse captcha")
	flag.Parse()

	// Check if file provided
	if *fileNamePtr == "" {
		flag.Usage()
		os.Exit(1)
	}

	// Setup graph and int->node map
	g := graph.New(graph.Undirected)
	nodes := make(map[int]graph.Node)

	// Read file into graph
	lines, errc := fileutils.ByLine(*fileNamePtr)
	for line := range lines {
		srcID, dstIDs, err := parseLine(line)
		if err != nil {
			log.Fatal(err)
		}

		src, ok := nodes[srcID]
		if !ok {
			src = g.MakeNode()
			nodes[srcID] = src
		}

		for _, dstID := range dstIDs {
			dst, ok := nodes[dstID]
			if !ok {
				dst = g.MakeNode()
				nodes[dstID] = dst
			}

			g.MakeEdge(src, dst)
		}
	}
	if err := <-errc; err != nil {
		log.Fatal(err)
	}

	// Invert the map so we can lookup by node instead of id
	names := make(map[graph.Node]int)
	for k, v := range nodes {
		names[v] = k
	}

	// Get the connected components of the graph
	components := g.StronglyConnectedComponents()
	for _, component := range components {

		// Print component size if it contains 0
		for _, node := range component {
			if names[node] == 0 {
				fmt.Println(len(component))
				break
			}
		}
	}

	fmt.Println(len(components))
}
