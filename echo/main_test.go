package main

import (
	"bytes"
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
		var buf bytes.Buffer
		echo(&buf, tt.input, tt.hasNoTrailingLine, tt.interpolates)

		if buf.String() != tt.want {
			t.Errorf("echo(%q, %v, %v) = %q, want %q", tt.input, tt.hasNoTrailingLine, tt.interpolates, buf.String(), tt.want)
		}
	}
}
