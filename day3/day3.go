package main

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"

	"github.com/tkajder/adventofcode17/fileutils"
)

type Point struct {
	x int64
	y int64
}

func pointFromSpiralCoordinate(coordinate uint64) (Point, error) {
	if coordinate == 0 {
		return Point{}, errors.New("Spiral coordinate 0 is invalid, coordinates are 0-based")
	}

	roundedSqrt := int64(math.Sqrt(float64(coordinate)) + 0.5)

	closestSquared := roundedSqrt * roundedSqrt

	var x int64
	var y int64
	if closestSquared%2 == 0 {
		x = -int64(math.Floor(float64(roundedSqrt) / 2))
		y = int64(math.Floor(float64(roundedSqrt) / 2))

		distanceFromCorner := int64(coordinate) - (closestSquared + 1)
		if distanceFromCorner > 0 {
			y -= distanceFromCorner
		} else if distanceFromCorner < 0 {
			x -= distanceFromCorner
		}
	} else {
		x = int64(math.Floor(float64(roundedSqrt) / 2))
		y = -int64(math.Floor(float64(roundedSqrt) / 2))

		distanceFromCorner := int64(coordinate) - closestSquared
		if distanceFromCorner > 0 {
			x++
			y += distanceFromCorner - 1
		} else if distanceFromCorner < 0 {
			x += distanceFromCorner
		}
	}

	return Point{x: int64(x), y: y}, nil
}

func manhattanDistance(p1 Point, p2 Point) uint64 {
	distanceX := uint64(math.Abs(float64(p1.x) - float64(p2.x)))
	distanceY := uint64(math.Abs(float64(p1.y) - float64(p2.y)))

	return distanceX + distanceY
}

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

	coordinate, err := strconv.Atoi(fileContents)
	if err != nil {
		log.Fatal(err)
	}

	// Will overflow negative - don't care for this simple program
	point, err := pointFromSpiralCoordinate(uint64(coordinate))
	if err != nil {
		log.Fatal(err)
	}

	distance := manhattanDistance(point, Point{0, 0})
	fmt.Println(distance)
}
