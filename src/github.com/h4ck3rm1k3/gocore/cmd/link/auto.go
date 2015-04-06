// Copyright 2014 The Go Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Automatic symbol generation.

// TODO(rsc): Handle go.typelink, go.track symbols.
// TODO(rsc): Do not handle $f64. and $f32. symbols. Instead, generate those
// from the compiler and assemblers as dupok data, and then remove autoData below.
package main

import (
	"github.com/h4ck3rm1k3/gocore/cmd/internal/goobj"
	"github.com/h4ck3rm1k3/gocore/strconv"
	"github.com/h4ck3rm1k3/gocore/strings"
)

// linkerDefined lists the symbols supplied by other parts of the linker
// (run_time.go and layout.go).
var linkerDefined = map[string]bool{
	"run_time.bss":        true,
	"run_time.data":       true,
	"run_time.ebss":       true,
	"run_time.edata":      true,
	"run_time.efunctab":   true,
	"run_time.end":        true,
	"run_time.enoptrbss":  true,
	"run_time.enoptrdata": true,
	"run_time.erodata":    true,
	"run_time.etext":      true,
	"run_time.etypelink":  true,
	"run_time.functab":    true,
	"run_time.gcbss":      true,
	"run_time.gcdata":     true,
	"run_time.noptrbss":   true,
	"run_time.noptrdata":  true,
	"run_time.pclntab":    true,
	"run_time.rodata":     true,
	"run_time.text":       true,
	"run_time.typelink":   true,
}

// isAuto reports whether sym is an automatically-generated data or constant symbol.
func (p *Prog) isAuto(sym goobj.SymID) bool {
	return strings.HasPrefix(sym.Name, "go.weak.") ||
		strings.HasPrefix(sym.Name, "$f64.") ||
		strings.HasPrefix(sym.Name, "$f32.") ||
		linkerDefined[sym.Name]
}

// autoData defines the automatically generated data symbols needed by p.
func (p *Prog) autoData() {
	for sym := range p.Missing {
		switch {
		// Floating-point constants that need to be loaded from memory are
		// written as $f64.{16 hex digits} or $f32.{8 hex digits}; the hex digits
		// give the IEEE bit pattern of the constant. As far as the layout into
		// memory is concerned, we interpret these as uint64 or uint32 constants.
		case strings.HasPrefix(sym.Name, "$f64."), strings.HasPrefix(sym.Name, "$f32."):
			size := 64
			if sym.Name[2:4] == "32" {
				size = 32
			}
			delete(p.Missing, sym)
			fbits, err := strconv.ParseUint(sym.Name[len("$f64."):], 16, size)
			if err != nil {
				p.errorf("unexpected floating point symbol %s", sym)
				continue
			}
			data := make([]byte, size/8)
			if size == 64 {
				p.byteorder.PutUint64(data, fbits)
			} else {
				p.byteorder.PutUint32(data, uint32(fbits))
			}
			p.addSym(&Sym{
				Sym: &goobj.Sym{
					SymID: sym,
					Kind:  goobj.SRODATA,
					Size:  size / 8,
				},
				Bytes: data,
			})
		}
	}
}

// autoConst defines the automatically generated constant symbols needed by p.
func (p *Prog) autoConst() {
	for sym := range p.Missing {
		switch {
		case strings.HasPrefix(sym.Name, "go.weak."):
			// weak symbol resolves to actual symbol if present, or else nil.
			delete(p.Missing, sym)
			targ := sym
			targ.Name = sym.Name[len("go.weak."):]
			var addr Addr
			if s := p.Syms[targ]; s != nil {
				addr = s.Addr
			}
			p.defineConst(sym.Name, addr)
		}
	}
}

// defineConst defines a new symbol with the given name and constant address.
func (p *Prog) defineConst(name string, addr Addr) {
	sym := goobj.SymID{Name: name}
	p.addSym(&Sym{
		Sym: &goobj.Sym{
			SymID: sym,
			Kind:  goobj.SCONST,
		},
		Package: nil,
		Addr:    addr,
	})
}
