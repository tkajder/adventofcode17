package main

import (
	"testing"
)

func Test_computeDifference(t *testing.T) {
	type args struct {
		line string
	}
	tests := []struct {
		name    string
		args    args
		want    uint64
		wantErr bool
	}{
		{"test empty line", args{""}, 0, true},
		{"test single value line", args{"1"}, 0, false},
		{"test multi value line", args{"1 5 4"}, 4, false},
		{"test non-number value line", args{"1 & 3"}, 0, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := computeDifference(tt.args.line)
			if (err != nil) != tt.wantErr {
				t.Errorf("computeDifference() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("computeDifference() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_minMaxDifference(t *testing.T) {
	type args struct {
		nums []uint64
	}
	tests := []struct {
		name    string
		args    args
		want    uint64
		wantErr bool
	}{
		{"test empty", args{[]uint64{}}, 0, true},
		{"test single value", args{[]uint64{1}}, 0, false},
		{"test multi value", args{[]uint64{1, 4}}, 3, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := minMaxDifference(tt.args.nums)
			if (err != nil) != tt.wantErr {
				t.Errorf("minMaxDifference() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("minMaxDifference() = %v, want %v", got, tt.want)
			}
		})
	}
}
