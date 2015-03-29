// Copyright 2014 The Go Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package runtime

import "unsafe"

//go:noescape
func lwp_create(param *lwpparams) int32

//go:noescape
func sigaltstack(new, old *sigaltstackt)

//go:noescape
func sigaction(sig int32, new, old *sigactiont)

//go:noescape
func sigprocmask(new, old *sigset)

//go:noescape
func setitimer(mode int32, new, old *itimerval)

//go:noescape
func sysctl(mib *uint32, miblen uint32, out *byte, size *uintptr, dst *byte, ndst uintptr) int32

//go:noescape
func getrlimit(kind int32, limit unsafe.Pointer) int32

func raise(sig int32)
func raiseproc(sig int32)

//go:noescape
func sys_umtx_sleep(addr *uint32, val, timeout int32) int32

//go:noescape
func sys_umtx_wakeup(addr *uint32, val int32) int32

func osyield()

const stackSystem = 0
