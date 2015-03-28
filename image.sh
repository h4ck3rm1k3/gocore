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
mkdir -p $WORK/github.com/h4ck3rm1k3/gocore/errors/_obj/
mkdir -p $WORK/github.com/h4ck3rm1k3/gocore/
#COMPILER /home/h4ck3rm1k3/install/bin/gccgo
cd /home/h4ck3rm1k3/testgo2/src/github.com/h4ck3rm1k3/gocore/errors
/home/h4ck3rm1k3/install/bin/gccgo -I $WORK -c -g -fgo-pkgpath=github.com/h4ck3rm1k3/gocore/errors -fgo-relative-import-path=_/home/h4ck3rm1k3/testgo2/src/github.com/h4ck3rm1k3/gocore/errors -o $WORK/github.com/h4ck3rm1k3/gocore/errors/_obj/_go_.o -c --verbose -save-temps -O0 ./errors.go
/home/h4ck3rm1k3/install/bin/gccgo -I $WORK/github.com/h4ck3rm1k3/gocore/errors/_obj/ -o $WORK/github.com/h4ck3rm1k3/gocore/errors/_obj/errors.o -D GOOS_linux -D GOARCH_ppc64le -D "GOPKGPATH=\"github_com_h4ck3rm1k3_gocore_errors\"" ./errors.s
ar cru $WORK/github.com/h4ck3rm1k3/gocore/liberrors.a $WORK/github.com/h4ck3rm1k3/gocore/errors/_obj/_go_.o $WORK/github.com/h4ck3rm1k3/gocore/errors/_obj/errors.o

#
# github.com/h4ck3rm1k3/gocore/unsafe
#

#github.com/h4ck3rm1k3/gocore/unsafe
mkdir -p $WORK/github.com/h4ck3rm1k3/gocore/unsafe/_obj/
#COMPILER /home/h4ck3rm1k3/install/bin/gccgo
cd /home/h4ck3rm1k3/testgo2/src/github.com/h4ck3rm1k3/gocore/unsafe
/home/h4ck3rm1k3/install/bin/gccgo -I $WORK -c -g -fgo-pkgpath=github.com/h4ck3rm1k3/gocore/unsafe -fgo-relative-import-path=_/home/h4ck3rm1k3/testgo2/src/github.com/h4ck3rm1k3/gocore/unsafe -o $WORK/github.com/h4ck3rm1k3/gocore/unsafe/_obj/_go_.o -c --verbose -save-temps -O0 ./unsafe.go
/home/h4ck3rm1k3/install/bin/gccgo -I $WORK/github.com/h4ck3rm1k3/gocore/unsafe/_obj/ -o $WORK/github.com/h4ck3rm1k3/gocore/unsafe/_obj/unsafe.o -D GOOS_linux -D GOARCH_ppc64le -D "GOPKGPATH=\"github_com_h4ck3rm1k3_gocore_unsafe\"" ./unsafe.s
ar cru $WORK/github.com/h4ck3rm1k3/gocore/libunsafe.a $WORK/github.com/h4ck3rm1k3/gocore/unsafe/_obj/_go_.o $WORK/github.com/h4ck3rm1k3/gocore/unsafe/_obj/unsafe.o

#
# github.com/h4ck3rm1k3/gocore/runtime
#

#github.com/h4ck3rm1k3/gocore/runtime
mkdir -p $WORK/github.com/h4ck3rm1k3/gocore/runtime/_obj/
#COMPILER /home/h4ck3rm1k3/install/bin/gccgo
cd /home/h4ck3rm1k3/testgo2/src/github.com/h4ck3rm1k3/gocore/runtime
/home/h4ck3rm1k3/install/bin/gccgo -I $WORK -I /home/h4ck3rm1k3/testgo2/pkg/gccgo_linux_ppc64le -c -g -fgo-pkgpath=github.com/h4ck3rm1k3/gocore/runtime -fgo-relative-import-path=_/home/h4ck3rm1k3/testgo2/src/github.com/h4ck3rm1k3/gocore/runtime -o $WORK/github.com/h4ck3rm1k3/gocore/runtime/_obj/_go_.o -c --verbose -save-temps -O0 ./alg.go ./arch1_ppc64le.go ./arch_ppc64le.go ./atomic_pointer.go ./atomic_ppc64x.go ./cgo.go ./cgocall.go ./cgocallback.go ./chan.go ./compiler.go ./complex.go ./cpuprof.go ./cputicks.go ./debug.go ./defs_linux_ppc64le.go ./env_posix.go ./error.go ./extern.go ./hash64.go ./hashmap.go ./hashmap_fast.go ./heapdump.go ./iface.go ./lfstack.go ./lfstack_linux_ppc64x.go ./lock_futex.go ./malloc.go ./mbarrier.go ./mbitmap.go ./mcache.go ./mcentral.go ./mem_linux.go ./mfinal.go ./mfixalloc.go ./mgc.go ./mgcmark.go ./mgcsweep.go ./mgcwork.go ./mheap.go ./mprof.go ./msize.go ./mstats.go ./netpoll.go ./netpoll_epoll.go ./noasm.go ./os1_linux.go ./os2_linux.go ./os_linux.go ./panic.go ./panic1.go ./parfor.go ./print1.go ./print1_write.go ./proc.go ./proc1.go ./race0.go ./rdebug.go ./rune.go ./runtime.go ./runtime1.go ./runtime2.go ./select.go ./sema.go ./signal1_unix.go ./signal_linux.go ./signal_linux_ppc64x.go ./signal_ppc64x.go ./signal_unix.go ./sigpanic_unix.go ./sigqueue.go ./slice.go ./softfloat64.go ./sqrt.go ./stack1.go ./stack2.go ./string.go ./string1.go ./stubs.go ./stubs2.go ./symtab.go ./sys_ppc64x.go ./time.go ./trace.go ./traceback.go ./type.go ./typekind.go ./typekind1.go ./unaligned2.go ./vdso_none.go ./wbfat.go ./zgoarch_ppc64le.go ./zgoos_linux.go
/home/h4ck3rm1k3/install/bin/gccgo -I $WORK/github.com/h4ck3rm1k3/gocore/runtime/_obj/ -o $WORK/github.com/h4ck3rm1k3/gocore/runtime/_obj/alg.o -D GOOS_linux -D GOARCH_ppc64le -D "GOPKGPATH=\"github_com_h4ck3rm1k3_gocore_runtime\"" ./alg.s
ar cru $WORK/github.com/h4ck3rm1k3/gocore/libruntime.a $WORK/github.com/h4ck3rm1k3/gocore/runtime/_obj/_go_.o $WORK/github.com/h4ck3rm1k3/gocore/runtime/_obj/alg.o

#
# github.com/h4ck3rm1k3/gocore/sync/atomic
#

#github.com/h4ck3rm1k3/gocore/sync/atomic
mkdir -p $WORK/github.com/h4ck3rm1k3/gocore/sync/atomic/_obj/
mkdir -p $WORK/github.com/h4ck3rm1k3/gocore/sync/
#COMPILER /home/h4ck3rm1k3/install/bin/gccgo
cd /home/h4ck3rm1k3/testgo2/src/github.com/h4ck3rm1k3/gocore/sync/atomic
/home/h4ck3rm1k3/install/bin/gccgo -I $WORK -I /home/h4ck3rm1k3/testgo2/pkg/gccgo_linux_ppc64le -c -g -fgo-pkgpath=github.com/h4ck3rm1k3/gocore/sync/atomic -fgo-relative-import-path=_/home/h4ck3rm1k3/testgo2/src/github.com/h4ck3rm1k3/gocore/sync/atomic -o $WORK/github.com/h4ck3rm1k3/gocore/sync/atomic/_obj/_go_.o -c --verbose -save-temps -O0 ./doc.go ./value.go
/home/h4ck3rm1k3/install/bin/gccgo -I $WORK/github.com/h4ck3rm1k3/gocore/sync/atomic/_obj/ -o $WORK/github.com/h4ck3rm1k3/gocore/sync/atomic/_obj/doc.o -D GOOS_linux -D GOARCH_ppc64le -D "GOPKGPATH=\"github_com_h4ck3rm1k3_gocore_sync_atomic\"" ./doc.s
ar cru $WORK/github.com/h4ck3rm1k3/gocore/sync/libatomic.a $WORK/github.com/h4ck3rm1k3/gocore/sync/atomic/_obj/_go_.o $WORK/github.com/h4ck3rm1k3/gocore/sync/atomic/_obj/doc.o

#
# github.com/h4ck3rm1k3/gocore/sync
#

#github.com/h4ck3rm1k3/gocore/sync
mkdir -p $WORK/github.com/h4ck3rm1k3/gocore/sync/_obj/
#COMPILER /home/h4ck3rm1k3/install/bin/gccgo
cd /home/h4ck3rm1k3/testgo2/src/github.com/h4ck3rm1k3/gocore/sync
/home/h4ck3rm1k3/install/bin/gccgo -I $WORK -I /home/h4ck3rm1k3/testgo2/pkg/gccgo_linux_ppc64le -c -g -fgo-pkgpath=github.com/h4ck3rm1k3/gocore/sync -fgo-relative-import-path=_/home/h4ck3rm1k3/testgo2/src/github.com/h4ck3rm1k3/gocore/sync -o $WORK/github.com/h4ck3rm1k3/gocore/sync/_obj/_go_.o -c --verbose -save-temps -O0 ./cond.go ./mutex.go ./once.go ./pool.go ./race0.go ./runtime.go ./rwmutex.go ./waitgroup.go
ar cru $WORK/github.com/h4ck3rm1k3/gocore/libsync.a $WORK/github.com/h4ck3rm1k3/gocore/sync/_obj/_go_.o

#
# github.com/h4ck3rm1k3/gocore/io
#

#github.com/h4ck3rm1k3/gocore/io
mkdir -p $WORK/github.com/h4ck3rm1k3/gocore/io/_obj/
#COMPILER /home/h4ck3rm1k3/install/bin/gccgo
cd /home/h4ck3rm1k3/testgo2/src/github.com/h4ck3rm1k3/gocore/io
/home/h4ck3rm1k3/install/bin/gccgo -I $WORK -I /home/h4ck3rm1k3/testgo2/pkg/gccgo_linux_ppc64le -c -g -fgo-pkgpath=github.com/h4ck3rm1k3/gocore/io -fgo-relative-import-path=_/home/h4ck3rm1k3/testgo2/src/github.com/h4ck3rm1k3/gocore/io -o $WORK/github.com/h4ck3rm1k3/gocore/io/_obj/_go_.o -c --verbose -save-temps -O0 ./io.go ./multi.go ./pipe.go
ar cru $WORK/github.com/h4ck3rm1k3/gocore/libio.a $WORK/github.com/h4ck3rm1k3/gocore/io/_obj/_go_.o

#
# github.com/h4ck3rm1k3/gocore/unicode
#

#github.com/h4ck3rm1k3/gocore/unicode
mkdir -p $WORK/github.com/h4ck3rm1k3/gocore/unicode/_obj/
#COMPILER /home/h4ck3rm1k3/install/bin/gccgo
cd /home/h4ck3rm1k3/testgo2/src/github.com/h4ck3rm1k3/gocore/unicode
/home/h4ck3rm1k3/install/bin/gccgo -I $WORK -c -g -fgo-pkgpath=github.com/h4ck3rm1k3/gocore/unicode -fgo-relative-import-path=_/home/h4ck3rm1k3/testgo2/src/github.com/h4ck3rm1k3/gocore/unicode -o $WORK/github.com/h4ck3rm1k3/gocore/unicode/_obj/_go_.o -c --verbose -save-temps -O0 ./casetables.go ./digit.go ./graphic.go ./letter.go ./tables.go
/home/h4ck3rm1k3/install/bin/gccgo -I $WORK/github.com/h4ck3rm1k3/gocore/unicode/_obj/ -o $WORK/github.com/h4ck3rm1k3/gocore/unicode/_obj/casetables.o -D GOOS_linux -D GOARCH_ppc64le -D "GOPKGPATH=\"github_com_h4ck3rm1k3_gocore_unicode\"" ./casetables.s
ar cru $WORK/github.com/h4ck3rm1k3/gocore/libunicode.a $WORK/github.com/h4ck3rm1k3/gocore/unicode/_obj/_go_.o $WORK/github.com/h4ck3rm1k3/gocore/unicode/_obj/casetables.o

#
# github.com/h4ck3rm1k3/gocore/unicode/utf8
#

#github.com/h4ck3rm1k3/gocore/unicode/utf8
mkdir -p $WORK/github.com/h4ck3rm1k3/gocore/unicode/utf8/_obj/
mkdir -p $WORK/github.com/h4ck3rm1k3/gocore/unicode/
#COMPILER /home/h4ck3rm1k3/install/bin/gccgo
cd /home/h4ck3rm1k3/testgo2/src/github.com/h4ck3rm1k3/gocore/unicode/utf8
/home/h4ck3rm1k3/install/bin/gccgo -I $WORK -c -g -fgo-pkgpath=github.com/h4ck3rm1k3/gocore/unicode/utf8 -fgo-relative-import-path=_/home/h4ck3rm1k3/testgo2/src/github.com/h4ck3rm1k3/gocore/unicode/utf8 -o $WORK/github.com/h4ck3rm1k3/gocore/unicode/utf8/_obj/_go_.o -c --verbose -save-temps -O0 ./utf8.go
/home/h4ck3rm1k3/install/bin/gccgo -I $WORK/github.com/h4ck3rm1k3/gocore/unicode/utf8/_obj/ -o $WORK/github.com/h4ck3rm1k3/gocore/unicode/utf8/_obj/utf8.o -D GOOS_linux -D GOARCH_ppc64le -D "GOPKGPATH=\"github_com_h4ck3rm1k3_gocore_unicode_utf8\"" ./utf8.s
ar cru $WORK/github.com/h4ck3rm1k3/gocore/unicode/libutf8.a $WORK/github.com/h4ck3rm1k3/gocore/unicode/utf8/_obj/_go_.o $WORK/github.com/h4ck3rm1k3/gocore/unicode/utf8/_obj/utf8.o

#
# github.com/h4ck3rm1k3/gocore/bytes
#

#github.com/h4ck3rm1k3/gocore/bytes
mkdir -p $WORK/github.com/h4ck3rm1k3/gocore/bytes/_obj/
#COMPILER /home/h4ck3rm1k3/install/bin/gccgo
cd /home/h4ck3rm1k3/testgo2/src/github.com/h4ck3rm1k3/gocore/bytes
/home/h4ck3rm1k3/install/bin/gccgo -I $WORK -I /home/h4ck3rm1k3/testgo2/pkg/gccgo_linux_ppc64le -c -g -fgo-pkgpath=github.com/h4ck3rm1k3/gocore/bytes -fgo-relative-import-path=_/home/h4ck3rm1k3/testgo2/src/github.com/h4ck3rm1k3/gocore/bytes -o $WORK/github.com/h4ck3rm1k3/gocore/bytes/_obj/_go_.o -c --verbose -save-temps -O0 ./buffer.go ./bytes.go ./bytes_decl.go ./reader.go
ar cru $WORK/github.com/h4ck3rm1k3/gocore/libbytes.a $WORK/github.com/h4ck3rm1k3/gocore/bytes/_obj/_go_.o

#
# github.com/h4ck3rm1k3/gocore/bufio
#

#github.com/h4ck3rm1k3/gocore/bufio
mkdir -p $WORK/github.com/h4ck3rm1k3/gocore/bufio/_obj/
#COMPILER /home/h4ck3rm1k3/install/bin/gccgo
cd /home/h4ck3rm1k3/testgo2/src/github.com/h4ck3rm1k3/gocore/bufio
/home/h4ck3rm1k3/install/bin/gccgo -I $WORK -I /home/h4ck3rm1k3/testgo2/pkg/gccgo_linux_ppc64le -c -g -fgo-pkgpath=github.com/h4ck3rm1k3/gocore/bufio -fgo-relative-import-path=_/home/h4ck3rm1k3/testgo2/src/github.com/h4ck3rm1k3/gocore/bufio -o $WORK/github.com/h4ck3rm1k3/gocore/bufio/_obj/_go_.o -c --verbose -save-temps -O0 ./bufio.go ./scan.go
ar cru $WORK/github.com/h4ck3rm1k3/gocore/libbufio.a $WORK/github.com/h4ck3rm1k3/gocore/bufio/_obj/_go_.o

#
# github.com/h4ck3rm1k3/gocore/image/color
#

#github.com/h4ck3rm1k3/gocore/image/color
mkdir -p $WORK/github.com/h4ck3rm1k3/gocore/image/color/_obj/
mkdir -p $WORK/github.com/h4ck3rm1k3/gocore/image/
#COMPILER /home/h4ck3rm1k3/install/bin/gccgo
cd /home/h4ck3rm1k3/testgo2/src/github.com/h4ck3rm1k3/gocore/image/color
/home/h4ck3rm1k3/install/bin/gccgo -I $WORK -c -g -fgo-pkgpath=github.com/h4ck3rm1k3/gocore/image/color -fgo-relative-import-path=_/home/h4ck3rm1k3/testgo2/src/github.com/h4ck3rm1k3/gocore/image/color -o $WORK/github.com/h4ck3rm1k3/gocore/image/color/_obj/_go_.o -c --verbose -save-temps -O0 ./color.go ./ycbcr.go
/home/h4ck3rm1k3/install/bin/gccgo -I $WORK/github.com/h4ck3rm1k3/gocore/image/color/_obj/ -o $WORK/github.com/h4ck3rm1k3/gocore/image/color/_obj/color.o -D GOOS_linux -D GOARCH_ppc64le -D "GOPKGPATH=\"github_com_h4ck3rm1k3_gocore_image_color\"" ./color.s
ar cru $WORK/github.com/h4ck3rm1k3/gocore/image/libcolor.a $WORK/github.com/h4ck3rm1k3/gocore/image/color/_obj/_go_.o $WORK/github.com/h4ck3rm1k3/gocore/image/color/_obj/color.o

#
# github.com/h4ck3rm1k3/gocore/math
#

#github.com/h4ck3rm1k3/gocore/math
mkdir -p $WORK/github.com/h4ck3rm1k3/gocore/math/_obj/
#COMPILER /home/h4ck3rm1k3/install/bin/gccgo
cd /home/h4ck3rm1k3/testgo2/src/github.com/h4ck3rm1k3/gocore/math
/home/h4ck3rm1k3/install/bin/gccgo -I $WORK -I /home/h4ck3rm1k3/testgo2/pkg/gccgo_linux_ppc64le -c -g -fgo-pkgpath=github.com/h4ck3rm1k3/gocore/math -fgo-relative-import-path=_/home/h4ck3rm1k3/testgo2/src/github.com/h4ck3rm1k3/gocore/math -o $WORK/github.com/h4ck3rm1k3/gocore/math/_obj/_go_.o -c --verbose -save-temps -O0 ./abs.go ./acosh.go ./asin.go ./asinh.go ./atan.go ./atan2.go ./atanh.go ./bits.go ./cbrt.go ./const.go ./copysign.go ./dim.go ./erf.go ./exp.go ./expm1.go ./floor.go ./frexp.go ./gamma.go ./hypot.go ./j0.go ./j1.go ./jn.go ./ldexp.go ./lgamma.go ./log.go ./log10.go ./log1p.go ./logb.go ./mod.go ./modf.go ./nextafter.go ./pow.go ./pow10.go ./remainder.go ./signbit.go ./sin.go ./sincos.go ./sinh.go ./sqrt.go ./tan.go ./tanh.go ./unsafe.go
/home/h4ck3rm1k3/install/bin/gccgo -I $WORK/github.com/h4ck3rm1k3/gocore/math/_obj/ -o $WORK/github.com/h4ck3rm1k3/gocore/math/_obj/abs.o -D GOOS_linux -D GOARCH_ppc64le -D "GOPKGPATH=\"github_com_h4ck3rm1k3_gocore_math\"" ./abs.s
ar cru $WORK/github.com/h4ck3rm1k3/gocore/libmath.a $WORK/github.com/h4ck3rm1k3/gocore/math/_obj/_go_.o $WORK/github.com/h4ck3rm1k3/gocore/math/_obj/abs.o

#
# github.com/h4ck3rm1k3/gocore/strconv
#

#github.com/h4ck3rm1k3/gocore/strconv
mkdir -p $WORK/github.com/h4ck3rm1k3/gocore/strconv/_obj/
#COMPILER /home/h4ck3rm1k3/install/bin/gccgo
cd /home/h4ck3rm1k3/testgo2/src/github.com/h4ck3rm1k3/gocore/strconv
/home/h4ck3rm1k3/install/bin/gccgo -I $WORK -I /home/h4ck3rm1k3/testgo2/pkg/gccgo_linux_ppc64le -c -g -fgo-pkgpath=github.com/h4ck3rm1k3/gocore/strconv -fgo-relative-import-path=_/home/h4ck3rm1k3/testgo2/src/github.com/h4ck3rm1k3/gocore/strconv -o $WORK/github.com/h4ck3rm1k3/gocore/strconv/_obj/_go_.o -c --verbose -save-temps -O0 ./atob.go ./atof.go ./atoi.go ./decimal.go ./extfloat.go ./ftoa.go ./isprint.go ./itoa.go ./quote.go
ar cru $WORK/github.com/h4ck3rm1k3/gocore/libstrconv.a $WORK/github.com/h4ck3rm1k3/gocore/strconv/_obj/_go_.o

#
# github.com/h4ck3rm1k3/gocore/image
#

#github.com/h4ck3rm1k3/gocore/image
mkdir -p $WORK/github.com/h4ck3rm1k3/gocore/image/_obj/
#COMPILER /home/h4ck3rm1k3/install/bin/gccgo
cd /home/h4ck3rm1k3/testgo2/src/github.com/h4ck3rm1k3/gocore/image
/home/h4ck3rm1k3/install/bin/gccgo -I $WORK -I /home/h4ck3rm1k3/testgo2/pkg/gccgo_linux_ppc64le -c -g -fgo-pkgpath=github.com/h4ck3rm1k3/gocore/image -fgo-relative-import-path=_/home/h4ck3rm1k3/testgo2/src/github.com/h4ck3rm1k3/gocore/image -o $WORK/github.com/h4ck3rm1k3/gocore/image/_obj/_go_.o -c --verbose -save-temps -O0 ./format.go ./geom.go ./image.go ./names.go ./ycbcr.go
ar cru $WORK/github.com/h4ck3rm1k3/gocore/libimage.a $WORK/github.com/h4ck3rm1k3/gocore/image/_obj/_go_.o
