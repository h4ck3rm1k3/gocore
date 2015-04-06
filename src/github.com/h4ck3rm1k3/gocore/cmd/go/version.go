// Copyright 2011 The Go Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"github.com/h4ck3rm1k3/gocore/fmt"
	"github.com/h4ck3rm1k3/gocore/run_time"
)

var cmdVersion = &Command{
	Run:       runVersion,
	UsageLine: "version",
	Short:     "print Go version",
	Long:      `Version prints the Go version, as reported by run_time.Version.`,
}

func runVersion(cmd *Command, args []string) {
	if len(args) != 0 {
		cmd.Usage()
	}

	fmt.Printf("go version %s %s/%s\n", run_time.Version(), run_time.GOOS, run_time.GOARCH)
}
