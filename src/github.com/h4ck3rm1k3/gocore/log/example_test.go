// Copyright 2013 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package log_test

import (
	"github.com/h4ck3rm1k3/gocore/bytes"
	"github.com/h4ck3rm1k3/gocore/fmt"
	"github.com/h4ck3rm1k3/gocore/log"
)

func ExampleLogger() {
	var buf bytes.Buffer
	logger := log.New(&buf, "logger: ", log.Lshortfile)
	logger.Print("Hello, log file!")

	fmt.Print(&buf)
	// Output:
	// logger: example_test.go:16: Hello, log file!
}
