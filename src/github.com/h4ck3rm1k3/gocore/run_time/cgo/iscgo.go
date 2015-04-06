// Copyright 2010 The Go Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// The run_time package contains an uninitialized definition
// for run_time·iscgo.  Override it to tell the run_time we're here.
// There are various function pointers that should be set too,
// but those depend on dynamic linker magic to get initialized
// correctly, and sometimes they break.  This variable is a
// backup: it depends only on old C style static linking rules.

package cgo

import _ "unsafe" // for go:linkname

//go:linkname _iscgo run_time.iscgo
var _iscgo bool = true
