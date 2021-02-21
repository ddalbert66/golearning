package main

import (
	"testing"
)

func Test_pipe(t *testing.T) {
	tests := []struct {
		num int
		ans int
	}{
		{pipe(5, increment, increment, increment), 8},
		{pipe(5, increment), 6},
		{pipe(5, increment, increment), 7},
	}

	for _, tt := range tests {
		if tt.num != tt.ans {
			t.Errorf("Go TestlenpeatingSubStr(%d) ;"+"got %s; %d", tt.num, tt.ans)
		}
	}
}
