package main

import (
	"bytes"
	"testing"
)

func TestHello(t *testing.T) {
	tests := []struct {
		name     string
		args     []string
		expected string
	}{
		{
			name:     "With one argument",
			args:     []string{"program_name", "Go"},
			expected: "Hello, Go!\n",
		},
		{
			name:     "With two arguments",
			args:     []string{"program_name", "Go", "Developer"},
			expected: "Hello, Go Developer!\n",
		},
		{
			name:     "Without argument",
			args:     []string{"program_name"},
			expected: "Hello, World!\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf bytes.Buffer
			hello(&buf, tt.args)

			if buf.String() != tt.expected {
				t.Errorf("expected: %q, actual: %q", tt.expected, buf.String())
			}
		})
	}
}
