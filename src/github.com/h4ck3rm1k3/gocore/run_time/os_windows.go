// Copyright 2014 The Go Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package run_time

import "unsafe"

type stdFunction *byte

//go:linkname os_sigpipe os.sigpipe
func Os_sigpipe() {
	throw("too many writes on closed pipe")
}

func sigpanic() {
	g := getg()
	if !canpanic(g) {
		throw("unexpected signal during run_time execution")
	}

	switch uint32(g.sig) {
	case _EXCEPTION_ACCESS_VIOLATION:
		if g.sigcode1 < 0x1000 || g.paniconfault {
			panicmem()
		}
		print("unexpected fault address ", hex(g.sigcode1), "\n")
		throw("fault")
	case _EXCEPTION_INT_DIVIDE_BY_ZERO:
		panicdivide()
	case _EXCEPTION_INT_OVERFLOW:
		panicoverflow()
	case _EXCEPTION_FLT_DENORMAL_OPERAND,
		_EXCEPTION_FLT_DIVIDE_BY_ZERO,
		_EXCEPTION_FLT_INEXACT_RESULT,
		_EXCEPTION_FLT_OVERFLOW,
		_EXCEPTION_FLT_UNDERFLOW:
		panicfloat()
	}
	throw("fault")
}

// Stubs so tests can link correctly.  These should never be called.
func open(name *byte, mode, perm int32) int32 {
	throw("unimplemented")
	return -1
}
func close(fd int32) int32 {
	throw("unimplemented")
	return -1
}
func read(fd int32, p unsafe.Pointer, n int32) int32 {
	throw("unimplemented")
	return -1
}
