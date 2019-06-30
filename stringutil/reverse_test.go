package stringutil

import "testing"

func TestReverse(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"should reverse string",
			args{s: "abcdefg"},
			"gfedcba",
		},
		{
			"empty string",
			args{s:""},
			"",
		},
		{
			"should do nothing with single letter",
			args{s: "a"},
			"a",
		},
		{
			"should reverse multi word string",
			args{s: "ahoj benky jedem!"},
			"!medej ykneb joha",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Reverse(tt.args.s); got != tt.want {
				t.Errorf("Reverse() = %v, want %v", got, tt.want)
			}
		})
	}
}
