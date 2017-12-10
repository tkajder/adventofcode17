package main

import (
	"reflect"
	"testing"
)

func Test_pointFromSpiralCoordinate(t *testing.T) {
	type args struct {
		coordinate uint64
	}
	tests := []struct {
		name    string
		args    args
		want    Point
		wantErr bool
	}{
		{"test error case 0", args{0}, Point{}, true},
		{"test origin", args{1}, Point{0, 0}, false},
		{"test even squared+1", args{5}, Point{-1, 1}, false},
		{"test gt even squared+1", args{20}, Point{-2, -1}, false},
		{"test lt even squared+1", args{33}, Point{1, 3}, false},
		{"test odd squared", args{9}, Point{1, -1}, false},
		{"test gt odd squared", args{53}, Point{4, 0}, false},
		{"test lt odd squared", args{47}, Point{1, -3}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := pointFromSpiralCoordinate(tt.args.coordinate)
			if (err != nil) != tt.wantErr {
				t.Errorf("pointFromSpiralCoordinate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("pointFromSpiralCoordinate() = %v, want %v", got, tt.want)
			}
		})
	}
}
func Test_manhattanDistance(t *testing.T) {
	type args struct {
		p1 Point
		p2 Point
	}
	tests := []struct {
		name string
		args args
		want uint64
	}{
		{"test same point", args{Point{1, 2}, Point{1, 2}}, 0},
		{"test different x", args{Point{1, 2}, Point{4, 2}}, 3},
		{"test different y", args{Point{1, 2}, Point{1, -2}}, 4},
		{"test different x and y", args{Point{1, 2}, Point{6, -14}}, 21},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := manhattanDistance(tt.args.p1, tt.args.p2); got != tt.want {
				t.Errorf("manhattanDistance() = %v, want %v", got, tt.want)
			}
		})
	}
}
