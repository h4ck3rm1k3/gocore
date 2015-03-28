// Copyright 2013 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package os

import "github.com/h4ck3rm1k3/gocore/syscall"

func hostname() (name string, err error) {
	return syscall.Gethostname()
}
