package main

import (
	"fmt"
	"testing"
)

func TestCommaToPeriod(t *testing.T) {
	var tests = []struct {
		flt  string
		want string
	}{
		{"123,20", "123.20"},
		{"10,5", "10.5"},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%s", tt.flt)
		t.Run(testname, func(t *testing.T) {
			ans := commaToPeriod(tt.flt)
			if ans != tt.want {
				t.Errorf("got %s, want %s", ans, tt.want)
			}
		})
	}
}
