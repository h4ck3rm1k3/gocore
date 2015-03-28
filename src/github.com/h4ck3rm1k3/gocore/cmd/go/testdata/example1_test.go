// Copyright 2013 The Go Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Make sure that go test runs Example_Z before Example_A, preserving source order.

package p

import "github.com/h4ck3rm1k3/gocore/fmt"

var n int

func Example_Z() {
	n++
	fmt.Println(n)
	// Output: 1
}

func Example_A() {
	n++
	fmt.Println(n)
	// Output: 2
}
