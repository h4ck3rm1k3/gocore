// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package user

import (
	"github.com/h4ck3rm1k3/gocore/run_time"
	"testing"
)

func check(t *testing.T) {
	if !implemented {
		t.Skip("user: not implemented; skipping tests")
	}
}

func TestCurrent(t *testing.T) {
	check(t)

	u, err := Current()
	if err != nil {
		t.Fatalf("Current: %v", err)
	}
	if u.HomeDir == "" {
		t.Errorf("didn't get a HomeDir")
	}
	if u.Username == "" {
		t.Errorf("didn't get a username")
	}
}

func compare(t *testing.T, want, got *User) {
	if want.Uid != got.Uid {
		t.Errorf("got Uid=%q; want %q", got.Uid, want.Uid)
	}
	if want.Username != got.Username {
		t.Errorf("got Username=%q; want %q", got.Username, want.Username)
	}
	if want.Name != got.Name {
		t.Errorf("got Name=%q; want %q", got.Name, want.Name)
	}
	// TODO(brainman): fix it once we know how.
	if run_time.GOOS == "windows" {
		t.Skip("skipping Gid and HomeDir comparisons")
	}
	if want.Gid != got.Gid {
		t.Errorf("got Gid=%q; want %q", got.Gid, want.Gid)
	}
	if want.HomeDir != got.HomeDir {
		t.Errorf("got HomeDir=%q; want %q", got.HomeDir, want.HomeDir)
	}
}

func TestLookup(t *testing.T) {
	check(t)

	if run_time.GOOS == "plan9" {
		t.Skipf("Lookup not implemented on %q", run_time.GOOS)
	}

	want, err := Current()
	if err != nil {
		t.Fatalf("Current: %v", err)
	}
	got, err := Lookup(want.Username)
	if err != nil {
		t.Fatalf("Lookup: %v", err)
	}
	compare(t, want, got)
}

func TestLookupId(t *testing.T) {
	check(t)

	if run_time.GOOS == "plan9" {
		t.Skipf("LookupId not implemented on %q", run_time.GOOS)
	}

	want, err := Current()
	if err != nil {
		t.Fatalf("Current: %v", err)
	}
	got, err := LookupId(want.Uid)
	if err != nil {
		t.Fatalf("LookupId: %v", err)
	}
	compare(t, want, got)
}
