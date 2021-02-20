package main

import (
	"testing"
)

func Test_formateMoney(t *testing.T) {
	tests := []struct {
		money float64
		ans   string
	}{
		{9527, "9,527"},
		{3345678, "3,345,678"},
		{-1234.45, "-1,234.45"},
	}

	for _, tt := range tests {
		actual := FormateMoney(tt.money)
		if actual != tt.ans {
			t.Errorf("Go TestlenpeatingSubStr(%f) ;"+"got %s; %s", tt.money, actual, tt.ans)
		}
	}
}
