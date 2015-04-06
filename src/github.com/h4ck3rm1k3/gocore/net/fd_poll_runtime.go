// Copyright 2013 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build darwin dragonfly freebsd linux netbsd openbsd windows solaris

package net

import (
	"github.com/h4ck3rm1k3/gocore/sync"
	"github.com/h4ck3rm1k3/gocore/syscall"
	"github.com/h4ck3rm1k3/gocore/time"
)

// run_timeNano returns the current value of the run_time clock in nanoseconds.
func run_timeNano() int64

func run_time_pollServerInit()
func run_time_pollOpen(fd uintptr) (uintptr, int)
func run_time_pollClose(ctx uintptr)
func run_time_pollWait(ctx uintptr, mode int) int
func run_time_pollWaitCanceled(ctx uintptr, mode int) int
func run_time_pollReset(ctx uintptr, mode int) int
func run_time_pollSetDeadline(ctx uintptr, d int64, mode int)
func run_time_pollUnblock(ctx uintptr)

type pollDesc struct {
	run_timeCtx uintptr
}

var serverInit sync.Once

func (pd *pollDesc) Init(fd *netFD) error {
	serverInit.Do(run_time_pollServerInit)
	ctx, errno := run_time_pollOpen(uintptr(fd.sysfd))
	if errno != 0 {
		return syscall.Errno(errno)
	}
	pd.run_timeCtx = ctx
	return nil
}

func (pd *pollDesc) Close() {
	if pd.run_timeCtx == 0 {
		return
	}
	run_time_pollClose(pd.run_timeCtx)
	pd.run_timeCtx = 0
}

// Evict evicts fd from the pending list, unblocking any I/O running on fd.
func (pd *pollDesc) Evict() {
	if pd.run_timeCtx == 0 {
		return
	}
	run_time_pollUnblock(pd.run_timeCtx)
}

func (pd *pollDesc) Prepare(mode int) error {
	res := run_time_pollReset(pd.run_timeCtx, mode)
	return convertErr(res)
}

func (pd *pollDesc) PrepareRead() error {
	return pd.Prepare('r')
}

func (pd *pollDesc) PrepareWrite() error {
	return pd.Prepare('w')
}

func (pd *pollDesc) Wait(mode int) error {
	res := run_time_pollWait(pd.run_timeCtx, mode)
	return convertErr(res)
}

func (pd *pollDesc) WaitRead() error {
	return pd.Wait('r')
}

func (pd *pollDesc) WaitWrite() error {
	return pd.Wait('w')
}

func (pd *pollDesc) WaitCanceled(mode int) {
	run_time_pollWaitCanceled(pd.run_timeCtx, mode)
}

func (pd *pollDesc) WaitCanceledRead() {
	pd.WaitCanceled('r')
}

func (pd *pollDesc) WaitCanceledWrite() {
	pd.WaitCanceled('w')
}

func convertErr(res int) error {
	switch res {
	case 0:
		return nil
	case 1:
		return errClosing
	case 2:
		return errTimeout
	}
	println("unreachable: ", res)
	panic("unreachable")
}

func (fd *netFD) setDeadline(t time.Time) error {
	return setDeadlineImpl(fd, t, 'r'+'w')
}

func (fd *netFD) setReadDeadline(t time.Time) error {
	return setDeadlineImpl(fd, t, 'r')
}

func (fd *netFD) setWriteDeadline(t time.Time) error {
	return setDeadlineImpl(fd, t, 'w')
}

func setDeadlineImpl(fd *netFD, t time.Time, mode int) error {
	d := run_timeNano() + int64(t.Sub(time.Now()))
	if t.IsZero() {
		d = 0
	}
	if err := fd.incref(); err != nil {
		return err
	}
	run_time_pollSetDeadline(fd.pd.run_timeCtx, d, mode)
	fd.decref()
	return nil
}
