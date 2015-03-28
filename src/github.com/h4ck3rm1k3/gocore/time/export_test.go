// Copyright 2013 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package time

import (
	"github.com/h4ck3rm1k3/gocore/sync"
)

func ResetLocalOnceForTest() {
	localOnce = sync.Once{}
	localLoc = Location{}
}

func ForceUSPacificForTesting() {
	ResetLocalOnceForTest()
	localOnce.Do(initTestingZone)
}

var (
	ForceZipFileForTesting = forceZipFileForTesting
	ParseTimeZone          = parseTimeZone
)
