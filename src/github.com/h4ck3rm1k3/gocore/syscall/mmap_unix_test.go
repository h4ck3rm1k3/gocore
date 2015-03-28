// Copyright 2014 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build darwin dragonfly freebsd linux netbsd openbsd

package syscall_test

import (
	"github.com/h4ck3rm1k3/gocore/syscall"
	"testing"
)

func TestMmap(t *testing.T) {
	b, err := syscall.Mmap(-1, 0, syscall.Getpagesize(), syscall.PROT_NONE, syscall.MAP_ANON|syscall.MAP_PRIVATE)
	if err != nil {
		t.Fatalf("Mmap: %v", err)
	}
	if err := syscall.Munmap(b); err != nil {
		t.Fatalf("Munmap: %v", err)
	}
}
