// Copyright 2013 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build ignore

package main

// This program generates palette.go. Invoke it as
//	go run gen.go -output palette.go

import (
	"github.com/h4ck3rm1k3/gocore/bytes"
	"github.com/h4ck3rm1k3/gocore/flag"
	"github.com/h4ck3rm1k3/gocore/fmt"
	"github.com/h4ck3rm1k3/gocore/go/format"
	"github.com/h4ck3rm1k3/gocore/io"
	"github.com/h4ck3rm1k3/gocore/io/ioutil"
	"github.com/h4ck3rm1k3/gocore/log"
)

var filename = flag.String("output", "palette.go", "output file name")

func main() {
	flag.Parse()

	var buf bytes.Buffer

	fmt.Fprintln(&buf, `// Copyright 2013 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.`)
	fmt.Fprintln(&buf)
	fmt.Fprintln(&buf, "// generated by go run gen.go -output palette.go; DO NOT EDIT")
	fmt.Fprintln(&buf)
	fmt.Fprintln(&buf, "package palette")
	fmt.Fprintln(&buf)
	fmt.Fprintln(&buf, `import "github.com/h4ck3rm1k3/gocore/image/color"`)
	fmt.Fprintln(&buf)
	printPlan9(&buf)
	printWebSafe(&buf)

	data, err := format.Source(buf.Bytes())
	if err != nil {
		log.Fatal(err)
	}
	err = ioutil.WriteFile(*filename, data, 0644)
	if err != nil {
		log.Fatal(err)
	}
}

func printPlan9(w io.Writer) {
	c, lines := [3]int{}, [256]string{}
	for r, i := 0, 0; r != 4; r++ {
		for v := 0; v != 4; v, i = v+1, i+16 {
			for g, j := 0, v-r; g != 4; g++ {
				for b := 0; b != 4; b, j = b+1, j+1 {
					den := r
					if g > den {
						den = g
					}
					if b > den {
						den = b
					}
					if den == 0 {
						c[0] = 0x11 * v
						c[1] = 0x11 * v
						c[2] = 0x11 * v
					} else {
						num := 17 * (4*den + v)
						c[0] = r * num / den
						c[1] = g * num / den
						c[2] = b * num / den
					}
					lines[i+(j&0x0f)] =
						fmt.Sprintf("\tcolor.RGBA{0x%02x, 0x%02x, 0x%02x, 0xff},", c[0], c[1], c[2])
				}
			}
		}
	}
	fmt.Fprintln(w, "// Plan9 is a 256-color palette that partitions the 24-bit RGB space")
	fmt.Fprintln(w, "// into 4×4×4 subdivision, with 4 shades in each subcube. Compared to the")
	fmt.Fprintln(w, "// WebSafe, the idea is to reduce the color resolution by dicing the")
	fmt.Fprintln(w, "// color cube into fewer cells, and to use the extra space to increase the")
	fmt.Fprintln(w, "// intensity resolution. This results in 16 gray shades (4 gray subcubes with")
	fmt.Fprintln(w, "// 4 samples in each), 13 shades of each primary and secondary color (3")
	fmt.Fprintln(w, "// subcubes with 4 samples plus black) and a reasonable selection of colors")
	fmt.Fprintln(w, "// covering the rest of the color cube. The advantage is better representation")
	fmt.Fprintln(w, "// of continuous tones.")
	fmt.Fprintln(w, "//")
	fmt.Fprintln(w, "// This palette was used in the Plan 9 Operating System, described at")
	fmt.Fprintln(w, "// http://plan9.bell-labs.com/magic/man2html/6/color")
	fmt.Fprintln(w, "var Plan9 = []color.Color{")
	for _, line := range lines {
		fmt.Fprintln(w, line)
	}
	fmt.Fprintln(w, "}")
	fmt.Fprintln(w)
}

func printWebSafe(w io.Writer) {
	lines := [6 * 6 * 6]string{}
	for r := 0; r < 6; r++ {
		for g := 0; g < 6; g++ {
			for b := 0; b < 6; b++ {
				lines[36*r+6*g+b] =
					fmt.Sprintf("\tcolor.RGBA{0x%02x, 0x%02x, 0x%02x, 0xff},", 0x33*r, 0x33*g, 0x33*b)
			}
		}
	}
	fmt.Fprintln(w, "// WebSafe is a 216-color palette that was popularized by early versions")
	fmt.Fprintln(w, "// of Netscape Navigator. It is also known as the Netscape Color Cube.")
	fmt.Fprintln(w, "//")
	fmt.Fprintln(w, "// See http://en.wikipedia.org/wiki/Web_colors#Web-safe_colors for details.")
	fmt.Fprintln(w, "var WebSafe = []color.Color{")
	for _, line := range lines {
		fmt.Fprintln(w, line)
	}
	fmt.Fprintln(w, "}")
	fmt.Fprintln(w)
}
