package hexagonal

import (
	"reflect"
	"testing"
)

func TestCoordinate_Move(t *testing.T) {
	type fields struct {
		x int
		y int
		z int
	}
	type args struct {
		direction Direction
	}
	tests := []struct {
		name     string
		fields   fields
		args     args
		expected fields
	}{
		{"test move N", fields{0, 0, 0}, args{N}, fields{0, 1, -1}},
		{"test move NE", fields{0, 0, 0}, args{NE}, fields{1, 0, -1}},
		{"test move SE", fields{0, 0, 0}, args{SE}, fields{1, -1, 0}},
		{"test move S", fields{0, 0, 0}, args{S}, fields{0, -1, 1}},
		{"test move SW", fields{0, 0, 0}, args{SW}, fields{-1, 0, 1}},
		{"test move NW", fields{0, 0, 0}, args{NW}, fields{-1, 1, 0}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			coordinate := &Coordinate{
				x: tt.fields.x,
				y: tt.fields.y,
				z: tt.fields.z,
			}
			coordinate.Move(tt.args.direction)

			expected := &Coordinate{
				x: tt.expected.x,
				y: tt.expected.y,
				z: tt.expected.z,
			}
			if !reflect.DeepEqual(coordinate, expected) {
				t.Errorf("Move() = %v, want %v", coordinate, expected)
			}
		})
	}
}

func TestNewCoordinate(t *testing.T) {
	type args struct {
		x int
		y int
		z int
	}
	tests := []struct {
		name    string
		args    args
		want    *Coordinate
		wantErr bool
	}{
		{"test origin", args{0, 0, 0}, &Coordinate{0, 0, 0}, false},
		{"test not origin", args{2, 0, -2}, &Coordinate{2, 0, -2}, false},
		{"test invalid", args{5, 0, 0}, nil, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewCoordinate(tt.args.x, tt.args.y, tt.args.z)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewCoordinate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewCoordinate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCoordinate_Distance(t *testing.T) {
	type fields struct {
		x int
		y int
		z int
	}
	type args struct {
		other *Coordinate
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   int
	}{
		{"test same coordinate", fields{0, 0, 0}, args{&Coordinate{0, 0, 0}}, 0},
		{"test other coordinate origin", fields{0, 0, 0}, args{&Coordinate{2, 0, -2}}, 2},
		{"test other coordinates", fields{-2, 2, 0}, args{&Coordinate{2, 0, -2}}, 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			coordinate := &Coordinate{
				x: tt.fields.x,
				y: tt.fields.y,
				z: tt.fields.z,
			}
			if got := coordinate.Distance(tt.args.other); got != tt.want {
				t.Errorf("Coordinate.Distance() = %v, want %v", got, tt.want)
			}
		})
	}
}
