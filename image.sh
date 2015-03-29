#IMPORT github.com/h4ck3rm1k3/gocore/image
#IMPORT github.com/h4ck3rm1k3/gocore/bufio
#IMPORT github.com/h4ck3rm1k3/gocore/bytes
#IMPORT github.com/h4ck3rm1k3/gocore/errors
#IMPORT runtime
#IMPORT unsafe
#IMPORT github.com/h4ck3rm1k3/gocore/io
#IMPORT github.com/h4ck3rm1k3/gocore/errors
#IMPORT github.com/h4ck3rm1k3/gocore/sync
#IMPORT github.com/h4ck3rm1k3/gocore/runtime
#IMPORT github.com/h4ck3rm1k3/gocore/unsafe
#IMPORT runtime
#IMPORT runtime
#IMPORT github.com/h4ck3rm1k3/gocore/sync/atomic
#IMPORT github.com/h4ck3rm1k3/gocore/unsafe
#IMPORT runtime
#IMPORT github.com/h4ck3rm1k3/gocore/unsafe
#IMPORT runtime
#IMPORT runtime
#IMPORT github.com/h4ck3rm1k3/gocore/unicode
#IMPORT runtime
#IMPORT github.com/h4ck3rm1k3/gocore/unicode/utf8
#IMPORT runtime
#IMPORT runtime
#IMPORT github.com/h4ck3rm1k3/gocore/errors
#IMPORT github.com/h4ck3rm1k3/gocore/io
#IMPORT github.com/h4ck3rm1k3/gocore/unicode/utf8
#IMPORT runtime
#IMPORT github.com/h4ck3rm1k3/gocore/errors
#IMPORT github.com/h4ck3rm1k3/gocore/image/color
#IMPORT runtime
#IMPORT github.com/h4ck3rm1k3/gocore/io
#IMPORT github.com/h4ck3rm1k3/gocore/strconv
#IMPORT github.com/h4ck3rm1k3/gocore/errors
#IMPORT github.com/h4ck3rm1k3/gocore/math
#IMPORT github.com/h4ck3rm1k3/gocore/unsafe
#IMPORT runtime
#IMPORT github.com/h4ck3rm1k3/gocore/unicode/utf8
#IMPORT runtime
#IMPORT runtime
#IMPORT github.com/h4ck3rm1k3/gocore/image

#
# github.com/h4ck3rm1k3/gocore/errors
#

#github.com/h4ck3rm1k3/gocore/errors
mkdir -p $WORK/github.com/h4ck3rm1k3/gocore/errors/
mkdir -p $WORK/github.com/h4ck3rm1k3/gocore/
#COMPILER /home/h4ck3rm1k3/install/bin/gccgo
cd /home/h4ck3rm1k3/testgo2/src/github.com/h4ck3rm1k3/gocore/errors
/home/h4ck3rm1k3/install/bin/gccgo -I $WORK -c -g -fgo-pkgpath=github.com/h4ck3rm1k3/gocore/errors -fgo-relative-import-path=_/home/h4ck3rm1k3/testgo2/src/github.com/h4ck3rm1k3/gocore/errors -o $WORK/github.com/h4ck3rm1k3/gocore/errors/_go_.o -c  -O0 ./errors.go
/home/h4ck3rm1k3/install/bin/gccgo -I $WORK/github.com/h4ck3rm1k3/gocore/errors/ -o $WORK/github.com/h4ck3rm1k3/gocore/errors/errors.o -D GOOS_linux -D GOARCH_ppc64le -D "GOPKGPATH=\"github_com_h4ck3rm1k3_gocore_errors\"" ./errors.s
ar cru $WORK/github.com/h4ck3rm1k3/gocore/liberrors.a $WORK/github.com/h4ck3rm1k3/gocore/errors/_go_.o $WORK/github.com/h4ck3rm1k3/gocore/errors/errors.o

#
# github.com/h4ck3rm1k3/gocore/unsafe
#

#github.com/h4ck3rm1k3/gocore/unsafe
mkdir -p $WORK/github.com/h4ck3rm1k3/gocore/unsafe/
#COMPILER /home/h4ck3rm1k3/install/bin/gccgo
cd /home/h4ck3rm1k3/testgo2/src/github.com/h4ck3rm1k3/gocore/unsafe

/home/h4ck3rm1k3/install/bin/gccgo -I $WORK -c -g -fgo-pkgpath=github.com/h4ck3rm1k3/gocore/unsafe -fgo-relative-import-path=_/home/h4ck3rm1k3/testgo2/src/github.com/h4ck3rm1k3/gocore/unsafe -o $WORK/github.com/h4ck3rm1k3/gocore/unsafe/_go_.o -c  -O0 ./unsafe.go
/home/h4ck3rm1k3/install/bin/gccgo -I $WORK/github.com/h4ck3rm1k3/gocore/unsafe/ -o $WORK/github.com/h4ck3rm1k3/gocore/unsafe/unsafe.o -D GOOS_linux -D GOARCH_ppc64le -D "GOPKGPATH=\"github_com_h4ck3rm1k3_gocore_unsafe\"" ./unsafe.s
ar cru $WORK/github.com/h4ck3rm1k3/gocore/libunsafe.a $WORK/github.com/h4ck3rm1k3/gocore/unsafe/_go_.o $WORK/github.com/h4ck3rm1k3/gocore/unsafe/unsafe.o

#
# github.com/h4ck3rm1k3/gocore/runtime
#


#
# github.com/h4ck3rm1k3/gocore/sync/atomic
#

#github.com/h4ck3rm1k3/gocore/sync/atomic
mkdir -p $WORK/github.com/h4ck3rm1k3/gocore/sync/atomic/
mkdir -p $WORK/github.com/h4ck3rm1k3/gocore/sync/
#COMPILER /home/h4ck3rm1k3/install/bin/gccgo
cd /home/h4ck3rm1k3/testgo2/src/github.com/h4ck3rm1k3/gocore/sync/atomic
/home/h4ck3rm1k3/install/bin/gccgo -I $WORK -I /home/h4ck3rm1k3/testgo2/pkg/gccgo_linux_ppc64le -c -g -fgo-pkgpath=github.com/h4ck3rm1k3/gocore/sync/atomic -fgo-relative-import-path=_/home/h4ck3rm1k3/testgo2/src/github.com/h4ck3rm1k3/gocore/sync/atomic -o $WORK/github.com/h4ck3rm1k3/gocore/sync/atomic/_go_.o -c  -O0 ./doc.go ./value.go
/home/h4ck3rm1k3/install/bin/gccgo -I $WORK/github.com/h4ck3rm1k3/gocore/sync/atomic/ -o $WORK/github.com/h4ck3rm1k3/gocore/sync/atomic/doc.o -D GOOS_linux -D GOARCH_ppc64le -D "GOPKGPATH=\"github_com_h4ck3rm1k3_gocore_sync_atomic\"" ./doc.s
ar cru $WORK/github.com/h4ck3rm1k3/gocore/sync/libatomic.a $WORK/github.com/h4ck3rm1k3/gocore/sync/atomic/_go_.o $WORK/github.com/h4ck3rm1k3/gocore/sync/atomic/doc.o

#
# github.com/h4ck3rm1k3/gocore/sync
#

#github.com/h4ck3rm1k3/gocore/sync
mkdir -p $WORK/github.com/h4ck3rm1k3/gocore/sync/
#COMPILER /home/h4ck3rm1k3/install/bin/gccgo
cd /home/h4ck3rm1k3/testgo2/src/github.com/h4ck3rm1k3/gocore/sync
/home/h4ck3rm1k3/install/bin/gccgo -I $WORK -I /home/h4ck3rm1k3/testgo2/pkg/gccgo_linux_ppc64le -c -g -fgo-pkgpath=github.com/h4ck3rm1k3/gocore/sync -fgo-relative-import-path=_/home/h4ck3rm1k3/testgo2/src/github.com/h4ck3rm1k3/gocore/sync -o $WORK/github.com/h4ck3rm1k3/gocore/sync/_go_.o -c  -O0 ./cond.go ./mutex.go ./once.go ./pool.go ./race0.go ./runtime.go ./rwmutex.go ./waitgroup.go
ar cru $WORK/github.com/h4ck3rm1k3/gocore/libsync.a $WORK/github.com/h4ck3rm1k3/gocore/sync/_go_.o

#
# github.com/h4ck3rm1k3/gocore/io
#

#github.com/h4ck3rm1k3/gocore/io
mkdir -p $WORK/github.com/h4ck3rm1k3/gocore/io/
#COMPILER /home/h4ck3rm1k3/install/bin/gccgo
cd /home/h4ck3rm1k3/testgo2/src/github.com/h4ck3rm1k3/gocore/io
/home/h4ck3rm1k3/install/bin/gccgo -I $WORK -I /home/h4ck3rm1k3/testgo2/pkg/gccgo_linux_ppc64le -c -g -fgo-pkgpath=github.com/h4ck3rm1k3/gocore/io -fgo-relative-import-path=_/home/h4ck3rm1k3/testgo2/src/github.com/h4ck3rm1k3/gocore/io -o $WORK/github.com/h4ck3rm1k3/gocore/io/_go_.o -c  -O0 ./io.go ./multi.go ./pipe.go
ar cru $WORK/github.com/h4ck3rm1k3/gocore/libio.a $WORK/github.com/h4ck3rm1k3/gocore/io/_go_.o

#
# github.com/h4ck3rm1k3/gocore/unicode
#

#github.com/h4ck3rm1k3/gocore/unicode
mkdir -p $WORK/github.com/h4ck3rm1k3/gocore/unicode/
#COMPILER /home/h4ck3rm1k3/install/bin/gccgo
cd /home/h4ck3rm1k3/testgo2/src/github.com/h4ck3rm1k3/gocore/unicode
/home/h4ck3rm1k3/install/bin/gccgo -I $WORK -c -g -fgo-pkgpath=github.com/h4ck3rm1k3/gocore/unicode -fgo-relative-import-path=_/home/h4ck3rm1k3/testgo2/src/github.com/h4ck3rm1k3/gocore/unicode -o $WORK/github.com/h4ck3rm1k3/gocore/unicode/_go_.o -c  -O0 ./casetables.go ./digit.go ./graphic.go ./letter.go ./tables.go
/home/h4ck3rm1k3/install/bin/gccgo -I $WORK/github.com/h4ck3rm1k3/gocore/unicode/ -o $WORK/github.com/h4ck3rm1k3/gocore/unicode/casetables.o -D GOOS_linux -D GOARCH_ppc64le -D "GOPKGPATH=\"github_com_h4ck3rm1k3_gocore_unicode\"" ./casetables.s
ar cru $WORK/github.com/h4ck3rm1k3/gocore/libunicode.a $WORK/github.com/h4ck3rm1k3/gocore/unicode/_go_.o $WORK/github.com/h4ck3rm1k3/gocore/unicode/casetables.o

#
# github.com/h4ck3rm1k3/gocore/unicode/utf8
#

#github.com/h4ck3rm1k3/gocore/unicode/utf8
mkdir -p $WORK/github.com/h4ck3rm1k3/gocore/unicode/utf8/
mkdir -p $WORK/github.com/h4ck3rm1k3/gocore/unicode/
#COMPILER /home/h4ck3rm1k3/install/bin/gccgo
cd /home/h4ck3rm1k3/testgo2/src/github.com/h4ck3rm1k3/gocore/unicode/utf8
/home/h4ck3rm1k3/install/bin/gccgo -I $WORK -c -g -fgo-pkgpath=github.com/h4ck3rm1k3/gocore/unicode/utf8 -fgo-relative-import-path=_/home/h4ck3rm1k3/testgo2/src/github.com/h4ck3rm1k3/gocore/unicode/utf8 -o $WORK/github.com/h4ck3rm1k3/gocore/unicode/utf8/_go_.o -c  -O0 ./utf8.go
/home/h4ck3rm1k3/install/bin/gccgo -I $WORK/github.com/h4ck3rm1k3/gocore/unicode/utf8/ -o $WORK/github.com/h4ck3rm1k3/gocore/unicode/utf8/utf8.o -D GOOS_linux -D GOARCH_ppc64le -D "GOPKGPATH=\"github_com_h4ck3rm1k3_gocore_unicode_utf8\"" ./utf8.s
ar cru $WORK/github.com/h4ck3rm1k3/gocore/unicode/libutf8.a $WORK/github.com/h4ck3rm1k3/gocore/unicode/utf8/_go_.o $WORK/github.com/h4ck3rm1k3/gocore/unicode/utf8/utf8.o

#
# github.com/h4ck3rm1k3/gocore/bytes
#

#github.com/h4ck3rm1k3/gocore/bytes
mkdir -p $WORK/github.com/h4ck3rm1k3/gocore/bytes/
#COMPILER /home/h4ck3rm1k3/install/bin/gccgo
cd /home/h4ck3rm1k3/testgo2/src/github.com/h4ck3rm1k3/gocore/bytes
/home/h4ck3rm1k3/install/bin/gccgo -I $WORK -I /home/h4ck3rm1k3/testgo2/pkg/gccgo_linux_ppc64le -c -g -fgo-pkgpath=github.com/h4ck3rm1k3/gocore/bytes -fgo-relative-import-path=_/home/h4ck3rm1k3/testgo2/src/github.com/h4ck3rm1k3/gocore/bytes -o $WORK/github.com/h4ck3rm1k3/gocore/bytes/_go_.o -c  -O0 ./buffer.go ./bytes.go ./bytes_decl.go ./reader.go
ar cru $WORK/github.com/h4ck3rm1k3/gocore/libbytes.a $WORK/github.com/h4ck3rm1k3/gocore/bytes/_go_.o

#
# github.com/h4ck3rm1k3/gocore/bufio
#

#github.com/h4ck3rm1k3/gocore/bufio
mkdir -p $WORK/github.com/h4ck3rm1k3/gocore/bufio/
#COMPILER /home/h4ck3rm1k3/install/bin/gccgo
cd /home/h4ck3rm1k3/testgo2/src/github.com/h4ck3rm1k3/gocore/bufio
/home/h4ck3rm1k3/install/bin/gccgo -I $WORK -I /home/h4ck3rm1k3/testgo2/pkg/gccgo_linux_ppc64le -c -g -fgo-pkgpath=github.com/h4ck3rm1k3/gocore/bufio -fgo-relative-import-path=_/home/h4ck3rm1k3/testgo2/src/github.com/h4ck3rm1k3/gocore/bufio -o $WORK/github.com/h4ck3rm1k3/gocore/bufio/_go_.o -c  -O0 ./bufio.go ./scan.go
ar cru $WORK/github.com/h4ck3rm1k3/gocore/libbufio.a $WORK/github.com/h4ck3rm1k3/gocore/bufio/_go_.o

#
# github.com/h4ck3rm1k3/gocore/image/color
#

#github.com/h4ck3rm1k3/gocore/image/color
mkdir -p $WORK/github.com/h4ck3rm1k3/gocore/image/color/
mkdir -p $WORK/github.com/h4ck3rm1k3/gocore/image/
#COMPILER /home/h4ck3rm1k3/install/bin/gccgo
cd /home/h4ck3rm1k3/testgo2/src/github.com/h4ck3rm1k3/gocore/image/color
/home/h4ck3rm1k3/install/bin/gccgo -I $WORK -c -g -fgo-pkgpath=github.com/h4ck3rm1k3/gocore/image/color -fgo-relative-import-path=_/home/h4ck3rm1k3/testgo2/src/github.com/h4ck3rm1k3/gocore/image/color -o $WORK/github.com/h4ck3rm1k3/gocore/image/color/_go_.o -c  -O0 ./color.go ./ycbcr.go
/home/h4ck3rm1k3/install/bin/gccgo -I $WORK/github.com/h4ck3rm1k3/gocore/image/color/ -o $WORK/github.com/h4ck3rm1k3/gocore/image/color/color.o -D GOOS_linux -D GOARCH_ppc64le -D "GOPKGPATH=\"github_com_h4ck3rm1k3_gocore_image_color\"" ./color.s
ar cru $WORK/github.com/h4ck3rm1k3/gocore/image/libcolor.a $WORK/github.com/h4ck3rm1k3/gocore/image/color/_go_.o $WORK/github.com/h4ck3rm1k3/gocore/image/color/color.o

#
# github.com/h4ck3rm1k3/gocore/math
#

#github.com/h4ck3rm1k3/gocore/math
mkdir -p $WORK/github.com/h4ck3rm1k3/gocore/math/
#COMPILER /home/h4ck3rm1k3/install/bin/gccgo
cd /home/h4ck3rm1k3/testgo2/src/github.com/h4ck3rm1k3/gocore/math
/home/h4ck3rm1k3/install/bin/gccgo -I $WORK -I /home/h4ck3rm1k3/testgo2/pkg/gccgo_linux_ppc64le -c -g -fgo-pkgpath=github.com/h4ck3rm1k3/gocore/math -fgo-relative-import-path=_/home/h4ck3rm1k3/testgo2/src/github.com/h4ck3rm1k3/gocore/math -o $WORK/github.com/h4ck3rm1k3/gocore/math/_go_.o -c  -O0 ./abs.go ./acosh.go ./asin.go ./asinh.go ./atan.go ./atan2.go ./atanh.go ./bits.go ./cbrt.go ./const.go ./copysign.go ./dim.go ./erf.go ./exp.go ./expm1.go ./floor.go ./frexp.go ./gamma.go ./hypot.go ./j0.go ./j1.go ./jn.go ./ldexp.go ./lgamma.go ./log.go ./log10.go ./log1p.go ./logb.go ./mod.go ./modf.go ./nextafter.go ./pow.go ./pow10.go ./remainder.go ./signbit.go ./sin.go ./sincos.go ./sinh.go ./sqrt.go ./tan.go ./tanh.go ./unsafe.go
/home/h4ck3rm1k3/install/bin/gccgo -I $WORK/github.com/h4ck3rm1k3/gocore/math/ -o $WORK/github.com/h4ck3rm1k3/gocore/math/abs.o -D GOOS_linux -D GOARCH_ppc64le -D "GOPKGPATH=\"github_com_h4ck3rm1k3_gocore_math\"" ./abs.s
ar cru $WORK/github.com/h4ck3rm1k3/gocore/libmath.a $WORK/github.com/h4ck3rm1k3/gocore/math/_go_.o $WORK/github.com/h4ck3rm1k3/gocore/math/abs.o

#
# github.com/h4ck3rm1k3/gocore/strconv
#

#github.com/h4ck3rm1k3/gocore/strconv
mkdir -p $WORK/github.com/h4ck3rm1k3/gocore/strconv/
#COMPILER /home/h4ck3rm1k3/install/bin/gccgo
cd /home/h4ck3rm1k3/testgo2/src/github.com/h4ck3rm1k3/gocore/strconv
/home/h4ck3rm1k3/install/bin/gccgo -I $WORK -I /home/h4ck3rm1k3/testgo2/pkg/gccgo_linux_ppc64le -c -g -fgo-pkgpath=github.com/h4ck3rm1k3/gocore/strconv -fgo-relative-import-path=_/home/h4ck3rm1k3/testgo2/src/github.com/h4ck3rm1k3/gocore/strconv -o $WORK/github.com/h4ck3rm1k3/gocore/strconv/_go_.o -c  -O0 ./atob.go ./atof.go ./atoi.go ./decimal.go ./extfloat.go ./ftoa.go ./isprint.go ./itoa.go ./quote.go
ar cru $WORK/github.com/h4ck3rm1k3/gocore/libstrconv.a $WORK/github.com/h4ck3rm1k3/gocore/strconv/_go_.o

#
# github.com/h4ck3rm1k3/gocore/image
#

#github.com/h4ck3rm1k3/gocore/image
mkdir -p $WORK/github.com/h4ck3rm1k3/gocore/image/
#COMPILER /home/h4ck3rm1k3/install/bin/gccgo
cd /home/h4ck3rm1k3/testgo2/src/github.com/h4ck3rm1k3/gocore/image
/home/h4ck3rm1k3/install/bin/gccgo -I $WORK -I /home/h4ck3rm1k3/testgo2/pkg/gccgo_linux_ppc64le -c -g -fgo-pkgpath=github.com/h4ck3rm1k3/gocore/image -fgo-relative-import-path=_/home/h4ck3rm1k3/testgo2/src/github.com/h4ck3rm1k3/gocore/image -o $WORK/github.com/h4ck3rm1k3/gocore/image/_go_.o -c  -O0 ./format.go ./geom.go ./image.go ./names.go ./ycbcr.go
ar cru $WORK/github.com/h4ck3rm1k3/gocore/libimage.a $WORK/github.com/h4ck3rm1k3/gocore/image/_go_.o
