package sliceutils

import (
	"reflect"
	"testing"
)

func Test_sliceAtoi64(t *testing.T) {
	type args struct {
		strings []string
	}
	tests := []struct {
		name    string
		args    args
		want    []int64
		wantErr bool
	}{
		{"test empty slice", args{[]string{}}, []int64{}, false},
		{"test valid slice", args{[]string{"12", "44", "-354", "0"}}, []int64{12, 44, -354, 0}, false},
		{"test invalid slice", args{[]string{"3", "scx"}}, nil, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Atoi64(tt.args.strings)
			if (err != nil) != tt.wantErr {
				t.Errorf("sliceAtoi64() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sliceAtoi64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_sliceAtoui64(t *testing.T) {
	type args struct {
		strings []string
	}
	tests := []struct {
		name    string
		args    args
		want    []uint64
		wantErr bool
	}{
		{"test empty slice", args{[]string{}}, []uint64{}, false},
		{"test valid slice", args{[]string{"12", "44", "354", "0"}}, []uint64{12, 44, 354, 0}, false},
		{"test negative num slice", args{[]string{"-5", "44"}}, nil, true},
		{"test invalid str slice", args{[]string{"3", "scx"}}, nil, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Atoui64(tt.args.strings)
			if (err != nil) != tt.wantErr {
				t.Errorf("sliceAtoui64() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sliceAtoui64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_sliceAtoi(t *testing.T) {
	type args struct {
		strings []string
	}
	tests := []struct {
		name    string
		args    args
		want    []int
		wantErr bool
	}{
		{"test empty slice", args{[]string{}}, []int{}, false},
		{"test valid slice", args{[]string{"12", "44", "-354", "0"}}, []int{12, 44, -354, 0}, false},
		{"test invalid slice", args{[]string{"3", "scx"}}, nil, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Atoi(tt.args.strings)
			if (err != nil) != tt.wantErr {
				t.Errorf("sliceAtoi() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sliceAtoi() = %v, want %v", got, tt.want)
			}
		})
	}
}
