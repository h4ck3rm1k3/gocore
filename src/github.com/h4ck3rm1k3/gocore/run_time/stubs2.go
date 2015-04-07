// Copyright 2014 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build !plan9
// +build !solaris
// +build !windows
// +build !nacl

package run_time

import "unsafe"
//import "runtime"

func read(fd int32, p unsafe.Pointer, n int32) int32{
	panic("not implemented")
}
func close(fd int32) int32{
	panic("not implemented")
}

func exit(code int32){
	panic("not implemented")
}
func nanotime() int64{
	panic("not implemented")
}
func usleep(usec uint32){
	panic("not implemented")
}

func mmap(addr unsafe.Pointer, n uintptr, prot, flags, fd int32, off uint32) unsafe.Pointer{
	panic("not implemented")
}
func munmap(addr unsafe.Pointer, n uintptr){
	panic("not implemented")
}

//go:noescape
func write(fd uintptr, p unsafe.Pointer, n int32) int32{
	panic("not implemented")
}

//go:noescape
func open(name *byte, mode, perm int32) int32{
	panic("not implemented")
}

func madvise(addr unsafe.Pointer, n uintptr, flags int32){
	panic("not implemented")
}
