package rotatematrix

import (
	"reflect"
	"testing"
)

func Test_replace(t *testing.T) {
	type args struct {
		m matrix
		x uint
		y uint
		v uint
	}
	tests := []struct {
		name string
		args args
		want uint
	}{
		{
			name: "Simple Replace",
			args: args{
				m: matrix{
					{1, 2},
					{3, 4},
				},
				x: 1,
				y: 0,
				v: 5,
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := replace(tt.args.m, tt.args.x, tt.args.y, tt.args.v); got != tt.want {
				t.Errorf("replace() = %v, want %v", got, tt.want)
			}

			if tt.args.m[tt.args.y][tt.args.x] != tt.args.v {
				t.Errorf("matrix[%v][%v] = %v, want %v", tt.args.x, tt.args.y, tt.args.m[tt.args.y][tt.args.x], tt.args.v)
			}
		})
	}
}

func Test_getNextPoint(t *testing.T) {
	type args struct {
		x uint
		y uint
		s uint
		d direction
	}
	tests := []struct {
		name string
		args args
		want nextPoint
	}{
		{
			name: "First Position 2x2",
			args: args{0, 0, 2, right},
			want: nextPoint{1, 0, up},
		},
		{
			name: "Second Position 2x2",
			args: args{1, 0, 2, up},
			want: nextPoint{1, 1, left},
		},
		{
			name: "Third Position 2x2",
			args: args{1, 1, 2, left},
			want: nextPoint{0, 1, down},
		},
		{
			name: "Fourth Position 2x2",
			args: args{0, 1, 2, down},
			want: nextPoint{0, 0, right},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getNextPoint(tt.args.x, tt.args.y, tt.args.s, tt.args.d); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getNextPoint() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_rotateMatrix(t *testing.T) {
	type args struct {
		input matrix
	}
	tests := []struct {
		name string
		args args
		want matrix
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := rotateMatrix(tt.args.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("rotateMatrix() = %v, want %v", got, tt.want)
			}
		})
	}
}
