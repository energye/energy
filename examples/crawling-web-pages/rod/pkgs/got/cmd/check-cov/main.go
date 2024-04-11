// Package main ...
package main

import (
	"flag"
	"fmt"
	"github.com/energye/energy/v2/examples/crawling-web-pages/rod/pkgs/got"
	"os"
)

var covFile = flag.String("cov-file", "coverage.out", "the path of the coverage report")
var min = flag.Float64("min", 100, "min coverage rate or exit code with 1")

func main() {
	if !flag.Parsed() {
		flag.Parse()
	}

	err := got.EnsureCoverage(*covFile, *min)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
