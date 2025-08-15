package main

import (
	"testing"
)

func TestEcho(t *testing.T) {
	table := []struct {
		input             string
		hasNoTrailingLine bool
		interpolates      bool
		want              string
	}{
		{"Hello, World!", false, false, "Hello, World!\n"},
		{"Hello, World!", true, false, "Hello, World!"},
		{"Hello, \\nWorld!", false, true, "Hello, \nWorld!\n"},
		{"Hello, \\nWorld!", true, true, "Hello, \nWorld!"},
	}

	for _, tt := range table {
		got := echo(tt.input, tt.hasNoTrailingLine, tt.interpolates)
		if got != tt.want {
			t.Errorf("echo(%q, %v, %v) = %q, want %q", tt.input, tt.hasNoTrailingLine, tt.interpolates, got, tt.want)
		}
	}
}
