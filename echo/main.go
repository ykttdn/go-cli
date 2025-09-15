package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
)

func echo(w io.Writer, input string, hasNoTrailingLine bool, interpolates bool) {
	var result string

	if len(input) > 0 {
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
			result = replacer.Replace(input)
		} else {
			result = input
		}
	}

	if !hasNoTrailingLine {
		result += "\n"
	}

	fmt.Fprintf(w, "%s", result)
}

func main() {
	hasNoTrailingLine := flag.Bool("n", false, "do not output the trailing newline")
	interpolates := flag.Bool("e", false, "enable interpretation of backslash escapes")

	flag.Parse()

	input := strings.Join(flag.Args(), " ")

	echo(os.Stdout, input, *hasNoTrailingLine, *interpolates)
}
