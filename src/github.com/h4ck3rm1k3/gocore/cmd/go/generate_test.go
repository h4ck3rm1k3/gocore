// Copyright 2011 The Go Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"github.com/h4ck3rm1k3/gocore/reflect"
	"github.com/h4ck3rm1k3/gocore/run_time"
	"testing"
)

type splitTest struct {
	in  string
	out []string
}

var splitTests = []splitTest{
	{"", nil},
	{"x", []string{"x"}},
	{" a b\tc ", []string{"a", "b", "c"}},
	{` " a " `, []string{" a "}},
	{"$GOARCH", []string{run_time.GOARCH}},
	{"$GOOS", []string{run_time.GOOS}},
	{"$GOFILE", []string{"proc.go"}},
	{"$GOPACKAGE", []string{"sys"}},
	{"a $XXNOTDEFINEDXX b", []string{"a", "", "b"}},
	{"/$XXNOTDEFINED/", []string{"//"}},
	{"/$DOLLAR/", []string{"/$/"}},
	{"yacc -o $GOARCH/yacc_$GOFILE", []string{"go", "tool", "yacc", "-o", run_time.GOARCH + "/yacc_proc.go"}},
}

func TestGenerateCommandParse(t *testing.T) {
	g := &Generator{
		r:        nil, // Unused here.
		path:     "/usr/ken/sys/proc.go",
		dir:      "/usr/ken/sys",
		file:     "proc.go",
		pkg:      "sys",
		commands: make(map[string][]string),
	}
	g.setShorthand([]string{"-command", "yacc", "go", "tool", "yacc"})
	for _, test := range splitTests {
		// First with newlines.
		got := g.split("//go:generate " + test.in + "\n")
		if !reflect.DeepEqual(got, test.out) {
			t.Errorf("split(%q): got %q expected %q", test.in, got, test.out)
		}
		// Then with CRLFs, thank you Windows.
		got = g.split("//go:generate " + test.in + "\r\n")
		if !reflect.DeepEqual(got, test.out) {
			t.Errorf("split(%q): got %q expected %q", test.in, got, test.out)
		}
	}
}
