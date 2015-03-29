export WORK=/tmp/go
#github.com/h4ck3rm1k3/gocore/unsafe
mkdir -p $WORK/github.com/h4ck3rm1k3/gocore/unsafe/_obj/
#COMPILER /home/h4ck3rm1k3/install/bin/gccgo
cd /home/h4ck3rm1k3/testgo2/src/github.com/h4ck3rm1k3/gocore/unsafe

/home/h4ck3rm1k3/install/bin/gccgo -I $WORK -c -g -fgo-pkgpath=github.com/h4ck3rm1k3/gocore/unsafe -fgo-relative-import-path=_/home/h4ck3rm1k3/testgo2/src/github.com/h4ck3rm1k3/gocore/unsafe -o $WORK/github.com/h4ck3rm1k3/gocore/unsafe/_obj/_go_.o -c --verbose -save-temps -O0 ./unsafe.go
/home/h4ck3rm1k3/install/bin/gccgo -I $WORK/github.com/h4ck3rm1k3/gocore/unsafe/_obj/ -o $WORK/github.com/h4ck3rm1k3/gocore/unsafe/unsafe.o -D GOOS_linux -D GOARCH_ppc64le -D "GOPKGPATH=\"github_com_h4ck3rm1k3_gocore_unsafe\"" ./unsafe.s
ar cru $WORK/github.com/h4ck3rm1k3/gocore/libunsafe.a $WORK/github.com/h4ck3rm1k3/gocore/unsafe/_go_.o $WORK/github.com/h4ck3rm1k3/gocore/unsafe/unsafe.o
