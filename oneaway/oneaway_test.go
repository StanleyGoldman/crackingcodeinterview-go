package oneaway

import "testing"

func Test_oneAway(t *testing.T) {
	type args struct {
		x string
		y string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Test Case 1",
			args: args{x: "pale", y: "ple"},
			want: true,
		},
		{
			name: "Test Case 2",
			args: args{x: "pales", y: "pale"},
			want: true,
		},
		{
			name: "Test Case 3",
			args: args{x: "pale", y: "bale"},
			want: true,
		},
		{
			name: "Test Case 4",
			args: args{x: "pale", y: "bake"},
			want: false,
		},
		{
			name: "Custom Test Case: Too Different",
			args: args{x: "pale", y: "paleee"},
			want: false,
		},
		{
			name: "Custom Test Case 1",
			args: args{x: "pale", y: "bake"},
			want: false,
		},
		{
			name: "Custom Test Case 2",
			args: args{x: "bbake", y: "bakee"},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := oneAway(tt.args.x, tt.args.y); got != tt.want {
				t.Errorf("oneAway() = %v, want %v", got, tt.want)
			}
		})
	}
}
