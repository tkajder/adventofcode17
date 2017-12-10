package main

import (
	"reflect"
	"testing"
)

func Test_balance(t *testing.T) {
	type args struct {
		mb []uint
	}
	tests := []struct {
		name string
		args args
		want []uint
	}{
		{"test empty", args{[]uint{}}, []uint{}},
		{"test example 1", args{[]uint{0, 2, 7, 0}}, []uint{2, 4, 1, 2}},
		{"test example 2", args{[]uint{2, 4, 1, 2}}, []uint{3, 1, 2, 3}},
		{"test example 3", args{[]uint{3, 1, 2, 3}}, []uint{0, 2, 3, 4}},
		{"test example 4", args{[]uint{0, 2, 3, 4}}, []uint{1, 3, 4, 1}},
		{"test example 4", args{[]uint{1, 3, 4, 1}}, []uint{2, 4, 1, 2}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := balance(tt.args.mb); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("balance() = %v, want %v", got, tt.want)
			}
		})
	}
}
