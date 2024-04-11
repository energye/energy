package main

import (
	"bytes"
	"compress/gzip"
	"crypto/md5"
	"encoding/base64"
	"fmt"
	"github.com/energye/energy/v2/examples/crawling-web-pages/rod/pkgs/leakless/pkg/utils"
	"os"
	"os/exec"
	"path/filepath"
)

func main() {
	setVersion()

	utils.E(os.RemoveAll("dist"))

	for _, target := range targets {
		pack(target)
	}
}

func pack(target utils.Target) {
	var bin []byte
	var err error
	name := target.BinName()

	build(target)

	bin, err = utils.ReadFile(filepath.FromSlash("dist/leakless-" + name))
	utils.E(err)

	buf := bytes.Buffer{}
	gw, err := gzip.NewWriterLevel(&buf, 9)
	utils.E(err)
	utils.E(gw.Write(bin))
	utils.E(gw.Close())

	tpl := `package leakless

func init() {
	leaklessBinaries["%s"] = "%s"
}
`
	code := fmt.Sprintf(tpl, name, base64.StdEncoding.EncodeToString(buf.Bytes()))

	utils.E(utils.OutputFile(fmt.Sprintf("bin_%s.go", name), code, nil))
}

func setVersion() {
	a, err := filepath.Glob("cmd/leakless/*.go")
	utils.E(err)

	b, err := filepath.Glob("cmd/pack/*.go")
	utils.E(err)

	files := append(a, b...)

	args := append([]string{"hash-object"}, files...)

	raw, err := exec.Command("git", args...).CombinedOutput()
	utils.E(err)

	hash := md5.Sum(raw)

	utils.E(utils.OutputFile("pkg/shared/version.go", fmt.Sprintf(`package shared

// Version ...
const Version = "%x"
`, hash), nil))
}

func build(target utils.Target) {
	flags := []string{
		"build",
		"-trimpath",
		"-o", filepath.FromSlash("dist/leakless-" + target.BinName()),
	}

	ldFlags := "-ldflags=-w -s"
	if target.OS() == "windows" {
		// On Windows, -H windowsgui writes a "GUI binary" instead of a "console binary."
		ldFlags += " -H=windowsgui"
	}
	flags = append(flags, ldFlags)

	flags = append(flags, filepath.FromSlash("./cmd/leakless"))

	cmd := exec.Command("go", flags...)
	cmd.Env = append(os.Environ(), []string{
		"GOOS=" + target.OS(),
		"GOARCH=" + target.ARCH(),
	}...)
	utils.E(cmd.Run())
}
