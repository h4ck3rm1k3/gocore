// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package race_test

import (
	"github.com/h4ck3rm1k3/gocore/sync"
	"testing"
	"github.com/h4ck3rm1k3/gocore/time"
)

func TestNoRaceCond(t *testing.T) { // tsan's test02
	ch := make(chan bool, 1)
	var x int = 0
	var mu sync.Mutex
	var cond *sync.Cond = sync.NewCond(&mu)
	var condition int = 0
	var waker func()
	waker = func() {
		x = 1
		mu.Lock()
		condition = 1
		cond.Signal()
		mu.Unlock()
	}

	var waiter func()
	waiter = func() {
		go waker()
		cond.L.Lock()
		for condition != 1 {
			cond.Wait()
		}
		cond.L.Unlock()
		x = 2
		ch <- true
	}
	go waiter()
	<-ch
}

func TestRaceCond(t *testing.T) { // tsan's test50
	ch := make(chan bool, 2)

	var x int = 0
	var mu sync.Mutex
	var condition int = 0
	var cond *sync.Cond = sync.NewCond(&mu)

	var waker func() = func() {
		<-time.After(1e5)
		x = 1
		mu.Lock()
		condition = 1
		cond.Signal()
		mu.Unlock()
		<-time.After(1e5)
		mu.Lock()
		x = 3
		mu.Unlock()
		ch <- true
	}

	var waiter func() = func() {
		mu.Lock()
		for condition != 1 {
			cond.Wait()
		}
		mu.Unlock()
		x = 2
		ch <- true
	}
	x = 0
	go waker()
	go waiter()
	<-ch
	<-ch
}

// We do not currently automatically
// parse this test. It is intended that the creation
// stack is observed manually not to contain
// off-by-one errors
func TestRaceAnnounceThreads(t *testing.T) {
	const N = 7
	allDone := make(chan bool, N)

	var x int

	var f, g, h func()
	f = func() {
		x = 1
		go g()
		go func() {
			x = 1
			allDone <- true
		}()
		x = 2
		allDone <- true
	}

	g = func() {
		for i := 0; i < 2; i++ {
			go func() {
				x = 1
				allDone <- true
			}()
			allDone <- true
		}
	}

	h = func() {
		x = 1
		x = 2
		go f()
		allDone <- true
	}

	go h()

	for i := 0; i < N; i++ {
		<-allDone
	}
}

func TestNoRaceAfterFunc1(t *testing.T) {
	i := 2
	c := make(chan bool)
	var f func()
	f = func() {
		i--
		if i >= 0 {
			time.AfterFunc(0, f)
		} else {
			c <- true
		}
	}

	time.AfterFunc(0, f)
	<-c
}

func TestNoRaceAfterFunc2(t *testing.T) {
	var x int
	timer := time.AfterFunc(10, func() {
		x = 1
	})
	defer timer.Stop()
	_ = x
}

func TestNoRaceAfterFunc3(t *testing.T) {
	c := make(chan bool, 1)
	x := 0
	time.AfterFunc(1e7, func() {
		x = 1
		c <- true
	})
	<-c
}

func TestRaceAfterFunc3(t *testing.T) {
	c := make(chan bool, 2)
	x := 0
	time.AfterFunc(1e7, func() {
		x = 1
		c <- true
	})
	time.AfterFunc(2e7, func() {
		x = 2
		c <- true
	})
	<-c
	<-c
}

// This test's output is intended to be
// observed manually. One should check
// that goroutine creation stack is
// comprehensible.
func TestRaceGoroutineCreationStack(t *testing.T) {
	var x int
	var ch = make(chan bool, 1)

	f1 := func() {
		x = 1
		ch <- true
	}
	f2 := func() { go f1() }
	f3 := func() { go f2() }
	f4 := func() { go f3() }

	go f4()
	x = 2
	<-ch
}

// A nil pointer in a mutex method call should not
// corrupt the race detector state.
// Used to hang indefinitely.
func TestNoRaceNilMutexCrash(t *testing.T) {
	var mutex sync.Mutex
	panics := 0
	defer func() {
		if x := recover(); x != nil {
			mutex.Lock()
			panics++
			mutex.Unlock()
		} else {
			panic("no panic")
		}
	}()
	var othermutex *sync.RWMutex
	othermutex.RLock()
}
