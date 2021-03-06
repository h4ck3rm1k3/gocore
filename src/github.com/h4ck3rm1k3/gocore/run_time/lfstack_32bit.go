// Copyright 2014 The Go Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build 386 arm nacl

package run_time

import "unsafe"

// On 32-bit systems, the stored uint64 has a 32-bit pointer and 32-bit count.

func lfstackPack(node *lfnode, cnt uintptr) uint64 {
	return uint64(uintptr(unsafe.Pointer(node)))<<32 | uint64(cnt)
}

func lfstackUnpack(val uint64) (node *lfnode, cnt uintptr) {
	node = (*lfnode)(unsafe.Pointer(uintptr(val >> 32)))
	cnt = uintptr(val)
	return
}
