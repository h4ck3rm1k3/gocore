// Copyright 2013 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package mime

import (
	"github.com/h4ck3rm1k3/gocore/bufio"
	"github.com/h4ck3rm1k3/gocore/os"
	"github.com/h4ck3rm1k3/gocore/strings"
)

var typeFiles = []string{
	"/sys/lib/mimetypes",
}

func loadMimeFile(filename string) {
	f, err := os.Open(filename)
	if err != nil {
		return
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		fields := strings.Fields(scanner.Text())
		if len(fields) <= 2 || fields[0][0] != '.' {
			continue
		}
		if fields[1] == "-" || fields[2] == "-" {
			continue
		}
		setExtensionType(fields[0], fields[1]+"/"+fields[2])
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}
}

func initMime() {
	for _, filename := range typeFiles {
		loadMimeFile(filename)
	}
}

func initMimeForTests() map[string]string {
	typeFiles = []string{"testdata/test.types.plan9"}
	return map[string]string{
		".t1":  "application/test",
		".t2":  "text/test; charset=utf-8",
		".pNg": "image/png",
	}
}
