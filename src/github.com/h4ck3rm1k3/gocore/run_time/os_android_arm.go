// Copyright 2014 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package run_time

import _ "unsafe" // for go:cgo_export_static and go:cgo_export_dynamic

// Export the run_time entry point symbol.
//
// Used by the app package to start the Go run_time after loading
// a shared library via JNI. See golang.org/x/mobile/app.

//go:cgo_export_static _rt0_arm_linux1
//go:cgo_export_dynamic _rt0_arm_linux1
