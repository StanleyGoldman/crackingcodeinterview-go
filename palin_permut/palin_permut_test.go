package palinpermut

import "testing"

func Test_isPalindromePermutation(t *testing.T) {
	type args struct {
		i string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Example case",
			args: args{i: "Tact Coa"},
			want: true,
		},
		{
			name: "Custom test case 1",
			args: args{i: "abb ba"},
			want: true,
		},
		{
			name: "Custom test case 2",
			args: args{i: "abb dba"},
			want: false,
		},
		{
			name: "Custom test case 3",
			args: args{i: "Tactt Coa"},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPalindromePermutation(tt.args.i); got != tt.want {
				t.Errorf("isPalindromePermutation() = %v, want %v", got, tt.want)
			}
		})
	}
}
