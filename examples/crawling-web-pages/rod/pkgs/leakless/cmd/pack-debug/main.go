package main

import (
	"fmt"
	"github.com/energye/energy/v2/examples/crawling-web-pages/rod/pkgs/leakless"
	"os/exec"
)

func main() {
	path := leakless.GetLeaklessBin()
	out, err := exec.Command("go", "build", "-o", path, "./cmd/leakless").CombinedOutput()
	fmt.Println(path, string(out), err)
}
