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

type ValuedPoint struct {
	point Point
	value uint64
}

func (p *Point) isAdjacent(other *Point) (bool, error) {
	if other == nil {
		return false, errors.New("Cannot be adjacent to nil Point")
	}

	return (p.x-other.x >= -1 && p.x-other.x <= 1) && (p.y-other.y >= -1 && p.y-other.y <= 1), nil
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

func firstGreaterCoordinate(value uint64) (uint64, error) {
	if value == 0 {
		return 1, nil
	}

	points := []ValuedPoint{ValuedPoint{Point{0, 0}, 1}}

	var coordinate uint64
	for coordinate = 2; ; coordinate++ {
		point, err := pointFromSpiralCoordinate(coordinate)
		if err != nil {
			return 0, err
		}

		var pointValue uint64
		for _, prevPoint := range points {
			adjacent, err := point.isAdjacent(&prevPoint.point)
			if err != nil {
				return 0, err
			}
			if adjacent {
				pointValue += prevPoint.value
			}
		}

		if pointValue > value {
			return pointValue, nil
		}

		points = append(points, ValuedPoint{point, pointValue})
	}
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

	value, err := firstGreaterCoordinate(uint64(coordinate))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(value)
}
