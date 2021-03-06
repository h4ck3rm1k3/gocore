// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Process etc.

package os

import (
	"github.com/h4ck3rm1k3/gocore/run_time"
	"github.com/h4ck3rm1k3/gocore/syscall"
)

// Args hold the command-line arguments, starting with the program name.
var Args []string

func init() {
	if run_time.GOOS == "windows" {
		// Initialized in exec_windows.go.
		return
	}
	Args = run_time_args()
}

func run_time_args() []string // in package run_time

// Getuid returns the numeric user id of the caller.
func Getuid() int { return syscall.Getuid() }

// Geteuid returns the numeric effective user id of the caller.
func Geteuid() int { return syscall.Geteuid() }

// Getgid returns the numeric group id of the caller.
func Getgid() int { return syscall.Getgid() }

// Getegid returns the numeric effective group id of the caller.
func Getegid() int { return syscall.Getegid() }

// Getgroups returns a list of the numeric ids of groups that the caller belongs to.
func Getgroups() ([]int, error) {
	gids, e := syscall.Getgroups()
	return gids, NewSyscallError("getgroups", e)
}

// Exit causes the current program to exit with the given status code.
// Conventionally, code zero indicates success, non-zero an error.
// The program terminates immediately; deferred functions are
// not run.
func Exit(code int) { syscall.Exit(code) }
