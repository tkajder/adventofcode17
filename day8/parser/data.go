package parser

type Instruction struct {
	Directive Directive
	Condition Condition
}

type Directive struct {
	RegisterName string
	Directive    func(int) int
}

type Condition struct {
	RegisterName string
	Comparison   func(int) bool
}
