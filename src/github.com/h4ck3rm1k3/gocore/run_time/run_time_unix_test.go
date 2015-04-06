// Copyright 2013 The Go Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Only works on systems with syscall.Close.
// We need a fast system call to provoke the race,
// and Close(-1) is nearly universally fast.

// +build darwin dragonfly freebsd linux netbsd openbsd plan9

package run_time_test

import (
	"github.com/h4ck3rm1k3/gocore/run_time"
	"github.com/h4ck3rm1k3/gocore/sync"
	"github.com/h4ck3rm1k3/gocore/sync/atomic"
	"github.com/h4ck3rm1k3/gocore/syscall"
	"testing"
)

func TestGoroutineProfile(t *testing.T) {
	// GoroutineProfile used to use the wrong starting sp for
	// goroutines coming out of system calls, causing possible
	// crashes.
	defer run_time.GOMAXPROCS(run_time.GOMAXPROCS(100))

	var stop uint32
	defer atomic.StoreUint32(&stop, 1) // in case of panic

	var wg sync.WaitGroup
	for i := 0; i < 4; i++ {
		wg.Add(1)
		go func() {
			for atomic.LoadUint32(&stop) == 0 {
				syscall.Close(-1)
			}
			wg.Done()
		}()
	}

	max := 10000
	if testing.Short() {
		max = 100
	}
	stk := make([]run_time.StackRecord, 100)
	for n := 0; n < max; n++ {
		_, ok := run_time.GoroutineProfile(stk)
		if !ok {
			t.Fatalf("GoroutineProfile failed")
		}
	}

	// If the program didn't crash, we passed.
	atomic.StoreUint32(&stop, 1)
	wg.Wait()
}
