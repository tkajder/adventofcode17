package hexagonal

import "testing"

func TestParseDirection(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name    string
		args    args
		want    Direction
		wantErr bool
	}{
		{"test N", args{"N"}, N, false},
		{"test NE", args{"Ne"}, NE, false},
		{"test SE", args{"sE"}, SE, false},
		{"test S", args{"S"}, S, false},
		{"test SW", args{"Sw"}, SW, false},
		{"test NW", args{"nw"}, NW, false},
		{"test invalid", args{"iNvaLid"}, -1, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseDirection(tt.args.s)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseDirection() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ParseDirection() = %v, want %v", got, tt.want)
			}
		})
	}
}
