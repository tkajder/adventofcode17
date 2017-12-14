package parser

import "testing"

func TestToken_String(t *testing.T) {
	tests := []struct {
		name string
		t    Token
		want string
	}{
		{"test illegal", ILLEGAL, "ILLEGAL"},
		{"test eof", EOF, "EOF"},
		{"test groupstart", GROUPSTART, "GROUPSTART"},
		{"test groupend", GROUPEND, "GROUPEND"},
		{"test garbage", GARBAGE, "GARBAGE"},
		{"test comma", COMMA, "COMMA"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.t.String(); got != tt.want {
				t.Errorf("Token.String() = %v, want %v", got, tt.want)
			}
		})
	}
}
