package got

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"regexp"
	"strconv"
	"strings"
)

// EnsureCoverage via report file generated from, for example:
//
//	go test -coverprofile=coverage.out
//
// Return error if any file's coverage is less than min, min is a percentage value.
func EnsureCoverage(path string, min float64) error {
	tmp, _ := ioutil.TempFile("", "")
	report := tmp.Name()
	defer func() { _ = os.Remove(report) }()
	_ = tmp.Close()
	_, err := exec.Command("go", "tool", "cover", "-html", path, "-o", report).CombinedOutput()
	if err != nil {
		return err
	}

	list := parseReport(report)
	rejected := []string{}
	for _, c := range list {
		if c.coverage < min {
			rejected = append(rejected, fmt.Sprintf("  %s (%0.1f%%)", c.path, c.coverage))
		}
	}

	if len(rejected) > 0 {
		return fmt.Errorf(
			"Test coverage for these files should be greater than %.2f%%:\n%s",
			min,
			strings.Join(rejected, "\n"),
		)
	}
	return nil
}

type cov struct {
	path     string
	coverage float64
}

var regCov = regexp.MustCompile(`<option value="file\d+">(.+) \((\d+\.\d+)%\)</option>`)

func parseReport(path string) []cov {
	out, _ := ioutil.ReadFile(path)

	ms := regCov.FindAllStringSubmatch(string(out), -1)

	list := []cov{}
	for _, m := range ms {
		c, _ := strconv.ParseFloat(m[2], 32)
		list = append(list, cov{m[1], c})
	}
	return list
}
