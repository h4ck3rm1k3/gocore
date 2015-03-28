// Copyright 2013 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// DO NOT EDIT.
// Generate with: go run gen.go -full -output md5block.go

package md5

import (
	"github.com/h4ck3rm1k3/gocore/runtime"
	"github.com/h4ck3rm1k3/gocore/unsafe"
)

const x86 = runtime.GOARCH == "amd64" || runtime.GOARCH == "386"

var littleEndian bool

func init() {
	x := uint32(0x04030201)
	y := [4]byte{0x1, 0x2, 0x3, 0x4}
	littleEndian = *(*[4]byte)(unsafe.Pointer(&x)) == y
}

func blockGeneric(dig *digest, p []byte) {
	a := dig.s[0]
	b := dig.s[1]
	c := dig.s[2]
	d := dig.s[3]
	var X *[16]uint32
	var xbuf [16]uint32
	for len(p) >= chunk {
		aa, bb, cc, dd := a, b, c, d

		// This is a constant condition - it is not evaluated on each iteration.
		if x86 {
			// MD5 was designed so that x86 processors can just iterate
			// over the block data directly as uint32s, and we generate
			// less code and run 1.3x faster if we take advantage of that.
			// My apologies.
			X = (*[16]uint32)(unsafe.Pointer(&p[0]))
		} else if littleEndian && uintptr(unsafe.Pointer(&p[0]))&(unsafe.Alignof(uint32(0))-1) == 0 {
			X = (*[16]uint32)(unsafe.Pointer(&p[0]))
		} else {
			X = &xbuf
			j := 0
			for i := 0; i < 16; i++ {
				X[i&15] = uint32(p[j]) | uint32(p[j+1])<<8 | uint32(p[j+2])<<16 | uint32(p[j+3])<<24
				j += 4
			}
		}

		// Round 1.

		a += (((c ^ d) & b) ^ d) + X[0] + 3614090360
		a = a<<7 | a>>(32-7) + b

		d += (((b ^ c) & a) ^ c) + X[1] + 3905402710
		d = d<<12 | d>>(32-12) + a

		c += (((a ^ b) & d) ^ b) + X[2] + 606105819
		c = c<<17 | c>>(32-17) + d

		b += (((d ^ a) & c) ^ a) + X[3] + 3250441966
		b = b<<22 | b>>(32-22) + c

		a += (((c ^ d) & b) ^ d) + X[4] + 4118548399
		a = a<<7 | a>>(32-7) + b

		d += (((b ^ c) & a) ^ c) + X[5] + 1200080426
		d = d<<12 | d>>(32-12) + a

		c += (((a ^ b) & d) ^ b) + X[6] + 2821735955
		c = c<<17 | c>>(32-17) + d

		b += (((d ^ a) & c) ^ a) + X[7] + 4249261313
		b = b<<22 | b>>(32-22) + c

		a += (((c ^ d) & b) ^ d) + X[8] + 1770035416
		a = a<<7 | a>>(32-7) + b

		d += (((b ^ c) & a) ^ c) + X[9] + 2336552879
		d = d<<12 | d>>(32-12) + a

		c += (((a ^ b) & d) ^ b) + X[10] + 4294925233
		c = c<<17 | c>>(32-17) + d

		b += (((d ^ a) & c) ^ a) + X[11] + 2304563134
		b = b<<22 | b>>(32-22) + c

		a += (((c ^ d) & b) ^ d) + X[12] + 1804603682
		a = a<<7 | a>>(32-7) + b

		d += (((b ^ c) & a) ^ c) + X[13] + 4254626195
		d = d<<12 | d>>(32-12) + a

		c += (((a ^ b) & d) ^ b) + X[14] + 2792965006
		c = c<<17 | c>>(32-17) + d

		b += (((d ^ a) & c) ^ a) + X[15] + 1236535329
		b = b<<22 | b>>(32-22) + c

		// Round 2.

		a += (((b ^ c) & d) ^ c) + X[(1+5*0)&15] + 4129170786
		a = a<<5 | a>>(32-5) + b

		d += (((a ^ b) & c) ^ b) + X[(1+5*1)&15] + 3225465664
		d = d<<9 | d>>(32-9) + a

		c += (((d ^ a) & b) ^ a) + X[(1+5*2)&15] + 643717713
		c = c<<14 | c>>(32-14) + d

		b += (((c ^ d) & a) ^ d) + X[(1+5*3)&15] + 3921069994
		b = b<<20 | b>>(32-20) + c

		a += (((b ^ c) & d) ^ c) + X[(1+5*4)&15] + 3593408605
		a = a<<5 | a>>(32-5) + b

		d += (((a ^ b) & c) ^ b) + X[(1+5*5)&15] + 38016083
		d = d<<9 | d>>(32-9) + a

		c += (((d ^ a) & b) ^ a) + X[(1+5*6)&15] + 3634488961
		c = c<<14 | c>>(32-14) + d

		b += (((c ^ d) & a) ^ d) + X[(1+5*7)&15] + 3889429448
		b = b<<20 | b>>(32-20) + c

		a += (((b ^ c) & d) ^ c) + X[(1+5*8)&15] + 568446438
		a = a<<5 | a>>(32-5) + b

		d += (((a ^ b) & c) ^ b) + X[(1+5*9)&15] + 3275163606
		d = d<<9 | d>>(32-9) + a

		c += (((d ^ a) & b) ^ a) + X[(1+5*10)&15] + 4107603335
		c = c<<14 | c>>(32-14) + d

		b += (((c ^ d) & a) ^ d) + X[(1+5*11)&15] + 1163531501
		b = b<<20 | b>>(32-20) + c

		a += (((b ^ c) & d) ^ c) + X[(1+5*12)&15] + 2850285829
		a = a<<5 | a>>(32-5) + b

		d += (((a ^ b) & c) ^ b) + X[(1+5*13)&15] + 4243563512
		d = d<<9 | d>>(32-9) + a

		c += (((d ^ a) & b) ^ a) + X[(1+5*14)&15] + 1735328473
		c = c<<14 | c>>(32-14) + d

		b += (((c ^ d) & a) ^ d) + X[(1+5*15)&15] + 2368359562
		b = b<<20 | b>>(32-20) + c

		// Round 3.

		a += (b ^ c ^ d) + X[(5+3*0)&15] + 4294588738
		a = a<<4 | a>>(32-4) + b

		d += (a ^ b ^ c) + X[(5+3*1)&15] + 2272392833
		d = d<<11 | d>>(32-11) + a

		c += (d ^ a ^ b) + X[(5+3*2)&15] + 1839030562
		c = c<<16 | c>>(32-16) + d

		b += (c ^ d ^ a) + X[(5+3*3)&15] + 4259657740
		b = b<<23 | b>>(32-23) + c

		a += (b ^ c ^ d) + X[(5+3*4)&15] + 2763975236
		a = a<<4 | a>>(32-4) + b

		d += (a ^ b ^ c) + X[(5+3*5)&15] + 1272893353
		d = d<<11 | d>>(32-11) + a

		c += (d ^ a ^ b) + X[(5+3*6)&15] + 4139469664
		c = c<<16 | c>>(32-16) + d

		b += (c ^ d ^ a) + X[(5+3*7)&15] + 3200236656
		b = b<<23 | b>>(32-23) + c

		a += (b ^ c ^ d) + X[(5+3*8)&15] + 681279174
		a = a<<4 | a>>(32-4) + b

		d += (a ^ b ^ c) + X[(5+3*9)&15] + 3936430074
		d = d<<11 | d>>(32-11) + a

		c += (d ^ a ^ b) + X[(5+3*10)&15] + 3572445317
		c = c<<16 | c>>(32-16) + d

		b += (c ^ d ^ a) + X[(5+3*11)&15] + 76029189
		b = b<<23 | b>>(32-23) + c

		a += (b ^ c ^ d) + X[(5+3*12)&15] + 3654602809
		a = a<<4 | a>>(32-4) + b

		d += (a ^ b ^ c) + X[(5+3*13)&15] + 3873151461
		d = d<<11 | d>>(32-11) + a

		c += (d ^ a ^ b) + X[(5+3*14)&15] + 530742520
		c = c<<16 | c>>(32-16) + d

		b += (c ^ d ^ a) + X[(5+3*15)&15] + 3299628645
		b = b<<23 | b>>(32-23) + c

		// Round 4.

		a += (c ^ (b | ^d)) + X[(7*0)&15] + 4096336452
		a = a<<6 | a>>(32-6) + b

		d += (b ^ (a | ^c)) + X[(7*1)&15] + 1126891415
		d = d<<10 | d>>(32-10) + a

		c += (a ^ (d | ^b)) + X[(7*2)&15] + 2878612391
		c = c<<15 | c>>(32-15) + d

		b += (d ^ (c | ^a)) + X[(7*3)&15] + 4237533241
		b = b<<21 | b>>(32-21) + c

		a += (c ^ (b | ^d)) + X[(7*4)&15] + 1700485571
		a = a<<6 | a>>(32-6) + b

		d += (b ^ (a | ^c)) + X[(7*5)&15] + 2399980690
		d = d<<10 | d>>(32-10) + a

		c += (a ^ (d | ^b)) + X[(7*6)&15] + 4293915773
		c = c<<15 | c>>(32-15) + d

		b += (d ^ (c | ^a)) + X[(7*7)&15] + 2240044497
		b = b<<21 | b>>(32-21) + c

		a += (c ^ (b | ^d)) + X[(7*8)&15] + 1873313359
		a = a<<6 | a>>(32-6) + b

		d += (b ^ (a | ^c)) + X[(7*9)&15] + 4264355552
		d = d<<10 | d>>(32-10) + a

		c += (a ^ (d | ^b)) + X[(7*10)&15] + 2734768916
		c = c<<15 | c>>(32-15) + d

		b += (d ^ (c | ^a)) + X[(7*11)&15] + 1309151649
		b = b<<21 | b>>(32-21) + c

		a += (c ^ (b | ^d)) + X[(7*12)&15] + 4149444226
		a = a<<6 | a>>(32-6) + b

		d += (b ^ (a | ^c)) + X[(7*13)&15] + 3174756917
		d = d<<10 | d>>(32-10) + a

		c += (a ^ (d | ^b)) + X[(7*14)&15] + 718787259
		c = c<<15 | c>>(32-15) + d

		b += (d ^ (c | ^a)) + X[(7*15)&15] + 3951481745
		b = b<<21 | b>>(32-21) + c

		a += aa
		b += bb
		c += cc
		d += dd

		p = p[chunk:]
	}

	dig.s[0] = a
	dig.s[1] = b
	dig.s[2] = c
	dig.s[3] = d
}
