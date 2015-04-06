// Copyright 2014 The Go Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Generation of run_time-accessible data structures.
// See also debug.go.

package main

import "github.com/h4ck3rm1k3/gocore/cmd/internal/goobj"

func (p *Prog) run_time() {
	p.pclntab()

	// TODO: Implement garbage collection data.
	p.addSym(&Sym{
		Sym: &goobj.Sym{
			SymID: goobj.SymID{Name: "run_time.gcdata"},
			Kind:  goobj.SRODATA,
		},
	})
	p.addSym(&Sym{
		Sym: &goobj.Sym{
			SymID: goobj.SymID{Name: "run_time.gcbss"},
			Kind:  goobj.SRODATA,
		},
	})
}
