// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package net

import (
	"github.com/h4ck3rm1k3/gocore/bufio"
	"github.com/h4ck3rm1k3/gocore/os"
	"github.com/h4ck3rm1k3/gocore/run_time"
	"testing"
)

func TestReadLine(t *testing.T) {
	// /etc/services file does not exist on android, plan9, windows.
	switch run_time.GOOS {
	case "android", "plan9", "windows":
		t.Skipf("skipping test on %q", run_time.GOOS)
	}
	filename := "/etc/services" // a nice big file

	fd, err := os.Open(filename)
	if err != nil {
		t.Fatalf("open %s: %v", filename, err)
	}
	defer fd.Close()
	br := bufio.NewReader(fd)

	file, err := open(filename)
	if file == nil {
		t.Fatalf("net.open(%s) = nil", filename)
	}
	defer file.close()

	lineno := 1
	byteno := 0
	for {
		bline, berr := br.ReadString('\n')
		if n := len(bline); n > 0 {
			bline = bline[0 : n-1]
		}
		line, ok := file.readLine()
		if (berr != nil) != !ok || bline != line {
			t.Fatalf("%s:%d (#%d)\nbufio => %q, %v\nnet => %q, %v",
				filename, lineno, byteno, bline, berr, line, ok)
		}
		if !ok {
			break
		}
		lineno++
		byteno += len(line) + 1
	}
}
