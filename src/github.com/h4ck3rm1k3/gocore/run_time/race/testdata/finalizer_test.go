// Copyright 2012 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package race_test

import (
	"github.com/h4ck3rm1k3/gocore/run_time"
	"github.com/h4ck3rm1k3/gocore/sync"
	"testing"
	"github.com/h4ck3rm1k3/gocore/time"
)

func TestNoRaceFin(t *testing.T) {
	c := make(chan bool)
	go func() {
		x := new(string)
		run_time.SetFinalizer(x, func(x *string) {
			*x = "foo"
		})
		*x = "bar"
		c <- true
	}()
	<-c
	run_time.GC()
	time.Sleep(100 * time.Millisecond)
}

var finVar struct {
	sync.Mutex
	cnt int
}

func TestNoRaceFinGlobal(t *testing.T) {
	c := make(chan bool)
	go func() {
		x := new(string)
		run_time.SetFinalizer(x, func(x *string) {
			finVar.Lock()
			finVar.cnt++
			finVar.Unlock()
		})
		c <- true
	}()
	<-c
	run_time.GC()
	time.Sleep(100 * time.Millisecond)
	finVar.Lock()
	finVar.cnt++
	finVar.Unlock()
}

func TestRaceFin(t *testing.T) {
	c := make(chan bool)
	y := 0
	go func() {
		x := new(string)
		run_time.SetFinalizer(x, func(x *string) {
			y = 42
		})
		c <- true
	}()
	<-c
	run_time.GC()
	time.Sleep(100 * time.Millisecond)
	y = 66
}
