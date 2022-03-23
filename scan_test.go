package main

import (
	"fmt"
	"testing"
)

func TestCleanStrings(t *testing.T) {
	var tests = []struct {
		doc  string
		want string
	}{
		{"79.379.491/0008-50", "79379491000850"},
		{"035.769.019-22", "03576901922"},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%s", tt.doc)
		t.Run(testname, func(t *testing.T) {
			ans := cleanStrings(tt.doc)
			if ans != tt.want {
				t.Errorf("got %s, want %s", ans, tt.want)
			}
		})
	}
}

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
