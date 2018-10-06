# go-ipfsutil

> Utilities functions for interfacing with the IPFS CLI

[![License](http://img.shields.io/badge/license-MIT-blue.svg)](https://raw.githubusercontent.com/miguelmota/go-ipfsutil/master/LICENSE.md) [![Build Status](https://travis-ci.org/miguelmota/go-ipfsutil.svg?branch=master)](https://travis-ci.org/miguelmota/go-ipfsutil) [![Go Report Card](https://goreportcard.com/badge/github.com/miguelmota/go-ipfsutil?)](https://goreportcard.com/report/github.com/miguelmota/go-ipfsutil) [![GoDoc](https://godoc.org/github.com/miguelmota/go-ipfsutil?status.svg)](https://godoc.org/github.com/miguelmota/go-ipfsutil)

## Why?

I needed a simple wrapper around the CLI but the existing [go-ipfs-api](https://github.com/ipfs/go-ipfs-api) package requires the daemon to be running which is not required for certain operations and [go-ipfs](https://github.com/ipfs/go-ipfs) requires too much set up.

## Getting started

```go
package main

import (
	"fmt"
	"log"

	"github.com/miguelmota/go-ipfsutil"
)

func main() {
	hash, err := ipfsutil.AddBytes([]byte("hello world"))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(hash) // Qmf412jQZiuVUtdgnB36FXFX7xg5V6KEbSJ4dpQuhkLyfD
}
```

Check out the [tests](pfsutil_test.go) for more examples

## Test

```bash
make test
```

## License

MIT
