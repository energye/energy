//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

// Package bindata 将资源以二进制形式生成go资源文件
package bindata

import (
	"github.com/cyber-xxm/energy/v2/cmd/internal/command"
	"path/filepath"
	"regexp"
	"strings"
)

func Bindata(c *command.Config) error {
	cfg := argsConfig(c.Bindata)
	return Translate(cfg)
}

// parseArgs create s a new, filled configuration instance
// by reading and parsing command line options.
//
// This function exits the program with an error, if
// any of the command line options are incorrect.
func argsConfig(bind command.Bindata) *Config {
	c := NewConfig()
	c.Debug = bind.Debug
	c.Dev = bind.Dev
	c.Tags = bind.Tags
	c.Prefix = bind.Prefix
	c.Package = bind.Package
	c.NoMemCopy = bind.NoMemCopy
	c.NoCompress = bind.NoCompress
	c.NoMetadata = bind.NoMetadata
	c.FSSystem = bind.FSSystem
	c.Mode = bind.Mode
	c.ModTime = bind.ModTime
	c.Output = bind.Output
	c.Paths = bind.Paths

	if bind.Ignore != "" {
		ignores := strings.Split(bind.Ignore, ",")
		patterns := make([]*regexp.Regexp, 0)
		for _, pattern := range ignores {
			patterns = append(patterns, regexp.MustCompile(pattern))
		}
		c.Ignore = patterns
	}
	if bind.Paths != "" {
		paths := strings.Split(bind.Paths, ",")
		c.Input = make([]InputConfig, len(paths))
		for i := range paths {
			c.Input[i] = parseInput(paths[i])
		}
	}
	return c
}

// parseRecursive determines whether the given path has a recursive indicator and
// returns a new path with the recursive indicator chopped off if it does.
//
//	ex:
//	    /path/to/foo/...    -> (/path/to/foo, true)
//	    /path/to/bar        -> (/path/to/bar, false)
func parseInput(path string) InputConfig {
	if strings.HasSuffix(path, "/...") {
		return InputConfig{
			Path:      filepath.Clean(path[:len(path)-4]),
			Recursive: true,
		}
	} else {
		return InputConfig{
			Path:      filepath.Clean(path),
			Recursive: false,
		}
	}

}
