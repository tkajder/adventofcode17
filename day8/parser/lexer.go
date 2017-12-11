package parser

import (
	"bufio"
	"io"
	"strings"
)

type Scanner struct {
	r *bufio.Reader
}

func NewScanner(r io.Reader) *Scanner {
	return &Scanner{r: bufio.NewReader(r)}
}

func (s *Scanner) Scan() (token Token, text string) {
	r := s.read()

	if isWhitespace(r) {
		s.unread()
		return WHITESPACE, s.scanContiguous(isWhitespace)
	} else if isNewline(r) {
		s.unread()
		return NEWLINE, s.scanContiguous(isNewline)
	} else if isNumeric(r) {
		s.unread()
		return NUMBER, s.scanContiguous(isNumeric)
	} else if isAlpha(r) {
		s.unread()
		return s.scanIdentifier()
	} else if isComparator(r) {
		s.unread()
		return s.scanComparator()
	}

	switch r {
	case rune(0):
		return EOF, ""
	default:
		return ILLEGAL, ""
	}
}

func (s *Scanner) scanContiguous(f func(rune) bool) string {
	runes := make([]rune, 0)

	for r := s.read(); f(r); r = s.read() {
		runes = append(runes, r)
	}
	s.unread()

	return string(runes)
}

func (s *Scanner) scanIdentifier() (Token, string) {
	text := s.scanContiguous(isAlpha)

	switch strings.ToUpper(text) {
	case "IF":
		return IF, text
	case "INC":
		return INC, text
	case "DEC":
		return DEC, text
	default:
		return IDENTIFER, text
	}
}

func (s *Scanner) scanComparator() (Token, string) {
	text := s.scanContiguous(isComparator)

	switch text {
	case "==":
		return EQ, text
	case "!=":
		return NEQ, text
	case "<=":
		return LTE, text
	case ">=":
		return GTE, text
	case "<":
		return LT, text
	case ">":
		return GT, text
	default:
		return ILLEGAL, text
	}
}

func (s *Scanner) read() rune {
	r, _, err := s.r.ReadRune()
	if err != nil {
		return rune(0)
	}

	return r
}

func (s *Scanner) unread() {
	s.r.UnreadRune()
}

func isWhitespace(r rune) bool {
	return r == ' ' || r == '\t'
}

func isNewline(r rune) bool {
	return r == '\r' || r == '\n'
}

func isEOF(r rune) bool {
	return r == 0
}

func isNumeric(r rune) bool {
	return r == '-' || r >= '0' && r <= '9'
}

func isAlpha(r rune) bool {
	return (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z')
}

func isComparator(r rune) bool {
	return r == '=' || r == '<' || r == '>' || r == '!'
}
