package main

import (
	"testing"
)

func TestSayHello(t *testing.T) {
	want := "Hello, World!"
	got := sayHello()

	if got != want {
		t.Errorf("sayHello() = %q, want %q", got, want)
	}
}
