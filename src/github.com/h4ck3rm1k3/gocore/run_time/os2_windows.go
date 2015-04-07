// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package run_time

import "unsafe"

// Call a Windows function with stdcall conventions,
// and switch to os stack during the call.
func asmstdcall(fn unsafe.Pointer){
	panic("not implemented")
}

func getlasterror() uint32{
	panic("not implemented")
}
func setlasterror(err uint32){
	panic("not implemented")
}

// Function to be called by windows CreateThread
// to start new os thread.
func tstart_stdcall(newm *m) uint32{
	panic("not implemented")
}

func ctrlhandler(_type uint32) uint32{
	panic("not implemented")
}

// TODO(brainman): should not need those
const (
	_NSIG = 65
)
