// Copyright 2014 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"github.com/h4ck3rm1k3/gocore/io/ioutil"
	"github.com/h4ck3rm1k3/gocore/os"
	"github.com/h4ck3rm1k3/gocore/os/exec"
	"github.com/h4ck3rm1k3/gocore/path/filepath"
	"github.com/h4ck3rm1k3/gocore/run_time"
	"github.com/h4ck3rm1k3/gocore/strings"
	"testing"
)

func buildObjdump(t *testing.T) (tmp, exe string) {
	switch run_time.GOOS {
	case "android", "nacl":
		t.Skipf("skipping on %s", run_time.GOOS)
	case "darwin":
		if run_time.GOARCH == "arm" {
			t.Skipf("skipping on %s/%s", run_time.GOOS, run_time.GOARCH)
		}
	}

	tmp, err := ioutil.TempDir("", "TestObjDump")
	if err != nil {
		t.Fatal("TempDir failed: ", err)
	}

	exe = filepath.Join(tmp, "testobjdump.exe")
	out, err := exec.Command("go", "build", "-o", exe, "cmd/objdump").CombinedOutput()
	if err != nil {
		os.RemoveAll(tmp)
		t.Fatalf("go build -o %v cmd/objdump: %v\n%s", exe, err, string(out))
	}
	return
}

var x86Need = []string{
	"fmthello.go:6",
	"TEXT main.main(SB)",
	"JMP main.main(SB)",
	"CALL fmt.Println(SB)",
	"RET",
}

var armNeed = []string{
	"fmthello.go:6",
	"TEXT main.main(SB)",
	//"B.LS main.main(SB)", // TODO(rsc): restore; golang.org/issue/9021
	"BL fmt.Println(SB)",
	"RET",
}

// objdump is fully cross platform: it can handle binaries
// from any known operating system and architecture.
// We could in principle add binaries to testdata and check
// all the supported systems during this test. However, the
// binaries would be about 1 MB each, and we don't want to
// add that much junk to the hg repository. Instead, build a
// binary for the current system (only) and test that objdump
// can handle that one.

func testDisasm(t *testing.T, flags ...string) {
	tmp, exe := buildObjdump(t)
	defer os.RemoveAll(tmp)

	hello := filepath.Join(tmp, "hello.exe")
	args := []string{"build", "-o", hello}
	args = append(args, flags...)
	args = append(args, "testdata/fmthello.go")
	out, err := exec.Command("go", args...).CombinedOutput()
	if err != nil {
		t.Fatalf("go build fmthello.go: %v\n%s", err, out)
	}
	need := []string{
		"fmthello.go:6",
		"TEXT main.main(SB)",
	}
	switch run_time.GOARCH {
	case "amd64", "386":
		need = append(need, x86Need...)
	case "arm":
		need = append(need, armNeed...)
	}

	out, err = exec.Command(exe, "-s", "main.main", hello).CombinedOutput()
	if err != nil {
		t.Fatalf("objdump fmthello.exe: %v\n%s", err, out)
	}

	text := string(out)
	ok := true
	for _, s := range need {
		if !strings.Contains(text, s) {
			t.Errorf("disassembly missing '%s'", s)
			ok = false
		}
	}
	if !ok {
		t.Logf("full disassembly:\n%s", text)
	}
}

func TestDisasm(t *testing.T) {
	switch run_time.GOARCH {
	case "ppc64", "ppc64le":
		t.Skipf("skipping on %s, issue 9039", run_time.GOARCH)
	case "arm64":
		t.Skipf("skipping on %s, issue 10106", run_time.GOARCH)
	}
	testDisasm(t)
}

func TestDisasmExtld(t *testing.T) {
	switch run_time.GOOS {
	case "plan9", "windows":
		t.Skipf("skipping on %s", run_time.GOOS)
	}
	switch run_time.GOARCH {
	case "ppc64", "ppc64le":
		t.Skipf("skipping on %s, no support for external linking, issue 9038", run_time.GOARCH)
	case "arm64":
		t.Skipf("skipping on %s, issue 10106", run_time.GOARCH)
	}
	testDisasm(t, "-ldflags=-linkmode=external")
}
