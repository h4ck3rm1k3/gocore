// Copyright 2011 The Go Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import "github.com/h4ck3rm1k3/gocore/go/ast"

func init() {
	addTestCases(importTests, nil)
}

var importTests = []testCase{
	{
		Name: "import.0",
		Fn:   addImportFn("github.com/h4ck3rm1k3/gocore/os"),
		In: `package main

import (
	"github.com/h4ck3rm1k3/gocore/os"
)
`,
		Out: `package main

import (
	"github.com/h4ck3rm1k3/gocore/os"
)
`,
	},
	{
		Name: "import.1",
		Fn:   addImportFn("github.com/h4ck3rm1k3/gocore/os"),
		In: `package main
`,
		Out: `package main

import "github.com/h4ck3rm1k3/gocore/os"
`,
	},
	{
		Name: "import.2",
		Fn:   addImportFn("github.com/h4ck3rm1k3/gocore/os"),
		In: `package main

// Comment
import "C"
`,
		Out: `package main

// Comment
import "C"
import "github.com/h4ck3rm1k3/gocore/os"
`,
	},
	{
		Name: "import.3",
		Fn:   addImportFn("github.com/h4ck3rm1k3/gocore/os"),
		In: `package main

// Comment
import "C"

import (
	"github.com/h4ck3rm1k3/gocore/io"
	"utf8"
)
`,
		Out: `package main

// Comment
import "C"

import (
	"github.com/h4ck3rm1k3/gocore/io"
	"github.com/h4ck3rm1k3/gocore/os"
	"utf8"
)
`,
	},
	{
		Name: "import.4",
		Fn:   deleteImportFn("github.com/h4ck3rm1k3/gocore/os"),
		In: `package main

import (
	"github.com/h4ck3rm1k3/gocore/os"
)
`,
		Out: `package main
`,
	},
	{
		Name: "import.5",
		Fn:   deleteImportFn("github.com/h4ck3rm1k3/gocore/os"),
		In: `package main

// Comment
import "C"
import "github.com/h4ck3rm1k3/gocore/os"
`,
		Out: `package main

// Comment
import "C"
`,
	},
	{
		Name: "import.6",
		Fn:   deleteImportFn("github.com/h4ck3rm1k3/gocore/os"),
		In: `package main

// Comment
import "C"

import (
	"github.com/h4ck3rm1k3/gocore/io"
	"github.com/h4ck3rm1k3/gocore/os"
	"utf8"
)
`,
		Out: `package main

// Comment
import "C"

import (
	"github.com/h4ck3rm1k3/gocore/io"
	"utf8"
)
`,
	},
	{
		Name: "import.7",
		Fn:   deleteImportFn("github.com/h4ck3rm1k3/gocore/io"),
		In: `package main

import (
	"github.com/h4ck3rm1k3/gocore/io"   // a
	"github.com/h4ck3rm1k3/gocore/os"   // b
	"utf8" // c
)
`,
		Out: `package main

import (
	// a
	"github.com/h4ck3rm1k3/gocore/os"   // b
	"utf8" // c
)
`,
	},
	{
		Name: "import.8",
		Fn:   deleteImportFn("github.com/h4ck3rm1k3/gocore/os"),
		In: `package main

import (
	"github.com/h4ck3rm1k3/gocore/io"   // a
	"github.com/h4ck3rm1k3/gocore/os"   // b
	"utf8" // c
)
`,
		Out: `package main

import (
	"github.com/h4ck3rm1k3/gocore/io" // a
	// b
	"utf8" // c
)
`,
	},
	{
		Name: "import.9",
		Fn:   deleteImportFn("utf8"),
		In: `package main

import (
	"github.com/h4ck3rm1k3/gocore/io"   // a
	"github.com/h4ck3rm1k3/gocore/os"   // b
	"utf8" // c
)
`,
		Out: `package main

import (
	"github.com/h4ck3rm1k3/gocore/io" // a
	"github.com/h4ck3rm1k3/gocore/os" // b
	// c
)
`,
	},
	{
		Name: "import.10",
		Fn:   deleteImportFn("github.com/h4ck3rm1k3/gocore/io"),
		In: `package main

import (
	"github.com/h4ck3rm1k3/gocore/io"
	"github.com/h4ck3rm1k3/gocore/os"
	"utf8"
)
`,
		Out: `package main

import (
	"github.com/h4ck3rm1k3/gocore/os"
	"utf8"
)
`,
	},
	{
		Name: "import.11",
		Fn:   deleteImportFn("github.com/h4ck3rm1k3/gocore/os"),
		In: `package main

import (
	"github.com/h4ck3rm1k3/gocore/io"
	"github.com/h4ck3rm1k3/gocore/os"
	"utf8"
)
`,
		Out: `package main

import (
	"github.com/h4ck3rm1k3/gocore/io"
	"utf8"
)
`,
	},
	{
		Name: "import.12",
		Fn:   deleteImportFn("utf8"),
		In: `package main

import (
	"github.com/h4ck3rm1k3/gocore/io"
	"github.com/h4ck3rm1k3/gocore/os"
	"utf8"
)
`,
		Out: `package main

import (
	"github.com/h4ck3rm1k3/gocore/io"
	"github.com/h4ck3rm1k3/gocore/os"
)
`,
	},
	{
		Name: "import.13",
		Fn:   rewriteImportFn("utf8", "encoding/utf8"),
		In: `package main

import (
	"github.com/h4ck3rm1k3/gocore/io"
	"github.com/h4ck3rm1k3/gocore/os"
	"utf8" // thanks ken
)
`,
		Out: `package main

import (
	"encoding/utf8" // thanks ken
	"github.com/h4ck3rm1k3/gocore/io"
	"github.com/h4ck3rm1k3/gocore/os"
)
`,
	},
	{
		Name: "import.14",
		Fn:   rewriteImportFn("asn1", "github.com/h4ck3rm1k3/gocore/encoding/asn1"),
		In: `package main

import (
	"asn1"
	"github.com/h4ck3rm1k3/gocore/crypto"
	"github.com/h4ck3rm1k3/gocore/crypto/rsa"
	_ "github.com/h4ck3rm1k3/gocore/crypto/sha1"
	"github.com/h4ck3rm1k3/gocore/crypto/x509"
	"github.com/h4ck3rm1k3/gocore/crypto/x509/pkix"
	"github.com/h4ck3rm1k3/gocore/time"
)

var x = 1
`,
		Out: `package main

import (
	"github.com/h4ck3rm1k3/gocore/crypto"
	"github.com/h4ck3rm1k3/gocore/crypto/rsa"
	_ "github.com/h4ck3rm1k3/gocore/crypto/sha1"
	"github.com/h4ck3rm1k3/gocore/crypto/x509"
	"github.com/h4ck3rm1k3/gocore/crypto/x509/pkix"
	"github.com/h4ck3rm1k3/gocore/encoding/asn1"
	"github.com/h4ck3rm1k3/gocore/time"
)

var x = 1
`,
	},
	{
		Name: "import.15",
		Fn:   rewriteImportFn("url", "github.com/h4ck3rm1k3/gocore/net/url"),
		In: `package main

import (
	"github.com/h4ck3rm1k3/gocore/bufio"
	"github.com/h4ck3rm1k3/gocore/net"
	"github.com/h4ck3rm1k3/gocore/path"
	"url"
)

var x = 1 // comment on x, not on url
`,
		Out: `package main

import (
	"github.com/h4ck3rm1k3/gocore/bufio"
	"github.com/h4ck3rm1k3/gocore/net"
	"github.com/h4ck3rm1k3/gocore/net/url"
	"github.com/h4ck3rm1k3/gocore/path"
)

var x = 1 // comment on x, not on url
`,
	},
	{
		Name: "import.16",
		Fn:   rewriteImportFn("http", "github.com/h4ck3rm1k3/gocore/net/http", "template", "github.com/h4ck3rm1k3/gocore/text/template"),
		In: `package main

import (
	"github.com/h4ck3rm1k3/gocore/flag"
	"http"
	"github.com/h4ck3rm1k3/gocore/log"
	"template"
)

var addr = flag.String("addr", ":1718", "http service address") // Q=17, R=18
`,
		Out: `package main

import (
	"github.com/h4ck3rm1k3/gocore/flag"
	"github.com/h4ck3rm1k3/gocore/log"
	"github.com/h4ck3rm1k3/gocore/net/http"
	"github.com/h4ck3rm1k3/gocore/text/template"
)

var addr = flag.String("addr", ":1718", "http service address") // Q=17, R=18
`,
	},
	{
		Name: "import.17",
		Fn:   addImportFn("x/y/z", "x/a/c"),
		In: `package main

// Comment
import "C"

import (
	"a"
	"b"

	"x/w"

	"d/f"
)
`,
		Out: `package main

// Comment
import "C"

import (
	"a"
	"b"

	"x/a/c"
	"x/w"
	"x/y/z"

	"d/f"
)
`,
	},
	{
		Name: "import.18",
		Fn:   addDelImportFn("e", "o"),
		In: `package main

import (
	"f"
	"o"
	"z"
)
`,
		Out: `package main

import (
	"e"
	"f"
	"z"
)
`,
	},
}

func addImportFn(path ...string) func(*ast.File) bool {
	return func(f *ast.File) bool {
		fixed := false
		for _, p := range path {
			if !imports(f, p) {
				addImport(f, p)
				fixed = true
			}
		}
		return fixed
	}
}

func deleteImportFn(path string) func(*ast.File) bool {
	return func(f *ast.File) bool {
		if imports(f, path) {
			deleteImport(f, path)
			return true
		}
		return false
	}
}

func addDelImportFn(p1 string, p2 string) func(*ast.File) bool {
	return func(f *ast.File) bool {
		fixed := false
		if !imports(f, p1) {
			addImport(f, p1)
			fixed = true
		}
		if imports(f, p2) {
			deleteImport(f, p2)
			fixed = true
		}
		return fixed
	}
}

func rewriteImportFn(oldnew ...string) func(*ast.File) bool {
	return func(f *ast.File) bool {
		fixed := false
		for i := 0; i < len(oldnew); i += 2 {
			if imports(f, oldnew[i]) {
				rewriteImport(f, oldnew[i], oldnew[i+1])
				fixed = true
			}
		}
		return fixed
	}
}
