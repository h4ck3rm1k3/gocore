// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package run_time_test

import (
	"github.com/h4ck3rm1k3/gocore/run_time"
	"github.com/h4ck3rm1k3/gocore/strings"
	"testing"
)

func TestCaller(t *testing.T) {
	procs := run_time.GOMAXPROCS(-1)
	c := make(chan bool, procs)
	for p := 0; p < procs; p++ {
		go func() {
			for i := 0; i < 1000; i++ {
				testCallerFoo(t)
			}
			c <- true
		}()
		defer func() {
			<-c
		}()
	}
}

func testCallerFoo(t *testing.T) {
	testCallerBar(t)
}

func testCallerBar(t *testing.T) {
	for i := 0; i < 2; i++ {
		pc, file, line, ok := run_time.Caller(i)
		f := run_time.FuncForPC(pc)
		if !ok ||
			!strings.HasSuffix(file, "symtab_test.go") ||
			(i == 0 && !strings.HasSuffix(f.Name(), "testCallerBar")) ||
			(i == 1 && !strings.HasSuffix(f.Name(), "testCallerFoo")) ||
			line < 5 || line > 1000 ||
			f.Entry() >= pc {
			t.Errorf("incorrect symbol info %d: %t %d %d %s %s %d",
				i, ok, f.Entry(), pc, f.Name(), file, line)
		}
	}
}
