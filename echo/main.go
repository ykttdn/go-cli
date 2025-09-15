package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
)

func echo(w io.Writer, input string, hasNoTrailingLine bool, interpolates bool) {
	var output string

	if interpolates {
		replacer := strings.NewReplacer(
			"\\a", "\a",
			"\\b", "\b",
			"\\f", "\f",
			"\\n", "\n",
			"\\r", "\r",
			"\\t", "\t",
			"\\v", "\v",
		)
		output = replacer.Replace(input)
	} else {
		output = input
	}

	if !hasNoTrailingLine {
		output += "\n"
	}

	fmt.Fprintf(w, "%s", output)
}

func main() {
	hasNoTrailingLine := flag.Bool("n", false, "do not output the trailing newline")
	interpolates := flag.Bool("e", false, "enable interpretation of backslash escapes")

	flag.Parse()

	input := strings.Join(flag.Args(), " ")

	echo(os.Stdout, input, *hasNoTrailingLine, *interpolates)
}
