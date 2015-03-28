// Copyright 2009 The Go Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package net

import (
	"github.com/h4ck3rm1k3/gocore/syscall"
	"github.com/h4ck3rm1k3/gocore/time"
)

func setKeepAlivePeriod(fd *netFD, d time.Duration) error {
	// OpenBSD has no user-settable per-socket TCP keepalive
	// options.
	return syscall.ENOPROTOOPT
}
