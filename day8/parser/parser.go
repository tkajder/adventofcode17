package parser

import (
	"errors"
	"fmt"
	"io"
	"strconv"
)

type Parser struct {
	s         *Scanner
	lastToken Token
	lastText  string
	hasLast   bool
}

func NewParser(r io.Reader) *Parser {
	return &Parser{s: NewScanner(r)}
}

func (p *Parser) Parse() (<-chan Instruction, <-chan error) {
	instructions := make(chan Instruction)
	errc := make(chan error)

	go func() {
		defer close(instructions)
		defer close(errc)

		for token, _ := p.scan(); token != EOF; token, _ = p.scan() {
			p.unscan()

			instruction, err := p.parseInstruction()
			if err != nil {
				errc <- err
				break
			}
			instructions <- instruction
		}
	}()

	return instructions, errc
}

func (p *Parser) parseInstruction() (Instruction, error) {
	directive, err := p.parseDirective()
	if err != nil {
		return Instruction{}, err
	}

	condition, err := p.parseCondition()
	if err != nil {
		return Instruction{}, err
	}

	if token, text := p.scan(); token != NEWLINE {
		// Accept no new line at EOF
		if token == EOF {
			p.unscan()
		} else {
			return Instruction{}, fmt.Errorf("Expected end of instruction; found token %s with value %s", token.String(), text)
		}
	}
	return Instruction{directive, condition}, nil
}

func (p *Parser) parseDirective() (Directive, error) {
	var (
		token         Token
		text          string
		register      string
		directiveType Token
		value         int
		directive     func(int) int
		err           error
	)

	// Parse register
	token, text = p.scanSkipToken(WHITESPACE)
	if token != IDENTIFER {
		return Directive{}, fmt.Errorf("Expected register name; found token %s with text %s", token.String(), text)
	}
	register = text

	// Parse directive type
	token, text = p.scanSkipToken(WHITESPACE)
	if !isDirectiveToken(token) {
		return Directive{}, fmt.Errorf("Expected directive; found token %s with text %s", token.String(), text)
	}
	directiveType = token

	// Parse value
	token, text = p.scanSkipToken(WHITESPACE)
	if token != NUMBER {
		return Directive{}, fmt.Errorf("Expected directive value; found token %s with text %s", token.String(), text)
	}
	value, err = strconv.Atoi(text)
	if err != nil {
		return Directive{}, fmt.Errorf("Expected directive value; found non-integer value %s", text)
	}

	// Resolve directive
	directive, err = getDirective(directiveType, value)
	if err != nil {
		return Directive{}, err
	}

	return Directive{register, directive}, nil
}

func (p *Parser) parseCondition() (Condition, error) {
	var (
		token          Token
		text           string
		register       string
		comparatorType Token
		value          int
		comparison     func(int) bool
		err            error
	)

	// Parse if
	token, text = p.scanSkipToken(WHITESPACE)
	if token != IF {
		return Condition{}, fmt.Errorf("Expected if; found token %s with text %s", token.String(), text)
	}

	// Parse register
	token, text = p.scanSkipToken(WHITESPACE)
	if token != IDENTIFER {
		return Condition{}, fmt.Errorf("Expected register name; found token %s with text %s", token.String(), text)
	}
	register = text

	// Parse comparator type
	token, text = p.scanSkipToken(WHITESPACE)
	if !isComparatorToken(token) {
		return Condition{}, fmt.Errorf("Expected comparator; found token %s with text %s", token.String(), text)
	}
	comparatorType = token

	// Parse comparison value
	token, text = p.scanSkipToken(WHITESPACE)
	if token != NUMBER {
		return Condition{}, fmt.Errorf("Expected comparison value; found token %s with text %s", token.String(), text)
	}
	value, err = strconv.Atoi(text)
	if err != nil {
		return Condition{}, fmt.Errorf("Expected comparison value; found non-integer value %s", text)
	}

	// Resolve comparison
	comparison, err = getComparison(comparatorType, value)
	if err != nil {
		return Condition{}, err
	}

	return Condition{register, comparison}, nil
}

func (p *Parser) scan() (Token, string) {
	if p.hasLast {
		p.hasLast = false
		return p.lastToken, p.lastText
	}

	p.lastToken, p.lastText = p.s.Scan()
	return p.lastToken, p.lastText
}

func (p *Parser) unscan() {
	p.hasLast = true
}

func (p *Parser) scanSkipToken(tokenType Token) (Token, string) {
	var (
		token Token
		text  string
	)

	for token, text = p.scan(); token == tokenType; token, text = p.scan() {
		// Skip tokens
	}

	return token, text
}

func getDirective(directiveType Token, value int) (func(int) int, error) {
	switch directiveType {
	case INC:
		return func(x int) int { return x + value }, nil
	case DEC:
		return func(x int) int { return x - value }, nil
	default:
		return nil, errors.New("Developer error; failed precondition of directive token types")
	}
}

func getComparison(comparatorType Token, value int) (func(int) bool, error) {
	switch comparatorType {
	case LT:
		return func(x int) bool { return x < value }, nil
	case GT:
		return func(x int) bool { return x > value }, nil
	case LTE:
		return func(x int) bool { return x <= value }, nil
	case GTE:
		return func(x int) bool { return x >= value }, nil
	case EQ:
		return func(x int) bool { return x == value }, nil
	case NEQ:
		return func(x int) bool { return x != value }, nil
	default:
		return nil, errors.New("Developer error; failed precondition of comparator token types")
	}
}
