// Copyright 2014 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package atomic

import (
	//"runtime"
	"unsafe"
	rtatom "sync/atomic"
)

// A Value provides an atomic load and store of a consistently typed value.
// Values can be created as part of other data structures.
// The zero value for a Value returns nil from Load.
// Once Store has been called, a Value must not be copied.
type Value struct {
	v interface{}
}

// ifaceWords is interface{} internal representation.
type ifaceWords struct {
	typ  unsafe.Pointer
	data unsafe.Pointer
}


// Load returns the value set by the most recent Store.
// It returns nil if there has been no call to Store for this Value.
func (v *Value) Load() (x interface{}) {
	vp := (*ifaceWords)(unsafe.Pointer(v))
	typ := LoadPointer(&vp.typ)
	if typ == nil || uintptr(typ) == ^uintptr(0) {
		// First store not yet completed.
		return nil
	}
	data := LoadPointer(&vp.data)
	xp := (*ifaceWords)(unsafe.Pointer(&x))
	xp.typ = typ
	xp.data = data
	return
}


// Store sets the value of the Value to x.
// All calls to Store for a given Value must use values of the same concrete type.
// Store of an inconsistent type panics, as does Store(nil).
func (v *Value) Store(x interface{}) {
	if x == nil {
		panic("sync/atomic: store of nil value into Value")
	}
	vp := (*ifaceWords)(unsafe.Pointer(v))
	xp := (*ifaceWords)(unsafe.Pointer(&x))
	for {
		typ := LoadPointer(&vp.typ)
		if typ == nil {
			// Attempt to start first store.
			// Disable preemption so that other goroutines can use
			// active spin wait to wait for completion; and so that
			// GC does not see the fake type accidentally.
			run_time_procPin()
			if !CompareAndSwapPointer(&vp.typ, nil, unsafe.Pointer(^uintptr(0))) {
				run_time_procUnpin()
				continue
			}
			// Complete first store.
			StorePointer(&vp.data, xp.data)
			StorePointer(&vp.typ, xp.typ)
			run_time_procUnpin()
			return
		}
		if uintptr(typ) == ^uintptr(0) {
			// First store in progress. Wait.
			// Since we disable preemption around the first store,
			// we can wait with active spinning.
			continue
		}
		// First store completed. Check type and overwrite data.
		if typ != xp.typ {
			panic("sync/atomic: store of inconsistently typed value into Value")
		}
		StorePointer(&vp.data, xp.data)
		return
	}
}

// Disable/enable preemption, implemented in run_time
//func run_time_procPin() { panic("todo") }
func run_time_procPin() {panic("undef")}
func run_time_procUnpin() {panic("undef")}

// hack
func LoadPointer(addr *unsafe.Pointer) (val unsafe.Pointer) {
	return rtatom.LoadPointer(addr)
}

//OAfunc run_time_procUnpin() { panic("todo") }
func StorePointer(addr *unsafe.Pointer, val unsafe.Pointer) {
	rtatom.StorePointer(addr, val)

}
func CompareAndSwapPointer(addr *unsafe.Pointer, old, new unsafe.Pointer) (swapped bool) {
	return rtatom.CompareAndSwapPointer(addr, old, new)
}

func AddInt64(addr *int64, delta int64) (new int64) {
	return rtatom.AddInt64(addr, delta)
}

func SwapUint64(addr *uint64, new uint64) (old uint64) {
	return rtatom.SwapUint64(addr, new)
}

// SwapUintptr atomically stores new into *addr and returns the previous *addr value.
func SwapUintptr(addr *uintptr, new uintptr) (old uintptr) {
	return rtatom.SwapUintptr(addr, new)
}

func SwapUint32(addr *uint32, new uint32) (old uint32) {
	return rtatom.SwapUint32(addr, new)
}

func LoadInt32(addr *int32) (val int32) {
	//return rtatom.SwapUint32(addr, new)
	return rtatom.LoadInt32(addr)
}

// LoadInt64 atomically loads *addr.
func LoadInt64(addr *int64) (val int64) {
	return rtatom.LoadInt64(addr)
}

// LoadUint32 atomically loads *addr.
func LoadUint32(addr *uint32) (val uint32) {
	return rtatom.LoadUint32(addr)
}

// LoadUint64 atomically loads *addr.
func LoadUint64(addr *uint64) (val uint64) {
	return rtatom.LoadUint64(addr)
}

func AddInt32(addr *int32, delta int32) (new int32) {
	return rtatom.AddInt32(addr, delta)
}
func AddUint32(addr *uint32, delta uint32) (new uint32) {
	return rtatom.AddUint32(addr, delta)
}

func AddUint64(addr *uint64, delta uint64) (new uint64) {
	return rtatom.AddUint64(addr,delta)
}
func AddUintptr(addr *uintptr, delta uintptr) (new uintptr) {
	return rtatom.AddUintptr(addr, delta)
}

// CompareAndSwapInt32 executes the compare-and-swap operation for an int32 value.
func CompareAndSwapInt32(addr *int32, old, new int32) (swapped bool) {
	return rtatom.CompareAndSwapInt32(addr,old,new)
}

// CompareAndSwapInt64 executes the compare-and-swap operation for an int64 value.
func CompareAndSwapInt64(addr *int64, old, new int64) (swapped bool) {
	return rtatom.CompareAndSwapInt64(addr,old, new)
}

// CompareAndSwapUint32 executes the compare-and-swap operation for a uint32 value.
func CompareAndSwapUint32(addr *uint32, old, new uint32) (swapped bool) {
	return rtatom.CompareAndSwapUint32(addr, old, new)
}

// CompareAndSwapUint64 executes the compare-and-swap operation for a uint64 value.
func CompareAndSwapUint64(addr *uint64, old, new uint64) (swapped bool) {
	return rtatom.CompareAndSwapUint64(addr, old, new)
}

// CompareAndSwapUintptr executes the compare-and-swap operation for a uintptr value.
func CompareAndSwapUintptr(addr *uintptr, old, new uintptr) (swapped bool) {
	return rtatom.CompareAndSwapUintptr(addr, old, new)
}

func LoadUintptr(addr *uintptr) (val uintptr) { 
	return rtatom.LoadUintptr(addr)
}

func StoreInt32(addr *int32, val int32){
	rtatom.StoreInt32(addr,val)
}

// StoreInt64 atomically stores val into *addr.
func StoreInt64(addr *int64, val int64){
	rtatom.StoreInt64(addr, val)

}

// StoreUint32 atomically stores val into *addr.
func StoreUint32(addr *uint32, val uint32){
	rtatom.StoreUint32(addr, val)

}

// StoreUint64 atomically stores val into *addr.
func StoreUint64(addr *uint64, val uint64){
	rtatom.StoreUint64(addr,val)
}

func SwapInt32(addr *int32, new int32) (old int32){
	return rtatom.SwapInt32(addr,new)
}

// SwapInt64 atomically stores new into *addr and returns the previous *addr value.
func SwapInt64(addr *int64, new int64) (old int64){
	return rtatom.SwapInt64(addr, new)
}

func SwapPointer(addr *unsafe.Pointer, new unsafe.Pointer) (old unsafe.Pointer) {
	return rtatom.SwapPointer(addr, new)
} 



func StoreUintptr(addr *uintptr, val uintptr) {
	rtatom.StoreUintptr(addr, val)
}
