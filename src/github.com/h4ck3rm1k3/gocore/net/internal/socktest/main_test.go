// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build !plan9

package socktest_test

import (
	"net/internal/socktest"
	"github.com/h4ck3rm1k3/gocore/os"
	"github.com/h4ck3rm1k3/gocore/syscall"
	"testing"
)

var sw socktest.Switch

func TestMain(m *testing.M) {
	installTestHooks()

	st := m.Run()

	for s := range sw.Sockets() {
		closeFunc(s)
	}
	uninstallTestHooks()
	os.Exit(st)
}

func TestSocket(t *testing.T) {
	for _, f := range []socktest.Filter{
		func(st *socktest.Status) (socktest.AfterFilter, error) { return nil, nil },
		nil,
	} {
		sw.Set(socktest.FilterSocket, f)
		for _, family := range []int{syscall.AF_INET, syscall.AF_INET6} {
			socketFunc(family, syscall.SOCK_STREAM, syscall.IPPROTO_TCP)
		}
	}
}
