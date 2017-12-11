package hexagonal

import (
	"errors"
	"log"
	"math"
)

type Coordinate struct {
	x int
	y int
	z int
}

func NewCoordinate(x int, y int, z int) (*Coordinate, error) {
	if x+y+z != 0 {
		return nil, errors.New("X, Y, Z must add up to 0 in hexagonal coordinates")
	}

	return &Coordinate{x, y, z}, nil
}

func (coordinate *Coordinate) Move(direction Direction) {
	switch direction {
	case N:
		coordinate.y++
		coordinate.z--
	case NE:
		coordinate.x++
		coordinate.z--
	case SE:
		coordinate.x++
		coordinate.y--
	case S:
		coordinate.y--
		coordinate.z++
	case SW:
		coordinate.x--
		coordinate.z++
	case NW:
		coordinate.x--
		coordinate.y++
	default:
		log.Fatal("Developer error: invalid direction")
	}
}

func (coordinate *Coordinate) Distance(other *Coordinate) int {
	xDistance := int(math.Abs(float64(coordinate.x - other.x)))
	yDistance := int(math.Abs(float64(coordinate.y - other.y)))
	zDistance := int(math.Abs(float64(coordinate.z - other.z)))

	return max(xDistance, yDistance, zDistance)
}

func max(nums ...int) int {
	var max int

	for _, num := range nums {
		if num > max {
			max = num
		}
	}

	return max
}
