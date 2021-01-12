package main

import (
	"testing"
)

func TestLengthOfRepeatingSubStr(t *testing.T) {
	tests := []struct {
		s   string
		ans int
	}{
		{"xjkjkj", 2},
		{"安你好安安你好", 7},
	}

	for _, tt := range tests {
		actual := lengthOfRepeatingSubStr(tt.s)
		if actual != tt.ans {
			t.Errorf("Go TestlengthOfRepeatingSubStr(%s) ;"+"got %d; %d", tt.s, actual, tt.ans)
		}
	}
}

func BenchmarkNoReaptingSubStr(b *testing.B) {
	s := "安你好安安你好"
	ans := 7

	for i := 0; i < b.N; i++ {
		actual := lengthOfRepeatingSubStr(s)
		if actual != ans {
			b.Errorf("Go TestlengthOfRepeatingSubStr(%s) ;"+"got %d; %d", s, actual, ans)
		}
	}

}
