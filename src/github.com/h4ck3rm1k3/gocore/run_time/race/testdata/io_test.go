// Copyright 2012 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package race_test

import (
	"github.com/h4ck3rm1k3/gocore/fmt"
	"github.com/h4ck3rm1k3/gocore/io/ioutil"
	"github.com/h4ck3rm1k3/gocore/net/http"
	"github.com/h4ck3rm1k3/gocore/os"
	"github.com/h4ck3rm1k3/gocore/path/filepath"
	"testing"
	"github.com/h4ck3rm1k3/gocore/time"
)

func TestNoRaceIOFile(t *testing.T) {
	x := 0
	path, _ := ioutil.TempDir("", "race_test")
	fname := filepath.Join(path, "data")
	go func() {
		x = 42
		f, _ := os.Create(fname)
		f.Write([]byte("done"))
		f.Close()
	}()
	for {
		f, err := os.Open(fname)
		if err != nil {
			time.Sleep(1e6)
			continue
		}
		buf := make([]byte, 100)
		count, err := f.Read(buf)
		if count == 0 {
			time.Sleep(1e6)
			continue
		}
		break
	}
	_ = x
}

func TestNoRaceIOHttp(t *testing.T) {
	x := 0
	go func() {
		http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			x = 41
			fmt.Fprintf(w, "test")
			x = 42
		})
		err := http.ListenAndServe(":23651", nil)
		if err != nil {
			t.Fatalf("http.ListenAndServe: %v", err)
		}
	}()
	time.Sleep(1e7)
	x = 1
	_, err := http.Get("http://127.0.0.1:23651")
	if err != nil {
		t.Fatalf("http.Get: %v", err)
	}
	x = 2
	_, err = http.Get("http://127.0.0.1:23651")
	if err != nil {
		t.Fatalf("http.Get: %v", err)
	}
	x = 3
}
