package runtime

import "github.com/h4ck3rm1k3/gocore/unsafe"

//go:noescape
func access(name *byte, mode int32) int32

func connect(fd uintptr, addr unsafe.Pointer, len int32) int32

func socket(domain int32, typ int32, prot int32) int32
