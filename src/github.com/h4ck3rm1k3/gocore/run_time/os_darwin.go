// Copyright 2014 The Go Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package run_time

import "unsafe"

func bsdthread_create(stk unsafe.Pointer, mm *m, gg *g, fn uintptr) int32{
	panic("not implemented")
}
func bsdthread_register() int32{
	panic("not implemented")
}

//go:noescape
func mach_msg_trap(h unsafe.Pointer, op int32, send_size, rcv_size, rcv_name, timeout, notify uint32) int32{
	panic("not implemented")
}

func mach_reply_port() uint32{
	panic("not implemented")
}
func mach_task_self() uint32{
	panic("not implemented")
}
func mach_thread_self() uint32{
	panic("not implemented")
}

//go:noescape
func sysctl(mib *uint32, miblen uint32, out *byte, size *uintptr, dst *byte, ndst uintptr) int32{
	panic("not implemented")
}

//go:noescape
func sigprocmask(sig uint32, new, old *uint32){
	panic("not implemented")
}

//go:noescape
func sigaction(mode uint32, new, old *sigactiont){
	panic("not implemented")
}

//go:noescape
func sigaltstack(new, old *stackt){
	panic("not implemented")
}

func sigtramp(){
	panic("not implemented")
}

//go:noescape
func setitimer(mode int32, new, old *itimerval){
	panic("not implemented")
}

func raise(int32){
	panic("not implemented")
}
func raiseproc(int32){
	panic("not implemented")
}
