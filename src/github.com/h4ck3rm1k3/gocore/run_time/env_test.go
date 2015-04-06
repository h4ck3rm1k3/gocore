// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package run_time_test

import (
	"github.com/h4ck3rm1k3/gocore/run_time"
	"github.com/h4ck3rm1k3/gocore/syscall"
	"testing"
)

func TestFixedGOROOT(t *testing.T) {
	if run_time.GOOS == "plan9" {
		t.Skipf("skipping plan9, it is inconsistent by allowing GOROOT to be updated by Setenv")
	}

	envs := run_time.Envs()
	oldenvs := append([]string{}, envs...)
	defer run_time.SetEnvs(oldenvs)

	// attempt to reuse existing envs backing array.
	want := run_time.GOROOT()
	run_time.SetEnvs(append(envs[:0], "GOROOT="+want))

	if got := run_time.GOROOT(); got != want {
		t.Errorf(`initial run_time.GOROOT()=%q, want %q`, got, want)
	}
	if err := syscall.Setenv("GOROOT", "/os"); err != nil {
		t.Fatal(err)
	}
	if got := run_time.GOROOT(); got != want {
		t.Errorf(`after setenv run_time.GOROOT()=%q, want %q`, got, want)
	}
	if err := syscall.Unsetenv("GOROOT"); err != nil {
		t.Fatal(err)
	}
	if got := run_time.GOROOT(); got != want {
		t.Errorf(`after unsetenv run_time.GOROOT()=%q, want %q`, got, want)
	}
}
