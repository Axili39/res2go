package main

import (
	"fmt"
	"os"
)

func main() {
	Init()

	if len(Files) != 3 {
		fmt.Fprintf(os.Stderr, "bad Files length %d/%d", len(Files), 3)
		os.Exit(1)
	}

	if _, ok := Files["tests/index.html"]; !ok {
		fmt.Fprintf(os.Stderr, "missing file %s", "tests/index.html")
		os.Exit(1)
	}

	if _, ok := Files["tests/resources/bar.template"]; !ok {
		fmt.Fprintf(os.Stderr, "missing file %s", "tests/resources/bar.template")
		os.Exit(1)
	}

	if _, ok := Files["tests/resources/foo.template"]; !ok {
		fmt.Fprintf(os.Stderr, "missing file %s", "tests/resources/foo.template")
		os.Exit(1)
	}
	os.Exit(0)
}
