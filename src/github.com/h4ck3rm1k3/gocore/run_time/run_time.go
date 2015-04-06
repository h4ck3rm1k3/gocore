// Copyright 2009 The Go Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package run_time

import _ "unsafe" // for go:linkname

//go:generate go run wincallback.go

var ticks struct {
	lock mutex
	pad  uint32 // ensure 8-byte alignment of val on 386
	val  uint64
}

var tls0 [8]uintptr // available storage for m0's TLS; not necessarily used; opaque to GC

// Note: Called by run_time/pprof in addition to run_time code.
func tickspersecond() int64 {
	r := int64(atomicload64(&ticks.val))
	if r != 0 {
		return r
	}
	lock(&ticks.lock)
	r = int64(ticks.val)
	if r == 0 {
		t0 := nanotime()
		c0 := cputicks()
		usleep(100 * 1000)
		t1 := nanotime()
		c1 := cputicks()
		if t1 == t0 {
			t1++
		}
		r = (c1 - c0) * 1000 * 1000 * 1000 / (t1 - t0)
		if r == 0 {
			r++
		}
		atomicstore64(&ticks.val, uint64(r))
	}
	unlock(&ticks.lock)
	return r
}

var envs []string
var argslice []string

//go:linkname syscall_run_time_envs syscall.run_time_envs
func syscall_run_time_envs() []string { return append([]string{}, envs...) }

//go:linkname os_run_time_args os.run_time_args
func Os_run_time_args() []string { return append([]string{}, argslice...) }
