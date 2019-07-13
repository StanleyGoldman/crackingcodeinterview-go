package urlify

import "testing"

func Test_urlify(t *testing.T) {
	type args struct {
		i string
		l int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Example Input",
			args: args{
				i: "Mr John Smith    ",
				l: 13,
			},
			want: "Mr%20John%20Smith",
		}, {
			name: "Custom Input 1",
			args: args{
				i: "Stanley Goldman  ",
				l: 15,
			},
			want: "Stanley%20Goldman",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := urlify(tt.args.i, tt.args.l); got != tt.want {
				t.Errorf("urlify() = %q, want %q", got, tt.want)
			}
		})
	}
}
