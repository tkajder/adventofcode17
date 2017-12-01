/*
 * Copyright 2017 Thomas Kajder
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
 */

package main

import (
	"container/ring"
	"reflect"
	"testing"
)

func generateRing(nums ...int) *ring.Ring {
	r := ring.New(len(nums))

	for _, x := range nums {
		r.Value = x
		r = r.Next()
	}

	return r
}

func Test_convertToSingleRuneIntRing(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name    string
		args    args
		want    *ring.Ring
		wantErr bool
	}{
		{"test empty string", args{""}, generateRing(), false},
		{"test invalid character string", args{"13a"}, nil, true},
		{"test valid string", args{"14236"}, generateRing(1, 4, 2, 3, 6), false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := convertToSingleRuneIntRing(tt.args.s)
			if (err != nil) != tt.wantErr {
				t.Errorf("convertToSingleRuneIntRing() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("convertToSingleRuneIntRing() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_calculateCaptcha(t *testing.T) {
	type args struct {
		r *ring.Ring
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test case 1122", args{generateRing(1, 1, 2, 2)}, 3},
		{"test case 1111", args{generateRing(1, 1, 1, 1)}, 4},
		{"test case 1234", args{generateRing(1, 2, 3, 4)}, 0},
		{"test case 91212129", args{generateRing(9, 1, 2, 1, 2, 1, 2, 9)}, 9},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := calculateCaptcha(tt.args.r); got != tt.want {
				t.Errorf("calculateCaptcha() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_matchesNext(t *testing.T) {
	type args struct {
		r *ring.Ring
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"test single element ring", args{generateRing(1)}, true},
		{"test matching element ring", args{generateRing(5, 5)}, true},
		{"test not matching element ring", args{generateRing(1, 2)}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := matchesNext(tt.args.r); got != tt.want {
				t.Errorf("matchesNext() = %v, want %v", got, tt.want)
			}
		})
	}
}
