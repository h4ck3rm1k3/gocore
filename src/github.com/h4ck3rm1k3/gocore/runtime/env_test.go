// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package runtime_test

import (
	"github.com/h4ck3rm1k3/gocore/runtime"
	"github.com/h4ck3rm1k3/gocore/syscall"
	"testing"
)

func TestFixedGOROOT(t *testing.T) {
	if runtime.GOOS == "plan9" {
		t.Skipf("skipping plan9, it is inconsistent by allowing GOROOT to be updated by Setenv")
	}

	envs := runtime.Envs()
	oldenvs := append([]string{}, envs...)
	defer runtime.SetEnvs(oldenvs)

	// attempt to reuse existing envs backing array.
	want := runtime.GOROOT()
	runtime.SetEnvs(append(envs[:0], "GOROOT="+want))

	if got := runtime.GOROOT(); got != want {
		t.Errorf(`initial runtime.GOROOT()=%q, want %q`, got, want)
	}
	if err := syscall.Setenv("GOROOT", "/os"); err != nil {
		t.Fatal(err)
	}
	if got := runtime.GOROOT(); got != want {
		t.Errorf(`after setenv runtime.GOROOT()=%q, want %q`, got, want)
	}
	if err := syscall.Unsetenv("GOROOT"); err != nil {
		t.Fatal(err)
	}
	if got := runtime.GOROOT(); got != want {
		t.Errorf(`after unsetenv runtime.GOROOT()=%q, want %q`, got, want)
	}
}
