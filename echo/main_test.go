package main

import (
	"bytes"
	"testing"
)

func TestEcho(t *testing.T) {
	tests := []struct {
		name              string
		input             string
		hasNoTrailingLine bool
		interpolates      bool
		want              string
	}{
		{
			name:              "Without options",
			input:             "Hello, World!",
			hasNoTrailingLine: false,
			interpolates:      false,
			want:              "Hello, World!\n",
		},
		{
			name:              "Without trailing newline",
			input:             "Hello, World!",
			hasNoTrailingLine: true,
			interpolates:      false,
			want:              "Hello, World!",
		},
		{
			name:              "With interpolations",
			input:             "Hello, \\nWorld!",
			hasNoTrailingLine: false,
			interpolates:      true,
			want:              "Hello, \nWorld!\n",
		},
		{
			name:              "With all options",
			input:             "Hello, \\nWorld!",
			hasNoTrailingLine: true,
			interpolates:      true,
			want:              "Hello, \nWorld!",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf bytes.Buffer
			echo(&buf, tt.input, tt.hasNoTrailingLine, tt.interpolates)

			if buf.String() != tt.want {
				t.Errorf("echo(%q, %v, %v) = %q, want %q", tt.input, tt.hasNoTrailingLine, tt.interpolates, buf.String(), tt.want)
			}
		})
	}
}
