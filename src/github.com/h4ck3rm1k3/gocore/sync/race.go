// Copyright 2012 The Go Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build race

package sync

import (
	"github.com/h4ck3rm1k3/gocore/run_time"
	"unsafe"
)

const raceenabled = true

func raceAcquire(addr unsafe.Pointer) {
	run_time.RaceAcquire(addr)
}

func raceRelease(addr unsafe.Pointer) {
	run_time.RaceRelease(addr)
}

func raceReleaseMerge(addr unsafe.Pointer) {
	run_time.RaceReleaseMerge(addr)
}

func raceDisable() {
	run_time.RaceDisable()
}

func raceEnable() {
	run_time.RaceEnable()
}

func raceRead(addr unsafe.Pointer) {
	run_time.RaceRead(addr)
}

func raceWrite(addr unsafe.Pointer) {
	run_time.RaceWrite(addr)
}
