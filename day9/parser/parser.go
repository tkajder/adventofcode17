package parser

import (
	"fmt"
	"io"
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

func (p *Parser) Parse() (<-chan Group, <-chan error) {
	groups := make(chan Group)
	errc := make(chan error, 1)

	go func() {
		defer close(groups)
		defer close(errc)

		group, err := p.parseGroup(0)
		if err != nil {
			errc <- err
			return
		}
		groups <- group

		token, text := p.scan()
		if token != EOF {
			errc <- fmt.Errorf("Expected EOF; got %s with value %s", token.String(), text)
		}
	}()

	return groups, errc
}

func (p *Parser) parseGroup(i int) (Group, error) {
	token, text := p.scan()
	if token != GROUPSTART {
		return Group{}, fmt.Errorf("Expected group start; encountered token %s with text %s", token.String(), text)
	}

	group := Group{}
	for {
		token, text = p.scan()
		switch token {
		case GROUPEND:
			return group, nil
		case GROUPSTART:
			p.unscan()
			child, err := p.parseGroup(i + 1)
			if err != nil {
				return Group{}, err
			}

			group.children = append(group.children, child)
		case GARBAGE:
			group.garbage += uint(len(text)) - 2
		case ILLEGAL:
			return Group{}, fmt.Errorf("Expected legal token; encountered token %s with text %s", token.String(), text)
		case EOF:
			return Group{}, fmt.Errorf("Expected group end; encountered EOF")
		default:
			// ignore
		}
	}
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
