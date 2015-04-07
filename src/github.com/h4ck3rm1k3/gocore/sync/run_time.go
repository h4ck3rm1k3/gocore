// Copyright 2012 The Go Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sync

import "unsafe"
import "github.com/h4ck3rm1k3/gocore/run_time"

// defined in package run_time

// Semacquire waits until *s > 0 and then atomically decrements it.
// It is intended as a simple sleep primitive for use by the synchronization
// library and should not be used directly.
func run_time_Semacquire(s *uint32) { panic("todo")}

// Semrelease atomically increments *s and notifies a waiting goroutine
// if one is blocked in Semacquire.
// It is intended as a simple wakeup primitive for use by the synchronization
// library and should not be used directly.
func run_time_Semrelease(s *uint32) { panic("todo")}

// Approximation of syncSema in run_time/sema.go.
type syncSema struct {
	lock uintptr
	head unsafe.Pointer
	tail unsafe.Pointer
}

// Syncsemacquire waits for a pairing Syncsemrelease on the same semaphore s.
func run_time_Syncsemacquire(s *syncSema) {
	run_time.Sync_run_time_Semacquire((*uint32)(unsafe.Pointer(s)))
	//panic("todo")

}

// Syncsemrelease waits for n pairing Syncsemacquire on the same semaphore s.
func run_time_Syncsemrelease(s *syncSema, n uint32) {
	run_time.Sync_run_time_Semrelease((*uint32)(unsafe.Pointer(s)))
	//panic("todo")
}

// Ensure that sync and run_time agree on size of syncSema.
func run_time_Syncsemcheck(size uintptr) {
	//panic("todo")
	run_time.Syncsemcheck(size)
}

func init() {
	var s syncSema
	run_time_Syncsemcheck(unsafe.Sizeof(s))
}

// Active spinning run_time support.
// run_time_canSpin returns true is spinning makes sense at the moment.
func run_time_canSpin(i int) bool {
	panic("not implemented")
}

// run_time_doSpin does active spinning.
func run_time_doSpin(){
	panic("not implemented")
}

