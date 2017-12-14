package parser

type Token int

const (
	ILLEGAL Token = iota
	EOF
	COMMA
	GROUPSTART
	GROUPEND
	GARBAGE
)

func (t Token) String() string {
	switch t {
	case ILLEGAL:
		return "ILLEGAL"
	case EOF:
		return "EOF"
	case GROUPSTART:
		return "GROUPSTART"
	case GROUPEND:
		return "GROUPEND"
	case COMMA:
		return "COMMA"
	case GARBAGE:
		return "GARBAGE"
	default:
		panic("Developer error; invalid token type")
	}
}
