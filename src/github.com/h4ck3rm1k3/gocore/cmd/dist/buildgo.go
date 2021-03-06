// Copyright 2012 The Go Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import "github.com/h4ck3rm1k3/gocore/fmt"

/*
 * Helpers for building cmd/go and cmd/cgo.
 */

// mkzdefaultcc writes zdefaultcc.go:
//
//	package main
//	const defaultCC = <defaultcc>
//	const defaultCXX = <defaultcxx>
//
// It is invoked to write cmd/go/zdefaultcc.go
// but we also write cmd/cgo/zdefaultcc.go
func mkzdefaultcc(dir, file string) {
	var out string

	out = fmt.Sprintf(
		"// auto generated by go tool dist\n"+
			"\n"+
			"package main\n"+
			"\n"+
			"const defaultCC = `%s`\n"+
			"const defaultCXX = `%s`\n",
		defaultcctarget, defaultcxxtarget)

	writefile(out, file, 0)

	// Convert file name to replace: turn go into cgo.
	i := len(file) - len("go/zdefaultcc.go")
	file = file[:i] + "c" + file[i:]
	writefile(out, file, 0)
}
