// Package got is an enjoyable golang test framework.
package got

import (
	"flag"
	"github.com/energye/energy/v2/examples/crawling-web-pages/rod/pkgs/gop"
	"github.com/energye/energy/v2/examples/crawling-web-pages/rod/pkgs/got/lib/diff"
	"os"
	"reflect"
	"regexp"
	"strings"
)

// Testable interface. Usually, you use *testing.T as it.
type Testable interface {
	Name() string                            // same as testing.common.Name
	Skipped() bool                           // same as testing.common.Skipped
	Failed() bool                            // same as testing.common.Failed
	Cleanup(func())                          // same as testing.common.Cleanup
	FailNow()                                // same as testing.common.FailNow
	Fail()                                   // same as testing.common.Fail
	Helper()                                 // same as testing.common.Helper
	Logf(format string, args ...interface{}) // same as testing.common.Logf
	SkipNow()                                // same as testing.common.Skip
}

// G is the helper context, it provides some handy helpers for testing
type G struct {
	Testable
	Assertions
	Utils
}

// Setup returns a helper to init G instance
func Setup(init func(g G)) func(t Testable) G {
	return func(t Testable) G {
		g := New(t)
		if init != nil {
			init(g)
		}
		return g
	}
}

// T is the shortcut for New
func T(t Testable) G {
	return New(t)
}

// New G instance
func New(t Testable) G {
	eh := NewDefaultAssertionError(gop.ThemeDefault, diff.ThemeDefault)
	return G{
		t,
		Assertions{Testable: t, ErrorHandler: eh},
		Utils{t},
	}
}

// DefaultFlags will set the "go test" flag if not yet presented.
// It must be executed in the init() function.
// Such as the timeout:
//
//	DefaultFlags("timeout=10s")
func DefaultFlags(flags ...string) {
	// remove default timeout from "go test"
	filtered := []string{}
	for _, arg := range os.Args {
		if arg != "-test.timeout=10m0s" {
			filtered = append(filtered, arg)
		}
	}
	os.Args = filtered

	list := map[string]struct{}{}
	reg := regexp.MustCompile(`^-test\.(\w+)`)
	for _, arg := range os.Args {
		ms := reg.FindStringSubmatch(arg)
		if ms != nil {
			list[ms[1]] = struct{}{}
		}
	}

	for _, flag := range flags {
		if _, has := list[strings.Split(flag, "=")[0]]; !has {
			os.Args = append(os.Args, "-test."+flag)
		}
	}
}

// Parallel config of "go test -parallel"
func Parallel() (n int) {
	flag.Parse()
	flag.Visit(func(f *flag.Flag) {
		if f.Name == "test.parallel" {
			v := reflect.ValueOf(f.Value).Elem().Convert(reflect.TypeOf(n))
			n = v.Interface().(int)
		}
	})
	return
}
