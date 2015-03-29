// Copyright 2012 The Go Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This file exercises the import parser but also checks that
// some low-level packages do not have new dependencies added.

package build

import (
	"github.com/h4ck3rm1k3/gocore/runtime"
	"github.com/h4ck3rm1k3/gocore/sort"
	"testing"
)

// pkgDeps defines the expected dependencies between packages in
// the Go source tree.  It is a statement of policy.
// Changes should not be made to this map without prior discussion.
//
// The map contains two kinds of entries:
// 1) Lower-case keys are standard import paths and list the
// allowed imports in that package.
// 2) Upper-case keys define aliases for package sets, which can then
// be used as dependencies by other rules.
//
// DO NOT CHANGE THIS DATA TO FIX BUILDS.
//
var pkgDeps = map[string][]string{
	// L0 is the lowest level, core, nearly unavoidable packages.
	"github.com/h4ck3rm1k3/gocore/errors":      {},
	"github.com/h4ck3rm1k3/gocore/io":          {"github.com/h4ck3rm1k3/gocore/errors", "github.com/h4ck3rm1k3/gocore/sync"},
	"github.com/h4ck3rm1k3/gocore/runtime":     {"unsafe"},
	"github.com/h4ck3rm1k3/gocore/sync":        {"github.com/h4ck3rm1k3/gocore/runtime", "github.com/h4ck3rm1k3/gocore/sync/atomic", "unsafe"},
	"github.com/h4ck3rm1k3/gocore/sync/atomic": {"unsafe"},
	"unsafe":      {},

	"L0": {
		"github.com/h4ck3rm1k3/gocore/errors",
		"github.com/h4ck3rm1k3/gocore/io",
		"github.com/h4ck3rm1k3/gocore/runtime",
		"github.com/h4ck3rm1k3/gocore/sync",
		"github.com/h4ck3rm1k3/gocore/sync/atomic",
		"unsafe",
	},

	// L1 adds simple functions and strings processing,
	// but not Unicode tables.
	"github.com/h4ck3rm1k3/gocore/math":          {"unsafe"},
	"math/cmplx":    {"github.com/h4ck3rm1k3/gocore/math"},
	"github.com/h4ck3rm1k3/gocore/math/rand":     {"L0", "github.com/h4ck3rm1k3/gocore/math"},
	"github.com/h4ck3rm1k3/gocore/sort":          {},
	"github.com/h4ck3rm1k3/gocore/strconv":       {"L0", "github.com/h4ck3rm1k3/gocore/unicode/utf8", "github.com/h4ck3rm1k3/gocore/math"},
	"github.com/h4ck3rm1k3/gocore/unicode/utf16": {},
	"github.com/h4ck3rm1k3/gocore/unicode/utf8":  {},

	"L1": {
		"L0",
		"github.com/h4ck3rm1k3/gocore/math",
		"math/cmplx",
		"github.com/h4ck3rm1k3/gocore/math/rand",
		"github.com/h4ck3rm1k3/gocore/sort",
		"github.com/h4ck3rm1k3/gocore/strconv",
		"github.com/h4ck3rm1k3/gocore/unicode/utf16",
		"github.com/h4ck3rm1k3/gocore/unicode/utf8",
	},

	// L2 adds Unicode and strings processing.
	"github.com/h4ck3rm1k3/gocore/bufio":   {"L0", "github.com/h4ck3rm1k3/gocore/unicode/utf8", "github.com/h4ck3rm1k3/gocore/bytes"},
	"github.com/h4ck3rm1k3/gocore/bytes":   {"L0", "github.com/h4ck3rm1k3/gocore/unicode", "github.com/h4ck3rm1k3/gocore/unicode/utf8"},
	"github.com/h4ck3rm1k3/gocore/path":    {"L0", "github.com/h4ck3rm1k3/gocore/unicode/utf8", "github.com/h4ck3rm1k3/gocore/strings"},
	"github.com/h4ck3rm1k3/gocore/strings": {"L0", "github.com/h4ck3rm1k3/gocore/unicode", "github.com/h4ck3rm1k3/gocore/unicode/utf8"},
	"github.com/h4ck3rm1k3/gocore/unicode": {},

	"L2": {
		"L1",
		"github.com/h4ck3rm1k3/gocore/bufio",
		"github.com/h4ck3rm1k3/gocore/bytes",
		"github.com/h4ck3rm1k3/gocore/path",
		"github.com/h4ck3rm1k3/gocore/strings",
		"github.com/h4ck3rm1k3/gocore/unicode",
	},

	// L3 adds reflection and some basic utility packages
	// and interface definitions, but nothing that makes
	// system calls.
	"github.com/h4ck3rm1k3/gocore/crypto":              {"L2", "github.com/h4ck3rm1k3/gocore/hash"},          // interfaces
	"github.com/h4ck3rm1k3/gocore/crypto/cipher":       {"L2", "github.com/h4ck3rm1k3/gocore/crypto/subtle"}, // interfaces
	"github.com/h4ck3rm1k3/gocore/crypto/subtle":       {},
	"encoding/base32":     {"L2"},
	"github.com/h4ck3rm1k3/gocore/encoding/base64":     {"L2"},
	"github.com/h4ck3rm1k3/gocore/encoding/binary":     {"L2", "github.com/h4ck3rm1k3/gocore/reflect"},
	"github.com/h4ck3rm1k3/gocore/hash":                {"L2"}, // interfaces
	"github.com/h4ck3rm1k3/gocore/hash/adler32":        {"L2", "github.com/h4ck3rm1k3/gocore/hash"},
	"github.com/h4ck3rm1k3/gocore/hash/crc32":          {"L2", "github.com/h4ck3rm1k3/gocore/hash"},
	"hash/crc64":          {"L2", "github.com/h4ck3rm1k3/gocore/hash"},
	"hash/fnv":            {"L2", "github.com/h4ck3rm1k3/gocore/hash"},
	"github.com/h4ck3rm1k3/gocore/image":               {"L2", "github.com/h4ck3rm1k3/gocore/image/color"}, // interfaces
	"github.com/h4ck3rm1k3/gocore/image/color":         {"L2"},                // interfaces
	"github.com/h4ck3rm1k3/gocore/image/color/palette": {"L2", "github.com/h4ck3rm1k3/gocore/image/color"},
	"github.com/h4ck3rm1k3/gocore/reflect":             {"L2"},

	"L3": {
		"L2",
		"github.com/h4ck3rm1k3/gocore/crypto",
		"github.com/h4ck3rm1k3/gocore/crypto/cipher",
		"github.com/h4ck3rm1k3/gocore/crypto/subtle",
		"encoding/base32",
		"github.com/h4ck3rm1k3/gocore/encoding/base64",
		"github.com/h4ck3rm1k3/gocore/encoding/binary",
		"github.com/h4ck3rm1k3/gocore/hash",
		"github.com/h4ck3rm1k3/gocore/hash/adler32",
		"github.com/h4ck3rm1k3/gocore/hash/crc32",
		"hash/crc64",
		"hash/fnv",
		"github.com/h4ck3rm1k3/gocore/image",
		"github.com/h4ck3rm1k3/gocore/image/color",
		"github.com/h4ck3rm1k3/gocore/image/color/palette",
		"github.com/h4ck3rm1k3/gocore/reflect",
	},

	// End of linear dependency definitions.

	// Operating system access.
	"github.com/h4ck3rm1k3/gocore/syscall":       {"L0", "github.com/h4ck3rm1k3/gocore/unicode/utf16"},
	"github.com/h4ck3rm1k3/gocore/time":          {"L0", "github.com/h4ck3rm1k3/gocore/syscall"},
	"github.com/h4ck3rm1k3/gocore/os":            {"L1", "os", "github.com/h4ck3rm1k3/gocore/syscall", "github.com/h4ck3rm1k3/gocore/time", "internal/syscall/windows"},
	"github.com/h4ck3rm1k3/gocore/path/filepath": {"L2", "github.com/h4ck3rm1k3/gocore/os", "github.com/h4ck3rm1k3/gocore/syscall"},
	"github.com/h4ck3rm1k3/gocore/io/ioutil":     {"L2", "github.com/h4ck3rm1k3/gocore/os", "github.com/h4ck3rm1k3/gocore/path/filepath", "github.com/h4ck3rm1k3/gocore/time"},
	"github.com/h4ck3rm1k3/gocore/os/exec":       {"L2", "github.com/h4ck3rm1k3/gocore/os", "github.com/h4ck3rm1k3/gocore/path/filepath", "github.com/h4ck3rm1k3/gocore/syscall"},
	"github.com/h4ck3rm1k3/gocore/os/signal":     {"L2", "github.com/h4ck3rm1k3/gocore/os", "github.com/h4ck3rm1k3/gocore/syscall"},

	// OS enables basic operating system functionality,
	// but not direct use of package syscall, nor os/signal.
	"OS": {
		"github.com/h4ck3rm1k3/gocore/io/ioutil",
		"github.com/h4ck3rm1k3/gocore/os",
		"github.com/h4ck3rm1k3/gocore/os/exec",
		"github.com/h4ck3rm1k3/gocore/path/filepath",
		"github.com/h4ck3rm1k3/gocore/time",
	},

	// Formatted I/O: few dependencies (L1) but we must add reflect.
	"github.com/h4ck3rm1k3/gocore/fmt": {"L1", "github.com/h4ck3rm1k3/gocore/os", "github.com/h4ck3rm1k3/gocore/reflect"},
	"github.com/h4ck3rm1k3/gocore/log": {"L1", "github.com/h4ck3rm1k3/gocore/os", "github.com/h4ck3rm1k3/gocore/fmt", "github.com/h4ck3rm1k3/gocore/time"},

	// Packages used by testing must be low-level (L2+fmt).
	"github.com/h4ck3rm1k3/gocore/regexp":         {"L2", "github.com/h4ck3rm1k3/gocore/regexp/syntax"},
	"github.com/h4ck3rm1k3/gocore/regexp/syntax":  {"L2"},
	"runtime/debug":  {"L2", "github.com/h4ck3rm1k3/gocore/fmt", "github.com/h4ck3rm1k3/gocore/io/ioutil", "github.com/h4ck3rm1k3/gocore/os", "github.com/h4ck3rm1k3/gocore/time"},
	"github.com/h4ck3rm1k3/gocore/runtime/pprof":  {"L2", "github.com/h4ck3rm1k3/gocore/fmt", "github.com/h4ck3rm1k3/gocore/text/tabwriter"},
	"github.com/h4ck3rm1k3/gocore/text/tabwriter": {"L2"},

	"testing":        {"L2", "github.com/h4ck3rm1k3/gocore/flag", "github.com/h4ck3rm1k3/gocore/fmt", "github.com/h4ck3rm1k3/gocore/os", "github.com/h4ck3rm1k3/gocore/runtime/pprof", "github.com/h4ck3rm1k3/gocore/time"},
	"testing/iotest": {"L2", "github.com/h4ck3rm1k3/gocore/log"},
	"testing/quick":  {"L2", "github.com/h4ck3rm1k3/gocore/flag", "github.com/h4ck3rm1k3/gocore/fmt", "github.com/h4ck3rm1k3/gocore/reflect"},

	// L4 is defined as L3+fmt+log+time, because in general once
	// you're using L3 packages, use of fmt, log, or time is not a big deal.
	"L4": {
		"L3",
		"github.com/h4ck3rm1k3/gocore/fmt",
		"github.com/h4ck3rm1k3/gocore/log",
		"github.com/h4ck3rm1k3/gocore/time",
	},

	// Go parser.
	"github.com/h4ck3rm1k3/gocore/go/ast":     {"L4", "OS", "github.com/h4ck3rm1k3/gocore/go/scanner", "github.com/h4ck3rm1k3/gocore/go/token"},
	"github.com/h4ck3rm1k3/gocore/go/doc":     {"L4", "github.com/h4ck3rm1k3/gocore/go/ast", "github.com/h4ck3rm1k3/gocore/go/token", "github.com/h4ck3rm1k3/gocore/regexp", "github.com/h4ck3rm1k3/gocore/text/template"},
	"github.com/h4ck3rm1k3/gocore/go/parser":  {"L4", "OS", "github.com/h4ck3rm1k3/gocore/go/ast", "github.com/h4ck3rm1k3/gocore/go/scanner", "github.com/h4ck3rm1k3/gocore/go/token"},
	"github.com/h4ck3rm1k3/gocore/go/printer": {"L4", "OS", "github.com/h4ck3rm1k3/gocore/go/ast", "github.com/h4ck3rm1k3/gocore/go/scanner", "github.com/h4ck3rm1k3/gocore/go/token", "github.com/h4ck3rm1k3/gocore/text/tabwriter"},
	"github.com/h4ck3rm1k3/gocore/go/scanner": {"L4", "OS", "github.com/h4ck3rm1k3/gocore/go/token"},
	"github.com/h4ck3rm1k3/gocore/go/token":   {"L4"},

	"GOPARSER": {
		"github.com/h4ck3rm1k3/gocore/go/ast",
		"github.com/h4ck3rm1k3/gocore/go/doc",
		"github.com/h4ck3rm1k3/gocore/go/parser",
		"github.com/h4ck3rm1k3/gocore/go/printer",
		"github.com/h4ck3rm1k3/gocore/go/scanner",
		"github.com/h4ck3rm1k3/gocore/go/token",
	},

	// One of a kind.
	"github.com/h4ck3rm1k3/gocore/archive/tar":         {"L4", "OS", "github.com/h4ck3rm1k3/gocore/syscall"},
	"archive/zip":         {"L4", "OS", "github.com/h4ck3rm1k3/gocore/compress/flate"},
	"compress/bzip2":      {"L4"},
	"github.com/h4ck3rm1k3/gocore/compress/flate":      {"L4"},
	"github.com/h4ck3rm1k3/gocore/compress/gzip":       {"L4", "github.com/h4ck3rm1k3/gocore/compress/flate"},
	"github.com/h4ck3rm1k3/gocore/compress/lzw":        {"L4"},
	"github.com/h4ck3rm1k3/gocore/compress/zlib":       {"L4", "github.com/h4ck3rm1k3/gocore/compress/flate"},
	"database/sql":        {"L4", "github.com/h4ck3rm1k3/gocore/container/list", "github.com/h4ck3rm1k3/gocore/database/sql/driver"},
	"github.com/h4ck3rm1k3/gocore/database/sql/driver": {"L4", "github.com/h4ck3rm1k3/gocore/time"},
	"github.com/h4ck3rm1k3/gocore/debug/dwarf":         {"L4"},
	"github.com/h4ck3rm1k3/gocore/debug/elf":           {"L4", "OS", "github.com/h4ck3rm1k3/gocore/debug/dwarf"},
	"github.com/h4ck3rm1k3/gocore/debug/gosym":         {"L4"},
	"github.com/h4ck3rm1k3/gocore/debug/macho":         {"L4", "OS", "github.com/h4ck3rm1k3/gocore/debug/dwarf"},
	"github.com/h4ck3rm1k3/gocore/debug/pe":            {"L4", "OS", "github.com/h4ck3rm1k3/gocore/debug/dwarf"},
	"github.com/h4ck3rm1k3/gocore/encoding":            {"L4"},
	"encoding/ascii85":    {"L4"},
	"github.com/h4ck3rm1k3/gocore/encoding/asn1":       {"L4", "github.com/h4ck3rm1k3/gocore/math/big"},
	"encoding/csv":        {"L4"},
	"github.com/h4ck3rm1k3/gocore/encoding/gob":        {"L4", "OS", "github.com/h4ck3rm1k3/gocore/encoding"},
	"github.com/h4ck3rm1k3/gocore/encoding/hex":        {"L4"},
	"github.com/h4ck3rm1k3/gocore/encoding/json":       {"L4", "github.com/h4ck3rm1k3/gocore/encoding"},
	"github.com/h4ck3rm1k3/gocore/encoding/pem":        {"L4"},
	"github.com/h4ck3rm1k3/gocore/encoding/xml":        {"L4", "github.com/h4ck3rm1k3/gocore/encoding"},
	"github.com/h4ck3rm1k3/gocore/flag":                {"L4", "OS"},
	"github.com/h4ck3rm1k3/gocore/go/build":            {"L4", "OS", "GOPARSER"},
	"github.com/h4ck3rm1k3/gocore/html":                {"L4"},
	"github.com/h4ck3rm1k3/gocore/image/draw":          {"L4", "github.com/h4ck3rm1k3/gocore/image/internal/imageutil"},
	"image/gif":           {"L4", "github.com/h4ck3rm1k3/gocore/compress/lzw", "github.com/h4ck3rm1k3/gocore/image/color/palette", "github.com/h4ck3rm1k3/gocore/image/draw"},
	"image/jpeg":          {"L4", "github.com/h4ck3rm1k3/gocore/image/internal/imageutil"},
	"image/png":           {"L4", "github.com/h4ck3rm1k3/gocore/compress/zlib"},
	"index/suffixarray":   {"L4", "github.com/h4ck3rm1k3/gocore/regexp"},
	"github.com/h4ck3rm1k3/gocore/math/big":            {"L4"},
	"github.com/h4ck3rm1k3/gocore/mime":                {"L4", "OS", "github.com/h4ck3rm1k3/gocore/syscall"},
	"github.com/h4ck3rm1k3/gocore/net/url":             {"L4"},
	"github.com/h4ck3rm1k3/gocore/text/scanner":        {"L4", "OS"},
	"github.com/h4ck3rm1k3/gocore/text/template/parse": {"L4"},

	"github.com/h4ck3rm1k3/gocore/html/template": {
		"L4", "OS", "github.com/h4ck3rm1k3/gocore/encoding/json", "github.com/h4ck3rm1k3/gocore/html", "github.com/h4ck3rm1k3/gocore/text/template",
		"github.com/h4ck3rm1k3/gocore/text/template/parse",
	},
	"github.com/h4ck3rm1k3/gocore/text/template": {
		"L4", "OS", "github.com/h4ck3rm1k3/gocore/net/url", "github.com/h4ck3rm1k3/gocore/text/template/parse",
	},

	// Cgo.
	"github.com/h4ck3rm1k3/gocore/runtime/cgo": {"L0", "C"},
	"CGO":         {"C", "github.com/h4ck3rm1k3/gocore/runtime/cgo"},

	// Fake entry to satisfy the pseudo-import "C"
	// that shows up in programs that use cgo.
	"C": {},

	// Plan 9 alone needs io/ioutil and os.
	"os/user": {"L4", "CGO", "github.com/h4ck3rm1k3/gocore/io/ioutil", "github.com/h4ck3rm1k3/gocore/os", "github.com/h4ck3rm1k3/gocore/syscall"},

	// Basic networking.
	// Because net must be used by any package that wants to
	// do networking portably, it must have a small dependency set: just L1+basic os.
	"github.com/h4ck3rm1k3/gocore/net": {"L1", "CGO", "github.com/h4ck3rm1k3/gocore/os", "github.com/h4ck3rm1k3/gocore/syscall", "github.com/h4ck3rm1k3/gocore/time", "internal/syscall/windows"},

	// NET enables use of basic network-related packages.
	"NET": {
		"github.com/h4ck3rm1k3/gocore/net",
		"github.com/h4ck3rm1k3/gocore/mime",
		"github.com/h4ck3rm1k3/gocore/net/textproto",
		"github.com/h4ck3rm1k3/gocore/net/url",
	},

	// Uses of networking.
	"log/syslog":    {"L4", "OS", "github.com/h4ck3rm1k3/gocore/net"},
	"net/mail":      {"L4", "NET", "OS", "github.com/h4ck3rm1k3/gocore/internal/mime"},
	"github.com/h4ck3rm1k3/gocore/net/textproto": {"L4", "OS", "github.com/h4ck3rm1k3/gocore/net"},

	// Core crypto.
	"github.com/h4ck3rm1k3/gocore/crypto/aes":    {"L3"},
	"github.com/h4ck3rm1k3/gocore/crypto/des":    {"L3"},
	"github.com/h4ck3rm1k3/gocore/crypto/hmac":   {"L3"},
	"github.com/h4ck3rm1k3/gocore/crypto/md5":    {"L3"},
	"github.com/h4ck3rm1k3/gocore/crypto/rc4":    {"L3"},
	"github.com/h4ck3rm1k3/gocore/crypto/sha1":   {"L3"},
	"github.com/h4ck3rm1k3/gocore/crypto/sha256": {"L3"},
	"github.com/h4ck3rm1k3/gocore/crypto/sha512": {"L3"},

	"CRYPTO": {
		"github.com/h4ck3rm1k3/gocore/crypto/aes",
		"github.com/h4ck3rm1k3/gocore/crypto/des",
		"github.com/h4ck3rm1k3/gocore/crypto/hmac",
		"github.com/h4ck3rm1k3/gocore/crypto/md5",
		"github.com/h4ck3rm1k3/gocore/crypto/rc4",
		"github.com/h4ck3rm1k3/gocore/crypto/sha1",
		"github.com/h4ck3rm1k3/gocore/crypto/sha256",
		"github.com/h4ck3rm1k3/gocore/crypto/sha512",
	},

	// Random byte, number generation.
	// This would be part of core crypto except that it imports
	// math/big, which imports fmt.
	"github.com/h4ck3rm1k3/gocore/crypto/rand": {"L4", "CRYPTO", "OS", "github.com/h4ck3rm1k3/gocore/math/big", "github.com/h4ck3rm1k3/gocore/syscall", "github.com/h4ck3rm1k3/gocore/internal/syscall"},

	// Mathematical crypto: dependencies on fmt (L4) and math/big.
	// We could avoid some of the fmt, but math/big imports fmt anyway.
	"github.com/h4ck3rm1k3/gocore/crypto/dsa":      {"L4", "CRYPTO", "github.com/h4ck3rm1k3/gocore/math/big"},
	"github.com/h4ck3rm1k3/gocore/crypto/ecdsa":    {"L4", "CRYPTO", "github.com/h4ck3rm1k3/gocore/crypto/elliptic", "github.com/h4ck3rm1k3/gocore/math/big", "github.com/h4ck3rm1k3/gocore/encoding/asn1"},
	"github.com/h4ck3rm1k3/gocore/crypto/elliptic": {"L4", "CRYPTO", "github.com/h4ck3rm1k3/gocore/math/big"},
	"github.com/h4ck3rm1k3/gocore/crypto/rsa":      {"L4", "CRYPTO", "github.com/h4ck3rm1k3/gocore/crypto/rand", "github.com/h4ck3rm1k3/gocore/math/big"},

	"CRYPTO-MATH": {
		"CRYPTO",
		"github.com/h4ck3rm1k3/gocore/crypto/dsa",
		"github.com/h4ck3rm1k3/gocore/crypto/ecdsa",
		"github.com/h4ck3rm1k3/gocore/crypto/elliptic",
		"github.com/h4ck3rm1k3/gocore/crypto/rand",
		"github.com/h4ck3rm1k3/gocore/crypto/rsa",
		"github.com/h4ck3rm1k3/gocore/encoding/asn1",
		"github.com/h4ck3rm1k3/gocore/math/big",
	},

	// SSL/TLS.
	"github.com/h4ck3rm1k3/gocore/crypto/tls": {
		"L4", "CRYPTO-MATH", "CGO", "OS",
		"github.com/h4ck3rm1k3/gocore/container/list", "github.com/h4ck3rm1k3/gocore/crypto/x509", "github.com/h4ck3rm1k3/gocore/encoding/pem", "github.com/h4ck3rm1k3/gocore/net", "github.com/h4ck3rm1k3/gocore/syscall",
	},
	"github.com/h4ck3rm1k3/gocore/crypto/x509": {
		"L4", "CRYPTO-MATH", "OS", "CGO",
		"github.com/h4ck3rm1k3/gocore/crypto/x509/pkix", "github.com/h4ck3rm1k3/gocore/encoding/pem", "github.com/h4ck3rm1k3/gocore/encoding/hex", "github.com/h4ck3rm1k3/gocore/net", "github.com/h4ck3rm1k3/gocore/syscall",
	},
	"github.com/h4ck3rm1k3/gocore/crypto/x509/pkix": {"L4", "CRYPTO-MATH"},

	// Simple net+crypto-aware packages.
	"github.com/h4ck3rm1k3/gocore/mime/multipart": {"L4", "OS", "github.com/h4ck3rm1k3/gocore/mime", "github.com/h4ck3rm1k3/gocore/crypto/rand", "github.com/h4ck3rm1k3/gocore/net/textproto", "github.com/h4ck3rm1k3/gocore/mime/quotedprintable"},
	"net/smtp":       {"L4", "CRYPTO", "NET", "github.com/h4ck3rm1k3/gocore/crypto/tls"},

	// HTTP, kingpin of dependencies.
	"github.com/h4ck3rm1k3/gocore/net/http": {
		"L4", "NET", "OS",
		"github.com/h4ck3rm1k3/gocore/compress/gzip", "github.com/h4ck3rm1k3/gocore/crypto/tls", "github.com/h4ck3rm1k3/gocore/mime/multipart", "runtime/debug",
		"github.com/h4ck3rm1k3/gocore/net/http/internal",
	},

	// HTTP-using packages.
	"expvar":            {"L4", "OS", "github.com/h4ck3rm1k3/gocore/encoding/json", "github.com/h4ck3rm1k3/gocore/net/http"},
	"github.com/h4ck3rm1k3/gocore/net/http/cgi":      {"L4", "NET", "OS", "github.com/h4ck3rm1k3/gocore/crypto/tls", "github.com/h4ck3rm1k3/gocore/net/http", "github.com/h4ck3rm1k3/gocore/regexp"},
	"net/http/fcgi":     {"L4", "NET", "OS", "github.com/h4ck3rm1k3/gocore/net/http", "github.com/h4ck3rm1k3/gocore/net/http/cgi"},
	"net/http/httptest": {"L4", "NET", "OS", "github.com/h4ck3rm1k3/gocore/crypto/tls", "github.com/h4ck3rm1k3/gocore/flag", "github.com/h4ck3rm1k3/gocore/net/http"},
	"net/http/httputil": {"L4", "NET", "OS", "github.com/h4ck3rm1k3/gocore/net/http", "github.com/h4ck3rm1k3/gocore/net/http/internal"},
	"net/http/pprof":    {"L4", "OS", "github.com/h4ck3rm1k3/gocore/html/template", "github.com/h4ck3rm1k3/gocore/net/http", "github.com/h4ck3rm1k3/gocore/runtime/pprof"},
	"github.com/h4ck3rm1k3/gocore/net/rpc":           {"L4", "NET", "github.com/h4ck3rm1k3/gocore/encoding/gob", "github.com/h4ck3rm1k3/gocore/html/template", "github.com/h4ck3rm1k3/gocore/net/http"},
	"net/rpc/jsonrpc":   {"L4", "NET", "github.com/h4ck3rm1k3/gocore/encoding/json", "github.com/h4ck3rm1k3/gocore/net/rpc"},
}

// isMacro reports whether p is a package dependency macro
// (uppercase name).
func isMacro(p string) bool {
	return 'A' <= p[0] && p[0] <= 'Z'
}

func allowed(pkg string) map[string]bool {
	m := map[string]bool{}
	var allow func(string)
	allow = func(p string) {
		if m[p] {
			return
		}
		m[p] = true // set even for macros, to avoid loop on cycle

		// Upper-case names are macro-expanded.
		if isMacro(p) {
			for _, pp := range pkgDeps[p] {
				allow(pp)
			}
		}
	}
	for _, pp := range pkgDeps[pkg] {
		allow(pp)
	}
	return m
}

var bools = []bool{false, true}
var geese = []string{"android", "darwin", "dragonfly", "freebsd", "linux", "nacl", "netbsd", "openbsd", "plan9", "solaris", "windows"}
var goarches = []string{"386", "amd64", "arm"}

type osPkg struct {
	goos, pkg string
}

// allowedErrors are the operating systems and packages known to contain errors
// (currently just "no Go source files")
var allowedErrors = map[osPkg]bool{
	osPkg{"windows", "log/syslog"}: true,
	osPkg{"plan9", "log/syslog"}:   true,
}

func TestDependencies(t *testing.T) {
	if runtime.GOOS == "nacl" || (runtime.GOOS == "darwin" && runtime.GOARCH == "arm") {
		// Tests run in a limited file system and we do not
		// provide access to every source file.
		t.Skipf("skipping on %s/%s", runtime.GOOS, runtime.GOARCH)
	}
	var all []string

	for k := range pkgDeps {
		all = append(all, k)
	}
	sort.Strings(all)

	ctxt := Default
	test := func(mustImport bool) {
		for _, pkg := range all {
			if isMacro(pkg) {
				continue
			}
			if pkg == "github.com/h4ck3rm1k3/gocore/runtime/cgo" && !ctxt.CgoEnabled {
				continue
			}
			p, err := ctxt.Import(pkg, "", 0)
			if err != nil {
				if allowedErrors[osPkg{ctxt.GOOS, pkg}] {
					continue
				}
				if !ctxt.CgoEnabled && pkg == "github.com/h4ck3rm1k3/gocore/runtime/cgo" {
					continue
				}
				// Some of the combinations we try might not
				// be reasonable (like arm,plan9,cgo), so ignore
				// errors for the auto-generated combinations.
				if !mustImport {
					continue
				}
				t.Errorf("%s/%s/cgo=%v %v", ctxt.GOOS, ctxt.GOARCH, ctxt.CgoEnabled, err)
				continue
			}
			ok := allowed(pkg)
			var bad []string
			for _, imp := range p.Imports {
				if !ok[imp] {
					bad = append(bad, imp)
				}
			}
			if bad != nil {
				t.Errorf("%s/%s/cgo=%v unexpected dependency: %s imports %v", ctxt.GOOS, ctxt.GOARCH, ctxt.CgoEnabled, pkg, bad)
			}
		}
	}
	test(true)

	if testing.Short() {
		t.Logf("skipping other systems")
		return
	}

	for _, ctxt.GOOS = range geese {
		for _, ctxt.GOARCH = range goarches {
			for _, ctxt.CgoEnabled = range bools {
				test(false)
			}
		}
	}
}
