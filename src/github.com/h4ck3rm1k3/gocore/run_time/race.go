// Copyright 2012 The Go Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build race

// Public race detection API, present iff build with -race.

package run_time

import (
	"unsafe"
)

func RaceRead(addr unsafe.Pointer){
	panic("not implemented")
}
func RaceWrite(addr unsafe.Pointer){
	panic("not implemented")
}
func RaceReadRange(addr unsafe.Pointer, len int){
	panic("not implemented")
}
func RaceWriteRange(addr unsafe.Pointer, len int){
	panic("not implemented")
}

func RaceSemacquire(s *uint32){
	panic("not implemented")
}
func RaceSemrelease(s *uint32){
	panic("not implemented")
}

// private interface for the run_time
const raceenabled = true

// For all functions accepting callerpc and pc,
// callerpc is a return PC of the function that calls this function,
// pc is start PC of the function that calls this function.
func raceReadObjectPC(t *_type, addr unsafe.Pointer, callerpc, pc uintptr) {
	kind := t.kind & kindMask
	if kind == kindArray || kind == kindStruct {
		// for composite objects we have to read every address
		// because a write might happen to any subobject.
		racereadrangepc(addr, t.size, callerpc, pc)
	} else {
		// for non-composite objects we can read just the start
		// address, as any write must write the first byte.
		racereadpc(addr, callerpc, pc)
	}
}

func raceWriteObjectPC(t *_type, addr unsafe.Pointer, callerpc, pc uintptr) {
	kind := t.kind & kindMask
	if kind == kindArray || kind == kindStruct {
		// for composite objects we have to write every address
		// because a write might happen to any subobject.
		racewriterangepc(addr, t.size, callerpc, pc)
	} else {
		// for non-composite objects we can write just the start
		// address, as any write must write the first byte.
		racewritepc(addr, callerpc, pc)
	}
}

//go:noescape
func racereadpc(addr unsafe.Pointer, callpc, pc uintptr){
	panic("not implemented")
}

//go:noescape
func racewritepc(addr unsafe.Pointer, callpc, pc uintptr){
	panic("not implemented")
}

type symbolizeContext struct {
	pc   uintptr
	fn   *byte
	file *byte
	line uintptr
	off  uintptr
	res  uintptr
}

var qq = [...]byte{'?', '?', 0}
var dash = [...]byte{'-', 0}

// Callback from C into Go, runs on g0.
func racesymbolize(ctx *symbolizeContext) {
	f := findfunc(ctx.pc)
	if f == nil {
		ctx.fn = &qq[0]
		ctx.file = &dash[0]
		ctx.line = 0
		ctx.off = ctx.pc
		ctx.res = 1
		return
	}

	ctx.fn = cfuncname(f)
	file, line := funcline(f, ctx.pc)
	ctx.line = uintptr(line)
	ctx.file = &bytes(file)[0] // assume NUL-terminated
	ctx.off = ctx.pc - f.entry
	ctx.res = 1
	return
}
