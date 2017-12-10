package main

import "testing"

func Test_numJumpsToExit(t *testing.T) {
	type args struct {
		jumps []int
	}
	tests := []struct {
		name string
		args args
		want uint
	}{
		{"test example case", args{[]int{0, 3, 0, 1, -3}}, 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numJumpsToExit(tt.args.jumps); got != tt.want {
				t.Errorf("numJumpsToExit() = %v, want %v", got, tt.want)
			}
		})
	}
}
