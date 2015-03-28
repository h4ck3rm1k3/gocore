// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build ignore

package main

// Need to compile package gob with debug.go to build this program.
// See comments in debug.go for how to do this.

import (
	"github.com/h4ck3rm1k3/gocore/encoding/gob"
	"github.com/h4ck3rm1k3/gocore/fmt"
	"github.com/h4ck3rm1k3/gocore/os"
)

func main() {
	var err error
	file := os.Stdin
	if len(os.Args) > 1 {
		file, err = os.Open(os.Args[1])
		if err != nil {
			fmt.Fprintf(os.Stderr, "dump: %s\n", err)
			os.Exit(1)
		}
	}
	gob.Debug(file)
}
