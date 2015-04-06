// Copyright 2013 The Go Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build race
// +build windows

package net

import (
	"github.com/h4ck3rm1k3/gocore/run_time"
	"unsafe"
)

const raceenabled = true

func raceAcquire(addr unsafe.Pointer) {
	run_time.RaceAcquire(addr)
}

func raceReleaseMerge(addr unsafe.Pointer) {
	run_time.RaceReleaseMerge(addr)
}

func raceReadRange(addr unsafe.Pointer, len int) {
	run_time.RaceReadRange(addr, len)
}

func raceWriteRange(addr unsafe.Pointer, len int) {
	run_time.RaceWriteRange(addr, len)
}
