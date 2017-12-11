package hexagonal

import (
	"fmt"
	"strings"
)

type Direction int

const (
	N Direction = iota
	NE
	SE
	S
	SW
	NW
)

func ParseDirection(s string) (Direction, error) {
	switch strings.ToUpper(s) {
	case "N":
		return N, nil
	case "NE":
		return NE, nil
	case "SE":
		return SE, nil
	case "S":
		return S, nil
	case "SW":
		return SW, nil
	case "NW":
		return NW, nil
	default:
		return -1, fmt.Errorf("Could not parse direction from %s", s)
	}
}
