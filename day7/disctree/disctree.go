package disctree

type DiscTree struct {
	Name     string
	Weight   uint
	Children []*DiscTree
}

func New(name string, weight uint) *DiscTree {
	return &DiscTree{Name: name, Weight: weight, Children: []*DiscTree{}}
}

func (tree *DiscTree) Walk(ch chan *DiscTree) {
	if tree == nil {
		return
	}

	ch <- tree
	for _, child := range tree.Children {
		child.Walk(ch)
	}
}

func (tree *DiscTree) Walker() chan *DiscTree {
	ch := make(chan *DiscTree)
	go func() {
		defer close(ch)
		tree.Walk(ch)
	}()
	return ch
}

func (tree *DiscTree) Insert(child *DiscTree) {
	if child == nil {
		return
	}

	tree.Children = append(tree.Children, child)
}

func (tree *DiscTree) Find(name string) *DiscTree {
	ch := tree.Walker()

	for subtree := range ch {
		if subtree.Name == name {
			return subtree
		}
	}

	return nil
}

func (tree *DiscTree) SumDirectChildrenWeights() uint {
	var sum uint

	for _, child := range tree.Children {
		sum += child.Weight
	}

	return sum
}
