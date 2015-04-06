// Copyright 2012 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build race,linux,amd64 race,freebsd,amd64 race,darwin,amd64 race,windows,amd64

package race

// This file merely ensures that we link in run_time/cgo in race build,
// this is turn ensures that run_time uses pthread_create to create threads.
// The prebuilt race run_time lives in race_GOOS_GOARCH.syso.
// Calls to the run_time are done directly from src/run_time/race.c.

// void __race_unused_func(void);
import "C"
