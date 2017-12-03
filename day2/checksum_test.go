package main

import (
	"testing"
)

func Test_largestDifference(t *testing.T) {
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
			got, err := largestDifference(tt.args.nums)
			if (err != nil) != tt.wantErr {
				t.Errorf("largestDifference() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("largestDifference() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_divisibleNumsSum(t *testing.T) {
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
		{"test single value", args{[]uint64{1}}, 0, true},
		{"test multi value with divisors", args{[]uint64{2, 3, 4, 5}}, 2, false},
		{"test multi value no divisors", args{[]uint64{3, 5, 7, 11}}, 0, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := divisibleNumsSum(tt.args.nums)
			if (err != nil) != tt.wantErr {
				t.Errorf("divisibleNumsSum() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("divisibleNumsSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
