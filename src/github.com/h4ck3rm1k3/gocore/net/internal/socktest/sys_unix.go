// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build darwin dragonfly freebsd linux nacl netbsd openbsd solaris

package socktest

import "github.com/h4ck3rm1k3/gocore/syscall"

// Socket wraps syscall.Socket.
func (sw *Switch) Socket(family, sotype, proto int) (s int, err error) {
	so := &Status{Cookie: cookie(family, sotype, proto)}
	sw.fmu.RLock()
	f, _ := sw.fltab[FilterSocket]
	sw.fmu.RUnlock()

	af, err := f.apply(so)
	if err != nil {
		return -1, err
	}
	s, so.Err = syscall.Socket(family, sotype, proto)
	if err = af.apply(so); err != nil {
		if so.Err == nil {
			syscall.Close(s)
		}
		return -1, err
	}

	if so.Err != nil {
		return -1, so.Err
	}
	sw.smu.Lock()
	nso := sw.addLocked(s, family, sotype, proto)
	sw.stats.getLocked(nso.Cookie).Opened++
	sw.smu.Unlock()
	return s, nil
}

// Close wraps syscall.Close.
func (sw *Switch) Close(s int) (err error) {
	so := sw.sockso(s)
	if so == nil {
		return syscall.Close(s)
	}
	sw.fmu.RLock()
	f, _ := sw.fltab[FilterClose]
	sw.fmu.RUnlock()

	af, err := f.apply(so)
	if err != nil {
		return err
	}
	so.Err = syscall.Close(s)
	if err = af.apply(so); err != nil {
		return err
	}

	if so.Err != nil {
		return so.Err
	}
	sw.smu.Lock()
	delete(sw.sotab, s)
	sw.stats.getLocked(so.Cookie).Closed++
	sw.smu.Unlock()
	return nil
}

// Connect wraps syscall.Connect.
func (sw *Switch) Connect(s int, sa syscall.Sockaddr) (err error) {
	so := sw.sockso(s)
	if so == nil {
		return syscall.Connect(s, sa)
	}
	sw.fmu.RLock()
	f, _ := sw.fltab[FilterConnect]
	sw.fmu.RUnlock()

	af, err := f.apply(so)
	if err != nil {
		return err
	}
	so.Err = syscall.Connect(s, sa)
	if err = af.apply(so); err != nil {
		return err
	}

	if so.Err != nil {
		return so.Err
	}
	sw.smu.Lock()
	sw.stats.getLocked(so.Cookie).Connected++
	sw.smu.Unlock()
	return nil
}

// Accept wraps syscall.Accept.
func (sw *Switch) Accept(s int) (ns int, sa syscall.Sockaddr, err error) {
	so := sw.sockso(s)
	if so == nil {
		return syscall.Accept(s)
	}
	sw.fmu.RLock()
	f, _ := sw.fltab[FilterAccept]
	sw.fmu.RUnlock()

	af, err := f.apply(so)
	if err != nil {
		return -1, nil, err
	}
	ns, sa, so.Err = syscall.Accept(s)
	if err = af.apply(so); err != nil {
		if so.Err == nil {
			syscall.Close(ns)
		}
		return -1, nil, err
	}

	if so.Err != nil {
		return -1, nil, so.Err
	}
	sw.smu.Lock()
	nso := sw.addLocked(ns, so.Cookie.Family(), so.Cookie.Type(), so.Cookie.Protocol())
	sw.stats.getLocked(nso.Cookie).Accepted++
	sw.smu.Unlock()
	return ns, sa, nil
}

// GetsockoptInt wraps syscall.GetsockoptInt.
func (sw *Switch) GetsockoptInt(s, level, opt int) (soerr int, err error) {
	so := sw.sockso(s)
	if so == nil {
		return syscall.GetsockoptInt(s, level, opt)
	}
	sw.fmu.RLock()
	f, _ := sw.fltab[FilterGetsockoptInt]
	sw.fmu.RUnlock()

	af, err := f.apply(so)
	if err != nil {
		return -1, err
	}
	so.SocketErr, so.Err = syscall.GetsockoptInt(s, level, opt)
	if err = af.apply(so); err != nil {
		return -1, err
	}

	if so.Err != nil {
		return -1, so.Err
	}
	if opt == syscall.SO_ERROR && (so.SocketErr == 0 || syscall.Errno(so.SocketErr) == syscall.EISCONN) {
		sw.smu.Lock()
		sw.stats.getLocked(so.Cookie).Connected++
		sw.smu.Unlock()
	}
	return so.SocketErr, nil
}