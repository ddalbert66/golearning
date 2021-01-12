package main

import "testing"

func lengthOfRepeatingSubStrTest(t *testing.T) {
	tests := []struct {
		s   string
		ans int
	}{
		{"xjkjkj", 3},
		{"sdfsdf", 3},
	}

	for _, tt := range tests {
		actual := lengthOfRepeatingSubStr(tt.s)
		if actual != tt.ans {
			t.Errorf("Go lengthOfRepeatingSubStrTest(%s) ;"+"got %d; %d", tt.s, actual, tt.ans)
		}
	}
}
