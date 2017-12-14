package parser

import (
	"bufio"
	"io"
)

const eof = rune(0)

type Scanner struct {
	r *bufio.Reader
}

func NewScanner(r io.Reader) *Scanner {
	return &Scanner{r: bufio.NewReader(r)}
}

func (s *Scanner) Scan() (Token, string) {
	r := s.read()

	switch r {
	case '\r', '\n', ' ', '\t':
		// Skip whitespace
		return s.Scan()
	case '!':
		s.skipNext()
		return s.Scan()
	case '<':
		s.unread()
		return s.readGarbage()
	case '{':
		return GROUPSTART, string(r)
	case '}':
		return GROUPEND, string(r)
	case ',':
		return COMMA, string(r)
	case eof:
		return EOF, ""
	default:
		return ILLEGAL, string(r)
	}
}

func (s *Scanner) skipNext() {
	r := s.read()

	if r == eof {
		s.unread()
	}
}

func (s *Scanner) readGarbage() (Token, string) {
	r := s.read()
	if r != '<' {
		s.unread()
		return s.Scan()
	}

	garbage := []rune{r}
	for r = s.read(); r != '>'; r = s.read() {
		switch r {
		case '!':
			s.skipNext()
		default:
			garbage = append(garbage, r)
		}
	}
	garbage = append(garbage, r)

	return GARBAGE, string(garbage)
}

func (s *Scanner) read() rune {
	r, _, err := s.r.ReadRune()
	if err != nil {
		return eof
	}
	return r
}

func (s *Scanner) unread() {
	s.r.UnreadRune()
}
