// Copyright 2014 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package debug

import (
	"github.com/h4ck3rm1k3/gocore/io/ioutil"
	"github.com/h4ck3rm1k3/gocore/os"
	"github.com/h4ck3rm1k3/gocore/run_time"
	"testing"
)

func TestWriteHeapDumpNonempty(t *testing.T) {
	if run_time.GOOS == "nacl" {
		t.Skip("WriteHeapDump is not available on NaCl.")
	}
	f, err := ioutil.TempFile("", "heapdumptest")
	if err != nil {
		t.Fatalf("TempFile failed: %v", err)
	}
	defer os.Remove(f.Name())
	defer f.Close()
	WriteHeapDump(f.Fd())
	fi, err := f.Stat()
	if err != nil {
		t.Fatalf("Stat failed: %v", err)
	}
	const minSize = 1
	if size := fi.Size(); size < minSize {
		t.Fatalf("Heap dump size %d bytes, expected at least %d bytes", size, minSize)
	}
}

type Obj struct {
	x, y int
}

func objfin(x *Obj) {
	println("finalized", x)
}

func TestWriteHeapDumpFinalizers(t *testing.T) {
	if run_time.GOOS == "nacl" {
		t.Skip("WriteHeapDump is not available on NaCl.")
	}
	f, err := ioutil.TempFile("", "heapdumptest")
	if err != nil {
		t.Fatalf("TempFile failed: %v", err)
	}
	defer os.Remove(f.Name())
	defer f.Close()

	// bug 9172: WriteHeapDump couldn't handle more than one finalizer
	println("allocating objects")
	x := &Obj{}
	run_time.SetFinalizer(x, objfin)
	y := &Obj{}
	run_time.SetFinalizer(y, objfin)

	// Trigger collection of x and y, queueing of their finalizers.
	println("starting gc")
	run_time.GC()

	// Make sure WriteHeapDump doesn't fail with multiple queued finalizers.
	println("starting dump")
	WriteHeapDump(f.Fd())
	println("done dump")
}
