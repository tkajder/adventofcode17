package main

import (
	"testing"
)

func Test_verifyPassphraseNoDuplicates(t *testing.T) {
	type args struct {
		passphrase []string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"test empty", args{[]string{}}, true},
		{"test no duplicates", args{[]string{"aaa", "bbb", "ccc", "ddd"}}, true},
		{"test duplicates", args{[]string{"aaa", "aaa", "bbb", "ccc"}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := verifyPassphraseNoDuplicates(tt.args.passphrase); got != tt.want {
				t.Errorf("verifyPassphraseNoDuplicates() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_verifyPassphraseNoAnagrams(t *testing.T) {
	type args struct {
		passphrase []string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"test empty", args{[]string{}}, true},
		{"test no anagrams", args{[]string{"aaa", "bbb", "ccc", "ddd"}}, true},
		{"test duplicates", args{[]string{"aaa", "aaa", "bbb", "ccc"}}, false},
		{"test anagrams", args{[]string{"aba", "aab", "bbb", "ccc"}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := verifyPassphraseNoAnagrams(tt.args.passphrase); got != tt.want {
				t.Errorf("verifyPassphraseNoAnagrams() = %v, want %v", got, tt.want)
			}
		})
	}
}
