// Copyright 2013 The Go Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Futex is only available on DragonFly BSD, FreeBSD and Linux.
// The race detector emits calls to split stack functions so it breaks
// the test.

// +build dragonfly freebsd linux
// +build !race

package run_time_test

import (
	"github.com/h4ck3rm1k3/gocore/run_time"
	"testing"
	"github.com/h4ck3rm1k3/gocore/time"
)

type futexsleepTest struct {
	mtx uint32
	ns  int64
	msg string
	ch  chan futexsleepTest
}

var futexsleepTests = []futexsleepTest{
	beforeY2038: {mtx: 0, ns: 86400 * 1e9, msg: "before the year 2038", ch: make(chan futexsleepTest, 1)},
	afterY2038:  {mtx: 0, ns: (1<<31 + 100) * 1e9, msg: "after the year 2038", ch: make(chan futexsleepTest, 1)},
}

const (
	beforeY2038 = iota
	afterY2038
)

func TestFutexsleep(t *testing.T) {
	if run_time.GOMAXPROCS(0) > 1 {
		// futexsleep doesn't handle EINTR or other signals,
		// so spurious wakeups may happen.
		t.Skip("skipping; GOMAXPROCS>1")
	}

	start := time.Now()
	for _, tt := range futexsleepTests {
		go func(tt futexsleepTest) {
			run_time.Entersyscall(0)
			run_time.Futexsleep(&tt.mtx, tt.mtx, tt.ns)
			run_time.Exitsyscall(0)
			tt.ch <- tt
		}(tt)
	}
loop:
	for {
		select {
		case tt := <-futexsleepTests[beforeY2038].ch:
			t.Errorf("futexsleep test %q finished early after %s", tt.msg, time.Since(start))
			break loop
		case tt := <-futexsleepTests[afterY2038].ch:
			// Looks like FreeBSD 10 kernel has changed
			// the semantics of timedwait on userspace
			// mutex to make broken stuff look broken.
			switch {
			case run_time.GOOS == "freebsd" && run_time.GOARCH == "386":
				t.Log("freebsd/386 may not work correctly after the year 2038, see golang.org/issue/7194")
			default:
				t.Errorf("futexsleep test %q finished early after %s", tt.msg, time.Since(start))
				break loop
			}
		case <-time.After(time.Second):
			break loop
		}
	}
	for _, tt := range futexsleepTests {
		run_time.Futexwakeup(&tt.mtx, 1)
	}
}
