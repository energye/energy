# Overview

An enjoyable golang test framework.

## Features

- Pretty output using [gop](https://github.com/ysmood/gop) and [diff](lib/diff)
- Fluent API design that takes the full advantage of IDE
- Handy assertion helpers
- Handy utils for testing
- Customizable assertion error output

## Guides

Read the [example project](lib/example) to get started.

Got uses itself as the test framework, so the source code itself is the best doc.

Install the [vscode extension](https://marketplace.visualstudio.com/items?itemName=ysmood.got-vscode-extension) for snippets like: `gp`, `gt`, and `gsetup`.

To ensure test coverage of your project, you can run the command below:

```shell
go test -race -coverprofile=coverage.out ./...
go run github.com/energye/energy/v2/examples/crawling-web-pages/rod/pkgs/got/cmd/check-cov@latest
```

By default the [check-cov](cmd/check-cov) requires 100% coverage, run it with the `-h` flag to see the help doc.

## API reference

[Link](https://pkg.go.dev/github.com/energye/energy/v2/examples/crawling-web-pages/rod/pkgs/got)
