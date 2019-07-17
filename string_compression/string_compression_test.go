package compression

import "testing"

func Test_compressString(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Test Case",
			args: args{input: "aabcccccaaa"},
			want: "a2b1c5a3",
		},
		{
			name: "Fall Through Case",
			args: args{input: "abcd"},
			want: "abcd",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := compressString(tt.args.input); got != tt.want {
				t.Errorf("compressString() = %q, want %q", got, tt.want)
			}
		})
	}
}
