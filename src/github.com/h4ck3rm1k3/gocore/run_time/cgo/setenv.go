// Copyright 2011 The Go Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build darwin dragonfly freebsd linux netbsd openbsd

package cgo

import _ "unsafe" // for go:linkname

//go:cgo_import_static x_cgo_setenv
//go:linkname x_cgo_setenv x_cgo_setenv
//go:linkname _cgo_setenv run_time._cgo_setenv
var x_cgo_setenv byte
var _cgo_setenv = &x_cgo_setenv

//go:cgo_import_static x_cgo_unsetenv
//go:linkname x_cgo_unsetenv x_cgo_unsetenv
//go:linkname _cgo_unsetenv run_time._cgo_unsetenv
var x_cgo_unsetenv byte
var _cgo_unsetenv = &x_cgo_unsetenv
