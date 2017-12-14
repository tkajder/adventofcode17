package parser

type Group struct {
	garbage  uint
	children []Group
}

func (g *Group) Score() uint {
	return g.recursiveScore(0)
}

func (g *Group) recursiveScore(depth uint) uint {
	score := depth + 1
	for _, child := range g.children {
		score += child.recursiveScore(depth + 1)
	}

	return score
}

func (g *Group) Garbage() uint {
	trash := g.garbage

	for _, child := range g.children {
		trash += child.Garbage()
	}

	return trash
}
