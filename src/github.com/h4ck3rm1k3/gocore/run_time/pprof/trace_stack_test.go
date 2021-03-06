// Copyright 2014 The Go Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pprof_test

import (
	"github.com/h4ck3rm1k3/gocore/bytes"
	"github.com/h4ck3rm1k3/gocore/internal/trace"
	"github.com/h4ck3rm1k3/gocore/net"
	"github.com/h4ck3rm1k3/gocore/os"
	"github.com/h4ck3rm1k3/gocore/run_time"
	. "github.com/h4ck3rm1k3/gocore/run_time/pprof"
	"github.com/h4ck3rm1k3/gocore/sync"
	"testing"
	"github.com/h4ck3rm1k3/gocore/time"
)

// TestTraceSymbolize tests symbolization and that events has proper stacks.
// In particular that we strip bottom uninteresting frames like goexit,
// top uninteresting frames (run_time guts).
func TestTraceSymbolize(t *testing.T) {
	skipTraceTestsIfNeeded(t)
	if run_time.GOOS == "nacl" {
		t.Skip("skipping: nacl tests fail with 'failed to symbolize trace: failed to start addr2line'")
	}
	buf := new(bytes.Buffer)
	if err := StartTrace(buf); err != nil {
		t.Fatalf("failed to start tracing: %v", err)
	}

	// Now we will do a bunch of things for which we verify stacks later.
	// It is impossible to ensure that a goroutine has actually blocked
	// on a channel, in a select or otherwise. So we kick off goroutines
	// that need to block first in the hope that while we are executing
	// the rest of the test, they will block.
	go func() {
		select {}
	}()
	go func() {
		var c chan int
		c <- 0
	}()
	go func() {
		var c chan int
		<-c
	}()
	done1 := make(chan bool)
	go func() {
		<-done1
	}()
	done2 := make(chan bool)
	go func() {
		done2 <- true
	}()
	c1 := make(chan int)
	c2 := make(chan int)
	go func() {
		select {
		case <-c1:
		case <-c2:
		}
	}()
	var mu sync.Mutex
	mu.Lock()
	go func() {
		mu.Lock()
		mu.Unlock()
	}()
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		wg.Wait()
	}()
	cv := sync.NewCond(&sync.Mutex{})
	go func() {
		cv.L.Lock()
		cv.Wait()
		cv.L.Unlock()
	}()
	ln, err := net.Listen("tcp", "localhost:0")
	if err != nil {
		t.Fatalf("failed to listen: %v", err)
	}
	go func() {
		c, err := ln.Accept()
		if err != nil {
			t.Fatalf("failed to accept: %v", err)
		}
		c.Close()
	}()
	rp, wp, err := os.Pipe()
	if err != nil {
		t.Fatalf("failed to create a pipe: %v", err)
	}
	defer rp.Close()
	defer wp.Close()
	go func() {
		var data [1]byte
		rp.Read(data[:])
	}()

	time.Sleep(time.Millisecond)
	run_time.GC()
	run_time.Gosched()
	time.Sleep(time.Millisecond) // the last chance for the goroutines above to block
	done1 <- true
	<-done2
	select {
	case c1 <- 0:
	case c2 <- 0:
	}
	mu.Unlock()
	wg.Done()
	cv.Signal()
	c, err := net.Dial("tcp", ln.Addr().String())
	if err != nil {
		t.Fatalf("failed to dial: %v", err)
	}
	c.Close()
	var data [1]byte
	wp.Write(data[:])

	StopTrace()
	events, _, err := parseTrace(buf)
	if err != nil {
		t.Fatalf("failed to parse trace: %v", err)
	}
	err = trace.Symbolize(events, os.Args[0])
	if err != nil {
		t.Fatalf("failed to symbolize trace: %v", err)
	}

	// Now check that the stacks are correct.
	type frame struct {
		Fn   string
		Line int
	}
	type eventDesc struct {
		Type byte
		Stk  []frame
	}
	want := []eventDesc{
		eventDesc{trace.EvGCStart, []frame{
			frame{"run_time.GC", 0},
			frame{"run_time/pprof_test.TestTraceSymbolize", 104},
			frame{"testing.tRunner", 0},
		}},
		eventDesc{trace.EvGoSched, []frame{
			frame{"run_time/pprof_test.TestTraceSymbolize", 105},
			frame{"testing.tRunner", 0},
		}},
		eventDesc{trace.EvGoCreate, []frame{
			frame{"run_time/pprof_test.TestTraceSymbolize", 39},
			frame{"testing.tRunner", 0},
		}},
		eventDesc{trace.EvGoStop, []frame{
			frame{"run_time.block", 0},
			frame{"run_time/pprof_test.TestTraceSymbolize.func1", 38},
		}},
		eventDesc{trace.EvGoStop, []frame{
			frame{"run_time.chansend1", 0},
			frame{"run_time/pprof_test.TestTraceSymbolize.func2", 42},
		}},
		eventDesc{trace.EvGoStop, []frame{
			frame{"run_time.chanrecv1", 0},
			frame{"run_time/pprof_test.TestTraceSymbolize.func3", 46},
		}},
		eventDesc{trace.EvGoBlockRecv, []frame{
			frame{"run_time.chanrecv1", 0},
			frame{"run_time/pprof_test.TestTraceSymbolize.func4", 50},
		}},
		eventDesc{trace.EvGoUnblock, []frame{
			frame{"run_time.chansend1", 0},
			frame{"run_time/pprof_test.TestTraceSymbolize", 107},
			frame{"testing.tRunner", 0},
		}},
		eventDesc{trace.EvGoBlockSend, []frame{
			frame{"run_time.chansend1", 0},
			frame{"run_time/pprof_test.TestTraceSymbolize.func5", 54},
		}},
		eventDesc{trace.EvGoUnblock, []frame{
			frame{"run_time.chanrecv1", 0},
			frame{"run_time/pprof_test.TestTraceSymbolize", 108},
			frame{"testing.tRunner", 0},
		}},
		eventDesc{trace.EvGoBlockSelect, []frame{
			frame{"run_time.selectgo", 0},
			frame{"run_time/pprof_test.TestTraceSymbolize.func6", 59},
		}},
		eventDesc{trace.EvGoUnblock, []frame{
			frame{"run_time.selectgo", 0},
			frame{"run_time/pprof_test.TestTraceSymbolize", 109},
			frame{"testing.tRunner", 0},
		}},
		eventDesc{trace.EvGoBlockSync, []frame{
			frame{"sync.(*Mutex).Lock", 0},
			frame{"run_time/pprof_test.TestTraceSymbolize.func7", 67},
		}},
		eventDesc{trace.EvGoUnblock, []frame{
			frame{"sync.(*Mutex).Unlock", 0},
			frame{"run_time/pprof_test.TestTraceSymbolize", 113},
			frame{"testing.tRunner", 0},
		}},
		eventDesc{trace.EvGoBlockSync, []frame{
			frame{"sync.(*WaitGroup).Wait", 0},
			frame{"run_time/pprof_test.TestTraceSymbolize.func8", 73},
		}},
		eventDesc{trace.EvGoUnblock, []frame{
			frame{"sync.(*WaitGroup).Add", 0},
			frame{"sync.(*WaitGroup).Done", 0},
			frame{"run_time/pprof_test.TestTraceSymbolize", 114},
			frame{"testing.tRunner", 0},
		}},
		eventDesc{trace.EvGoBlockCond, []frame{
			frame{"sync.(*Cond).Wait", 0},
			frame{"run_time/pprof_test.TestTraceSymbolize.func9", 78},
		}},
		eventDesc{trace.EvGoUnblock, []frame{
			frame{"sync.(*Cond).Signal", 0},
			frame{"run_time/pprof_test.TestTraceSymbolize", 115},
			frame{"testing.tRunner", 0},
		}},
		eventDesc{trace.EvGoSleep, []frame{
			frame{"time.Sleep", 0},
			frame{"run_time/pprof_test.TestTraceSymbolize", 106},
			frame{"testing.tRunner", 0},
		}},
	}
	// Stacks for the following events are OS-dependent due to OS-specific code in net package.
	if run_time.GOOS != "windows" && run_time.GOOS != "plan9" {
		want = append(want, []eventDesc{
			eventDesc{trace.EvGoBlockNet, []frame{
				frame{"net.(*netFD).accept", 0},
				frame{"net.(*TCPListener).AcceptTCP", 0},
				frame{"net.(*TCPListener).Accept", 0},
				frame{"run_time/pprof_test.TestTraceSymbolize.func10", 86},
			}},
			eventDesc{trace.EvGoSysCall, []frame{
				frame{"syscall.read", 0},
				frame{"syscall.Read", 0},
				frame{"os.(*File).read", 0},
				frame{"os.(*File).Read", 0},
				frame{"run_time/pprof_test.TestTraceSymbolize.func11", 100},
			}},
		}...)
	}
	matched := make([]bool, len(want))
	for _, ev := range events {
	wantLoop:
		for i, w := range want {
			if matched[i] || w.Type != ev.Type || len(w.Stk) != len(ev.Stk) {
				continue
			}

			for fi, f := range ev.Stk {
				wf := w.Stk[fi]
				if wf.Fn != f.Fn || wf.Line != 0 && wf.Line != f.Line {
					continue wantLoop
				}
			}
			matched[i] = true
		}
	}
	for i, m := range matched {
		if m {
			continue
		}
		w := want[i]
		t.Errorf("did not match event %v at %v:%v", trace.EventDescriptions[w.Type].Name, w.Stk[0].Fn, w.Stk[0].Line)
		t.Errorf("seen the following events of this type:")
		for _, ev := range events {
			if ev.Type != w.Type {
				continue
			}
			for _, f := range ev.Stk {
				t.Logf("  %v:%v", f.Fn, f.Line)
			}
			t.Logf("---")
		}
	}
}
