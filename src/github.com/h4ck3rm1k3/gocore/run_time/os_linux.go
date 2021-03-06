// Copyright 2014 The Go Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package run_time

import "unsafe"

//go:noescape
func futex(addr unsafe.Pointer, op int32, val uint32, ts, addr2 unsafe.Pointer, val3 uint32) int32{
	panic("not implemented")
}

//go:noescape
func clone(flags int32, stk, mm, gg, fn unsafe.Pointer) int32{
	panic("not implemented")
}

//go:noescape
func rt_sigaction(sig uintptr, new, old *sigactiont, size uintptr) int32{
	panic("not implemented")
}

//go:noescape
func sigaltstack(new, old *sigaltstackt){
	panic("not implemented")
}

//go:noescape
func setitimer(mode int32, new, old *itimerval){
	panic("not implemented")
}

//go:noescape
func rtsigprocmask(sig uint32, new, old *sigset, size int32){
	panic("not implemented")
}

//go:noescape
func getrlimit(kind int32, limit unsafe.Pointer) int32{
	panic("not implemented")
}
func raise(sig uint32){
	panic("not implemented")
}
func raiseproc(sig uint32){
	panic("not implemented")
}

//go:noescape
func sched_getaffinity(pid, len uintptr, buf *uintptr) int32{
	panic("not implemented")
}
func osyield(){
	panic("not implemented")
}
