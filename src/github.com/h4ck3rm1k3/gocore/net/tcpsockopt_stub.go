// Copyright 2009 The Go Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build nacl

package net

import (
	"github.com/h4ck3rm1k3/gocore/syscall"
	"github.com/h4ck3rm1k3/gocore/time"
)

func setNoDelay(fd *netFD, noDelay bool) error {
	return syscall.ENOPROTOOPT
}

func setKeepAlivePeriod(fd *netFD, d time.Duration) error {
	return syscall.ENOPROTOOPT
}
