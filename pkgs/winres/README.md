# winres


Put this file in your project directory, name it "something.syso" or, preferably,
"something_windows_amd64.syso", and you're done :
the `go build` command will detect it and automatically use it.

You should have a look at the [command line tool](https://github.com/tc-hib/go-winres) to try it. Using the library
gives you more control, though.

Here is a quick example:

```go
package main

import (
	"io/ioutil"
	"os"

	"github.com/tc-hib/winres"
)

func main() {
	// Start by creating an empty resource set
	rs := winres.ResourceSet{}

	// Add resources
	// This is a cursor named ID(1)
	cursorData, _ := ioutil.ReadFile("cursor.cur")
	rs.Set(winres.RT_CURSOR, winres.ID(1), 0, cursorData)

	// This is a custom data type, translated in english (0x409) and french (0x40C)
	// You can find more language IDs by searching for LCID
	rs.Set(winres.Name("CUSTOM"), winres.Name("COOLDATA"), 0x409, []byte("Hello World"))
	rs.Set(winres.Name("CUSTOM"), winres.Name("COOLDATA"), 0x40C, []byte("Bonjour Monde"))

	// Compile to a COFF object file
	// It is recommended to use the target suffix "_window_amd64"
	// so that `go build` knows when not to include it.
	out, _ := os.Create("rsrc_windows_amd64.syso")
	rs.WriteObject(out, winres.ArchAMD64)
}
```

# From: github.com/tc-hib/winres