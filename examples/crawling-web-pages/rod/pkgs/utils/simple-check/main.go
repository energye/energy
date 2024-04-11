// Package main ...
package main

import (
	"github.com/energye/energy/v2/examples/crawling-web-pages/rod/pkgs/utils"
)

func main() {
	utils.Exec("go run ./lib/utils/setup")

	utils.Exec("go run ./lib/utils/lint")

	utils.Exec("go test -coverprofile=coverage.out ./lib/launcher")
	utils.Exec("go run ./lib/utils/check-cov")

	utils.Exec("go test -coverprofile=coverage.out")
	utils.Exec("go run ./lib/utils/check-cov")
}
