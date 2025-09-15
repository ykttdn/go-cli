package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func hello(w io.Writer, args []string) {
	if len(args) > 1 {
		fmt.Fprintf(w, "Hello, %s!\n", strings.Join(args[1:], " "))
	} else {
		fmt.Fprintln(w, "Hello, World!")
	}
}

func main() {
	hello(os.Stdout, os.Args)
}
