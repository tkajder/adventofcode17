package parser

type Token uint

const (
	// Special
	ILLEGAL Token = iota
	EOF
	NEWLINE
	WHITESPACE

	// Literals
	IDENTIFER
	NUMBER

	// Reserved
	IF

	// Directives
	INC
	DEC

	// Comparators
	GT
	LT
	GTE
	LTE
	EQ
	NEQ
)

func isComparatorToken(token Token) bool {
	return token >= GT && token <= NEQ
}

func isDirectiveToken(token Token) bool {
	return token == INC || token == DEC
}

func (t *Token) String() string {
	switch *t {
	case ILLEGAL:
		return "ILLEGAL"
	case EOF:
		return "EOF"
	case NEWLINE:
		return "NEWLINE"
	case WHITESPACE:
		return "WHITESPACE"
	case IDENTIFER:
		return "IDENTIFIER"
	case NUMBER:
		return "NUMBER"
	case IF:
		return "IF"
	case INC:
		return "INC"
	case DEC:
		return "DEC"
	case GT:
		return "GT"
	case LT:
		return "LT"
	case GTE:
		return "GTE"
	case LTE:
		return "LTE"
	case EQ:
		return "EQ"
	case NEQ:
		return "NEQ"
	default:
		return ""
	}
}
