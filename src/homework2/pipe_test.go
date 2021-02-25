package main

import (
	"testing"
)

func Test_pipe(t *testing.T) {
	tests := []struct {
		name     string
		value    int
		funcList []func(int) int
		want     int
	}{
		{"1", 5, []func(int) int{increment, increment, increment}, 8},
		{"2", 5, []func(int) int{increment, increment}, 7},
		{"3", 5, []func(int) int{increment}, 6},
		{"4", 5, []func(int) int{increment, increment, increment}, 8},
		{"5", 5, []func(int) int{increment, increment, increment, increment}, 9},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := pipe(tt.value, tt.funcList...); got != tt.want {
				t.Errorf("pipe() = %v,vant %v", got, tt.want)
			}
		})
	}
}
